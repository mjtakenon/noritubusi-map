<template>
  <v-alert
    :type="type"
    dismissible
    transition="slide-y-transition"
    v-model="isVisible"
    >{{ message }}</v-alert
  >
</template>

<script lang="ts">
import Vue from "vue"

import { AlertType } from "@/entities/Alert"

export default Vue.extend({
  data() {
    return {}
  },
  computed: {
    isVisible: {
      get(): boolean {
        return this.$accessor.Sidebar.Alert.isVisible
      },
      set(value: boolean) {
        this.$accessor.Sidebar.Alert.setVisible(value)
      },
    },
    message(): string {
      return this.$accessor.Sidebar.Alert.message
    },
    type(): AlertType {
      return this.$accessor.Sidebar.Alert.type
    },
  },
})
</script>
