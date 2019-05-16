<style scoped></style>
<template>
  <v-navigation-drawer
    aboslute
    temporary
    v-model="isVisible"
  ></v-navigation-drawer>
</template>

<script>
import UserInfo from "./UserInfo"

export default {
  components: {
    UserInfo,
  },
  props: {},
  data() {
    return {
      userInfo: {
        name: "真白たん",
      },
    }
  },
  computed: {
    isVisible: {
      get() {
        console.log("<Sidebar/> getter::isVisible")
        return this.$store.getters.isVisible
      },
      set(value) {
        console.log(`<Sidebar/> setter::isVisible(value=${value})`)
        this.$store.commit("isVisible", value)
      },
    },
  },
}
</script>
