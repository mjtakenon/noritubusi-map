// Vuex::Sidebar::UserInfo -- UserInfo.vue に関するデータストア

import { actionTree, getterTree, mutationTree } from "typed-vuex"

import { accessor } from "@/vuex"

import { LoginInfo, SignupInfo, User } from "@/entities/User"
import { signup, login } from "@/utils/api/user"

export type State = {
  userInfo: User | null
}

const state = (): State => ({
  userInfo: null,
})

const getters = getterTree(state, {
  isLoggedIn(state): boolean {
    return state.userInfo != null
  },
})

const mutations = mutationTree(state, {
  setUserInfo(state, payload: User) {
    state.userInfo = payload
  },
  clearUserInfo(state) {
    state.userInfo = null
  },
})

const actions = actionTree(
  { state, getters, mutations },
  {
    async login({ commit }, payload: LoginInfo) {
      accessor.Sidebar.Alert.close()
      try {
        const loginUser = await login(payload)
        console.log(loginUser)

        accessor.Sidebar.Alert.setData({
          type: "success",
          message: "ログインに成功しました。",
        })
        commit("setUserInfo", loginUser)
        accessor.Sidebar.closeForm()
      } catch (error) {
        console.error(error)

        accessor.Sidebar.Alert.setData({
          type: "error",
          message: "ログインに失敗しました。",
        })
      }
    },

    // logout: ログアウト処理
    // userInfo を null にすることで実装
    // TODO: ログアウト時にcookieを消す等の処理が必要があれば実装する
    logout({ commit }) {
      accessor.Sidebar.Alert.close()
      accessor.Sidebar.Alert.setData({
        type: "success",
        message: "ログアウトしました。",
      })
      commit("clearUserInfo")
    },

    async signup({ commit }, payload: SignupInfo) {
      accessor.Sidebar.Alert.close()
      try {
        const signupUser = await signup(payload)
        console.log(signupUser)

        accessor.Sidebar.Alert.setData({
          type: "success",
          message: "サインアップに成功しました。",
        })
        commit("setUserInfo", signupUser)
        accessor.Sidebar.closeForm()
      } catch (error) {
        console.error(error)

        accessor.Sidebar.Alert.setData({
          type: "error",
          message: "サインアップに失敗しました。",
        })
      }
    },
  }
)

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
}
