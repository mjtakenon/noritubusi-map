// Vuex::Sidebar::AlertModal -- ログイン処理時のメッセージモーダル

const store = {
  namespaced: true,

  // データストア
  // 複数のインスタンスが生成されるモジュールの場合、state を
  // オブジェクトではなく関数として定義する必要がある。
  // そうしないと、すべてのインスタンスで state が共有されてしまう
  state() {
    return {
      isVisible: false,
      message: "テストメッセージです",
    }
  },

  // ゲッター
  getters: {
    isVisible(state) {
      return state.isVisible
    },
    message(state) {
      return state.message
    },
  },

  // ミューテーション
  mutations: {
    isVisible(state, payload) {
      state.isVisible = payload
    },
    message(state, payload) {
      state.message = payload
    },
  },

  // アクション
  actions: {
    isVisible({ commit }, payload) {
      commit("isVisible", payload)
    },
    setMessage({ commit }, payload) {
      commit("message", payload)
    },
  },
}
export default store
