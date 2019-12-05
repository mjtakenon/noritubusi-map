const axiosBase = require("axios")
const axios = axiosBase.create({
  baseURL: "http://localhost:1323",
  headers: {
    "Content-Type": "application/json",
    "X-Requested-With": "XMLHttpRequest",
  },
  responseType: "json",
})

export async function signup(username, password) {
  let params = new URLSearchParams()
  params.append("userid", username)
  params.append("password", password)

  // クエリを投げる
  return await axios
    .post("/signup", params, {
      headers: { "Content-Type": "application/x-www-form-urlencoded" },
    })
    .then(response => {
      // 201 Created なら成功なのでレスポンスを返す それ以外なら例外を投げる
      if (response.status === 201) {
        return response
      } else {
        throw response
      }
    })
    .catch(error => {
      throw error.response
    })
}
export function login(username, password) {}
