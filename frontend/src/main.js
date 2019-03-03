// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Vuetify from 'vuetify'
import './plugins/vuetify'
import App from './App'
import router from './router'

Vue.config.productionTip = false

// Module: vue2-leaflet
import {
  L
} from 'vue2-leaflet'
import 'leaflet/dist/leaflet.css'

// Import css
import './stylesheet/main.css'

/* eslint-disable no-new */
Vue.use( Vuetify, {
  theme: {
    primary: "#378640"
  }
} )
new Vue( {
  el: '#app',
  router,
  components: {
    App
  },
  template: '<App/>'
} )
