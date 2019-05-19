<template>
  <v-app>
    <!-- まずは、ユーザー管理機能があるサイドバーからリファクタリング -->
    <Sidebar/>
    <!-- サイドバーを開けるためのボタン(仮) -->
    <v-btn
      @click="onClickBtn"
      color="warning"
      right
      style="position: absolute;"
    >Test</v-btn>
  </v-app>
</template>

<script>
import Sidebar from "./components/Sidebar/Sidebar"

export default {
  // 使用するコンポーネントを宣言
  components: {
    Sidebar,
  },
  // データ
  data() {
    return {}
  },
  // メソッド
  methods: {
    onClickBtn() {
      this.showSidebar = !this.showSidebar
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
