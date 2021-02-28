import { getterTree, mutationTree } from "typed-vuex"

import { LatLng } from "@/entities/Common"
import { Bounds, MapOptions } from "@/entities/Map"
import { Pin } from "@/entities/Pin"
import { SuggestBuilding } from "@/entities/SuggestBuilding"

export type State = {
  lMapOptions: MapOptions
  pins: Array<Pin>
}

const state = (): State => ({
  lMapOptions: {
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
  pins: [],
})

// ゲッター
const getters = getterTree(state, {
  zoom(state): number {
    return state.lMapOptions.zoom
  },
  center(state): LatLng {
    return state.lMapOptions.center
  },
  bounds(state): Bounds {
    return state.lMapOptions.bounds
  },
})

// ミューテーション
const mutations = mutationTree(state, {
  setZoom(state, payload: number) {
    state.lMapOptions.zoom = payload
  },
  setCenter(state, payload: LatLng) {
    state.lMapOptions.center = payload
  },
  setBounds(state, payload: Bounds) {
    state.lMapOptions.bounds = payload
  },
  setPins(state, payload: Array<Pin>) {
    state.pins = payload
  },
  addPin(state, payload: Pin) {
    state.pins.push(payload)
  },
  addPins(state, payload: Array<Pin>) {
    state.pins.push(...payload)
  },
  setPinAndFocus(state, payload: SuggestBuilding) {
    state.pins = [
      {
        latLng: {
          lat: payload.latitude,
          lng: payload.longitude,
        },
        popup: {
          stationName: payload.name,
          lines: payload.lines,
        },
        openPopup: true,
        focusPin: true,
      },
    ]
  },
})

// アクション

export default {
  namespaced: true,
  state,
  getters,
  mutations,
} // as "typed-vuex/NuxtStore"
