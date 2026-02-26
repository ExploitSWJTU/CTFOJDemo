import { reactive } from 'vue'
import type { Category } from '@/types/challenge'

export type ChallengeType = 'static_attachment' | 'static_container' | 'dynamic_attachment' | 'dynamic_container'

export interface ContestChallenge {
  id: number
  title: string
  category: Category
  challengeType: ChallengeType
  description: string
  points: number
  solvedCount: number
  difficulty: string
  enabled: boolean
  firstBloodReward: string
}

const defaultChallenges = (): ContestChallenge[] => [
  {
    id: 1,
    title: 'Easy SQL Injection',
    category: 'Web',
    challengeType: 'dynamic_container',
    description: 'This is a basic SQL injection challenge. Your goal is to bypass the login page.',
    points: 100,
    solvedCount: 0,
    difficulty: 'Easy',
    enabled: true,
    firstBloodReward: '+10 积分',
  },
  {
    id: 2,
    title: 'Buffer Overflow Level 1',
    category: 'Pwn',
    challengeType: 'static_container',
    description: 'A simple buffer overflow vulnerability. Can you overwrite the return address?',
    points: 200,
    solvedCount: 0,
    difficulty: 'Medium',
    enabled: true,
    firstBloodReward: '',
  },
  {
    id: 3,
    title: 'Misc 隐写入门',
    category: 'Misc',
    challengeType: 'static_attachment',
    description: '在附件中找到隐藏的 flag。',
    points: 100,
    solvedCount: 0,
    difficulty: 'Easy',
    enabled: true,
    firstBloodReward: '',
  },
  {
    id: 4,
    title: '动态附件示例',
    category: 'Web',
    challengeType: 'dynamic_attachment',
    description: '每个队伍获得一份独立附件与对应 Flag。',
    points: 150,
    solvedCount: 0,
    difficulty: 'Medium',
    enabled: true,
    firstBloodReward: '',
  },
]

const state = reactive<{ byContestId: Record<string, ContestChallenge[]> }>({
  byContestId: {},
})

function getList(contestId: string): ContestChallenge[] {
  if (!state.byContestId[contestId]) {
    state.byContestId[contestId] = defaultChallenges()
  }
  return state.byContestId[contestId]
}

export function getContestChallenges(contestId: string): ContestChallenge[] {
  return getList(contestId)
}

export function removeContestChallenge(contestId: string, challengeId: number): boolean {
  const list = getList(contestId)
  const idx = list.findIndex((c) => c.id === challengeId)
  if (idx === -1) return false
  list.splice(idx, 1)
  return true
}

export function addContestChallenge(contestId: string, challenge: ContestChallenge): void {
  getList(contestId).push(challenge)
}

export function getNextChallengeId(contestId: string): number {
  const list = getList(contestId)
  return Math.max(0, ...list.map((c) => c.id)) + 1
}

export function getContestChallenge(contestId: string, challengeId: number): ContestChallenge | undefined {
  return getList(contestId).find((c) => c.id === challengeId)
}

export function updateContestChallenge(
  contestId: string,
  challengeId: number,
  data: Partial<Omit<ContestChallenge, 'id'>>
): boolean {
  const list = getList(contestId)
  const challenge = list.find((c) => c.id === challengeId)
  if (!challenge) return false
  if (data.title !== undefined) challenge.title = data.title
  if (data.category !== undefined) challenge.category = data.category
  if (data.challengeType !== undefined) challenge.challengeType = data.challengeType
  if (data.description !== undefined) challenge.description = data.description
  if (data.points !== undefined) challenge.points = data.points
  if (data.solvedCount !== undefined) challenge.solvedCount = data.solvedCount
  if (data.difficulty !== undefined) challenge.difficulty = data.difficulty
  if (data.enabled !== undefined) challenge.enabled = data.enabled
  if (data.firstBloodReward !== undefined) challenge.firstBloodReward = data.firstBloodReward
  return true
}
