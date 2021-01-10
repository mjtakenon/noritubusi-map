const store = {
  namespaced: true,
  state: {
    buildings: [],
  },
  getters: {
    buildings(state) {
      return state.buildings
    },
  },
  mutations: {
    buildings(state, buildings) {
      state.buildings = buildings
    },
  },
}
export default store
