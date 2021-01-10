<template>
  <div v-bind:style="style.searchBox">
    <v-toolbar dense extended>
      <v-btn @click="toggleSidebar" icon>
        <v-icon>menu</v-icon>
      </v-btn>

      <!-- テキストボックス(乗車駅) -->
      <v-text-field
        hide-details
        append-icon="search"
        single-line
        :readonly="isStationFromConfirmed"
        clearable
        v-bind:style="style.textField"
        @input="suggestFromBuildings"
        @click:clear="resetStationFrom"
        :value="stationFrom.name || inputStationFrom"
      ></v-text-field>

      <template v-slot:extension>
        <v-btn @click="toggleSidebar" icon>
          <v-icon>import_export</v-icon>
        </v-btn>
        <!-- テキストボックス(降車駅) -->
        <v-text-field
          hide-details
          append-icon="search"
          single-line
          clearable
          :readonly="isStationToConfirmed"
          v-bind:style="style.textField"
          @input="suggestToBuildings"
          @click:clear="resetStationTo"
          :value="stationTo.name || inputStationTo"
        ></v-text-field>
      </template>

      <!-- <v-btn icon>
        <v-icon>my_location</v-icon>
      </v-btn> -->
    </v-toolbar>
    <SuggestList></SuggestList>
  </div>
</template>

<script>
import { suggest } from "../../utils/api/search.js"

import SuggestList from "./SuggestList"

export default {
  components: {
    SuggestList,
  },
  data() {
    return {
      style: {
        searchBox: {
          // 適度にmarginを取ってfloatingさせる
          position: "absolute",
          padding: "10px",
        },
        textField: {
          // 上に謎のスペースがあるから消す
          padding: "0px",
        },
      },
      inputStationFrom: "",
      inputStationTo: "",
    }
  },
  // メソッド
  methods: {
    toggleSidebar() {
      this.showSidebar = !this.showSidebar
    },

    updateSuggest(stationName) {
      if (stationName.length === 0) {
        this.$store.commit("SuggestList/buildings", [])
        return
      }
      suggest(stationName)
        .then(response =>
          this.$store.commit("SuggestList/buildings", response.data)
        )
        .catch(error => console.error(error))
    },

    suggestFromBuildings(input) {
      this.inputStationFrom = input
      this.updateSuggest(input)
    },

    suggestToBuildings(input) {
      this.inputStationTo = input
      this.updateSuggest(input)
    },
    resetStationFrom() {
      this.$store.commit("TripRecord/resetStationFrom")
    },
    resetStationTo() {
      this.$store.commit("TripRecord/resetStationTo")
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
    stationFrom() {
      return this.$store.getters["TripRecord/stationFrom"]
    },
    stationTo() {
      return this.$store.getters["TripRecord/stationTo"]
    },
    isStationFromConfirmed() {
      return this.$store.getters["TripRecord/isStationFromConfirmed"]
    },
    isStationToConfirmed() {
      return this.$store.getters["TripRecord/isStationToConfirmed"]
    },
  },
}
</script>
