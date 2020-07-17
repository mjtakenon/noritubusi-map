<template>
  <div>
    <v-list two-line style="max-height: 200px; overflow-y: auto;">
      <v-list-item-group v-model="stations" color="primary">
        <v-list-item v-for="(station, i) in this.stations" :key="i">
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
  methods: {
    isShinkansen(railwayName) {
      return railwayName.indexOf('新幹線') != -1
    }
  },
  // 算出プロパティ
  computed: {}
};
</script>
