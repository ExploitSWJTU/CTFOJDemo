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
  isActive?: boolean // 是否激活，未激活的比赛在用户端不显示
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
      isActive: true,
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
      isActive: true,
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
      isActive: true,
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
      isActive: true,
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
      isActive: true,
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
      isActive: true,
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
      isActive: true,
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
      isActive: true,
    },
    {
      id: 9,
      name: 'Web 安全专项赛',
      brief: '专注于 Web 安全漏洞挖掘和利用的专项比赛。',
      description: '专注于 Web 安全漏洞挖掘和利用的专项比赛，包含 SQL 注入、XSS、CSRF 等常见漏洞类型。',
      startTime: '2025-04-01 09:00',
      endTime: '2025-04-01 18:00',
      status: 'upcoming',
      imageUrl: 'https://dummyimage.com/640x360/0f172a/3b82f6&text=Web+Security',
      countdownText: '19天后开始',
      participantCount: 203,
      type: 'individual',
      isActive: true,
    },
    {
      id: 10,
      name: '密码学挑战赛',
      brief: '深入密码学算法和协议的挑战赛。',
      description: '深入密码学算法和协议的挑战赛，涵盖对称加密、非对称加密、哈希函数等。',
      startTime: '2025-01-10 10:00',
      endTime: '2025-01-10 16:00',
      status: 'finished',
      imageUrl: 'https://dummyimage.com/640x360/111827/10b981&text=Crypto+Challenge',
      participantCount: 128,
      type: 'individual',
      isActive: true,
    },
    {
      id: 11,
      name: 'Pwn 入门训练赛',
      brief: '面向初学者的 Pwn 题目训练赛。',
      description: '面向初学者的 Pwn 题目训练赛，帮助新手掌握二进制漏洞利用基础。',
      startTime: '2025-03-28 19:00',
      endTime: '2025-03-28 23:00',
      status: 'upcoming',
      imageUrl: 'https://dummyimage.com/640x360/020617/f59e0b&text=Pwn+Training',
      countdownText: '15天后开始',
      participantCount: 95,
      type: 'individual',
      isActive: true,
    },
    {
      id: 12,
      name: 'Misc 综合挑战',
      brief: '涵盖多种 Misc 题型的综合挑战赛。',
      description: '涵盖多种 Misc 题型的综合挑战赛，包括隐写、取证、编码等。',
      startTime: '2025-02-15 14:00',
      endTime: '2025-02-15 20:00',
      status: 'finished',
      imageUrl: 'https://dummyimage.com/640x360/1e293b/8b5cf6&text=Misc+Challenge',
      participantCount: 176,
      type: 'team',
      isActive: true,
    },
    {
      id: 13,
      name: '逆向工程竞赛',
      brief: '专注于逆向分析和代码审计的竞赛。',
      description: '专注于逆向分析和代码审计的竞赛，考验选手的分析能力和耐心。',
      startTime: '2025-04-10 10:00',
      endTime: '2025-04-10 18:00',
      status: 'upcoming',
      imageUrl: 'https://dummyimage.com/640x360/0f172a/ef4444&text=Reverse+Eng',
      countdownText: '28天后开始',
      participantCount: 145,
      type: 'individual',
      isActive: true,
    },
    {
      id: 14,
      name: 'CTF 春季联赛',
      brief: '春季学期大型联赛，多轮次积分制。',
      description: '春季学期大型联赛，采用多轮次积分制，最终根据总积分排名。',
      startTime: '2025-03-01 09:00',
      endTime: '2025-03-01 21:00',
      status: 'finished',
      imageUrl: 'https://dummyimage.com/640x360/111827/06b6d4&text=Spring+League',
      participantCount: 312,
      type: 'team',
      isActive: true,
    },
    {
      id: 15,
      name: '网络安全知识竞赛',
      brief: '理论知识与实践结合的知识竞赛。',
      description: '理论知识与实践结合的知识竞赛，包含网络安全法律法规、安全标准等。',
      startTime: '2025-04-20 14:00',
      endTime: '2025-04-20 18:00',
      status: 'upcoming',
      imageUrl: 'https://dummyimage.com/640x360/020617/a5b4fc&text=Security+Quiz',
      countdownText: '38天后开始',
      participantCount: 267,
      type: 'individual',
      isActive: true,
    },
    {
      id: 16,
      name: 'CTF 周末挑战',
      brief: '周末举办的快速挑战赛，题目难度适中。',
      description: '周末举办的快速挑战赛，题目难度适中，适合日常练习和提升。',
      startTime: '2025-03-30 10:00',
      endTime: '2025-03-30 16:00',
      status: 'upcoming',
      imageUrl: 'https://dummyimage.com/640x360/1e293b/f97316&text=Weekend+CTF',
      countdownText: '17天后开始',
      participantCount: 189,
      type: 'individual',
      isActive: true,
    },
    {
      id: 17,
      name: '漏洞挖掘实战赛',
      brief: '真实环境下的漏洞挖掘实战比赛。',
      description: '真实环境下的漏洞挖掘实战比赛，模拟真实的安全测试场景。',
      startTime: '2025-02-05 09:00',
      endTime: '2025-02-05 17:00',
      status: 'finished',
      imageUrl: 'https://dummyimage.com/640x360/0f172a/10b981&text=Bug+Hunting',
      participantCount: 198,
      type: 'team',
      isActive: true,
    },
    {
      id: 18,
      name: 'CTF 新人赛',
      brief: '专门为新人设计的入门级比赛。',
      description: '专门为新人设计的入门级比赛，题目简单易懂，帮助新手快速入门。',
      startTime: '2025-04-05 19:00',
      endTime: '2025-04-05 22:00',
      status: 'upcoming',
      imageUrl: 'https://dummyimage.com/640x360/111827/3b82f6&text=Newbie+CTF',
      countdownText: '23天后开始',
      participantCount: 234,
      type: 'individual',
      isActive: true,
    },
    {
      id: 19,
      name: '高级渗透测试赛',
      brief: '面向高级选手的渗透测试挑战赛。',
      description: '面向高级选手的渗透测试挑战赛，题目难度较高，考验综合能力。',
      startTime: '2025-01-25 10:00',
      endTime: '2025-01-25 18:00',
      status: 'finished',
      imageUrl: 'https://dummyimage.com/640x360/020617/ef4444&text=Advanced+PenTest',
      participantCount: 87,
      type: 'team',
      isActive: true,
    },
    {
      id: 20,
      name: 'CTF 月度排位赛',
      brief: '每月一次的排位赛，根据排名获得积分。',
      description: '每月一次的排位赛，根据排名获得积分，累计积分可参与年度总决赛。',
      startTime: '2025-04-25 19:00',
      endTime: '2025-04-25 23:00',
      status: 'upcoming',
      imageUrl: 'https://dummyimage.com/640x360/1e293b/8b5cf6&text=Monthly+Rank',
      countdownText: '43天后开始',
      participantCount: 298,
      type: 'team',
      isActive: true,
    },
    {
      id: 21,
      name: '区块链安全挑战',
      brief: '专注于区块链和智能合约安全的挑战赛。',
      description: '专注于区块链和智能合约安全的挑战赛，包含智能合约审计、区块链协议分析等。',
      startTime: '2025-02-10 14:00',
      endTime: '2025-02-10 20:00',
      status: 'finished',
      imageUrl: 'https://dummyimage.com/640x360/0f172a/f59e0b&text=Blockchain+Sec',
      participantCount: 156,
      type: 'individual',
      isActive: true,
    },
    {
      id: 22,
      name: 'IoT 安全研究赛',
      brief: '物联网设备安全研究和漏洞挖掘比赛。',
      description: '物联网设备安全研究和漏洞挖掘比赛，包含嵌入式系统、IoT 协议等。',
      startTime: '2025-04-15 10:00',
      endTime: '2025-04-15 18:00',
      status: 'upcoming',
      imageUrl: 'https://dummyimage.com/640x360/111827/06b6d4&text=IoT+Security',
      countdownText: '33天后开始',
      participantCount: 112,
      type: 'team',
      isActive: true,
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
    imageUrl?: string
    isActive?: boolean
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
    if (data.imageUrl !== undefined) contest.imageUrl = data.imageUrl
    if (data.isActive !== undefined) contest.isActive = data.isActive
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
      isActive: true, // 新创建的比赛默认为未激活
    }

  // 添加到列表
  contestStore.contests.push(newContest)

  // TODO: 实际应该调用 API 保存到后端
  return newContest
}

// 删除比赛
export function deleteContest(id: number): boolean {
  const index = contestStore.contests.findIndex((c) => c.id === id)
  if (index !== -1) {
    contestStore.contests.splice(index, 1)
    // TODO: 实际应该调用 API 删除后端数据
    return true
  }
  return false
}

