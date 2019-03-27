<template>
  <l-marker :lat-lng="latLong" v-on:click="onClick" @l-add="$event.target.openPopup()">
    <l-popup> {{ this.latLong.name }} </l-popup>
  </l-marker>
  
</template>

<script>
import { LMarker, LPopup } from "vue2-leaflet";

export default {
  components: {
    LMarker,
    LPopup
  },
  props: {
    latLong: {
      type: Object,
      required: true
    }
  },
  data() {
    return {};
  },
  methods: {
    onClick : function() {
      console.log(this)
    }
  }
};

</script>
