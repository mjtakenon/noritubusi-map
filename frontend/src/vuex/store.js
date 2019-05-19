import Vue from 'vue'
import Vuex from 'vuex'

import Sidebar from './modules/Sidebar/Sidebar'

Vue.use(Vuex)

// Vuex Store
const store = new Vuex.Store({
  strict: true,
  modules: {
    Sidebar,
  }
})
export default store
