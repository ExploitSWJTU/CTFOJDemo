import { reactive } from 'vue'

export type ContestStatus = 'ongoing' | 'upcoming' | 'finished'
export type ContestType = 'individual' | 'team'

export interface Contest {
  id: number
  name: string
  brief: string // 简介（简短描述）
  description: string
  startTime: string
  endTime: string
  status: ContestStatus
  imageUrl: string
  countdownText?: string
  participantCount: number
  type: ContestType
}

// 共享的比赛数据存储
export const contestStore = reactive<{
  contests: Contest[]
}>({
  contests: [
    {
      id: 1,
      name: '第八届西南交通大学 CTF 新秀杯',
      brief: '面向零基础新生的入门赛，涵盖基础题目。',
      description:
        '# 赛事背景\n\n面向零基础新生的入门赛，涵盖基础 Web / Crypto / Misc 等题目，帮助同学们快速熟悉 CTF 比赛形式。\n\n## 比赛形式\n\n- 赛制：个人赛\n- 题型：Web / Crypto / Misc\n- 评分：动态积分制\n',
      startTime: '2025-03-10 19:00',
      endTime: '2025-03-10 22:00',
      status: 'ongoing',
      imageUrl: 'https://dummyimage.com/640x360/0f172a/38bdf8&text=Newbee+CTF',
      participantCount: 156,
      type: 'individual',
    },
    {
      id: 2,
      name: '校内常规训练赛',
      brief: '每周一次的校内训练赛，用于巩固日常练习内容。',
      description:
        '# 赛事介绍\n\n每周一次的校内训练赛，用于巩固日常练习内容，题目难度适中，适合有一定基础的同学。\n\n## 适合人群\n\n- 已经完成基础新生赛\n- 希望保持刷题节奏的同学\n',
      startTime: '2025-03-15 19:00',
      endTime: '2025-03-15 23:00',
      status: 'upcoming',
      imageUrl: 'https://dummyimage.com/640x360/020617/a5b4fc&text=Weekly+Training',
      countdownText: '2天后开始',
      participantCount: 89,
      type: 'team',
    },
    {
      id: 3,
      name: 'SWJTU CTF 校赛',
      brief: '正式选拔赛，成绩将作为集训队选拔的重要参考。',
      description:
        '# 赛事介绍\n\n正式选拔赛，成绩将作为集训队选拔的重要参考。\n\n## 注意事项\n\n- 比赛全程禁止作弊与共享 Flag\n- 需要独立完成题目\n',
      startTime: '2025-02-01 09:00',
      endTime: '2025-02-01 17:00',
      status: 'finished',
      imageUrl: 'https://dummyimage.com/640x360/111827/f97316&text=SWJTU+CTF+Final',
      participantCount: 234,
      type: 'team',
    },
    {
      id: 4,
      name: '春季 CTF 挑战赛',
      brief: '春季学期大型 CTF 比赛，包含 Web、Crypto、Pwn 等多个方向。',
      description: '春季学期大型 CTF 比赛，包含 Web、Crypto、Pwn 等多个方向。',
      startTime: '2025-03-20 10:00',
      endTime: '2025-03-20 18:00',
      status: 'upcoming',
      imageUrl: 'https://dummyimage.com/640x360/1e293b/10b981&text=Spring+CTF',
      countdownText: '7天后开始',
      participantCount: 201,
      type: 'individual',
    },
    {
      id: 5,
      name: '网络安全周 CTF 竞赛',
      brief: '配合网络安全宣传周举办的 CTF 竞赛，提升网络安全意识。',
      description: '配合网络安全宣传周举办的 CTF 竞赛，提升网络安全意识。',
      startTime: '2025-04-15 09:00',
      endTime: '2025-04-15 21:00',
      status: 'upcoming',
      imageUrl: 'https://dummyimage.com/640x360/0f172a/f59e0b&text=Cyber+Security',
      countdownText: '33天后开始',
      participantCount: 178,
      type: 'team',
    },
    {
      id: 6,
      name: 'CTF 新人训练营',
      brief: '面向新手的训练营，提供详细讲解和指导。',
      description: '面向新手的训练营，提供详细讲解和指导。',
      startTime: '2025-01-15 14:00',
      endTime: '2025-01-15 17:00',
      status: 'finished',
      imageUrl: 'https://dummyimage.com/640x360/111827/8b5cf6&text=Training+Camp',
      participantCount: 95,
      type: 'individual',
    },
    {
      id: 7,
      name: 'CTF 月度挑战',
      brief: '每月一次的挑战赛，题目难度适中，适合日常练习。',
      description: '每月一次的挑战赛，题目难度适中，适合日常练习。',
      startTime: '2025-03-25 19:00',
      endTime: '2025-03-25 23:00',
      status: 'upcoming',
      imageUrl: 'https://dummyimage.com/640x360/020617/06b6d4&text=Monthly+Challenge',
      countdownText: '12天后开始',
      participantCount: 142,
      type: 'individual',
    },
    {
      id: 8,
      name: 'CTF 团队对抗赛',
      brief: '团队形式的对抗赛，考验团队协作和综合能力。',
      description: '团队形式的对抗赛，考验团队协作和综合能力。',
      startTime: '2025-01-20 10:00',
      endTime: '2025-01-20 18:00',
      status: 'finished',
      imageUrl: 'https://dummyimage.com/640x360/1e293b/ef4444&text=Team+Battle',
      participantCount: 167,
      type: 'team',
    },
  ],
})

