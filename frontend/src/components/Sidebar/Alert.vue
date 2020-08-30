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
        return this.$store.getters[`Sidebar/Alert/isVisible`]
      },
      set(value) {
        this.$store.dispatch(`Sidebar/Alert/setVisible`, value)
      },
    },
    // ゲッターのみでいい場合は、関数として定義
    message() {
      return this.$store.getters[`Sidebar/Alert/message`]
    },
    type() {
      return this.$store.getters[`Sidebar/Alert/type`]
    },
  },
}
</script>
