// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from "vue";
import Vuetify from "vuetify";
import "./plugins/vuetify";
import App from "./App";
import router from "./router";

Vue.config.productionTip = false;

// Module: vue2-leaflet

// Icon fix
import { Icon } from "leaflet";
delete Icon.Default.prototype._getIconUrl;
Icon.Default.mergeOptions({
  iconRetinaUrl: require("leaflet/dist/images/marker-icon-2x.png"),
  iconUrl: require("leaflet/dist/images/marker-icon.png"),
  shadowUrl: require("leaflet/dist/images/marker-shadow.png")
});

// Import css
import "./stylesheet/main.css";

/* eslint-disable no-new */
Vue.use(Vuetify, {
  theme: {
    primary: "#378640"
  }
});
const vm = new Vue({
  el: "#app",
  router,
  components: {
    App
  },
  template: "<App/>"
});
