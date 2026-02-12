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
}

