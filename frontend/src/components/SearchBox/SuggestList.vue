<template>
  <div>
    <v-list two-line style="max-height: 200px; overflow-y: auto;">
      <v-list-item-group color="primary">
        <v-list-item v-for="(building, i) in this.buildings" :key="i" @click="onClickSuggestedBuilding(building)">
          <v-list-item-avatar style="font-size: 150%">
            <v-icon>train</v-icon>
            <!-- <div v-if="isShinkansen(building.railway_line_name)">🚅</div> -->
            <!-- <div v-else>🚃</div> -->
          </v-list-item-avatar>
          <v-list-item-content>
            <v-list-item-title> {{ building.name }} </v-list-item-title>
            <v-list-item-subtitle>
              {{ building.lines[0].railway_name + ((building.lines.length !== 1) ? "..." : "") }}
            </v-list-item-subtitle>
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
    buildings: {
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
    focusBuilding(station) {
      this.$store.dispatch("Map/updateCenter", {
        lat: station.latitude,
        lng: station.longitude,
      })
      return 
    },
    setPin(building) {
      this.$store.dispatch("Map/setPins", [
        { 
          // lat: building.latitude,
          // lng: building.longitude,
          latLng: [building.latitude, building.longitude],
          popup: {
            name: building.name,
            lines: building.lines,
          },
          autoOpenPopup: true,
        }
      ])
      return
    },
    onClickSuggestedBuilding(building) {
      this.setPin(building)
      this.focusBuilding(building)
      return
    },
  },
  // 算出プロパティ
  computed: {}
};
</script>
