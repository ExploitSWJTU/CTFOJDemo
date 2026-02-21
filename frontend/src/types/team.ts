export interface TeamMember {
  id: number;
  username: string;
  avatar?: string;
}

export interface Team {
  id: number;
  name: string;
  avatar?: string;
  members: TeamMember[];
  description: string;
  inviteCode?: string; // 邀请码
  creatorId?: number; // 创建者ID
}

