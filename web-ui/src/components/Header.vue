<template>
  <n-layout-header class="app-header" bordered>
    <n-space justify="space-between">
      <n-popover trigger="click" placement="right">
        <template #trigger>
          <div class="logo">
            <div>{{ t('title') }}</div>
          </div>
        </template>
        <span>Banana</span>
      </n-popover>
      <Locale class="header-locale" />
    </n-space>
  </n-layout-header>
</template>

<script lang="ts" setup>
import Locale from './Locale.vue'
import { NSpace, NLayoutHeader, NPopover } from 'naive-ui'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
</script>

<style lang="scss" scoped>
.app-header {
  width: 100%;
  height: 82px;
  padding: 10px;
  .header-locale {
    padding-top: 13px;
  }
  .logo {
    width: 100%;
    user-select: none;
    div {
      height: 60px;
      line-height: 60px;
      font-size: 32px;
      cursor: pointer;
      color: var(--text-color);
      text-align: center;
      transition: color 0.5s;
      transform: translate(0%);
    }

    div::before {
      content: '';
      position: absolute;
      left: 0%;
      bottom: 0;
      width: 100%;
      transform: scaleX(0);
      height: 2px;
      background: #ff6266;
      z-index: -1;
      transition: transform 0.5s;
      transform-origin: 100% 0;
    }

    div:hover::before {
      transform: scaleX(1);
      transform-origin: 0 0;
    }
  }
}
</style>
