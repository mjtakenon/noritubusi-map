<template>
  <l-marker :lat-lng="data" v-on:click="onClick" @l-add="$event.target.openPopup()">
    <!-- <l-popup> <H1> {{ this.data.id }} : {{ this.data.name }} <br>{{ this.data.companyName }} : {{ this.data.railwayName }} </H1> </l-popup> -->
    <TPopup :v-show="isVisible" :data="data"></TPopup>
  </l-marker>
</template>

<script>
import { LMarker } from "vue2-leaflet";
import TPopup from "./Popup";

export default {
  components: {
    LMarker,
    TPopup
  },
  props: {
    data: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      isVisible: false,
    };
  },
  methods: {
    onClick: function() {
      console.log(this);
      // this.popup = new TPopup();
      this.isVisible = true;
    }
  },
};
</script>
