import Vue from "vue"
import { Icon } from "leaflet"

import App from "@/components/App.vue"
import vuetify from "./plugins/vuetify" // Vuetify
import store from "./vuex" // Vuex

// Import stylesheet
import "./stylesheet/style.scss"

Vue.config.productionTip = false

// Icon fix
type D = Icon.Default & {
  _getIconUrl?: string
}
delete (Icon.Default.prototype as D)._getIconUrl
Icon.Default.mergeOptions({
  iconRetinaUrl: require("leaflet/dist/images/marker-icon-2x.png"),
  iconUrl: require("leaflet/dist/images/marker-icon.png"),
  shadowUrl: require("leaflet/dist/images/marker-shadow.png"),
})

// Vue インスタンスの生成
new Vue({
  store,
  vuetify,
  render: h => h(App),
}).$mount("#app")
