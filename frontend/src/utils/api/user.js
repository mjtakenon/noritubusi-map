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
  let res = null
  params.append("userid", username)
  params.append("password", password)
  console.log("/signup?" + params.toString())

  await axios
    .post("/signup?" + params.toString())
    .then(response => {
      res = response
    })
    .catch(error => {
      res = error
    })
  return res
}
export function login(username, password) {}
