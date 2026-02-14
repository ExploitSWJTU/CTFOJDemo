import { reactive } from 'vue'
import type { Team } from '../types/team'
import { teams as mockTeams } from '../mock/teams'

// 共享的队伍数据存储
export const teamStore = reactive<{
  teams: Team[]
}>({
  teams: [...mockTeams],
})

// 生成邀请码（12位，包含数字和大写字母）
function generateInviteCode(): string {
  const chars = 'ABCDEFGHJKLMNPQRSTUVWXYZ23456789' // 排除容易混淆的字符：0, O, I, 1
  let code = ''
  let attempts = 0
  const maxAttempts = 100 // 防止无限循环
  let found = false

  while (!found && attempts < maxAttempts) {
    // 生成新的邀请码
    let rawCode = ''
    for (let i = 0; i < 12; i++) {
      rawCode += chars.charAt(Math.floor(Math.random() * chars.length))
    }
    // 格式化为 XXXX-XXXX-XXXX
    code = `${rawCode.substring(0, 4)}-${rawCode.substring(4, 8)}-${rawCode.substring(8, 12)}`
    
    attempts++
    // 检查是否已存在
    const exists = teamStore.teams.some((team) => team.inviteCode === code)
    if (!exists) {
      found = true
    }
  }
  
  // 如果尝试次数过多仍未找到唯一码，添加随机后缀
  if (!found && attempts >= maxAttempts) {
    const suffix = Math.floor(Math.random() * 1000).toString().padStart(3, '0')
    code = `${code}-${suffix}`
  }

  return code
}

// 创建队伍
export function createTeam(data: {
  name: string
  avatar?: string
  description: string
  creatorId: number
  creatorUsername: string
  creatorAvatar?: string
}): Team | null {
  const maxId = teamStore.teams.length > 0
    ? Math.max(...teamStore.teams.map((t) => t.id))
    : 0
  const newId = maxId + 1

  const newTeam: Team = {
    id: newId,
    name: data.name,
    avatar: data.avatar || `https://api.dicebear.com/7.x/shapes/svg?seed=${data.name}`,
    description: data.description,
    members: [
      {
        id: data.creatorId,
        username: data.creatorUsername,
        avatar: data.creatorAvatar,
      },
    ],
    inviteCode: generateInviteCode(),
    creatorId: data.creatorId,
  }

  teamStore.teams.push(newTeam)
  return newTeam
}

// 通过邀请码加入队伍
export function joinTeamByInviteCode(inviteCode: string, userId: number, username: string, avatar?: string): boolean {
  const team = teamStore.teams.find((t) => t.inviteCode === inviteCode)
  if (!team) {
    return false
  }

  // 检查是否已经是成员
  if (team.members.some((m) => m.id === userId)) {
    return false
  }

  // 添加成员
  team.members.push({
    id: userId,
    username,
    avatar,
  })

  return true
}

// 更新队伍
export function updateTeam(
  id: number,
  data: {
    name?: string
    avatar?: string
    description?: string
    members?: Team['members']
  }
): boolean {
  const team = teamStore.teams.find((t) => t.id === id)
  if (team) {
    if (data.name !== undefined) team.name = data.name
    if (data.avatar !== undefined) team.avatar = data.avatar
    if (data.description !== undefined) team.description = data.description
    if (data.members !== undefined) team.members = data.members
    return true
  }
  return false
}

// 删除队伍
export function deleteTeam(id: number): boolean {
  const index = teamStore.teams.findIndex((t) => t.id === id)
  if (index !== -1) {
    teamStore.teams.splice(index, 1)
    return true
  }
  return false
}

// 获取队伍
export function getTeam(id: number): Team | undefined {
  return teamStore.teams.find((t) => t.id === id)
}

// 获取用户所在的队伍
export function getUserTeams(userId: number): Team[] {
  return teamStore.teams.filter((t) => t.members.some((m) => m.id === userId))
}

// 离开队伍
export function leaveTeam(teamId: number, userId: number): boolean {
  const team = teamStore.teams.find((t) => t.id === teamId)
  if (!team) {
    return false
  }

  // 如果是创建者，解散队伍（删除整个队伍）
  if (team.creatorId === userId) {
    return deleteTeam(teamId)
  }

  // 普通成员离开队伍
  const index = team.members.findIndex((m) => m.id === userId)
  if (index !== -1) {
    team.members.splice(index, 1)
    return true
  }
  return false
}

