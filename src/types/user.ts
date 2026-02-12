export type UserRole = 'admin' | 'user';

export interface User {
  id: number;
  username: string;
  email: string;
  realName: string;
  studentId: string;
  role: UserRole;
  avatar?: string;
}

