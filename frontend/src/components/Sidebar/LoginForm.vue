<template>
  <v-container fluid text-xs-center>
    <v-layout row wrap>
      <v-flex>
        <ValidatableInput
          :form-value.sync="username"
          :rules="rules.username"
          @keydown.enter="onClickLogin"
        >
        </ValidatableInput>
        <PasswordInput
          :form-value.sync="password"
          :rules="rules.password"
          @keydown.enter="onClickLogin"
        ></PasswordInput>
        <v-form>
          <v-btn
            block
            class="text-font-bold"
            color="#2196f3"
            large
            :disabled="!isAllValid"
            @click="onClickLogin"
          >
            ログイン
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
}

type RuleSet = {
  username: InputValidationRules
  password: InputValidationRules
}

export default Vue.extend({
  name: "LoginForm",
  components: {
    ValidatableInput,
    PasswordInput,
  },
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
    }
  },
  computed: {
    isAllValid(): boolean {
      return !this.username.hasError && !this.password.hasError
    },
    rules(): RuleSet {
      return {
        username: createInputValidationRules(hasAnyValue),
        password: createInputValidationRules(hasAnyValue, hasLength(8)),
      }
    },
  },
  methods: {
    onClickLogin() {
      if (!this.isAllValid) {
        return
      }
      this.$accessor.Sidebar.UserInfo.login({
        username: this.username.value,
        password: this.password.value,
      })
    },
    onClickCancel() {
      this.$accessor.Sidebar.closeForm
    },
  },
})
</script>
