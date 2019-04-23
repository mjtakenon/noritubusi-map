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
      <TMarker v-for="m in markers" :key="m.id" :data="m"/>
    </l-map>
    <v-layout align-start justify-start row/>
    <br>
    <v-flex xs4 offset-xs1 sm3 offset-sm1 md2 offset-md1>
      <v-card>
        <v-toolbar>
          <!-- <v-btn icon @click="searchStation"> -->
          <v-btn icon>
            <v-icon>search</v-icon>
          </v-btn>
          <v-text-field
            clearable
            label="駅名を入力"
            single-line
            v-model="textField"
            @keyup.enter="searchStation"
          ></v-text-field>
          <v-btn icon @click="getCurrentRect">
            <v-icon>my_location</v-icon>
          </v-btn>
        </v-toolbar>
        <v-list v-show="stationList.length">
          <!-- array.slice はバックエンドでやるべき -->
          <v-list-tile
            v-for="l in stationList.slice(0, 5)"
            :key="l.id"
            @click="onClickStationList(l.id)"
          >
            <v-list-tile-content>
              <v-list-tile-title v-text="l.name"></v-list-tile-title>
              <v-list-tile-sub-title v-text="l.railwayName"></v-list-tile-sub-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-card>
    </v-flex>
  </div>
</template>

<script>
// Module: vue2-leaflet
import { LMap, LTileLayer } from "vue2-leaflet";
import TMarker from "./Marker";
import "leaflet/dist/leaflet.css";

import axios from "axios";

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
      markers: [],
      stationList: [],
      textField: "",
      hasResult: false
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
        .get(`http://${window.location.hostname}:1323/buildings`, {
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
            lng: elem.longitude,
            name: elem.name,
            id: elem.id
          }));
        })
        .catch(err => {
          console.log(`Axios ERROR!: ${err}`);
        });
    },

    getStationListByKeyword(keyword) {
      console.log(`Keyword: ${keyword}`);

      axios.get(`http://${window.location.hostname}:1323/stations/suggest?keyword=${keyword}`).then(resp => {
        let stationList = Array();

        stationList = resp.data.map(elem => ({
          lat: elem.latitude,
          lng: elem.longitude,
          stationName: elem.name,
          railwayName: elem.railwayName,
          orderInRailway: elem.orderInRailway,
          stationId: elem.station_id,
          buildingId: elem.building_id
        }));

        return stationList;
      });
    },

    getStationById(stationId) {
      let stationInfo;

      axios.get(`http://${window.location.hostname}:1323/stations/${stationId}`).then(resp => {
        stationInfo = resp.data.map(elem => ({
          name: elem.name,
          lat: elem.latitude,
          lng: elem.longitude,
          railwayName: elem.railwayName,
          orderInRailway: elem.orderInRailway
        }));
      });

      return stationInfo;
    },

    forcusToStation(stationInfo) {
      this.$refs.mainMap.mapObject.panTo([stationInfo.lat, stationInfo.lng]);
    },

    checkCompleteMatchAndForcus(keyword) {
      const matchedToKeywordCompletely = this.stationList.filter(elem => elem.stationName == keyword);

      if (matchedToKeywordCompletely.length > 0) {
        this.forcusToStation(matchedToKeywordCompletely[0]);
      }
    },

    // DBからstrをキーワードに駅検索
    // 駅名で検索・移動
    // 駅検索

    searchStation() {
      console.log("::Called:: searchStation");
      let keyword = this.textField;

      this.stationList = this.getStationListByKeyword(keyword);

      // もし完全一致する駅が存在すれば検索結果の1つ目の駅にフォーカス
      this.checkCompleteMatchAndForcus(keyword);
    },

    onClickStationList(stationId) {
      console.log("::Called:: onClickStationList");
      const stationInfo = this.getStationById(stationId);

      this.forcusToStation(stationInfo);
    }
  },
  mounted: function() {
    this.$nextTick(function() {
      // 初期位置・ズームの設定
      this.bounds = this.$refs.mainMap.mapObject.getBounds();
    });
  },
  watch: {
    textField(str) {
      console.log("::Called:: textField");
      // 何も入力されてなければリストは非表示
      if (isEmpty(str)) {
        this.hasResult = false;
      }
      // 入力された文字列が駅名に一致すればリストに表示
      // 文字列が入力されていて一致がなければ最後のリストを表示し続ける
      else {
        const stationList = this.getStationListByKeyword(this.textField);

        console.log(stationList);

        if (stationList.length >= 1) {
          this.stationList = stationList;
          this.hasResult = true;
        }
      }
    }
  }
};
function isEmpty(val) {
  return !val ? (!val === false ? true : false) : false;
}
</script>
