<style scoped>
</style>

<template>
  <v-list>
    <v-list-tile>
      <!-- 「未ログイン」時に表示されるテンプレート -->
      <template id="default" v-if="!isLoggedIn">
        <!-- アバターアイコン -->
        <v-list-tile-avatar>
          <v-icon x-large>portrait</v-icon>
        </v-list-tile-avatar>
        <!-- ユーザー名表示 -->
        <v-list-tile-content>
          <v-list-tile-title>未ログイン</v-list-tile-title>
        </v-list-tile-content>
      </template>
      <!-- 「ログイン」時に表示されるテンプレート -->
      <template id="logged-in" v-else>
        <!-- アバターアイコン -->
        <v-list-tile-avatar>
          <v-icon x-large>portrait</v-icon>
        </v-list-tile-avatar>
        <!-- ユーザー名表示 -->
        <v-list-tile-content>
          <v-list-tile-title>{{ userInfo.name }}</v-list-tile-title>
        </v-list-tile-content>
      </template>
    </v-list-tile>
  </v-list>
</template>

<script>
export default {
  data() {
    return {}
  },
  computed: {
    // [Vuex] isLoggedIn: ログイン状態の確認
    isLoggedIn() {
      return this.$store.getters["Sidebar/UserInfo/isLoggedIn"]
    },
    userInfo() {
      return this.$store.getters["Sidebar/UserInfo/userInfo"]
    },
  },
}
</script>
