<template>
  <div>
    <v-list two-line style="max-height: 200px; overflow-y: auto;">
      <v-list-item-group color="primary">
        <v-list-item v-for="(station, i) in this.stations" :key="i" @click="focusStation(station)">
          <v-list-item-avatar style="font-size: 150%">
            <div v-if="isShinkansen(station.railway_line_name)">🚅</div>
            <div v-else>🚃</div>
          </v-list-item-avatar>
          <v-list-item-content>
            <v-list-item-title> {{ station.name }} </v-list-item-title>
            <v-list-item-subtitle> {{ station.railway_line_name }} </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </v-list-item-group>
    </v-list>
  </div>
</template>

<script>
import { suggest } from "../../utils/api/search.js";

export default {
  props: {
    stations: {
      type: Array,
      required: true,
    },
  },
  data() {
    return {
    }
  },
  // メソッド
  // {
  //   "building_id": 2103,
  //   "station_id": 4635,
  //   "name": "東京",
  //   "latitude": "35.681391",
  //   "longitude": "139.766103",
  //   "railway_line_name": "JR総武本線",
  //   "order_in_railway": 1
  // }
  methods: {
    isShinkansen(railwayName) {
      return railwayName.indexOf('新幹線') != -1
    },
    focusStation(station) {
      this.$store.dispatch("Map/updateCenter", {
        lat: station.latitude,
        lng: station.longitude,
      })
      return 
    }
  },
  // 算出プロパティ
  computed: {}
};
</script>
