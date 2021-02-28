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
          <v-list-item-title>{{ userName }}</v-list-item-title>
        </v-list-item-content>
      </template>
    </v-list-item>
  </v-list>
</template>

<script lang="ts">
import Vue from "vue"

export default Vue.extend({
  data() {
    return {}
  },
  computed: {
    isLoggedIn(): boolean {
      return this.$accessor.Sidebar.UserInfo.isLoggedIn
    },
    userName(): string | undefined {
      return this.$accessor.Sidebar.UserInfo.userInfo?.userId
    },
  },
})
</script>
