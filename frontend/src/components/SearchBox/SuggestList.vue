<template>
  <div>
    <v-list style="max-height: 200px; overflow-y: auto;">
      <v-list-item-group v-model="stations" color="primary">
        <v-list-item v-for="(station, i) in this.stations" :key="i">
          <v-list-item-icon>
            <v-icon>train</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title> {{ station.name }} </v-list-item-title>
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
  methods: {},
  // 算出プロパティ
  computed: {}
};
</script>
