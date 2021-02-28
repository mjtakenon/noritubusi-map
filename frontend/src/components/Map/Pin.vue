<template>
  <l-marker ref="marker" :lat-lng="latLng" @add="$nextTick(() => updateMap())">
    <Popup v-bind="popup"></Popup>
  </l-marker>
</template>

<script lang="ts">
import Vue, { PropType } from "vue"

import { LMarker } from "vue2-leaflet"

import { LatLng } from "@/entities/Common"
import { Popup as PopupObject } from "@/entities/Pin"

import Popup from "@/components/Map/Popup.vue"

export default Vue.extend({
  components: {
    LMarker,
    Popup,
  },
  props: {
    latLng: {
      type: Object as PropType<LatLng>,
      required: true,
    },
    popup: {
      type: Object as PropType<PopupObject>,
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
  computed: {
    marker(): LMarker {
      return this.$refs.marker as LMarker
    },
  },
  methods: {
    updateMap() {
      if (this.openPopup) {
        this.marker.mapObject.openPopup()
      } else {
        this.marker.mapObject.closePopup()
      }
      if (this.focusPin) {
        this.focusAtPin()
      }
    },
    focusAtPin() {
      this.$accessor.Map.setCenter(this.latLng)
    },
  },
  updated() {
    if (!this.openPopup) {
      this.marker.mapObject.closePopup()
    }
    this.updateMap()
  },
})
</script>
