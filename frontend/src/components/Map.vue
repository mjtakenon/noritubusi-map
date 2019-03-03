<template>
  <div>
    <l-map
      class="l-map"
      ref="mainMap"
      :options="{zoomControl: false}"
      :zoom="zoom"
      :center="center"
      @update:zoom="zoomUpdated"
      @update:center="centerUpdated"
      @update:bounds="boundsUpdated"
    >
      <l-tile-layer :url="url"></l-tile-layer>
    </l-map>
    <v-toolbar class="float-toolbar" dense floating>
      <v-btn icon>
        <v-icon>search</v-icon>
      </v-btn>
      <v-text-field clearable label="駅名を入力" single-line></v-text-field>
      <v-btn icon>
        <v-icon>my_location</v-icon>
      </v-btn>
    </v-toolbar>
  </div>
</template>

<script>
// Module: vue2-leaflet
import { LMap, LTileLayer, LMarker } from "vue2-leaflet";
import "leaflet/dist/leaflet.css";

export default {
  components: {
    LMap,
    LTileLayer,
    LMarker
  },
  data() {
    return {
      url: "http://{s}.tile.osm.org/{z}/{x}/{y}.png",
      zoom: 14,
      center: [35.680446, 139.761801],
      bounds: {
        _southWest: {
          lat: 35.63532680480169,
          lng: 139.73910056054595
        },
        _northEast: {
          lat: 35.691113860493594,
          lng: 139.79489050805572
        }
      }
    };
  },
  methods: {
    zoomUpdated(zoom) {
      this.zoom = zoom;
    },
    centerUpdated(center) {
      this.center = center;
    },
    boundsUpdated(bounds) {
      this.bounds = bounds;
    }
  }
};
</script>
