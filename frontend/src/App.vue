<template>
  <v-app>
    <Sidebar/>
    <v-btn
      @click="onClickBtn"
      color="warning"
      right
      style="position: absolute;"
    >Test</v-btn>
  </v-app>
</template>

<script>
import Sidebar from "./components/Sidebar/Sidebar"

export default {
  components: {
    Sidebar,
  },
  data() {
    return {}
  },
  methods: {
    onClickBtn() {
      this.showSidebar = !this.showSidebar
    },
  },
  computed: {
    showSidebar: {
      get() {
        console.log("<App/> getter::showSidebar")
        return this.$store.getters["Sidebar/isVisible"]
      },
      set(value) {
        console.log(`<App/> setter::showSidebar(value=${value})`)
        this.$store.dispatch("Sidebar/isVisible", value)
      },
    },
  },
}
</script>
