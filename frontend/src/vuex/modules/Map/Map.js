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
    // markerList: マップ上のマーカー
    markerList: [],
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
    markerList(state) {
      return state.markerList
    },
  },

  // ミューテーション
  mutations: {
    mapObject(state, payload) {
      state.mapObject = payload
    },
    zoom(state, payload) {
      state.mapProps.zoom = payload
    },
    center(state, payload) {
      state.mapProps.center = payload
    },
    bounds(state, payload) {
      state.mapProps.bounds = payload
    },
    markerList(state, payload) {
      state.markerList = payload
    },
  },

  // アクション
  actions: {
    updateZoom({
      commit
    }, payload) {
      commit('zoom', payload)
    },
    updateCenter({
      commit
    }, payload) {
      commit('center', payload)
    },
    updateBounds({
      commit
    }, payload) {
      commit('bounds', payload)
    },
    markerList({
      commit
    }, payload) {
      commit('markerList', payload)
    }
  }
}
export default store
