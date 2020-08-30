const axiosBase = require("axios")
const axios = axiosBase.create({
  baseURL: "http://localhost:1323",
  headers: {
    "Content-Type": "application/json",
    "X-Requested-With": "XMLHttpRequest",
  },
  responseType: "json",
})

export async function suggest(keyword) {
  // TODO: リクエスト間隔を空けるような実装に (例: 100ms ごと)

  return await axios
    .get(
      "/buildings/suggest",
      {
        params: { keyword },
      },
      {
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
      }
    )
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
