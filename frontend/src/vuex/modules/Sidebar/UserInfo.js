// Vuex::Sidebar::UserInfo -- UserInfo.vue に関するデータストア

const store = {
  namespaced: true,
  state: {
    userInfo: null
  },
  getters: {
    isLoggedIn(state) {
      return state.userInfo != null
    },
    userInfo(state) {
      return state.userInfo
    }
  },
  mutations: {
    userInfo(state, payload) {
      state.userInfo = payload
      if (state.userInfo != null) state.isLoggedIn = true;
    },
  },
  actions: {
    login({
      commit
    }, payload) {
      commit('userInfo', payload)
    },

    logout({
      getters,
      commit
    }) {
      commit('userInfo', null)
    }
  }
}
export default store
