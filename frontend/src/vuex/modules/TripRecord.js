const store = {
  namespaced: true,

  // ステート(データ)
  state: {
    stationFrom: {
      id: 0,
      name: "",
      lines: [],
    },
    stationTo: {
      id: 0,
      name: "",
      lines: [],
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
    isStationFromConfirmed(state) {
      return state.stationFrom.name.length != 0
    },
    isStationToConfirmed(state) {
      return state.stationTo.name.length != 0
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
    resetStationFrom(state) {
      state.stationFrom = {
        id: 0,
        name: "",
      }
    },
    resetStationTo(state) {
      state.stationTo = {
        id: 0,
        name: "",
      }
    },
    resetRailway(state) {
      state.railway = {
        id: 0,
        name: "",
      }
    },
  },
}
export default store
