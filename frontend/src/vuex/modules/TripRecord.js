const store = {
  namespaced: true,

  // ステート(データ)
  state: {
    stationFrom: {
      id: 0,
      name: "",
    },
    stationTo: {
      id: 0,
      name: "",
    },
    railway: {
      id: 0,
      name: "",
    },
  },

  // ゲッター
  getters: {
    stationFrom(state) {
      return state.stationFrom
    },
    stationTo(state) {
      return state.stationTo
    },
    railway(state) {
      return state.railway
    },
  },

  // ミューテーション(セッター)
  mutations: {
    stationFrom(state, station) {
      state.stationFrom = station
    },
    stationTo(state, station) {
      state.stationTo = station
    },
    railway(state, railway) {
      state.railway = railway
    },
  },
}
export default store
