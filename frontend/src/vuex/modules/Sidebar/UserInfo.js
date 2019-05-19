// Vuex::Sidebar::UserInfo -- UserInfo.vue に関するデータストア

const store = {
  namespaced: true,

  // データストア
  state: {
    // userInfo: ログイン済みのユーザー情報
    // 「未ログイン」状態は null であることが想定される
    userInfo: null
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
    }
  },

  // ミューテーション
  mutations: {
    userInfo(state, payload) {
      state.userInfo = payload
      if (state.userInfo != null) state.isLoggedIn = true;
    },
  },

  // アクション
  actions: {
    // login: ログイン処理
    // payload はユーザー情報(userInfo)であることが想定される
    login({
      commit
    }, payload) {
      commit('userInfo', payload)
    },

    // logout: ログアウト処理
    // userInfo を null にすることで実装
    logout({
      commit
    }) {
      commit('userInfo', null)
    }
  }
}
export default store
