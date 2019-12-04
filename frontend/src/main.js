// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from "vue"
import Vuetify from "vuetify"
import App from "./App"

// Vuetify
import "./plugins/vuetify"

// Vuex
import store from "./vuex/store"

Vue.config.productionTip = false

// Icon fix
import { Icon } from "leaflet"
delete Icon.Default.prototype._getIconUrl
Icon.Default.mergeOptions({
  iconRetinaUrl: require("leaflet/dist/images/marker-icon-2x.png"),
  iconUrl: require("leaflet/dist/images/marker-icon.png"),
  shadowUrl: require("leaflet/dist/images/marker-shadow.png"),
})

// Import css
import "./stylesheet/main.css"

/* eslint-disable no-new */
Vue.use(Vuetify, {
  theme: {
    primary: "#378640",
  },
})
// Vue インスタンスの生成
const vm = new Vue({
  // el: マウント先のDOMセレクタ
  el: "#app",
  // Vuex データストアを Vue インスタンスから参照できるように
  store,
  // 使用するコンポーネントを宣言
  components: {
    App,
  },
  // マウントされるテンプレート
  template: "<App/>",
})
