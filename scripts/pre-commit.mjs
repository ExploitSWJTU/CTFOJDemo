#!/usr/bin/env node

/**
 * Smart pre-commit hook for Go + Vue Monorepo
 * Only runs checks for changed files (frontend or backend)
 */

import { execSync } from 'child_process';
import { existsSync } from 'fs';

const FRONTEND_EXTENSIONS = ['.ts', '.tsx', '.vue', '.js', '.jsx', '.css'];
const BACKEND_EXTENSIONS = ['.go'];

// Get staged files
function getStagedFiles() {
  try {
    const output = execSync('git diff --cached --name-only --diff-filter=ACM', {
      encoding: 'utf-8'
    });
    return output.trim().split('\n').filter(f => f.length > 0);
  } catch (e) {
    console.error('Failed to get staged files:', e.message);
    return [];
  }
}

// Categorize files
function categorizeFiles(files) {
  const frontend = [];
  const backend = [];
  
  for (const file of files) {
    if (file.startsWith('frontend/')) {
      frontend.push(file);
    } else if (file.startsWith('backend/')) {
      backend.push(file);
    } else if (FRONTEND_EXTENSIONS.some(ext => file.endsWith(ext))) {
      frontend.push(file);
    } else if (BACKEND_EXTENSIONS.some(ext => file.endsWith(ext))) {
      backend.push(file);
    }
  }
  
  return { frontend, backend };
}

// Run checks
function runFrontendCheck() {
  console.log('🔍 Running frontend checks (lint + type-check)...');
  try {
    execSync('npx lint-staged', { stdio: 'inherit' });
    console.log('✅ Frontend checks passed');
    return true;
  } catch (e) {
    console.error('❌ Frontend checks failed');
    return false;
  }
}

function runBackendCheck() {
  console.log('🔍 Running backend checks (go fmt + go vet)...');
  try {
    // go fmt is run by lint-staged, just run go vet here
    execSync('cd backend && go vet ./...', { stdio: 'inherit' });
    console.log('✅ Backend checks passed');
    return true;
  } catch (e) {
    console.error('❌ Backend checks failed');
    return false;
  }
}

// Main
const stagedFiles = getStagedFiles();
const { frontend, backend } = categorizeFiles(stagedFiles);

console.log('📋 Staged files analysis:');
console.log(`   Frontend: ${frontend.length} files`);
console.log(`   Backend:  ${backend.length} files`);
console.log('');

if (frontend.length === 0 && backend.length === 0) {
  console.log('✨ No frontend/backend files to check');
  process.exit(0);
}

let success = true;

if (frontend.length > 0) {
  success = runFrontendCheck() && success;
}

if (backend.length > 0) {
  success = runBackendCheck() && success;
}

if (success) {
  console.log('\n✅ All checks passed!');
  process.exit(0);
} else {
  console.error('\n❌ Some checks failed. Please fix the issues above.');
  process.exit(1);
}
