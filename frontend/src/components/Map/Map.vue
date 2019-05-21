<template>
  <div>
    <l-map
      class="l-map"
      ref="mainMap"
      :options="{ zoomControl: false }"
      :zoom="zoom"
      :center="center"
      @update:zoom="onUpdateZoom"
      @update:center="onUpdateCenter"
      @update:bounds="onUpdateBounds"
    >
      <l-tile-layer :url="tileMapUrl"></l-tile-layer>
      <Marker v-for="m in markerList" :key="m.id" :data="marker"/>
    </l-map>
  </div>
</template>

<script>
// Module: vue2-leaflet
import { LMap, LTileLayer } from "vue2-leaflet";
import Marker from "./Marker";
import "leaflet/dist/leaflet.css";

export default {
  components: {
    LMap,
    LTileLayer,
    Marker
  },
  data() {
    return {
      // tileMapUrl: Leaflet.js のタイルマップのURL
      tileMapUrl: "https://cyberjapandata.gsi.go.jp/xyz/std/{z}/{x}/{y}.png", // 地理院地図
      // tileMapUrl: "http://{s}.tile.osm.org/{z}/{x}/{y}.png",                   // OpenStreetMap

      // zoom: Leaflet.js Map のズームスケール
      zoom: 14,
      // center: Leaflet.js Map の中心座標
      center: {
        lat: 35.680446,
        lng: 139.761801
      },
      // bounds: Leaflet.js Mapの表示範囲
      bounds: {
        // 左下の座標
        _southWest: {
          lat: 35.63532680480169,
          lng: 139.73910056054595
        },
        // 右上の座標
        _northEast: {
          lat: 35.691113860493594,
          lng: 139.79489050805572
        }
      },
      // markerList: 地図上にプロットされるマーカーのリスト
      markerList: []
    };
  },
  methods: {
    // ズームスケールが変更されたとき
    onUpdateZoom(zoom) {
      this.zoom = zoom;
    },

    // 中心座標が変更されたとき
    onUpdateCenter(center) {
      this.center = center;
    },

    // 表示範囲が変更されたとき
    onUpdateBounds(bounds) {
      this.bounds = bounds;
    }
  },

  // このコンポーネントがマウントされたときに実行される処理
  mounted: function() {
    this.$nextTick(function() {
      // 初期位置・ズームの設定
      this.bounds = this.$refs.mainMap.mapObject.getBounds();
    });
  }
};
</script>
