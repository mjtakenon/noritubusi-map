// store.js -- Vuex ルートデータストア

// Vue, Vuex のインポート
import Vue from "vue"
import Vuex from "vuex"

// データストアのインポート
import Sidebar from "./modules/Sidebar/Sidebar"
import Map from "./modules/Map/Map"
import SuggestList from "./modules/SearchBox/SuggestList"
import TripRecord from "./modules/TripRecord"

// Vue に Vuex の機能を統合
Vue.use(Vuex)

// Vuex Store
const store = new Vuex.Store({
  // strict モード: 未定義の変数に対する代入処理を無効
  strict: true,
  // 使用するモジュールを宣言
  // 各モジュールへのデータアクセスは、ファイルパスのように参照
  modules: {
    Map,
    Sidebar,
    SuggestList,
    TripRecord,
  },
})
export default store