// 更新比赛描述
export function updateContestDescription(id: number, description: string) {
  const contest = contestStore.contests.find((c) => c.id === id)
  if (contest) {
    contest.description = description
    // TODO: 实际应该调用 API 保存到后端
    return true
  }
  return false
}

// 更新比赛信息
export function updateContest(
  id: number,
  data: {
    name?: string
    brief?: string
    description?: string
    startTime?: string
    endTime?: string
    type?: ContestType
  }
) {
  const contest = contestStore.contests.find((c) => c.id === id)
  if (contest) {
    if (data.name !== undefined) contest.name = data.name
    if (data.brief !== undefined) contest.brief = data.brief
    if (data.description !== undefined) contest.description = data.description
    if (data.startTime !== undefined) contest.startTime = data.startTime
    if (data.endTime !== undefined) contest.endTime = data.endTime
    if (data.type !== undefined) contest.type = data.type
    // TODO: 实际应该调用 API 保存到后端
    return true
  }
  return false
}

// 获取比赛
export function getContest(id: number): Contest | undefined {
  return contestStore.contests.find((c) => c.id === id)
}

// 创建比赛
export function createContest(data: {
  name: string
  brief: string
  description: string
  startTime: string
  endTime: string
  type: ContestType
  imageUrl: string
}): Contest | null {
  // 生成新 ID
  const maxId = contestStore.contests.length > 0
    ? Math.max(...contestStore.contests.map((c) => c.id))
    : 0
  const newId = maxId + 1

  // 根据时间计算状态
  const now = new Date()
  const startDate = new Date(data.startTime)
  const endDate = new Date(data.endTime)

  let status: ContestStatus
  if (now < startDate) {
    status = 'upcoming'
  } else if (now >= startDate && now <= endDate) {
    status = 'ongoing'
  } else {
    status = 'finished'
  }

  // 创建新比赛
  const newContest: Contest = {
    id: newId,
    name: data.name,
    brief: data.brief,
    description: data.description,
    startTime: data.startTime,
    endTime: data.endTime,
    status: status,
    imageUrl: data.imageUrl,
    participantCount: 0,
    type: data.type,
  }

  // 添加到列表
  contestStore.contests.push(newContest)

  // TODO: 实际应该调用 API 保存到后端
  return newContest
}

