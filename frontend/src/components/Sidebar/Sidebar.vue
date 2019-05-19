<template>
  <!-- サイドバーとして Navigation Drawer を利用 -->
  <v-navigation-drawer :style="style" aboslute temporary v-model="isVisible">
    <UserInfo/>
    <!-- サインアップ・ログイン用のボタン -->
    <div class="text-xs-center" id="account-controls" v-show="!isFormVisible">
      <v-btn @click="onClickSignup">アカウント登録</v-btn>
      <v-btn @click="onClickLogin">ログイン</v-btn>
    </div>
    <!-- ログイン処理通知用 -->
    <AlertModal module-name="successModal" type="success"/>
    <AlertModal module-name="errorModal" type="error"/>
    <!-- サインアップ用フォーム -->
    <SignupForm v-if="isSignupVisible"/>
    <!-- ログイン用フォーム -->
    <LoginForm v-if="isLoginVisible"/>
  </v-navigation-drawer>
</template>

<script>
// 各 Vue コンポーネントのインポート
import UserInfo from "./UserInfo"
import AlertModal from "./AlertModal"
import SignupForm from "./SignupForm"
import LoginForm from "./LoginForm"

export default {
  // 使用するコンポーネントを宣言
  components: {
    UserInfo,
    AlertModal,
    SignupForm,
    LoginForm,
  },
  // データ
  data() {
    return {
      // style: サイドバーのスタイル指定
      // テンプレート上に書くと見通しが悪くなるので、分離
      style: {
        // サイドバーの横幅は、画面全体に対して 30％ に
        width: "30%",
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
  },
}
</script>
