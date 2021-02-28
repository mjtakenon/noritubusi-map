<template>
  <v-container fluid text-xs-center>
    <v-layout row wrap>
      <v-flex>
        <ValidatableInput
          :form-value.sync="username"
          :rules="rules.username"
          @keydown.enter="onClickSignup"
        >
        </ValidatableInput>
        <PasswordInput
          :form-value.sync="password"
          :rules="rules.password"
          @keydown.enter="onClickSignup"
        ></PasswordInput>
        <PasswordInput
          :form-value.sync="confirmPassword"
          :rules="rules.confirmPassword"
          @keydown.enter="onClickSignup"
        ></PasswordInput>
        <v-form>
          <v-btn
            block
            class="text-font-bold"
            color="#2196f3"
            large
            :disabled="!isAllValid"
            @click="onClickSignup"
          >
            登録
          </v-btn>
          <v-btn block class="text-font-bold" large @click="onClickCancel">
            キャンセル
          </v-btn>
        </v-form>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script lang="ts">
import Vue from "vue"

import { InputValidationRules } from "vuetify"

import {
  Validator,
  createInputValidationRules,
  hasAnyValue,
  hasLength,
} from "@/utils/validators"

import { FormValue } from "@/components/Sidebar/Form/common"

import ValidatableInput from "@/components/Sidebar/Form/ValidatableInput.vue"
import PasswordInput from "@/components/Sidebar/Form/PasswordInput.vue"

type Data = {
  username: FormValue
  password: FormValue
  confirmPassword: FormValue
}

type RuleSet = {
  username: InputValidationRules
  password: InputValidationRules
  confirmPassword: InputValidationRules
}

function isSamePassword(password: string): Validator {
  return {
    validate(confirmPassword: string) {
      return password === confirmPassword
    },
    errorMessage: "同じパスワードを入力してください",
  }
}

export default Vue.extend({
  name: "SignupForm",
  data(): Data {
    return {
      username: {
        value: "",
        hasError: true,
      },
      password: {
        value: "",
        hasError: true,
      },
      confirmPassword: {
        value: "",
        hasError: true,
      },
    }
  },
  computed: {
    isAllValid(): boolean {
      return (
        !this.username.hasError &&
        !this.password.hasError &&
        !this.confirmPassword.hasError
      )
    },
    rules(): RuleSet {
      return {
        username: createInputValidationRules(hasAnyValue),
        password: createInputValidationRules(hasAnyValue, hasLength(8)),
        confirmPassword: createInputValidationRules(
          hasAnyValue,
          hasLength(8),
          isSamePassword(this.password.value)
        ),
      }
    },
  },
  methods: {
    onClickSignup() {
      if (!this.isAllValid) {
        return
      }
      this.$accessor.Sidebar.UserInfo.signup({
        username: this.username.value,
        password: this.password.value,
      })
    },
    onClickCancel() {
      this.$accessor.Sidebar.closeForm
    },
    isSamePassword,
  },
})
</script>
