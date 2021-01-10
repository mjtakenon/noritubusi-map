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
        v-bind:style="style.textField"
        @input="onInput"
        @click:clear="resetStation"
        :value="inputStation"
      ></v-text-field>
      <v-btn @click="toggleSidebar" icon>
        <v-icon>swap_horiz</v-icon>
      </v-btn>
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
        textField: {
          // 上に謎のスペースがあるから消す
          padding: "0px",
        },
      },
      inputStation: "",
    }
  },
  // メソッド
  methods: {
    toggleSidebar() {
      this.showSidebar = !this.showSidebar
    },

    updateSuggest() {
      if (this.inputStation === null || this.inputStation.length === 0) {
        this.$store.commit("SuggestList/buildings", [])
        return
      }
      suggest(this.inputStation)
        .then(response =>
          this.$store.commit("SuggestList/buildings", response.data)
        )
        .catch(error => console.error(error))
    },

    onInput(input) {
      this.inputStation = input
      this.updateSuggest()
    },

    resetStation() {
      this.inputStation = ""
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
