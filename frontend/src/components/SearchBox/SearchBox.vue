<template>
  <div v-bind:style="style.searchBox">
    <v-toolbar dense>
      <v-btn @click="toggleSidebar" icon>
        <v-icon>fa-bars</v-icon>
      </v-btn>

      <v-text-field
        hide-details
        append-icon="search"
        single-line
        v-bind:style="style.textField"
        @input="suggestBuildings"
      ></v-text-field>

      <!-- <v-btn icon>
        <v-icon>my_location</v-icon>
      </v-btn> -->
    </v-toolbar>
    <SuggestList
      v-show="buildings.length !== 0"
      :buildings="buildings"
    ></SuggestList>
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
      buildings: [],
    }
  },
  // メソッド
  methods: {
    toggleSidebar() {
      this.showSidebar = !this.showSidebar
    },

    suggestBuildings(input) {
      if (input.length === 0) {
        this.buildings = []
        return
      }

      suggest(input)
        .then(response => (this.buildings = response.data))
        .catch(error => console.error(error))
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
