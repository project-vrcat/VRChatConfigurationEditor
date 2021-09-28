<template>
  <n-form :label-width="80">
    <n-form-item>
      <template #label>
        {{ t("cache-location.label") }}
        <HelpIcon :tooltip="t('cache-location.tooltip')" />
      </template>
      <n-input-group>
        <n-input
          :placeholder="t('cache-location.placeholder')"
          v-model:value="formValue.cache_directory"
          readonly
        />
        <n-button type="primary">{{ t("cache-location.button") }}</n-button>
      </n-input-group>
    </n-form-item>

    <n-grid cols="2 xs:1" :x-gap="10" responsive="screen">
      <n-grid-item>
        <n-form-item :label="t('cache-location.cache_size.label')">
          <n-input-number v-model:value="formValue.cache_size" :min="20">
            <template #suffix>
              {{ t("cache-location.cache_size.suffix") }}
            </template>
          </n-input-number>
        </n-form-item>
      </n-grid-item>
      <n-grid-item>
        <n-form-item :label="t('cache-location.cache_expiry_delay.label')">
          <n-input-number
            v-model:value="formValue.cache_expiry_delay"
            :min="30"
          >
            <template #suffix>
              {{ t("cache-location.cache_expiry_delay.suffix") }}
            </template>
          </n-input-number>
        </n-form-item>
      </n-grid-item>
    </n-grid>

    <n-form-item class="n-form-item-hide-content">
      <template #label>
        {{ t("rich-presence.label") }}
        <HelpIcon :tooltip="t('rich-presence.tooltip')" />
        <n-switch
          v-model:value="richPresence"
          size="small"
          style="--height: 17px; margin-left: 5px"
        />
      </template>
    </n-form-item>

    <n-grid cols="2 xs:1" :x-gap="10" responsive="screen">
      <n-grid-item>
        <n-form-item :label="t('camera_and_screenshot.camera.label')">
          <n-select
            v-model:value="cameraValue"
            :options="resolutionOptions"
            :render-label="resolutionRenderLabel"
          />
        </n-form-item>
      </n-grid-item>
      <n-grid-item>
        <n-form-item :label="t('camera_and_screenshot.screenshot.label')">
          <n-select
            v-model:value="screenshotValue"
            :options="resolutionOptions"
            :render-label="resolutionRenderLabel"
          />
        </n-form-item>
      </n-grid-item>
    </n-grid>

    <n-form-item :label="t('first_person_steadycam_fov.label')">
      <n-slider
        v-model:value="formValue.fpv_steadycam_fov"
        :step="1"
        :min="30"
        :max="110"
      />
    </n-form-item>

    <n-form-item :label="t('avatar_dynamic_bone_limits.label')">
      <span>{{ t("switch.disable") }}</span>
      <n-switch
        v-model:value="avatarDynamicBoneLimits"
        size="small"
        style="--height: 17px; margin-left: 5px"
      />
      <span style="padding-left: 5px">{{ t("switch.enable") }}</span>
    </n-form-item>

    <n-grid
      cols="2 xs:1"
      :x-gap="10"
      responsive="screen"
      v-show="avatarDynamicBoneLimits"
    >
      <n-grid-item>
        <n-form-item
          :label="t('avatar_dynamic_bone_limits.max_affected_transforms')"
        >
          <n-input-number
            v-model:value="formValue.dynamic_bone_max_affected_transform_count"
          />
        </n-form-item>
      </n-grid-item>
      <n-grid-item>
        <n-form-item
          :label="t('avatar_dynamic_bone_limits.max_collision_checks')"
        >
          <n-input-number
            v-model:value="formValue.dynamic_bone_max_collider_check_count"
          />
        </n-form-item>
      </n-grid-item>
    </n-grid>
  </n-form>
</template>

<script lang="ts" setup>
import { ref, watch, h } from "vue";
import {
  NButton,
  NInput,
  NInputGroup,
  NForm,
  NFormItem,
  NGrid,
  NGridItem,
  NInputNumber,
  NSwitch,
  NSelect,
  NTag,
  NSlider,
} from "naive-ui";
import HelpIcon from "../components/HelpIcon.vue";

import { useI18n } from "vue-i18n";
const { t } = useI18n();

interface resolution {
  width: number;
  height: number;
}

const formValue = ref({
  betas: [],
  disableRichPresence: false,
  cache_directory: "",
  cache_size: 20,
  cache_expiry_delay: 30,
  camera_res_height: 1080,
  camera_res_width: 1920,
  screenshot_res_height: 1080,
  screenshot_res_width: 1920,
  fpv_steadycam_fov: 55,
  dynamic_bone_max_affected_transform_count: 32,
  dynamic_bone_max_collider_check_count: 8,
});
const richPresence = ref(true);
watch(richPresence, (v) => {
  formValue.value.disableRichPresence = !v;
});

const cameraValue = ref("1080p");
const screenshotValue = ref("1080p");
const customCameraResolutionValue = ref({ width: 1920, height: 1080 });
const customScreenshotResolutionValue = ref({ width: 1920, height: 1080 });
const resolutionRenderLabel = ref((option: any, selected: boolean) => {
  if (option.default) {
    return [
      option.label,
      h(
        NTag,
        {
          style: {
            marginLeft: "8px",
          },
          round: true,
          size: "small",
        },
        { default: () => t("camera_and_screenshot.default") }
      ),
    ];
  }
  return option.label;
});
const resolutionOptions = ref([
  {
    label: "720p (1280x720)",
    value: "720p",
  },
  {
    label: "1080p (1920x1080)",
    value: "1080p",
    default: true,
  },
  {
    label: "2K (2560x1440)",
    value: "2k",
  },
  {
    label: "4K (3840x2160)",
    value: "4k",
  },
]);
watch(cameraValue, (v) => {
  let r: resolution;
  if (v !== "custom") {
    r = resolutionValue(v);
    customCameraResolutionValue.value = r;
  } else {
    r = customCameraResolutionValue.value;
  }
  formValue.value.camera_res_height = r.height;
  formValue.value.camera_res_width = r.width;
});
watch(screenshotValue, (v) => {
  let r: resolution;
  if (v !== "custom") {
    r = resolutionValue(v);
    customScreenshotResolutionValue.value = r;
  } else {
    r = customScreenshotResolutionValue.value;
  }
  formValue.value.screenshot_res_height = r.height;
  formValue.value.screenshot_res_width = r.width;
});
const resolutionValue = (name: string): resolution => {
  switch (name) {
    case "720p":
      return { width: 1280, height: 720 };
    case "2k":
      return { width: 2560, height: 1440 };
    case "4k":
      return { width: 3840, height: 2160 };
    default:
      return { width: 1920, height: 1080 };
  }
};

const avatarDynamicBoneLimits = ref(true);
</script>

<style lang="scss" scoped>
.n-input-number {
  width: 100%;
}
</style>

<style lang="scss">
:root {
  color-scheme: dark light;
}

.n-form-item-hide-content {
  .n-form-item-blank {
    display: none !important;
  }
  .n-form-item-feedback-wrapper {
    min-height: 16px !important;
  }
}
</style>
