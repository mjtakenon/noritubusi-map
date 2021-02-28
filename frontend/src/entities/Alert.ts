export type AlertType = "info" | "error" | "success" | "warning"

export interface Alert {
  isVisible: boolean
  type: AlertType
  message: string
}
