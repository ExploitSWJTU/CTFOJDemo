import type { Challenge } from '../types/challenge';

export const challenges: Challenge[] = [
  {
    id: 1,
    title: 'Easy SQL Injection',
    category: 'Web',
    description: `
### Description

This is a basic SQL injection challenge. Your goal is to bypass the login page.

\`\`\`sql
SELECT * FROM users WHERE username = '$username' AND password = '$password'
\`\`\`

### Goal

Find the flag in the database.
    `,
    points: 100,
    solvedCount: 120,
    difficulty: 'Easy',
    status: 'solved',
    containerState: 'running',
    containerInfo: {
      ip: '192.168.1.100',
      port: 8080,
      timeLeft: '00:58:20',
    },
  },
  {
    id: 2,
    title: 'Buffer Overflow Level 1',
    category: 'Pwn',
    description: `
### Description

A simple buffer overflow vulnerability. Can you overwrite the return address?

\`\`\`c
void vulnerable_function(char *input) {
    char buffer[64];
    strcpy(buffer, input);
}
\`\`\`
    `,
    points: 200,
    solvedCount: 45,
    difficulty: 'Medium',
    status: 'unsolved',
    containerState: 'idle',
  },
  {
    id: 3,
    title: 'RSA 101',
    category: 'Crypto',
    description: 'Decrypt the message using the given N and e.',
    points: 150,
    solvedCount: 80,
    difficulty: 'Easy',
    status: 'unsolved',
    containerState: 'idle',
  },
  {
    id: 4,
    title: 'Hidden Image',
    category: 'Misc',
    description: 'Find the hidden flag in the image metadata.',
    points: 100,
    solvedCount: 200,
    difficulty: 'Easy',
    status: 'solved',
    containerState: 'idle',
  },
  {
    id: 5,
    title: 'Advanced XSS',
    category: 'Web',
    description: 'Bypass the CSP to execute your JavaScript.',
    points: 300,
    solvedCount: 15,
    difficulty: 'Hard',
    status: 'unsolved',
    containerState: 'loading',
  },
  {
    id: 6,
    title: 'Kernel Exploit',
    category: 'Pwn',
    description: 'Exploit a race condition in the kernel module.',
    points: 500,
    solvedCount: 5,
    difficulty: 'Hard',
    status: 'unsolved',
    containerState: 'idle',
  },
  {
    id: 7,
    title: 'Easy Reverse',
    category: 'Reverse',
    description: 'Find the password hidden in the binary.',
    points: 100,
    solvedCount: 150,
    difficulty: 'Easy',
    status: 'unsolved',
    containerState: 'idle',
  },
  {
    id: 8,
    title: 'Android Backup',
    category: 'Mobile',
    description: 'Extract data from the Android backup file.',
    points: 200,
    solvedCount: 60,
    difficulty: 'Medium',
    status: 'unsolved',
    containerState: 'idle',
  },
  {
    id: 9,
    title: 'Smart Contract Vulnerability',
    category: 'Blockchain',
    description: 'Exploit the reentrancy bug in the smart contract.',
    points: 400,
    solvedCount: 20,
    difficulty: 'Hard',
    status: 'unsolved',
    containerState: 'idle',
  },
  {
    id: 10,
    title: 'Adversarial Example',
    category: 'AI',
    description: 'Create an image that fools the neural network.',
    points: 350,
    solvedCount: 10,
    difficulty: 'Medium',
    status: 'unsolved',
    containerState: 'idle',
  },
];
