<template>
  <n-tooltip trigger="hover" placement="right">
    <template #trigger>
      <n-icon style="margin-left: 6px" size="16">
        <help />
      </n-icon>
    </template>
    <span>{{ props.tooltip }}</span>
  </n-tooltip>
</template>

<script lang="ts" setup>
import { NIcon, NTooltip } from 'naive-ui'
import { Help } from '@vicons/tabler'
const props = defineProps({ tooltip: String })
</script>
