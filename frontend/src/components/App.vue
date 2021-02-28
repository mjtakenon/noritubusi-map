<template>
  <v-app>
    <Map />
    <!-- まずは、ユーザー管理機能があるサイドバーからリファクタリング -->
    <Sidebar />
    <FloatPanel />
    <!-- 位置情報を取得するためのボタン(仮) -->
    <v-btn
      @click="onClickTest"
      absolute
      color="info"
      right
      style="margin-top: 12px;"
      >Test</v-btn
    >
  </v-app>
</template>

<script lang="ts">
import Vue from "vue"

import Map from "@/components/Map/Map.vue"
import Sidebar from "@/components/Sidebar/Sidebar.vue"
import FloatPanel from "@/components/FloatPanel/FloatPanel.vue"

export default Vue.extend({
  components: {
    Sidebar,
    Map,
    FloatPanel,
  },
  methods: {
    onClickTest() {
      this.$accessor.Map.setCenter({
        lat: 34.9613557,
        lng: 137.1389376,
      })
    },
  },
  computed: {
    showSidebar(): boolean {
      return this.$accessor.Sidebar.isVisible
    },
  },
})
</script>
