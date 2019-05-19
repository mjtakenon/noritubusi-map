<template>
  <v-alert
    :type="type"
    dismissible
    transition="slide-y-transition"
    v-model="isVisible"
  >{{ message }}</v-alert>
</template>

<script>
export default {
  // 引数
  props: {
    // アラートの種別(success, error など)
    type: {
      type: String,
      required: true,
    },
    // [Vuex] Sidebar ストアのモジュール名
    moduleName: {
      type: String,
      required: true,
    },
  },
  // データ
  data() {
    return {}
  },
  // 算出プロパティ
  computed: {
    // セッター, ゲッター両方書く場合は、Object として定義
    isVisible: {
      get() {
        return this.$store.getters[`Sidebar/${this.moduleName}/isVisible`]
      },
      set(value) {
        this.$store.dispatch(`Sidebar/${this.moduleName}/isVisible`, value)
      },
    },
    // ゲッターのみでいい場合は、関数として定義
    message() {
      return this.$store.getters[`Sidebar/${this.moduleName}/message`]
    },
  },
}
</script>
