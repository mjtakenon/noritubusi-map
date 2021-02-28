// Vuex::Sidebar -- Sidebar.vue に関するデータストア
import { getterTree, mutationTree } from "typed-vuex"

import { FormType, State as SidebarState } from "@/entities/Sidebar"

import UserInfo from "./UserInfo"
import Alert from "./Alert"

export type State = SidebarState

const state = (): State => ({
  isVisible: false,
  visibleForm: "",
})

const getters = getterTree(state, {
  isVisible(state): boolean {
    return state.isVisible
  },
  isFormVisible(state): boolean {
    return state.visibleForm === "login" || state.visibleForm === "signup"
  },
  visibleForm(state): FormType {
    return state.visibleForm
  },
})

const mutations = mutationTree(state, {
  setVisible(state, payload: boolean) {
    state.isVisible = payload
  },
  setVisibleForm(state, payload: FormType) {
    state.visibleForm = payload
  },
  toggleVisiblity(state) {
    state.isVisible = !state.isVisible
  },
  closeForm(state) {
    state.visibleForm = ""
  },
})

export default {
  namespaced: true,
  modules: {
    UserInfo,
    Alert,
  },
  state,
  getters,
  mutations,
}
