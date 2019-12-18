// Vuex::Sidebar::UserInfo -- UserInfo.vue に関するデータストア

import { signup, login } from "../../../utils/api/user.js"

const store = {
  namespaced: true,

  // データストア
  state: {
    // userInfo: ログイン済みのユーザー情報
    // 「未ログイン」状態は null であることが想定される
    userInfo: null,
  },

  // ゲッター
  getters: {
    // isLoggedIn: ログイン済みかどうかのフラグ
    // userInfo の Null チェックにより実装
    isLoggedIn(state) {
      return state.userInfo != null
    },
    userInfo(state) {
      return state.userInfo
    },
  },

  // ミューテーション
  mutations: {
    userInfo(state, payload) {
      state.userInfo = payload
      if (state.userInfo != null) state.isLoggedIn = true
    },
  },

  // アクション
  actions: {
    // login: ログイン処理
    // payload はユーザー情報(userInfo)であることが想定される
    async login({ commit, dispatch }, payload) {
      const { username, password } = payload
      let response = null

      dispatch("Sidebar/Alert/close", null, { root: true })

      try {
        response = await login(username, password)
        console.log(response)

        dispatch(
          "Sidebar/Alert/setData",
          {
            type: "success",
            message: "ログインに成功しました。",
          },
          { root: true }
        )
        dispatch("Sidebar/closeForm", null, { root: true })
      } catch (error) {
        console.error(error)

        dispatch(
          "Sidebar/Alert/setData",
          {
            type: "error",
            message: "ログインに失敗しました。",
          },
          { root: true }
        )
      }
      commit("userInfo", response.data)
    },

    // logout: ログアウト処理
    // userInfo を null にすることで実装
    logout({ commit, dispatch }) {
      dispatch("Sidebar/Alert/close", null, { root: true })
      dispatch(
        "Sidebar/Alert/setData",
        {
          type: "success",
          message: "ログアウトしました。",
        },
        { root: true }
      )
      commit("userInfo", null)
    },

    async signup({ commit, dispatch }, payload) {
      const { username, password } = payload
      let response = null

      dispatch("Sidebar/Alert/close", null, { root: true })

      try {
        response = await signup(username, password)
        console.log(response)

        dispatch(
          "Sidebar/Alert/setData",
          {
            type: "success",
            message: "サインアップに成功しました。",
          },
          { root: true }
        )
        dispatch("Sidebar/closeForm", null, { root: true })
      } catch (error) {
        console.error(error)

        dispatch(
          "Sidebar/Alert/setData",
          {
            type: "error",
            message: "サインアップに失敗しました。",
          },
          { root: true }
        )
      }
      commit("userInfo", response.data)
    },
  },
}
export default store
