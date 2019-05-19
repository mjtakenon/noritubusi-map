// store.js -- Vuex ルートデータストア

// Vue, Vuex のインポート
import Vue from 'vue'
import Vuex from 'vuex'

// Sidebar データストアのインポート
import Sidebar from './modules/Sidebar/Sidebar'

// Vue に Vuex の機能を統合
Vue.use(Vuex)

// Vuex Store
const store = new Vuex.Store({
  // strict モード: 未定義の変数に対する代入処理を無効
  strict: true,
  // 使用するモジュールを宣言
  // 各モジュールへのデータアクセスは、ファイルパスのように参照
  modules: {
    Sidebar,
  }
})
export default store
