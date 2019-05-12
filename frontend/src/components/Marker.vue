<template>
  <l-marker :lat-lng="markerInfo" v-on:click="onClick" @l-add="$event.target.openPopup()">
    <l-icon :icon-anchor="markerInfo" :icon-size="[24, 24]" class-name="circle-icon">
      <!-- <v-icon color="#ff8c00">lens</v-icon> -->
      <svg viewBox="0 0 24 24">
        <circle cx="12" cy="12" r="10" stroke="#FFFFFF" stroke-width="2" :fill="fillColor"></circle>
      </svg>
    </l-icon>
    <TPopup :v-show="isVisible" :popupInfo="markerInfo"></TPopup>
  </l-marker>
</template>

<script>
import { LMarker, LIcon } from "vue2-leaflet";
import TPopup from "./Popup";

export default {
  components: {
    LMarker,
    LIcon,
    TPopup
  },
  props: {
    markerInfo: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      isVisible: false
    };
  },
  computed: {
    fillColor: function() {
      var nOfLines = this.markerInfo.lines.length;

      if (nOfLines >= 5) {
        return "red";
      } else if (nOfLines >= 2) {
        return "orange";
      } else {
        return "blue";
      }
    }
  },
  methods: {
    onClick: function() {
      this.isVisible = true;
    }
  }
};
</script>
