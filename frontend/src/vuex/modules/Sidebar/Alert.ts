// Vuex::Sidebar::Alert -- ログイン処理時のメッセージモーダル
import { actionTree, mutationTree } from "typed-vuex"

import { Alert, AlertType } from "@/entities/Alert"

export type State = Alert

const state = (): State => ({
  isVisible: false,
  type: "info",
  message: "テストメッセージです",
})

const mutations = mutationTree(state, {
  setVisible(state, payload: boolean) {
    state.isVisible = payload
  },
  setMessage(state, payload: string) {
    state.message = payload
  },
  setType(state, payload: AlertType) {
    state.type = payload
  },
})

const actions = actionTree(
  { state, mutations },
  {
    setVisible({ commit }, payload: boolean) {
      commit("setVisible", payload)
    },
    setData({ commit }, payload: { message: string; type: AlertType }) {
      commit("setMessage", payload.message)
      commit("setType", payload.type)
      commit("setVisible", true)
    },
    close({ commit }) {
      commit("setVisible", false)
    },
  }
)

export default {
  namespaced: true,
  state,
  mutations,
  actions,
}
