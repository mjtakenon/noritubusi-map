<template>
  <div>
    <!-- Leaflet.js マップ -->
    <l-map
      :center="center"
      :options="mapOptions"
      :zoom="zoom"
      @update:bounds="onUpdateBounds"
      @update:center="onUpdateCenter"
      @update:zoom="onUpdateZoom"
      class="l-map"
      ref="mainMap"
    >
      <!-- Leaflet.js タイルレイヤー -->
      <l-tile-layer :url="tileMapUrl"></l-tile-layer>
      <!-- マーカー -->
      <Pin v-bind="pin" :key="idx" v-for="(pin, idx) in pins" />
    </l-map>
  </div>
</template>

<script>
// Module: vue2-leaflet
import { LMap, LTileLayer } from "vue2-leaflet"
import Pin from "./Pin"
import "leaflet/dist/leaflet.css"

export default {
  // 使用するコンポーネント
  components: {
    LMap,
    LTileLayer,
    Pin,
  },
  // データ
  data() {
    return {
      // mapOptions: Leaflet.js Map のオプション
      mapOptions: {
        // ズーム操作用の「＋/−」ボタンは非表示に
        zoomControl: false,
      },

      // tileMapUrl: Leaflet.js のタイルマップのURL
      tileMapUrl: "https://cyberjapandata.gsi.go.jp/xyz/pale/{z}/{x}/{y}.png", // 地理院地図
      // tileMapUrl: "http://{s}.tile.osm.org/{z}/{x}/{y}.png",                   // OpenStreetMap
    }
  },
  // 算出プロパティ
  computed: {
    // [Vuex] zoom: マップのズーム率
    zoom: {
      get() {
        return this.$store.getters["Map/zoom"]
      },
      set(value) {
        this.$store.dispatch("Map/updateZoom", value)
      },
    },
    // [Vuex] center: マップの中心座標
    center: {
      get() {
        return this.$store.getters["Map/center"]
      },
      set(value) {
        this.$store.dispatch("Map/updateCenter", value)
      },
    },
    // [Vuex] bounds: マップの矩形座標(左上, 右下)
    bounds: {
      get() {
        return this.$store.getters["Map/bounds"]
      },
      set(value) {
        this.$store.dispatch("Map/updateBounds", value)
      },
    },
    // [Vuex] pins: マップ上にプロットされるマーカー
    pins: {
      get() {
        return this.$store.getters["Map/pins"]
      },
      set(value) {
        return this.$store.dispatch("Map/pins", value)
      },
    },
  },
  methods: {
    // ズームスケールが変更されたとき
    onUpdateZoom(zoom) {
      this.zoom = zoom
    },

    // 中心座標が変更されたとき
    onUpdateCenter(center) {
      this.center = center
    },

    // 表示範囲が変更されたとき
    onUpdateBounds(bounds) {
      this.bounds = bounds
    },
  },
  // このコンポーネントがマウントされたときに実行される処理
  mounted() {
    this.$nextTick(function() {
      // 初期位置・ズームの設定
      this.bounds = this.$refs.mainMap.mapObject.getBounds()
    })
  },
}
</script>

<style lang="scss" scoped>
.l-map {
  position: fixed;
}
</style>
