package model

// APIResponse represents the standard API response format
type APIResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// SuccessResponse creates a success response
func SuccessResponse(data interface{}) APIResponse {
	return APIResponse{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
}

// ErrorResponse creates an error response
func ErrorResponse(code int, msg string) APIResponse {
	return APIResponse{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

// Pagination represents pagination metadata
type Pagination struct {
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
}

// PaginatedResponse creates a paginated response
func PaginatedResponse(data interface{}, pagination Pagination) APIResponse {
	return APIResponse{
		Code: 200,
		Msg:  "success",
		Data: map[string]interface{}{
			"items":      data,
			"pagination": pagination,
		},
	}
}
