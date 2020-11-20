<template>
  <l-marker
    ref="marker"
    :lat-lng="latLng"
    @add="$nextTick(() => updateMap())"
  >
    <Popup v-bind="popup"></Popup>
  </l-marker>
</template>

<script>
import { LMarker } from "vue2-leaflet"
import Popup from "./Popup"

export default {
  components: {
    LMarker,
    Popup,
  },
  props: {
    latLng: {
      type: Array,
      required: true,
      validator(latLng) {
        // 配列長が2(緯度, 軽度) かつ 全要素が Number であること
        return latLng.length === 2 && latLng.every(elem => !isNaN(elem))
      },
    },
    popup: {
      type: Object,
      required: true,
    },
    openPopup: {
      type: Boolean,
      required: false,
      default: false,
    },
    focusPin: {
      type: Boolean,
      required: false,
      default: false,
    },
  },
  data() {
    return {}
  },
  methods: {
    updateMap() {
      if (this.openPopup) {
        this.$refs.marker.mapObject.openPopup()
      }
      if (this.focusPin) {
        this.focusAtPin()
      }
    },
    focusAtPin() {
      this.$store.dispatch("Map/updateCenter", {
        lat: this.latLng[0],
        lng: this.latLng[1],
      })
    },
  },
  updated() {
    this.updateMap()
  },
}
</script>
