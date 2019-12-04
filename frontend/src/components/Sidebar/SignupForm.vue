<template>
  <v-container fluid text-xs-center>
    <v-layout row wrap>
      <v-flex>
        <!-- 入力フォーム -->
        <v-form>
          <!-- ユーザー名 -->
          <v-text-field
            :rules="validators.username"
            clearable
            label="ユーザー名"
            v-model="username"
          ></v-text-field>
          <!-- パスワード -->
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
          <!-- パスワード確認欄 -->
          <v-text-field
            :rules="validators.passwordConfirm"
            @click:append="
              isFormVisible.passwordConfirm = !isFormVisible.passwordConfirm
            "
            clearable
            counter
            hint="8文字以上で入力してください。"
            label="パスワードの確認"
            v-bind="formAttributeByVisibility(isFormVisible.password)"
            v-model="passwordConfirm"
          ></v-text-field>
        </v-form>
        <!-- 「登録」ボタン -->
        <v-btn
          @click="onClickSignup"
          block
          color="#2196f3"
          large
          v-bind:disabled="!isPreparedSignup"
        >
          <strong>
            <font color="white">登録</font>
          </strong>
        </v-btn>
        <!-- 「キャンセル」ボタン -->
        <v-btn @click="onClickCancel" block large>
          <strong>キャンセル</strong>
        </v-btn>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      // username: 「ユーザー名」フォームに入力された値
      username: "",
      // password: 「パスワード」フォームに入力された値
      password: "",
      // passwordConfirm: 「パスワード確認」フォームに入力された値
      passwordConfirm: "",
      // isFormVisible: 「パスワード」「パスワード確認」欄の文字列表示・非表示のフラグ
      isFormVisible: {
        password: false,
        passwordConfirm: false,
      },
    }
  },
  methods: {
    /*************** イベントリスナ ***************/
    onClickSignup() {
      console.log("onClickSignup")
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
    // isEnoughLength は関数を返す関数(高階関数)
    // こうすることで、処理はそのまま、文字列長(strLength)
    // を変更するだけで、それにあった関数を生成し、
    // v-text-field のバリデータ(rules)に渡すことができる
    isEnoughLength(strLength) {
      return val =>
        val.length >= strLength ||
        `${strLength} 文字以上で入力する必要があります`
    },
    // 値一致のチェック
    isEqualToConstant(constVal, fieldName) {
      return val => val === constVal || `${fieldName} が一致しません`
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
    // [Vuex] isVisible: SignupForm の表示・非表示フラグ
    isVisible() {
      return this.$store.getters["Sidebar/visibleForm"] === "signup"
    },
    // Validators: 各フォームのバリデーション関数
    validators() {
      return {
        // ユーザー名
        username: [this.requireField],
        // パスワード
        password: [this.requireField, this.isEnoughLength(8)],
        // パスワード確認欄
        passwordConfirm: [
          this.requireField,
          this.isEqualToConstant(this.password, "パスワード"),
        ],
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
      const isPasswordConfirmPrepared = this.validators.passwordConfirm.every(
        validator => validator(this.passwordConfirm) === true
      )
      return (
        isUsernamePrepared && isPasswordPrepared && isPasswordConfirmPrepared
      )
    },
  },
}
</script>
