<template>
  <router-view />
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap');

body {
  font-family: 'Inter', sans-serif;
  /* 移除 Arco 下拉框/气泡的外层默认边框 */
  --color-border-1: transparent;
  --color-bg-popup: transparent;
}

/* 彻底消除 Arco Dropdown 的默认包裹样式 */
.arco-dropdown,
.arco-dropdown-list,
.arco-dropdown-list-wrapper {
  padding: 0 !important;
  background-color: transparent !important;
  border: none !important;
  box-shadow: none !important;
}

.dark body {
  /* 基础背景 (Slate 950) */
  --color-bg-1: #020617; 
  /* 卡片背景 (Slate 900) */
  --color-bg-2: #0f172a; 
  /* 弹出层/模态框背景 (Slate 800) - 确保这是纯正的深蓝色 */
  --color-bg-3: #1e293b; 
  /* 悬浮/激活背景 (Slate 700) */
  --color-bg-4: #334155;
  /* 边框辅助背景 (Slate 600) */
  --color-bg-5: #475569;
  
  --color-bg-white: #020617;
  --color-text-1: #f1f5f9;
  --color-text-2: #e2e8f0;
  --color-text-3: #cbd5e1;
  --color-text-4: #94a3b8;
  
  /* 边框体系同步 */
  --color-border-1: transparent;
  --color-border-2: #1e293b; /* Slate 800 作为边框 */
  --color-border-3: #334155;
  --color-border-4: #475569;
  
  --color-bg-popup: #1e293b; /* 强制弹出气泡使用 Slate 800 */
}

/* 全局模态框样式深度定制 */
.dark .arco-modal {
  background-color: #1e293b !important; /* 锁定 Slate 800 */
  border: 1px solid #334155 !important; /* 锁定 Slate 700 边框 */
}

.dark .arco-modal-header {
  background-color: #1e293b !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.04) !important;
}

.dark .arco-modal-body {
  background-color: #1e293b !important;
}

.dark .arco-modal-footer {
  background-color: #1e293b !important;
  border-top: 1px solid rgba(255, 255, 255, 0.04) !important;
}
</style>