// Vuex::Map -- Map.vue に関するデータストア

const store = {
  namespaced: true,

  // データ
  state: {
    // mapProps: マップ に関するデータ
    mapProps: {
      // zoom: ズームスケール
      zoom: 14,
      // center: 中心座標
      center: {
        lat: 35.680446,
        lng: 139.761801,
      },
      // bounds: 表示範囲
      bounds: {
        // 左下
        _southWest: {
          lat: 35.63532680480169,
          lng: 139.73910056054595,
        },
        // 右上
        _northEast: {
          lat: 35.691113860493594,
          lng: 139.79489050805572,
        },
      },
    },
    // pins: マップ上のマーカー
    /***** データ構造サンプル ***************
     * pins: [
     *   {
     *     latLng: [12.3333333, 45.6666666],
     *     popup: {
     *       name: "駅建物名",
     *       // lines は /buildings/suggest のレスポンスと同じ構造
     *       lines: [ 
     *         {
     *           railwayName: "路線名"
     *           station_id,
     *           order_in_railway
     *         }
     *       ]
     *     }
     *   }
     * ],
     ******************************************/
    pins: [],
  },

  // ゲッター
  getters: {
    zoom(state) {
      return state.mapProps.zoom
    },
    center(state) {
      return state.mapProps.center
    },
    bounds(state) {
      return state.mapProps.bounds
    },
    pins(state) {
      console.log(state.pins)
      return state.pins
    },
  },

  // ミューテーション
  mutations: {
    zoom(state, payload) {
      state.mapProps.zoom = payload
    },
    center(state, payload) {
      state.mapProps.center = payload
    },
    bounds(state, payload) {
      state.mapProps.bounds = payload
    },
    setPins(state, payload) {
      state.pins = payload
    },
    addPin(state, payload) {
      state.pins.push(payload)
    },
  },

  // アクション
  actions: {
    updateZoom({ commit }, payload) {
      commit("zoom", payload)
    },
    updateCenter({ commit }, payload) {
      commit("center", payload)
    },
    updateBounds({ commit }, payload) {
      commit("bounds", payload)
    },
    setPins({ commit }, payload) {
      commit("setPins", payload)
    },
    addPin({ commit }, payload) {
      commit("addPin", payload)
    },
  },
}
export default store
