<template>
  <div>
    <v-text-field
      clearable
      counter
      type="text"
      :label="label"
      :rules="rules"
      :value="formValue.value"
      @input="onInput"
      @update:error="updateError"
      @keydown.enter="keyDownEnter"
    >
    </v-text-field>
  </div>
</template>
<script lang="ts">
import Vue, { PropType } from "vue"

import { InputValidationRules } from "vuetify"

import { FormValue } from "@/components/Sidebar/Form/common"

type Data = {
  isVisible: boolean
}

export default Vue.extend({
  name: "ValidatableInput",
  props: {
    formValue: Object as PropType<FormValue>,
    hint: String,
    label: String,
    rules: {
      type: Array as PropType<InputValidationRules>,
      default: () => [],
    },
  },
  data(): Data {
    return {
      isVisible: false,
    }
  },
  methods: {
    onInput(value: string) {
      this.$emit("update:form-value", {
        value,
        hasError: this.formValue.hasError,
      } as FormValue)
    },
    updateError(status: boolean) {
      this.$emit("update:form-value", {
        value: this.formValue.value,
        hasError: status,
      } as FormValue)
    },
    keyDownEnter() {
      this.$emit("keydown.enter")
    },
  },
})
</script>
