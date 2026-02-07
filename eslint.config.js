import pluginVue from 'eslint-plugin-vue';
import tseslint from 'typescript-eslint';
import pluginJs from '@eslint/js';
import globals from 'globals';

export default tseslint.config(
  // 基础 JS 推荐配置
  pluginJs.configs.recommended,
  // TypeScript 推荐配置
  ...tseslint.configs.recommended,
  // Vue 推荐配置
  ...pluginVue.configs['flat/recommended'],
  {
    files: ['*.vue', '**/*.vue', '**/*.ts', '**/*.js'],
    languageOptions: {
      globals: {
        ...globals.browser,
        ...globals.node,
      },
      parserOptions: {
        parser: tseslint.parser,
      },
    },
  },
  {
    // 自定义规则
    rules: {
      'vue/multi-word-component-names': 'off',
      '@typescript-eslint/no-explicit-any': 'warn',
      '@typescript-eslint/no-unused-vars': 'warn',
      'no-useless-assignment': 'off',
      'vue/max-attributes-per-line': [
        'error',
        {
          singleline: { max: 3 },
          multiline: { max: 1 },
        },
      ],
    },
  },
  {
    // 忽略文件
    ignores: ['dist/**', 'node_modules/**', 'tmp/**', 'eslint.config.js'],
  },
);
