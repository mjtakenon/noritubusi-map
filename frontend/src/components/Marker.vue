<template>
  <l-marker :lat-lng="latLong"></l-marker>
</template>

<script>
import { LMarker } from "vue2-leaflet";

export default {
  components: {
    LMarker
  },
  props: {
    latLong: {
      type: Object,
      required: true
    }
  },
  data() {
    return {};
  }
};
</script>
