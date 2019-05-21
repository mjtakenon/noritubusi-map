<!-- 
  作りはほとんど SignupForm.vue と同じなので、各テンプレート・処理の詳細
  は SignupForm.vue を参照すること 
-->

<template>
  <v-container fluid text-xs-center>
    <v-layout row wrap>
      <v-flex>
        <v-form>
          <v-text-field
            :rules="validators.username"
            clearable
            label="ユーザー名"
            v-model="username"
          ></v-text-field>
          <v-text-field
            :rules="validators.password"
            @click:append="isFormVisible.password = !isFormVisible.password"
            clearable
            counter
            hint="8文字以上で入力してください。"
            label="パスワード"
            v-bind="formAttributeByVisibility(isFormVisible.password)"
            v-model="password"
          ></v-text-field>
        </v-form>
        <v-btn
          @click="onClickLogin"
          block
          color="#2196f3"
          large
          v-bind:disabled="!isPreparedSignup"
        >
          <strong>
            <font color="white">ログイン</font>
          </strong>
        </v-btn>
        <v-btn @click="onClickCancel" block large>
          <strong>キャンセル</strong>
        </v-btn>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  props: {},
  data() {
    return {
      username: "",
      password: "",
      isFormVisible: {
        password: false,
      },
    }
  },
  methods: {
    /*************** イベントリスナ ***************/
    onClickLogin() {
      console.log("onClickLogin")
    },
    onClickCancel() {
      this.$store.dispatch("Sidebar/closeForm")
    },
    /**********************************************/

    /********* バリデーション用の補助関数 *********/
    // 空文字列であるかどうか
    isEmpty(str) {
      return !str || /^\s*$/.test(str)
    },
    /**********************************************/
    /*********** 各種バリデーション処理 ***********/
    // 必須フィールドチェック
    requireField(val) {
      return !this.isEmpty(val) || "このフィールドは必須です"
    },
    // 文字列長チェック
    isEnoughLength(strLength) {
      return val =>
        val.length >= strLength ||
        `${strLength} 文字以上で入力する必要があります`
    },
    /**********************************************/
    // フォームの表示・非表示状態に基づきHTML属性を変更
    formAttributeByVisibility(isVisible) {
      return {
        appendIcon: isVisible ? "visibility" : "visibility_off",
        type: isVisible ? "text" : "password",
      }
    },
  },
  computed: {
    isVisible() {
      return this.$store.getters["Sidebar/visibleForm"] === "login"
    },
    validators() {
      return {
        username: [this.requireField],
        password: [this.requireField, this.isEnoughLength(8)],
      }
    },
    // すべてのフォームで問題がないかどうか
    isPreparedSignup() {
      const isUsernamePrepared = this.validators.username.every(
        validator => validator(this.username) === true
      )
      const isPasswordPrepared = this.validators.password.every(
        validator => validator(this.password) === true
      )
      return (
        isUsernamePrepared && isPasswordPrepared
      )
    },
  },
}
</script>
