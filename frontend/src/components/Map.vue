<template>
  <div>
    <l-map
      class="l-map"
      ref="mainMap"
      :options="{ zoomControl: false }"
      :zoom="zoom"
      :center="center"
      @update:zoom="zoomUpdated"
      @update:center="centerUpdated"
      @update:bounds="boundsUpdated"
    >
      <l-tile-layer :url="url"></l-tile-layer>
      <!-- <l-marker v-for="(m, i) in markers" :key="i" :lat-lng="m.latlong"></l-marker> -->
      <TMarker v-for="(m, i) in markers" :key="i" :lat-long="m" />
    </l-map>
    <v-toolbar class="float-toolbar" dense floating>
      <v-btn icon>
        <v-icon>search</v-icon>
      </v-btn>
      <v-text-field clearable label="駅名を入力" single-line></v-text-field>
      <v-btn icon @click="getCurrentRect">
        <v-icon>my_location</v-icon>
      </v-btn>
    </v-toolbar>
  </div>
</template>

<script>
// Module: vue2-leaflet
import { LMap, LTileLayer } from "vue2-leaflet";
import TMarker from "./Marker";
import "leaflet/dist/leaflet.css";

export default {
  components: {
    LMap,
    LTileLayer,
    TMarker
  },
  data() {
    return {
      // url: "http://{s}.tile.osm.org/{z}/{x}/{y}.png",
      url: "https://cyberjapandata.gsi.go.jp/xyz/std/{z}/{x}/{y}.png",
      zoom: 14,
      center: {
        lat: 35.680446,
        lng: 139.761801
      },
      bounds: {
        _southWest: {
          lat: 35.63532680480169,
          lng: 139.73910056054595
        },
        _northEast: {
          lat: 35.691113860493594,
          lng: 139.79489050805572
        }
      },
      markers: []
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
    },
    getCurrentRect() {
      axios
        .get(`http://${window.location.hostname}:1323/stations`, {
          params: {
            begin_latitude: this.bounds._northEast.lat,
            begin_longitude: this.bounds._northEast.lng,
            end_latitude: this.bounds._southWest.lat,
            end_longitude: this.bounds._southWest.lng
          }
        })
        .then(resp => {
          console.log("Axios SUCCESS!");
          console.log(`response: ${resp}`);
          console.log(`status: ${resp.status}`);
          console.log(resp.data);
          this.markers = resp.data.map(elem => ({
            lat: elem.latitude,
            lng: elem.longitude
          }));
        })
        .catch(err => {
          console.log(`Axios ERROR!: ${err}`);
        });
    }
  }
};
</script>
