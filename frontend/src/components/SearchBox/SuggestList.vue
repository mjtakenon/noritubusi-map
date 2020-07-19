<template>
  <div>
    <v-list two-line style="max-height: 200px; overflow-y: auto;">
      <v-list-item-group color="primary">
        <v-list-item v-for="(building, i) in this.buildings" :key="i" @click="onClickSuggestedBuilding(building)">
          <v-list-item-avatar style="font-size: 150%">
            <v-icon>train</v-icon>
          </v-list-item-avatar>
          <v-list-item-content>
            <v-list-item-title> {{ building.name }} </v-list-item-title>
            <v-list-item-subtitle>
              {{ railwayNameWithTrailing(building) }}
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
  methods: {
    railwayNameWithTrailing(building) {
      return building.lines[0].railway_name + ((building.lines.length !== 1) ? "..." : "")
    },
    setPin(building) {
      this.$store.dispatch("Map/setPins", [
        { 
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
      return
    },
  },
  // 算出プロパティ
  computed: {}
};
</script>
