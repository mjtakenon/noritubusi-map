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

  // クエリを投げる
  return await axios
    .post("/signup", {
      userid: username,
      password,
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

export async function login(username, password) {
  let params = new URLSearchParams()

  // クエリを投げる
  return await axios
    .post("/signin", {
      userid: username,
      password,
    })
    .then(response => {
      // 200 OK なら成功なのでレスポンスを返す それ以外なら例外を投げる
      if (response.status === 200) {
        return response
      } else {
        throw response
      }
    })
    .catch(error => {
      throw error.response
    })
}
