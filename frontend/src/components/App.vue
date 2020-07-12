<template>
  <v-app>
    <Map />
    <!-- まずは、ユーザー管理機能があるサイドバーからリファクタリング -->
    <Sidebar />
    <SearchBox />
    <!-- 位置情報を取得するためのボタン(仮) -->
    <v-btn
      @click="onClickBtn"
      absolute
      color="info"
      right
      style="margin-top: 12px;"
    >Test</v-btn>
  </v-app>
</template>

<script>
import Map from "./Map/Map"
import Sidebar from "./Sidebar/Sidebar"
import SearchBox from "./SearchBox/SearchBox"

export default {
  // 使用するコンポーネントを宣言
  components: {
    Sidebar,
    Map,
    SearchBox,
  },
  // データ
  data() {
    return {}
  },
  // メソッド
  methods: {
    toggleSidebar() {
      this.showSidebar = !this.showSidebar
    },
    onClickBtn() {
      this.$store.dispatch("Map/updateCenter", {
        lat: 34.9613557,
        lng: 137.1389376,
      })
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
