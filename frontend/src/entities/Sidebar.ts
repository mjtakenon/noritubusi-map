export type FormType = "" | "login" | "signup"

export interface State {
  isVisible: boolean
  visibleForm: FormType
}
