<template>
  <div>
    <!-- Leaflet.js マップ -->
    <l-map
      :center="center"
      :options="{ zoomControl: false }"
      :zoom="zoom"
      @update:bounds="onUpdateBounds"
      @update:center="onUpdateCenter"
      @update:zoom="onUpdateZoom"
      class="l-map"
      ref="map"
    >
      <!-- Leaflet.js タイルレイヤー -->
      <l-tile-layer
        url="https://cyberjapandata.gsi.go.jp/xyz/pale/{z}/{x}/{y}.png"
      ></l-tile-layer>
      <!-- マーカー -->
      <Pin v-bind="pin" :key="idx" v-for="(pin, idx) in pins" />
    </l-map>
  </div>
</template>

<script lang="ts">
import Vue from "vue"

// Module: vue2-leaflet
import { LMap, LTileLayer } from "vue2-leaflet"
import "leaflet/dist/leaflet.css"

import { Bounds } from "@/entities/Map"
import { LatLng } from "@/entities/Common"
import { Pin as PinObject } from "@/entities/Pin"

import Pin from "@/components/Map/Pin.vue"

export default Vue.extend({
  // 使用するコンポーネント
  components: {
    LMap,
    LTileLayer,
    Pin,
  },
  data() {
    return {}
  },
  computed: {
    zoom: {
      get(): number {
        return this.$accessor.Map.zoom
      },
      set(value: number) {
        this.$accessor.Map.setZoom(value)
      },
    },
    center: {
      get(): LatLng {
        return this.$accessor.Map.center
      },
      set(value: LatLng) {
        this.$accessor.Map.setCenter(value)
      },
    },
    bounds: {
      get(): Bounds {
        return this.$accessor.Map.bounds
      },
      set(value: Bounds) {
        this.$accessor.Map.setBounds(value)
      },
    },
    pins: {
      get(): Array<PinObject> {
        return this.$accessor.Map.pins
      },
      set(value: Array<PinObject>) {
        this.$accessor.Map.setPins(value)
      },
    },
    map(): LMap {
      return this.$refs.map as LMap
    },
  },
  methods: {
    // ズームスケールが変更されたとき
    onUpdateZoom(zoom: number) {
      this.zoom = zoom
    },

    // 中心座標が変更されたとき
    onUpdateCenter(center: LatLng) {
      this.center = center
    },
    onUpdateBounds(bounds: Bounds) {
      this.bounds = bounds
    },
  },
  // このコンポーネントがマウントされたときに実行される処理
  mounted() {
    this.$nextTick(function() {
      // 初期位置・ズームの設定
      const bounds = this.map.mapObject.getBounds()
      this.bounds = {
        _southWest: {
          lat: bounds.getSouthWest().lat,
          lng: bounds.getSouthWest().lng,
        },
        _northEast: {
          lat: bounds.getNorthEast().lat,
          lng: bounds.getNorthEast().lng,
        },
      }
    })
  },
})
</script>

<style lang="scss" scoped>
.l-map {
  position: fixed;
  font: inherit;
}
</style>
