<template>
  <v-alert
    :type="type"
    dismissible
    transition="slide-y-transition"
    v-model="isVisible"
    >{{ message }}</v-alert
  >
</template>

<script>
export default {
  // データ
  data() {
    return {}
  },
  // 算出プロパティ
  computed: {
    // セッター, ゲッター両方書く場合は、Object として定義
    isVisible: {
      get() {
        return this.$store.getters[`Sidebar/AlertModal/isVisible`]
      },
      set(value) {
        this.$store.dispatch(`Sidebar/AlertModal/setVisible`, value)
      },
    },
    // ゲッターのみでいい場合は、関数として定義
    message() {
      return this.$store.getters[`Sidebar/AlertModal/message`]
    },
    type() {
      return this.$store.getters[`Sidebar/AlertModal/type`]
    },
  },
}
</script>
