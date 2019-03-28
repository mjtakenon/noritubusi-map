<template>
  <l-marker :lat-lng="data" v-on:click="onClick" @l-add="$event.target.openPopup()">
    <l-popup> <H1> {{ this.data.id }} : {{ this.data.name }} <br>
     {{ this.data.companyName }} : {{ this.data.railwayName }} </H1> </l-popup>
    <!-- <Tpopup></Tpopup> -->
  </l-marker>
  
</template>

<script>
import { LMarker,LPopup } from "vue2-leaflet";
import TPopup from "./Popup";

export default {
  components: {
    LMarker,
    LPopup,
    // TPopup
  },
  props: {
    data: {
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
