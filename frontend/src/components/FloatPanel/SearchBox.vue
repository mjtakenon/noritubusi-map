<template>
  <div>
    <v-toolbar rounded>
      <v-app-bar-nav-icon @click="toggleSidebar"></v-app-bar-nav-icon>
      <!-- テキストボックス(乗車駅) -->
      <v-text-field
        hide-details
        append-icon="search"
        @click:append="focusOnFirstCandidate"
        single-line
        clearable
        :loading="isLoading"
        @input="onInput"
        @change="focusOnFirstCandidate"
        @click:clear="clearInput"
        :value="inputValue"
      ></v-text-field>
      <v-btn @click="onClickSwap" icon>
        <v-icon>swap_horiz</v-icon>
      </v-btn>
    </v-toolbar>
    <SuggestList></SuggestList>
  </div>
</template>

<script lang="ts">
import Vue from "vue"
import _ from "lodash"

import { SuggestBuilding } from "@/entities/SuggestBuilding"
import { searchBuildingsByKeyword } from "../../utils/api/search"

import SuggestList from "@/components/FloatPanel/SuggestList.vue"

type Data = {
  inputValue: string
  isLoading: boolean
  updateSuggestWithInterval: Function | null
}

export default Vue.extend({
  components: {
    SuggestList,
  },
  data(): Data {
    return {
      inputValue: "",
      isLoading: false,
      updateSuggestWithInterval: null,
    }
  },
  created() {
    // サジェストが正しく表示されないバグがあるため
    // クエリを送る頻度を500msに間引き
    this.updateSuggestWithInterval = _.debounce(
      () => this.updateSuggest(),
      500,
      { leading: true }
    )
  },
  // メソッド
  methods: {
    onInput(input: string) {
      // 変換が確定された時に二度イベントが
      // 発火されるのを抑止する
      if (this.inputValue == input) {
        return
      }
      this.inputValue = input
      if (this.updateSuggestWithInterval !== null) {
        this.updateSuggestWithInterval()
      }
    },
    onClickSwap() {
      this.$accessor.TripRecord.setStationFrom(
        this.$accessor.TripRecord.stationTo
      )
      this.$accessor.TripRecord.setStationTo(
        this.$accessor.TripRecord.stationFrom
      )
    },
    clearInput() {
      this.inputValue = ""
    },
    toggleSidebar() {
      this.showSidebar = !this.showSidebar
    },
    updateSuggest() {
      if (this.inputValue === null || this.inputValue.length === 0) {
        this.$accessor.SuggestList.setBuildings([])
        this.isLoading = false
        return
      }
      this.$accessor.SuggestList.setKeyword(this.inputValue)
      this.isLoading = true

      searchBuildingsByKeyword(this.inputValue)
        .then(suggestBuildings =>
          this.$accessor.SuggestList.setBuildings(suggestBuildings)
        )
        .catch(error => console.error(error))
        .finally(() => (this.isLoading = false))
    },
    focusOnFirstCandidate() {
      if (this.buildings.length < 1) {
        return
      }
      const first = this.buildings[0]
      this.$accessor.Map.setPinAndFocus(first)
    },
  },
  computed: {
    showSidebar: {
      get(): boolean {
        return this.$accessor.Sidebar.isVisible
      },
      set(value: boolean) {
        this.$accessor.Sidebar.setVisible(value)
      },
    },
    buildings(): Array<SuggestBuilding> {
      return this.$accessor.SuggestList.buildings
    },
  },
})
</script>
