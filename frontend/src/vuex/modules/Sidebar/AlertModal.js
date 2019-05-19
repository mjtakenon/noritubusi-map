// Vuex::Sidebar::AlertModal -- ログイン処理時のメッセージモーダル

const store = {
  namespaced: true,
  state() {
    return {
      isVisible: false,
      message: "テストメッセージです"
    }
  },
  getters: {
    isVisible(state) {
      return state.isVisible
    },
    message(state) {
      return state.message
    }
  },
  mutations: {
    isVisible(state, payload) {
      state.isVisible = payload
    },
    message(state, payload) {
      state.message = payload
    }
  },
  actions: {
    isVisible({
      commit
    }, payload) {
      commit('isVisible', payload)
    },
    setMessage({
      commit
    }, payload) {
      commit('message', payload)
    }
  }
}
export default store
