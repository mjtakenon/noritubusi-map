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
      type: "info",
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
    type(state) {
      return state.type
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
    type(state, payload) {
      state.type = payload
    },
  },

  // アクション
  actions: {
    setVisible({ commit }, payload) {
      commit("isVisible", payload)
    },
    setData({ commit }, payload) {
      commit("message", payload.message)
      commit("type", payload.type)
      commit("isVisible", true)
    },
  },
}
export default store
