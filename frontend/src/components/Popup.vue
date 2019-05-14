<template>
  <l-popup>
    <!-- ID:{{ this.data.id }} -->
    <H1>{{ this.data.name }}</H1>
    <v-list dense>
      <v-list-tile
        v-for="(l, idx) in this.data.lines"
        :key="idx"
        @click="railwayNameClicked(l)"
      >
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
    railwayNameClicked(railwayInfo) {
      this.$parent.$parent.$parent.$parent.onClickRideRailwayList(railwayInfo);
    }
  },
  mounted: function() {
    this.$nextTick(function() {
      this.icon = "🚊";
      // アイコンの指定
      // if (this.data.railwayName[0].indexOf("新幹線") !== -1) {
      //   // 新幹線
      //   this.icon = "🚆";
      // } else {
      //   // それ以外
      //   this.icon = "🚊";
      // }
    });
  }
};
</script>
