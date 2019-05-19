<template>
  <v-navigation-drawer :style="style" aboslute temporary v-model="isVisible">
    <UserInfo/>
    <div class="text-xs-center" id="account-controls" v-show="!isFormVisible">
      <v-btn @click="onClickSignup">アカウント登録</v-btn>
      <v-btn @click="onClickLogin">ログイン</v-btn>
    </div>
    <AlertModal module-name="successModal" type="success"/>
    <AlertModal module-name="errorModal" type="error"/>
    <SignupForm v-if="isSignupVisible"/>
    <LoginForm v-if="isLoginVisible"/>
  </v-navigation-drawer>
</template>

<script>
import UserInfo from "./UserInfo"
import AlertModal from "./AlertModal"
import SignupForm from "./SignupForm"
import LoginForm from "./LoginForm"

export default {
  components: {
    UserInfo,
    AlertModal,
    SignupForm,
    LoginForm,
  },
  props: {},
  data() {
    return {
      style: {
        width: "30%",
      },
    }
  },
  methods: {
    onClickSignup() {
      this.visibleForm = "signup"
      console.log("onClickSignup")
    },
    onClickLogin() {
      this.visibleForm = "login"
      console.log("onClickLogin")
    },
  },
  computed: {
    isVisible: {
      get() {
        return this.$store.getters["Sidebar/isVisible"]
      },
      set(value) {
        this.$store.dispatch("Sidebar/isVisible", value)
      },
    },
    visibleForm: {
      get() {
        return this.$store.getters["Sidebar/visibleForm"]
      },
      set(value) {
        this.$store.dispatch("Sidebar/visibleForm", value)
      },
    },
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
