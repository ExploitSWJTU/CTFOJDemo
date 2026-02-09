import pluginVue from 'eslint-plugin-vue';
import tseslint from 'typescript-eslint';
import pluginJs from '@eslint/js';
import globals from 'globals';
// import tailwind from 'eslint-plugin-tailwindcss';

export default tseslint.config(
  pluginJs.configs.recommended,
  ...tseslint.configs.recommended,
  ...pluginVue.configs['flat/recommended'],
  // ...tailwind.configs['flat/recommended'],
  {
    files: ['*.vue', '**/*.vue', '**/*.ts', '**/*.js'],
    languageOptions: {
      globals: {
        ...globals.browser,
        ...globals.node,
      },
      // 针对 Vue 文件使用正确的解析器
      parserOptions: {
        parser: tseslint.parser,
        extraFileExtensions: ['.vue'],
        sourceType: 'module',
      },
    },
    // 如果你使用的是 Tailwind v4，请取消下面 settings 的注释
    // settings: {
    //   tailwindcss: {
    //     config: {},
    //   },
    // },
  },
  {
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
      // --- Tailwind 核心修复规则 ---
      // 'tailwindcss/classnames-order': 'warn', // 自动排序类名
      // 'tailwindcss/enforces-shorthand': 'error', // 强制使用简写 (例如 px-2 py-2 -> p-2)
      // 'tailwindcss/no-arbitrary-value': 'warn', // 重点：将 [1536px] 识别为警告，从而允许自动修复
    },
  },
  {
    ignores: ['dist/**', 'node_modules/**', 'tmp/**', 'eslint.config.js'],
  },
);
