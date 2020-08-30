<template>
  <l-marker
    ref="marker"
    :lat-lng="latLng"
    @add="$nextTick(() => focusBuildingAndOpenPopup())"
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
    autoOpenPopup: {
      type: Boolean,
      required: false,
      default: false,
    },
  },
  data() {
    return {}
  },
  methods: {
    focusBuildingAndOpenPopup() {
      if (this.autoOpenPopup) {
        this.$refs.marker.mapObject.openPopup()
      }
      this.focusBuilding()
    },
    focusBuilding() {
      this.$store.dispatch("Map/updateCenter", {
        lat: this.latLng[0],
        lng: this.latLng[1],
      })
    },
  },
  updated() {
    this.focusBuildingAndOpenPopup()
  },
}
</script>
