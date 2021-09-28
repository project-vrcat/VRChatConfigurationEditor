<template>
  <div class="locale">
    <n-icon size="24">
      <language />
    </n-icon>
    <n-select v-model:value="locale" :options="languages" />
  </div>
</template>

<script lang="ts" setup>
import { watch } from "vue";
import { useI18n } from "vue-i18n";
import { NSelect, NIcon } from "naive-ui";
import { Language } from "@vicons/tabler";

const languages = [
  { label: "English", value: "en" },
  { label: "简体中文", value: "zh-CN" },
];

const { locale, availableLocales } = useI18n();

let localLanguage = localStorage.getItem("language");
if (localLanguage) locale.value = localLanguage;
else {
  const navigatorLocale = navigator.language;
  const trimmedLocale = navigatorLocale.trim().split(/-|_/)[0];

  const localLocale = availableLocales.find(
    (v) => v === navigatorLocale || v === trimmedLocale
  );

  if (localLocale) locale.value = localLocale;
}
watch(locale, (value, oldValue) => {
  localStorage.setItem("language", value);
});
</script>

<style lang="scss" scoped>
.locale {
  .n-icon {
    height: 33px;
    line-height: 40px;
  }
  .n-select {
    display: inline-block;
    width: 200px;
    padding-left: 4px;
  }
}
</style>
