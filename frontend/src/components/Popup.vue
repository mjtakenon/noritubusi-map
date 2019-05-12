<template>
  <l-popup>
    <H1 class="popup">{{ this.popupInfo.name }}</H1>
    <v-list dense>
      <v-list-tile
        v-for="(line, idx) in this.popupInfo.lines"
        :key="idx"
        @click="companyNameClicked(idx)"
      >
        <v-list-tile-avatar>
          <H2>{{listIcon}}</H2>
        </v-list-tile-avatar>
        <v-list-tile-content>
          <v-list-tile-title>{{ line.railway_name }}</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
    </v-list>
  </l-popup>
</template>

<script>
import { LPopup } from "vue2-leaflet";

export default {
  components: {
    LPopup
  },
  props: {
    popupInfo: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      listIcon: ""
    };
  },
  methods: {
    companyNameClicked(val) {
      console.log(val);
    }
  },
  mounted: function() {
    this.$nextTick(function() {
      this.listIcon = "🚊";
    });
  }
};
</script>
