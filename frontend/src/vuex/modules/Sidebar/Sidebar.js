// Vuex::Sidebar -- Sidebar.vue に関するデータストア

import AlertModal from './AlertModal'
import UserInfo from './UserInfo'

const store = {
  namespaced: true,
  modules: {
    userInfo: UserInfo,
    successModal: AlertModal,
    errorModal: AlertModal,
  },
  state: {
    isVisible: true,
    visibleForm: "",
  },
  getters: {
    isVisible(state) {
      return state.isVisible
    },
    visibleForm(state) {
      return state.visibleForm
    },
  },
  mutations: {
    isVisible(state, payload) {
      state.isVisible = payload
    },
    visibleForm(state, payload) {
      state.visibleForm = payload
    },
  },
  actions: {
    isVisible({
      commit
    }, payload) {
      commit('isVisible', payload)
    },
    toggleVisiblity({
      getters,
      commit
    }) {
      commit('isVisible', !getters.isVisible)
    },
    visibleForm({
      commit
    }, payload) {
      commit('visibleForm', payload)
    },
    closeForm({
      commit
    }) {
      commit('visibleForm', '')
    },
  },
}
export default store
