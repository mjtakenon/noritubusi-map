// store.js -- Vuex ルートデータストア

// Vue, Vuex のインポート
import Vue from "vue"
import Vuex from "vuex"
import { useAccessor } from "typed-vuex"

// データストアのインポート
import Sidebar from "./modules/Sidebar"
import Map from "./modules/Map"
import SuggestList from "./modules/SearchBox/SuggestList"
import TripRecord from "./modules/TripRecord"

const storePatterns = {
  // strict モード: 未定義の変数に対する代入処理を無効
  strict: true,
  modules: {
    Map,
    Sidebar,
    SuggestList,
    TripRecord,
  },
}

Vue.use(Vuex)
const store = new Vuex.Store(storePatterns)

export const accessor = useAccessor(store, storePatterns)
export default store

Vue.prototype.$accessor = accessor
