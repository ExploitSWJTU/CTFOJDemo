package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	dockercontainer "github.com/docker/docker/api/types/container"
	dockerstrslice "github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
	"github.com/docker/docker/errdefs"
	"github.com/docker/go-connections/nat"
	"gorm.io/gorm"

	"swjtu-ctf-oj/internal/model"
	"swjtu-ctf-oj/internal/repository"
)

var ErrActiveContainerNotFound = errors.New("active container not found")

type ChallengeService struct {
	db           *repository.Database
	mu           sync.Mutex
	dockerClient *client.Client
	dockerOnce   sync.Once
	dockerErr    error
	config       ChallengeRuntimeConfig
}

type ChallengeRuntimeConfig struct {
	DockerHost             string
	PublicHost             string
	PublicScheme           string
	PortRangeStart         int
	PortRangeEnd           int
	DefaultDurationSeconds int
}

type ChallengeContainerResponse struct {
	ID          uint      `json:"id"`
	ChallengeID uint     `json:"challenge_id"`
	Status      string    `json:"status"`
	Host        string    `json:"host"`
	Port        int       `json:"port"`
	AccessURL   string    `json:"access_url"`
	StartedAt   time.Time `json:"started_at"`
	ExpiresAt   time.Time `json:"expires_at"`
	TimeLeft    string    `json:"time_left"`
	LastError   string    `json:"last_error,omitempty"`
}

