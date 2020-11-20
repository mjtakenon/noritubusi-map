<style scoped></style>

<template>
  <v-list>
    <v-list-item>
      <!-- 「未ログイン」時に表示されるテンプレート -->
      <template id="default" v-if="!isLoggedIn">
        <!-- アバターアイコン -->
        <v-list-item-avatar>
          <v-icon x-large>portrait</v-icon>
        </v-list-item-avatar>
        <!-- ユーザー名表示 -->
        <v-list-item-content>
          <v-list-item-title>未ログイン</v-list-item-title>
        </v-list-item-content>
      </template>
      <!-- 「ログイン」時に表示されるテンプレート -->
      <template id="logged-in" v-else>
        <!-- アバターアイコン -->
        <v-list-item-avatar>
          <v-icon x-large>portrait</v-icon>
        </v-list-item-avatar>
        <!-- ユーザー名表示 -->
        <v-list-item-content>
          <v-list-item-title>{{ userInfo.userId }}</v-list-item-title>
        </v-list-item-content>
      </template>
    </v-list-item>
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
