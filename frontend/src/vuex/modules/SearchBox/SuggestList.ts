const store = {
  namespaced: true,
  state: {
    buildings: [],
    keyword: "",
  },
  getters: {
    buildings(state) {
      return state.buildings
    },
    keyword(state) {
      return state.keyword
    },
  },
  mutations: {
    buildings(state, buildings) {
      state.buildings = buildings
    },
    keyword(state, keyword) {
      state.keyword = keyword
    },
  },
}
export default store
