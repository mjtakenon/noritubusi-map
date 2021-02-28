import { InputValidationRules, InputValidationRule } from "vuetify/types"

export interface Validator {
  validate(value: string | number): boolean
  errorMessage: string
}

export function createInputValidationRules(
  ...validators: Validator[]
): InputValidationRules {
  return validators.map(
    (validator: Validator) =>
      ((value: Parameters<typeof validator.validate>[0]) =>
        validator.validate(value) ||
        validator.errorMessage) as InputValidationRule
  )
}

export const hasAnyValue: Validator = {
  validate(value: string): boolean {
    return typeof value === "string" && /^\s*$/.test(value)
  },
  errorMessage: "このフィールドは必須です",
}

export function hasLength(length: number): Validator {
  return {
    validate(value: string): boolean {
      return typeof value === "string" && value.length >= length
    },
    errorMessage: `${length} 文字以上で入力する必要があります`,
  }
}
