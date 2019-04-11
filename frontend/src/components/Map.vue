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
      <TMarker v-for="m in markers" :key="m.id" :data="m"/>
    </l-map>
    <!-- v-autocompleteはバグ(仕様)が多くて使いづらかったため没 -->
    <!-- v-toolbarのfloatingは表示が崩れて難しい -->
    <!-- <v-toolbar class="float-toolbar" dense floating extended height="200px"> -->
      <!-- <v-btn icon @click="searchStation">
        <v-icon>search</v-icon>
      </v-btn>
      <v-text-field clearable label="駅名を入力" single-line ></v-text-field> -->
      <!-- <v-autocomplete ref="autoComplete" label="駅名を入力" :single-line="true" :items="stationList" append-icon="" clearable :search-input.sync="searchText" no-data-text="検索結果なし"></v-autocomplete> -->
      <!-- <v-btn icon @click="getCurrentRect">
        <v-icon>my_location</v-icon>
      </v-btn> -->
      <!-- <v-list>
        <v-list-tile
          v-for="l in stationList"
          :key="l.id"
          @click="stationListClicked">
          <v-list-tile-content>
            <v-list-tile-title v-text="l.name"></v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-toolbar> -->
    <v-layout align-start justify-start row/> 
    <br>
    <v-flex xs4 offset-xs1 sm3 offset-sm1 md2 offset-md1>
      <v-card>
        <v-toolbar>
          <!-- <v-toolbar-side-icon></v-toolbar-side-icon> -->
          <v-btn icon @click="searchStation">
            <v-icon>search</v-icon>
          </v-btn>
          <!-- <v-toolbar-title>Inbox</v-toolbar-title> -->
          <v-text-field clearable label="駅名を入力" single-line v-model="textField" @keyup.enter="searchStation" ></v-text-field>
          <v-btn icon @click="getCurrentRect">
            <v-icon>my_location</v-icon>
          </v-btn>
        </v-toolbar>
        <v-list v-show="hasResult">
          <!-- array.slice はバックエンドでやるべき -->
          <v-list-tile
            v-for="l in stationList.slice(0, 5)"
            :key="l.id"
            @click="stationListClicked(l.id)">
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
      stationList: [
        {name: "東京", railwayName: "東海道本線", id: 1},
        {name: "大阪", railwayName: "東海道本線", id: 2},
        {name: "名古屋", railwayName: "東海道本線", id: 3},
      ],
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
            lng: elem.longitude,
            name: elem.name,
            companyName: [elem.company], // <- Arrayで帰ってくるのであれば不要
            railwayName: [elem.railwayName], // <- 同上
            id: elem.id
          }));
        })
        .catch(err => {
          console.log(`Axios ERROR!: ${err}`);
        });
    },
    // DBからstrをキーワードに駅検索
    getStationList(str) {
      axios.get(`http://${window.location.hostname}:1323/stations/suggest?keyword=` + str)
        .then(resp => {
          console.log(resp.data);
      })
    },
    // 駅名で検索・移動
    // 駅検索
    searchStation() {
      console.log("search:" + this.textField);
      // TODO: 別の関数を呼び出すとReferenceErrorる
      // getStationList(this.textField);
      axios.get(`http://${window.location.hostname}:1323/stations/suggest?keyword=` + this.textField)
        .then(resp => {
          this.markers = resp.data.map(elem => ({
            lat: elem.latitude,
            lng: elem.longitude,
            name: elem.name,
            companyName: [elem.company], // <- Arrayで帰ってくるのであれば不要
            railwayName: [elem.railwayName], // <- 同上
            id: elem.id }));
        // 駅名が完全一致でなければ表示をしない
        for(var n=resp.data.length-1; n>=0; n--) {
          if(this.markers[n].name !== this.textField) {
            this.markers.splice(n, 1);
          }
        }
        // もし完全一致する駅が存在すれば検索結果の1つ目の駅にフォーカス
        if(this.markers.length >= 1){
          this.$refs.mainMap.mapObject.panTo([this.markers[0].lat,this.markers[0].lng]);
        }
      })
    },
    stationListClicked(val) {
      axios.get(`http://${window.location.hostname}:1323/stations/` + val)
          .then(resp => {
            resp.data = resp.data;
            this.markers = resp.data.map(elem => ({
              lat: elem.latitude,
              lng: elem.longitude,
              name: elem.name,
              companyName: [elem.company], // <- Arrayで帰ってくるのであれば不要
              railwayName: [elem.railwayName], // <- 同上
              id: elem.id }));

            // 駅をフォーカスして表示
            this.$refs.mainMap.mapObject.panTo([this.markers[0].lat,this.markers[0].lng]);
          })
    }
  },
  mounted: function () {
  this.$nextTick(function () {
    // 初期位置・ズームの設定
    this.bounds = this.$refs.mainMap.mapObject.getBounds();
  })},
  watch: {
    textField(str) {
      // console.log("textField :" + str);
      // 何も入力されてなければリストは非表示
      if(isEmpty(str)) {
        this.hasResult = false;
      } 
      // 入力された文字列が駅名に一致すればリストに表示
      // 文字列が入力されていて一致がなければ最後のリストを表示し続ける
      else {
          axios.get(`http://${window.location.hostname}:1323/stations/suggest?keyword=` + this.textField)
          .then(resp => {
            var list = resp.data.map(elem => ({
            lat: elem.latitude,
            lng: elem.longitude,
            name: elem.name,
            companyName: elem.company,
            railwayName: elem.railwayName,
            id: elem.id }));
            if (list.length >= 1) {
              this.stationList = list;
              this.hasResult = true;
            }
          })
      }
    }
  }
};
function isEmpty(val){return !val?!(val===0||val===false)?true:false:false;}
</script>
