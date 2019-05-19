<style scoped>
</style>

<template>
  <v-list>
    <v-list-tile>
      <template id="default" v-if="!isLoggedIn">
        <v-list-tile-avatar>
          <v-icon x-large>portrait</v-icon>
        </v-list-tile-avatar>
        <v-list-tile-content>
          <v-list-tile-title>未ログイン</v-list-tile-title>
        </v-list-tile-content>
      </template>
      <template id="logged-in" v-else>
        <v-list-tile-avatar>
          <v-icon x-large>portrait</v-icon>
        </v-list-tile-avatar>
        <v-list-tile-content>
          <v-list-tile-title>{{ userInfo.name }}</v-list-tile-title>
        </v-list-tile-content>
      </template>
    </v-list-tile>
  </v-list>
</template>

<script>
export default {
  components: {},
  props: {},
  data() {
    return {}
  },
  computed: {
    isLoggedIn() {
      return this.$store.getters["Sidebar/UserInfo/isLoggedIn"]
    },
    userInfo() {
      return this.$store.getters["Sidebar/UserInfo/userInfo"]
    },
  },
}
</script>
