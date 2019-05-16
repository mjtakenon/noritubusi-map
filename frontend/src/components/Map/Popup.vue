<template>
  <l-popup>
    <H1>{{ this.data.name }}</H1>
    <v-list dense>
      <v-list-tile v-for="(l, idx) in this.data.lines" :key="idx" @click="companyNameClicked(idx)">
        <v-list-tile-avatar>
          <H2>{{icon}}</H2>
        </v-list-tile-avatar>
        <v-list-tile-content>
          <v-list-tile-title>{{ data.lines[idx].railway_name }}</v-list-tile-title>
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
    data: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      icon: ""
    };
  },
  methods: {
    companyNameClicked(val) {
      console.log(val);
    }
  },
  mounted: function() {
    this.$nextTick(function() {
      this.icon = "🚊";
    });
  }
};
</script>
