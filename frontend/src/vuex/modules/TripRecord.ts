import { getterTree, mutationTree } from "typed-vuex"
import { TripRecord, RecordValue } from "@/entities/TripRecord"

export type State = TripRecord

const state = (): State => ({
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
})

const getters = getterTree(state, {
  isStationFromConfirmed(state): boolean {
    return state.stationFrom.name.length != 0
  },
  isStationToConfirmed(state): boolean {
    return state.stationTo.name.length != 0
  },
})

const mutations = mutationTree(state, {
  setStationFrom(state, station: RecordValue) {
    state.stationFrom = station
  },
  setStationTo(state, station: RecordValue) {
    state.stationTo = station
  },
  setRailway(state, railway: RecordValue) {
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
})

export default {
  namespaced: true,
  state,
  getters,
  mutations,
} // as "typed-vuex/NuxtStore"
