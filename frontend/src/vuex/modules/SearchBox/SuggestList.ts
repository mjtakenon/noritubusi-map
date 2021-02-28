import { mutationTree } from "typed-vuex"

import { SuggestBuilding } from "@/entities/SuggestBuilding"

export type State = {
  buildings: Array<SuggestBuilding>
  keyword: string
}

const state = (): State => ({
  buildings: [],
  keyword: "",
})

const mutations = mutationTree(state, {
  setBuildings(state, buildings: Array<SuggestBuilding>) {
    state.buildings = buildings
  },
  setKeyword(state, keyword: string) {
    state.keyword = keyword
  },
})

export default {
  namespaced: true,
  state,
  mutations,
}
