<template>
  <div>
    <v-toolbar rounded>
      <v-app-bar-nav-icon @click="toggleSidebar"></v-app-bar-nav-icon>
      <!-- テキストボックス(乗車駅) -->
      <v-text-field
        hide-details
        append-icon="search"
        single-line
        clearable
        :loading="isLoading"
        @input="onInput"
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

<script>
import _ from "lodash"
import { suggest } from "../../utils/api/search.js"

import SuggestList from "./SuggestList"

export default {
  components: {
    SuggestList,
  },
  data() {
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
    onInput(input) {
      // 変換が確定された時に二度イベントが
      // 発火されるのを抑止する
      if (this.inputValue == input) {
        return
      }
      this.inputValue = input
      this.updateSuggestWithInterval()
    },
    onClickSwap() {
      const stationFrom = this.$store.getters["TripRecord/stationFrom"]
      const stationTo = this.$store.getters["TripRecord/stationTo"]

      this.$store.commit("TripRecord/stationFrom", stationTo)
      this.$store.commit("TripRecord/stationTo", stationFrom)
    },
    clearInput() {
      this.inputValue = ""
    },
    toggleSidebar() {
      this.showSidebar = !this.showSidebar
    },
    updateSuggest() {
      if (this.inputValue === null || this.inputValue.length === 0) {
        this.$store.commit("SuggestList/buildings", [])
        this.isLoading = false
        return
      }
      this.$store.commit("SuggestList/keyword", this.inputValue)
      this.isLoading = true
      suggest(this.inputValue)
        .then(response =>
          this.$store.commit("SuggestList/buildings", response.data)
        )
        .catch(error => console.error(error))
        .finally(() => (this.isLoading = false))
    },
  },
  // 算出プロパティ
  computed: {
    // showSidebar は Vuex 上で状態管理をしているため、
    // Vue の算出プロパティ機能を使ってゲッター／セッターを用意する
    // Vuex インスタンスへは this.$store でアクセスできる
    showSidebar: {
      get() {
        // Vuex 上で管理しているデータには $store.getters でアクセスできる
        // $store.getters['対象データへのパス']
        return this.$store.getters["Sidebar/isVisible"]
      },
      set(value) {
        // Vuex 上で管理しているデータに関する変更は $store.dispatch で行える
        // dispatch は非同期処理なので、データ変更を「予約」する形になる
        // $store.dispatch('対象データへのパス', '変更後の値')
        this.$store.dispatch("Sidebar/isVisible", value)
      },
    },
  },
}
</script>
