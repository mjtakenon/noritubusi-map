import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

// Vuex Store
const store = new Vuex.Store({
  state: {
    showSidebar: false
  },
  getters: {
    isVisible(state) {
      return state.showSidebar
    }
  },
  mutations: {
    isVisible(state, payload) {
      state.showSidebar = payload
    }
  }
})
export default store