type AdminInstanceResponse struct {
	ID             uint       `json:"id"`
	UserID         uint       `json:"user_id"`
	Username       string     `json:"username"`
	ChallengeID    uint       `json:"challenge_id"`
	ChallengeTitle string     `json:"challenge_title"`
	Status         string     `json:"status"`
	ContainerID    string     `json:"container_id"`
	Host           string     `json:"host"`
	Port           int        `json:"port"`
	AccessURL      string     `json:"access_url"`
	StartedAt      *time.Time `json:"started_at,omitempty"`
	ExpiresAt      time.Time  `json:"expires_at"`
	TimeLeft       string     `json:"time_left"`
	LastError      string     `json:"last_error,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type FlagSubmissionResult struct {
	Correct              bool   `json:"correct"`
	AlreadySolved        bool   `json:"already_solved"`
	Message              string `json:"message"`
	UserScore            int    `json:"user_score"`
	ChallengeSolvedCount int    `json:"challenge_solved_count"`
}

type AdminDeleteResult struct {
	Deleted              bool `json:"deleted"`
	DestroyedContainers  int  `json:"destroyed_containers"`
}

type challengeRuntimeSpec struct {
	Image           string
	InternalPort    int
	Command         []string
	Env             []string
	ExpectedFlag    string
	DurationSeconds int
}

type adminContainerQueryRow struct {
	ID             uint
	UserID         uint
	Username       string
	ChallengeID    uint
	ChallengeTitle string
	ContainerID    string
	HostPort       int
	AccessHost     string
	Status         string
	LastError      string
	StartedAt      *time.Time
	ExpiresAt      time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewChallengeService(db *repository.Database) *ChallengeService {
	return &ChallengeService{
		db: db,
		config: ChallengeRuntimeConfig{
			DockerHost:             os.Getenv("CHALLENGE_DOCKER_HOST"),
			PublicHost:             getEnvOrDefault("CHALLENGE_PUBLIC_HOST", "127.0.0.1"),
			PublicScheme:           getEnvOrDefault("CHALLENGE_PUBLIC_SCHEME", "http"),
			PortRangeStart:         getEnvIntOrDefault("DOCKER_PORT_RANGE_START", 30000),
			PortRangeEnd:           getEnvIntOrDefault("DOCKER_PORT_RANGE_END", 40000),
			DefaultDurationSeconds: getEnvIntOrDefault("CHALLENGE_CONTAINER_DURATION_SECONDS", 3600),
		},
	}
}

func (s *ChallengeService) StartContainer(userID, challengeID uint) (*model.Container, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	challenge, err := s.getChallenge(challengeID, false)
	if err != nil {
		return nil, err
	}

	spec, err := s.buildRuntimeSpec(challenge)
	if err != nil {
		return nil, err
	}

	existing, err := s.getLatestActiveContainer(userID, challengeID)
	if err == nil {
		if existing.IsExpired() {
			if destroyErr := s.destroyContainerLocked(existing, model.ContainerStatusExpired); destroyErr != nil {
				return nil, destroyErr
			}
		} else {
			return existing, nil
		}
	} else if !errors.Is(err, ErrActiveContainerNotFound) {
		return nil, err
	}

	cli, err := s.getDockerClient()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if _, _, err := cli.ImageInspectWithRaw(ctx, spec.Image); err != nil {
		return nil, fmt.Errorf("challenge image %q not found on challenge daemon: %w", spec.Image, err)
	}

	hostPort, err := s.allocateHostPortLocked()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	expiresAt := now.Add(time.Duration(spec.DurationSeconds) * time.Second)
	portKey := nat.Port(fmt.Sprintf("%d/tcp", spec.InternalPort))
	portMapping, _ := json.Marshal(model.PortMapping{
		ContainerPort: spec.InternalPort,
		HostPort:      hostPort,
	})

	record := model.Container{
		UserID:      userID,
		ChallengeID: challengeID,
		HostPort:    hostPort,
		AccessHost:  s.config.PublicHost,
		PortMapping: portMapping,
		Flag:        spec.ExpectedFlag,
		Status:      string(model.ContainerStatusCreating),
		StartedAt:   &now,
		ExpiresAt:   expiresAt,
	}
	if err := s.db.DB.Create(&record).Error; err != nil {
		return nil, err
	}

	containerName := fmt.Sprintf("ctfoj-u%d-c%d-%d", userID, challengeID, now.Unix())
	containerConfig := &dockercontainer.Config{
		Image:        spec.Image,
		Env:          spec.Env,
		ExposedPorts: nat.PortSet{portKey: struct{}{}},
	}
	if len(spec.Command) > 0 {
		containerConfig.Cmd = dockerstrslice.StrSlice(spec.Command)
	}

	hostConfig := &dockercontainer.HostConfig{
		PortBindings: nat.PortMap{
			portKey: []nat.PortBinding{{
				HostIP:   "0.0.0.0",
				HostPort: strconv.Itoa(hostPort),
			}},
		},
	}

	created, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, containerName)
	if err != nil {
		s.markContainerFailed(&record, err)
		return nil, err
	}

	record.ContainerID = created.ID
	if err := s.db.DB.Model(&record).Update("container_id", created.ID).Error; err != nil {
		return nil, err
	}

	if err := cli.ContainerStart(ctx, created.ID, dockercontainer.StartOptions{}); err != nil {
		_ = cli.ContainerRemove(context.Background(), created.ID, dockercontainer.RemoveOptions{Force: true})
		s.markContainerFailed(&record, err)
		return nil, err
	}

	record.Status = string(model.ContainerStatusRunning)
	record.LastError = ""
	if err := s.db.DB.Model(&record).Updates(map[string]any{
		"status":     record.Status,
		"last_error": "",
	}).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

func (s *ChallengeService) GetContainer(userID, challengeID uint) (*model.Container, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	record, err := s.getLatestActiveContainer(userID, challengeID)
	if err != nil {
		return nil, err
	}

	if record.IsExpired() {
		if err := s.destroyContainerLocked(record, model.ContainerStatusExpired); err != nil {
			return nil, err
		}
		return nil, ErrActiveContainerNotFound
	}

	if record.ContainerID == "" {
		return record, nil
	}

	cli, err := s.getDockerClient()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	info, err := cli.ContainerInspect(ctx, record.ContainerID)
	if err != nil {
		if errdefs.IsNotFound(err) {
			if err := s.destroyContainerLocked(record, model.ContainerStatusStopped); err != nil {
				return nil, err
			}
			return nil, ErrActiveContainerNotFound
		}
		return nil, err
	}

	if info.State != nil && !info.State.Running {
		if err := s.destroyContainerLocked(record, model.ContainerStatusStopped); err != nil {
			return nil, err
		}
		return nil, ErrActiveContainerNotFound
	}

	return record, nil
}

func (s *ChallengeService) ExtendContainer(userID, challengeID uint) (*model.Container, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	record, err := s.getLatestActiveContainer(userID, challengeID)
	if err != nil {
		return nil, err
	}

	challenge, err := s.getChallenge(challengeID, true)
	if err != nil {
		return nil, err
	}

	spec, err := s.buildRuntimeSpec(challenge)
	if err != nil {
		return nil, err
	}

	if record.IsExpired() {
		if err := s.destroyContainerLocked(record, model.ContainerStatusExpired); err != nil {
			return nil, err
		}
		return nil, ErrActiveContainerNotFound
	}

	record.ExpiresAt = record.ExpiresAt.Add(time.Duration(spec.DurationSeconds) * time.Second)
	if err := s.db.DB.Model(record).Update("expires_at", record.ExpiresAt).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (s *ChallengeService) DestroyContainer(userID, challengeID uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	record, err := s.getLatestActiveContainer(userID, challengeID)
	if err != nil {
		if errors.Is(err, ErrActiveContainerNotFound) {
			return nil
		}
		return err
	}

	return s.destroyContainerLocked(record, model.ContainerStatusStopped)
}

func (s *ChallengeService) CleanupExpiredContainers() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var records []model.Container
	if err := s.db.DB.
		Where("status IN ? AND expires_at <= ?", []string{string(model.ContainerStatusRunning), string(model.ContainerStatusCreating)}, time.Now()).
		Find(&records).Error; err != nil {
		return err
	}

	for i := range records {
		if err := s.destroyContainerLocked(&records[i], model.ContainerStatusExpired); err != nil {
			return err
		}
	}

	return nil
}

func (s *ChallengeService) DestroyChallengeContainers(challengeID uint) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var records []model.Container
	if err := s.db.DB.
		Where("challenge_id = ? AND status IN ?", challengeID, []string{string(model.ContainerStatusRunning), string(model.ContainerStatusCreating)}).
		Find(&records).Error; err != nil {
		return 0, err
	}

	destroyed := 0
	for i := range records {
		if err := s.destroyContainerLocked(&records[i], model.ContainerStatusStopped); err != nil {
			return destroyed, err
		}
		destroyed++
	}

	return destroyed, nil
}

func (s *ChallengeService) ListAdminContainers() ([]AdminInstanceResponse, error) {
	rows, err := s.loadAdminContainerRows(0)
	if err != nil {
		return nil, err
	}

	items := make([]AdminInstanceResponse, 0, len(rows))
	for i := range rows {
		items = append(items, s.buildAdminInstanceResponse(rows[i]))
	}
	return items, nil
}

func (s *ChallengeService) AdminDestroyContainer(containerRecordID uint) (*AdminInstanceResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var record model.Container
	if err := s.db.DB.First(&record, containerRecordID).Error; err != nil {
		return nil, err
	}

	nextStatus := adminStatusAfterDestroy(record.Status)
	if err := s.destroyContainerLocked(&record, nextStatus); err != nil {
		return nil, err
	}

	rows, err := s.loadAdminContainerRows(containerRecordID)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	response := s.buildAdminInstanceResponse(rows[0])
	return &response, nil
}

func (s *ChallengeService) SubmitFlag(userID, challengeID uint, submittedFlag string) (*FlagSubmissionResult, error) {
	challenge, err := s.getChallenge(challengeID, false)
	if err != nil {
		return nil, err
	}

	spec, err := s.buildRuntimeSpec(challenge)
	if err != nil {
		return nil, err
	}

	normalizedSubmitted := strings.TrimSpace(submittedFlag)
	normalizedExpected := strings.TrimSpace(spec.ExpectedFlag)
	isCorrect := normalizedExpected != "" && normalizedSubmitted == normalizedExpected

	result := &FlagSubmissionResult{
		Correct: isCorrect,
		Message: "Flag incorrect.",
	}

	tx := s.db.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	submission := model.Submission{
		UserID:        userID,
		ChallengeID:   challengeID,
		SubmittedFlag: normalizedSubmitted,
		IsCorrect:     isCorrect,
		IsSolve:       false,
	}
	if err := tx.Create(&submission).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var user model.User
	if err := tx.First(&user, userID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if isCorrect {
		var existingCorrect int64
		if err := tx.Model(&model.Submission{}).
			Where("user_id = ? AND challenge_id = ? AND is_correct = ?", userID, challengeID, true).
			Count(&existingCorrect).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		if existingCorrect > 1 {
			result.AlreadySolved = true
			result.Message = "Flag correct, but this challenge was already solved."
		} else {
			if err := tx.Model(&submission).Update("is_solve", true).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
			if err := tx.Model(&model.User{}).Where("id = ?", userID).
				Update("score", gorm.Expr("score + ?", challenge.Points)).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
			if err := tx.Model(&model.Challenge{}).Where("id = ?", challengeID).
				Update("solved_count", gorm.Expr("solved_count + 1")).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
			user.Score += challenge.Points
			challenge.SolvedCount++
			result.Message = "Flag correct."
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	result.UserScore = user.Score
	result.ChallengeSolvedCount = challenge.SolvedCount
	return result, nil
}

func (s *ChallengeService) BuildContainerResponse(record *model.Container) ChallengeContainerResponse {
	startedAt := record.CreatedAt
	if record.StartedAt != nil {
		startedAt = *record.StartedAt
	}
	host := record.AccessHost
	if strings.TrimSpace(host) == "" {
		host = s.config.PublicHost
	}

	return ChallengeContainerResponse{
		ID:          record.ID,
		ChallengeID: record.ChallengeID,
		Status:      record.Status,
		Host:        host,
		Port:        record.HostPort,
		AccessURL:   s.accessURL(record),
		StartedAt:   startedAt,
		ExpiresAt:   record.ExpiresAt,
		TimeLeft:    record.TimeLeft(),
		LastError:   record.LastError,
	}
}

func (s *ChallengeService) buildAdminInstanceResponse(row adminContainerQueryRow) AdminInstanceResponse {
	record := model.Container{
		HostPort:   row.HostPort,
		AccessHost: row.AccessHost,
		ExpiresAt:  row.ExpiresAt,
	}

	return AdminInstanceResponse{
		ID:             row.ID,
		UserID:         row.UserID,
		Username:       row.Username,
		ChallengeID:    row.ChallengeID,
		ChallengeTitle: row.ChallengeTitle,
		Status:         row.Status,
		ContainerID:    row.ContainerID,
		Host:           adminAccessHost(row.AccessHost, s.config.PublicHost),
		Port:           row.HostPort,
		AccessURL:      s.accessURL(&record),
		StartedAt:      row.StartedAt,
		ExpiresAt:      row.ExpiresAt,
		TimeLeft:       adminTimeLeft(row.Status, &record),
		LastError:      row.LastError,
		CreatedAt:      row.CreatedAt,
		UpdatedAt:      row.UpdatedAt,
	}
}

func (s *ChallengeService) loadAdminContainerRows(containerRecordID uint) ([]adminContainerQueryRow, error) {
	var rows []adminContainerQueryRow

	query := s.db.DB.Table("containers").
		Select(`
			containers.id,
			containers.user_id,
			COALESCE(users.username, '') AS username,
			containers.challenge_id,
			COALESCE(challenges.title, '') AS challenge_title,
			containers.container_id,
			containers.host_port,
			containers.access_host,
			containers.status,
			containers.last_error,
			containers.started_at,
			containers.expires_at,
			containers.created_at,
			containers.updated_at
		`).
		Joins("LEFT JOIN users ON users.id = containers.user_id").
		Joins("LEFT JOIN challenges ON challenges.id = containers.challenge_id").
		Where("containers.deleted_at IS NULL").
		Order("CASE WHEN containers.status IN ('running', 'creating') THEN 0 ELSE 1 END, containers.created_at DESC")

	if containerRecordID > 0 {
		query = query.Where("containers.id = ?", containerRecordID)
	}

	if err := query.Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (s *ChallengeService) getChallenge(challengeID uint, includeHidden bool) (*model.Challenge, error) {
	var challenge model.Challenge
	query := s.db.DB
	if !includeHidden {
		query = query.Where("is_visible = ?", true)
	}
	if err := query.First(&challenge, challengeID).Error; err != nil {
		return nil, err
	}
	return &challenge, nil
}

func (s *ChallengeService) buildRuntimeSpec(challenge *model.Challenge) (*challengeRuntimeSpec, error) {
	var containerConfig model.ContainerConfig
	if len(challenge.ContainerConfig) > 0 {
		if err := json.Unmarshal(challenge.ContainerConfig, &containerConfig); err != nil {
			return nil, fmt.Errorf("invalid container config: %w", err)
		}
	}

	image := challenge.Image
	if image == "" {
		image = containerConfig.Image
	}

	internalPort := challenge.InternalPort
	if internalPort <= 0 {
		internalPort = containerConfig.ExposedPort
	}

	expectedFlag := challenge.ExpectedFlag
	if expectedFlag == "" {
		expectedFlag = challenge.Flag
	}

	duration := challenge.ContainerDurationSeconds
	if duration <= 0 {
		duration = s.config.DefaultDurationSeconds
	}

	if image == "" {
		return nil, errors.New("challenge image is required")
	}
	if internalPort <= 0 {
		return nil, errors.New("challenge internal port must be greater than 0")
	}
	if expectedFlag == "" {
		return nil, errors.New("challenge expected flag is required")
	}

	spec := &challengeRuntimeSpec{
		Image:           image,
		InternalPort:    internalPort,
		Env:             append([]string{}, containerConfig.Env...),
		ExpectedFlag:    expectedFlag,
		DurationSeconds: duration,
	}
	spec.Env = append(spec.Env, "CHALLENGE_EXPECTED_FLAG="+expectedFlag)
	if command := strings.TrimSpace(containerConfig.Command); command != "" {
		spec.Command = strings.Fields(command)
	}
	return spec, nil
}

func (s *ChallengeService) getLatestActiveContainer(userID, challengeID uint) (*model.Container, error) {
	var record model.Container
	err := s.db.DB.
		Where("user_id = ? AND challenge_id = ? AND status IN ?", userID, challengeID, []string{string(model.ContainerStatusRunning), string(model.ContainerStatusCreating)}).
		Order("created_at DESC").
		First(&record).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrActiveContainerNotFound
		}
		return nil, err
	}
	return &record, nil
}

func (s *ChallengeService) allocateHostPortLocked() (int, error) {
	var records []model.Container
	if err := s.db.DB.
		Where("status IN ? AND expires_at > ? AND host_port >= ? AND host_port <= ?", []string{string(model.ContainerStatusRunning), string(model.ContainerStatusCreating)}, time.Now(), s.config.PortRangeStart, s.config.PortRangeEnd).
		Find(&records).Error; err != nil {
		return 0, err
	}

	used := make(map[int]struct{}, len(records))
	for _, record := range records {
		if record.HostPort > 0 {
			used[record.HostPort] = struct{}{}
		}
	}

	for port := s.config.PortRangeStart; port <= s.config.PortRangeEnd; port++ {
		if _, exists := used[port]; !exists {
			return port, nil
		}
	}

	return 0, errors.New("no free challenge port is available")
}

func (s *ChallengeService) destroyContainerLocked(record *model.Container, nextStatus model.ContainerStatus) error {
	if record.ContainerID != "" {
		cli, err := s.getDockerClient()
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		timeout := 5
		if err := cli.ContainerStop(ctx, record.ContainerID, dockercontainer.StopOptions{Timeout: &timeout}); err != nil && !errdefs.IsNotFound(err) {
			return err
		}
		if err := cli.ContainerRemove(ctx, record.ContainerID, dockercontainer.RemoveOptions{Force: true}); err != nil && !errdefs.IsNotFound(err) {
			return err
		}
	}

	updates := map[string]any{
		"status":       string(nextStatus),
		"container_id": "",
		"last_error":   "",
	}
	if nextStatus == model.ContainerStatusExpired {
		updates["last_error"] = "container expired"
	} else if nextStatus == model.ContainerStatusFailed && strings.TrimSpace(record.LastError) != "" {
		updates["last_error"] = record.LastError
	}

	return s.db.DB.Model(record).Updates(updates).Error
}

func (s *ChallengeService) markContainerFailed(record *model.Container, err error) {
	_ = s.db.DB.Model(record).Updates(map[string]any{
		"status":     string(model.ContainerStatusFailed),
		"last_error": err.Error(),
	}).Error
}

func (s *ChallengeService) getDockerClient() (*client.Client, error) {
	s.dockerOnce.Do(func() {
		var opts []client.Opt
		opts = append(opts, client.WithAPIVersionNegotiation())
		if s.config.DockerHost != "" {
			opts = append(opts, client.WithHost(s.config.DockerHost))
		} else {
			opts = append(opts, client.FromEnv)
		}
		s.dockerClient, s.dockerErr = client.NewClientWithOpts(opts...)
	})

	return s.dockerClient, s.dockerErr
}

func (s *ChallengeService) accessURL(record *model.Container) string {
	host := record.AccessHost
	if strings.TrimSpace(host) == "" {
		host = s.config.PublicHost
	}
	return fmt.Sprintf("%s://%s:%d", s.config.PublicScheme, host, record.HostPort)
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return defaultValue
}

func getEnvIntOrDefault(key string, defaultValue int) int {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return defaultValue
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return parsed
}

func adminTimeLeft(status string, record *model.Container) string {
	switch status {
	case string(model.ContainerStatusRunning), string(model.ContainerStatusCreating):
		return record.TimeLeft()
	case string(model.ContainerStatusExpired):
		return "00:00:00"
	default:
		return "-"
	}
}

func adminAccessHost(value string, fallback string) string {
	if strings.TrimSpace(value) != "" {
		return value
	}
	return fallback
}

func adminStatusAfterDestroy(currentStatus string) model.ContainerStatus {
	switch currentStatus {
	case string(model.ContainerStatusExpired):
		return model.ContainerStatusExpired
	case string(model.ContainerStatusFailed):
		return model.ContainerStatusFailed
	default:
		return model.ContainerStatusStopped
	}
}
