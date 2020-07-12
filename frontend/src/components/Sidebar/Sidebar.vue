<template>
  <!-- サイドバーとして Navigation Drawer を利用 -->
  <v-navigation-drawer
    aboslute
    temporary
    v-bind="style.navigationDrawer"
    v-model="isVisible"
  >
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
            <v-btn @click="onClickSignup">アカウント登録</v-btn>
            <v-btn @click="onClickLogin">ログイン</v-btn>
          </div>
        </v-col>
      </v-row>
      <!-- ログイン処理通知用 -->
      <Alert />
      <!-- サインアップ用フォーム -->
      <SignupForm v-if="isSignupVisible" />
      <!-- ログイン用フォーム -->
      <LoginForm v-if="isLoginVisible" />
    </v-container>
  </v-navigation-drawer>
</template>

<script>
// 各 Vue コンポーネントのインポート
import UserInfo from "./UserInfo.vue"
import Alert from "./Alert.vue"
import SignupForm from "./SignupForm.vue"
import LoginForm from "./LoginForm.vue"

export default {
  // 使用するコンポーネントを宣言
  components: {
    UserInfo,
    Alert,
    SignupForm,
    LoginForm,
  },
  // データ
  data() {
    return {
      // style: サイドバーのスタイル指定
      // テンプレート上に書くと見通しが悪くなるので、分離
      style: {
        navigationDrawer: {
          // サイドバーの横幅は、300px に
          width: "300px",
        },
      },
    }
  },
  // メソッド
  methods: {
    // 「アカウント登録」ボタンのイベントハンドラ
    onClickSignup() {
      this.visibleForm = "signup"
      console.log("onClickSignup")
    },
    // 「ログイン」ボタンのイベントハンドラ
    onClickLogin() {
      this.visibleForm = "login"
      console.log("onClickLogin")
    },
    // 「ログアウト」ボタンのイベントハンドラ
    onClickLogout() {
      this.$store.dispatch("Sidebar/UserInfo/logout")
      // TODO: Cookieを消す?
      console.log("onClickLogout")
    },
  },
  computed: {
    // [Vuex] isVisible: Sidebar の表示・非表示フラグ
    isVisible: {
      get() {
        return this.$store.getters["Sidebar/isVisible"]
      },
      set(value) {
        this.$store.dispatch("Sidebar/isVisible", value)
      },
    },
    // [Vuex] visibleForm: 表示されているフォームの種別
    visibleForm: {
      get() {
        return this.$store.getters["Sidebar/visibleForm"]
      },
      set(value) {
        this.$store.dispatch("Sidebar/visibleForm", value)
      },
    },
    // テンプレートで表示・非表示を切り替えるフラグ用
    isFormVisible() {
      return this.visibleForm === "signup" || this.visibleForm === "login"
    },
    isSignupVisible() {
      return this.visibleForm === "signup"
    },
    isLoginVisible() {
      return this.visibleForm === "login"
    },
    userInfo() {
      return this.$store.getters["Sidebar/UserInfo/userInfo"]
    },
  },
}
</script>
