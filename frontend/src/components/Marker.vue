<template>
  <l-marker :lat-lng="data" v-on:click="onClick" @l-add="$event.target.openPopup()">
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
      isVisible: false
    };
  },
  methods: {
    onClick: function() {
      this.isVisible = true;
      console.log(this.data.name);
      this.$parent.$parent.onClickStationList(this.data);
    }
  }
};
</script>
