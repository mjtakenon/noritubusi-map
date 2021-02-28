export type LoginInfo = {
  username: string
  password: string
}

export type SignupInfo = LoginInfo

export interface User {
  userId: string
}

export function isUser(obj: any): obj is User {
  return (
    obj !== null && typeof obj === "object" && typeof obj.userId === "string"
  )
}
