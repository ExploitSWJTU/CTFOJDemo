/**
 * Lint-staged configuration for Go + Vue Monorepo
 * Runs linters only on staged files
 */

export default {
  // Frontend files
  'frontend/**/*.{ts,tsx,vue}': [
    () => 'pnpm --dir frontend lint'
  ],
  'frontend/**/*.{ts,tsx,vue,js,jsx}': [
    () => 'pnpm --dir frontend type-check'
  ],
  
  // Backend files
  'backend/**/*.go': [
    () => 'cd backend && go fmt ./...',
    () => 'cd backend && go vet ./...'
  ],
  
  // Config files
  '*.{json,yaml,yml,md}': [
    () => 'prettier --check'
  ]
};
