import { User, LoginInfo, isUser, SignupInfo } from "@/entities/User"
import axios from "axios"

const httpClient = axios.create({
  baseURL: "http://localhost:1323",
  headers: {
    "Content-Type": "application/json",
    "X-Requested-With": "XMLHttpRequest",
  },
  responseType: "json",
})

export async function signup(signupInfo: SignupInfo): Promise<User> {
  // クエリを投げる
  return await httpClient
    .post("/signup", {
      userid: signupInfo.username,
      password: signupInfo.password,
    })
    .then(response => {
      // 201 Created なら成功なのでレスポンスを返す それ以外なら例外を投げる
      if (response.status === 201) {
        const responseData = response.data
        if (isUser(responseData)) {
          return responseData
        } else {
          console.error("Invalid response data structure")
          throw response
        }
      } else {
        throw response
      }
    })
    .catch(error => {
      throw error.response
    })
}

export async function login(loginInfo: LoginInfo): Promise<User> {
  // クエリを投げる
  return httpClient
    .post("/signin", {
      userid: loginInfo.username,
      password: loginInfo.password,
    })
    .then(response => {
      // 200 OK なら成功なのでレスポンスを返す それ以外なら例外を投げる
      if (response.status === 200) {
        const responseData = response.data
        if (isUser(responseData)) {
          return responseData
        } else {
          console.error("Invalid response data structure")
          throw response
        }
      } else {
        throw response
      }
    })
    .catch(error => {
      throw error.response
    })
}
