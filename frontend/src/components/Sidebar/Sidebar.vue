<template>
  <!-- サイドバーとして Navigation Drawer を利用 -->
  <v-navigation-drawer aboslute temporary width="300px" v-model="isVisible">
    <v-container>
      <v-row
        class="text-xs-center"
        id="account-controls"
        v-show="!isFormVisible"
      >
        <v-col>
          <UserInfo />
          <!-- サインアップ・ログイン用のボタン -->
          <div class="d-flex justify-space-around">
            <v-btn @click="visibleForm = 'signup'">アカウント登録</v-btn>
            <v-btn @click="visibleForm = 'login'">ログイン</v-btn>
          </div>
        </v-col>
      </v-row>
      <!-- ログイン処理通知用 -->
      <Alert />
      <!-- サインアップ用フォーム -->
      <SignupForm v-if="visibleForm === 'signup'" />
      <!-- ログイン用フォーム -->
      <LoginForm v-if="visibleForm === 'login'" />
    </v-container>
  </v-navigation-drawer>
</template>

<script lang="ts">
import Vue from "vue"

import { FormType } from "@/entities/Sidebar"

import UserInfo from "@/components/Sidebar/UserInfo.vue"
import Alert from "@/components/Sidebar/Alert.vue"
import SignupForm from "@/components/Sidebar/SignupForm.vue"
import LoginForm from "@/components/Sidebar/LoginForm.vue"

export default Vue.extend({
  components: {
    UserInfo,
    Alert,
    SignupForm,
    LoginForm,
  },
  methods: {
    onClickLogout() {
      this.$accessor.Sidebar.UserInfo.logout()
    },
  },
  computed: {
    isVisible: {
      get(): boolean {
        return this.$accessor.Sidebar.isVisible
      },
      set(value: boolean) {
        this.$accessor.Sidebar.setVisible(value)
      },
    },
    visibleForm: {
      get(): FormType {
        return this.$accessor.Sidebar.visibleForm
      },
      set(value: FormType) {
        this.$accessor.Sidebar.setVisibleForm(value)
      },
    },
    isFormVisible(): boolean {
      return this.$accessor.Sidebar.isFormVisible
    },
  },
})
</script>
