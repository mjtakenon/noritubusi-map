<template>
  <l-marker
    ref="marker"
    :lat-lng="latLng"
    @add="$nextTick(() => $event.target.openPopup())"
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
        return latLng.length === 2
              && latLng.every(
                elem => !isNaN(elem)
              )
      }
    },
    popup: {
      type: Object,
      required: true,
    },
    autoOpenPopup: {
      type: Boolean,
      required: false,
      default: false,
    }
  },
  data() {
    return {}
  },
  mounted() {
    console.log(this.$el)
    // console.log(this.$refs.marker.mapObject)
    // if (this.autoOpenPopup) {
    //   this.$refs.marker.mapObject.togglePopup()
    // }
    // this.isVisible = this.autoOpenPopup
  },
  methods: {},
}
</script>
