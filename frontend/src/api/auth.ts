import request from './request';
import type { User } from '@/types/user';

export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  token: string;
  user: User;
}

export interface RegisterRequest {
  username: string;
  password: string;
  email: string;
}

export interface RegisterResponse {
  token: string;
  user: User;
}

/**
 * User login
 */
export async function login(data: LoginRequest): Promise<LoginResponse> {
  return request.post('/auth/login', data);
}

/**
 * User registration
 */
export async function register(data: RegisterRequest): Promise<RegisterResponse> {
  return request.post('/auth/register', data);
}

/**
 * Get current user info
 */
export async function getCurrentUser(): Promise<User> {
  return request.get('/auth/me');
}
