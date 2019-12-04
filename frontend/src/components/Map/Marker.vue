<template>
  <l-marker
    :lat-lng="data"
    v-on:click="onClick"
    @l-add="$event.target.openPopup()"
  >
    <Popup :v-show="isVisible" :data="data"></Popup>
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
    data: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      isVisible: false,
    }
  },
  methods: {
    onClick: function() {
      this.isVisible = true
    },
  },
}
</script>
