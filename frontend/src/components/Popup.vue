<template>
  <l-popup :lat-lng="latLong"> Hello! </l-popup>
</template>

<script>
import { LPopup } from "vue2-leaflet";

export default {
  components: {
    LPopup
  },
  props: {
    latLng: {
      type: Object,
      required: true
    }
  },
  data() {
    return {};
  }
};
</script>
