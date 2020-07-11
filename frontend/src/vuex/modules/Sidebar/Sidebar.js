// Vuex::Sidebar -- Sidebar.vue に関するデータストア

import UserInfo from "./UserInfo"
import Alert from "./Alert"

const store = {
  // 名前空間を定義することで、変数名の重複
  // ができるようになる
  namespaced: true,

  // 使用するモジュールの宣言
  modules: {
    UserInfo,
    Alert,
  },

  // ステート(データ)
  state: {
    // isVisible: Sidebar の表示・非表示フラグ
    isVisible: true,
    // visibleForm: 表示されているフォームの種別
    visibleForm: "",
  },

  // Vuex は Vue と異なり、this インスタンスを介した
  // データアクセスは行わない。これにより、現在の状態
  // に依存しないため、より広範囲のモジュールから安定
  // したデータアクセスが可能となっている

  // ゲッター
  getters: {
    // 各関数で必要となる変数は引数で渡す
    // ゲッター関数は引数に以下のものを取ることができる
    // - state: データストア。現在の値を参照するために使用する
    //          モジュールの場合は、ローカルのデータストア
    // - getters: ゲッター関数一覧
    // - rootState: 親モジュールの データストア
    // - rootGetter: 親モジュールのゲッター関数一覧
    isVisible(state) {
      return state.isVisible
    },
    visibleForm(state) {
      return state.visibleForm
    },
  },

  // ミューテーション(セッター)
  mutations: {
    // ミューテーション関数は、引数に以下のものを取ることができる
    // - state: データストア。変更先の変数を参照するために使用される
    // - payload: 呼び出し側から与えられるデータペイロード。
    //            変更値そのものや、変更するために必要な値が渡される
    isVisible(state, payload) {
      state.isVisible = payload
    },
    visibleForm(state, payload) {
      state.visibleForm = payload
    },
  },

  // アクション
  // 外部からのデータ変更の場合、直接ミューテーションを
  // 呼び出すことでも可能だが、アクションは非同期処理
  // によってデータの変更を行うので、アクションを介して
  // のデータ変更が推奨される
  actions: {
    // アクション関数は、第1引数に Vuex を参照するためのオブジェクト、
    // 第2引数に呼び出し側から与えられるデータペイロードが指定できる
    // 第1引数は次のような構造になっている
    // {
    //   state,       ローカルのデータストア
    //   rootState,   親モジュールのデータストア
    //   commit,      ローカルのミューテーション
    //   dispatch,    ローカルのアクション
    //   getters,     ローカルのゲッター
    //   rootGetters  親モジュールのゲッター
    // }
    isVisible({ commit }, payload) {
      commit("isVisible", payload)
    },
    // isVisible をトグル切り替えするためのアクション
    toggleVisiblity({ getters, commit }) {
      commit("isVisible", !getters.isVisible)
    },
    visibleForm({ commit }, payload) {
      commit("visibleForm", payload)
    },
    // 各フォームの「キャンセル」ボタンで閉じるためのアクション
    closeForm({ commit }) {
      commit("visibleForm", "")
    },
  },
}
export default store
