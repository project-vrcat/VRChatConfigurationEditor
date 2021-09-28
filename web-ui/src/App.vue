<template>
  <n-config-provider :theme="theme">
    <n-message-provider>
      <n-layout style="height: 100vh;">
        <Header />
        <n-layout-content content-style="padding: 20px;" :native-scrollbar="false">
          <HomeView />
        </n-layout-content>
      </n-layout>
    </n-message-provider>
  </n-config-provider>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import Header from './components/Header.vue'
import HomeView from './views/Home.vue'
import { NConfigProvider, NMessageProvider, NLayout, NLayoutContent } from 'naive-ui'

import { useOsTheme, darkTheme } from 'naive-ui'
const osThemeRef = useOsTheme()
const theme = computed(() => osThemeRef.value === 'dark' ? darkTheme : null)
</script>