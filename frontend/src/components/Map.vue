<template>
  <div>
    <l-map
      class="l-map"
      ref="mainMap"
      :options="{ zoomControl: false }"
      :zoom="zoom"
      :center="center"
      @update:zoom="onUpdateZoom"
      @update:center="onUpdateCenter"
      @update:bounds="onUpdateBounds"
    >
      <l-tile-layer :url="urlTileMap"></l-tile-layer>
      <TMarker v-for="marker in markerList" :key="m.id" :data="marker"/>
    </l-map>
<!-- 元のやつ -->
    <!-- <v-layout align-start justify-start row/>
    <v-flex xs4 offset-xs1 sm3 offset-sm1 md2 offset-md1>
      <v-card>
        <v-toolbar>
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
          <v-btn icon @click="onClickMyLocationIcon">
            <v-icon>my_location</v-icon>
          </v-btn>
        </v-toolbar>
        <v-list v-show="hasResult">
          <v-list-tile
            v-for="stationInfo in stationList.slice(0, 5)"
            :key="stationInfo.stationId"
            @click="onClickStationList(stationInfo)"
          >
            <v-list-tile-content>
              <v-list-tile-title v-text="stationInfo.stationName"></v-list-tile-title>
              <v-list-tile-sub-title v-text="stationInfo.railwayName"></v-list-tile-sub-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-card>
    </v-flex> -->

    <div class="pa-3">
      <v-card-text style="width: 320px; position: relative;">
        <v-toolbar absolute flat height="50px" style="border-radius:10px 10px 0px 0px;">
          <v-toolbar-side-icon></v-toolbar-side-icon>
          <v-text-field
            clearable
            label="乗車駅を入力"
            v-model="textField"
            @keyup.enter="searchStation"
          ></v-text-field>
          <v-btn icon>
            <v-icon>search</v-icon>
          </v-btn>
          <v-btn icon @click="onClickMyLocationIcon">
            <v-icon>my_location</v-icon>
          </v-btn>
        </v-toolbar>
      </v-card-text>
      <v-card-text style="width: 352px; position: relative; top: -20px; left:-16px; padding-top:30px;">
        <v-list subheader absolute avatar v-show="hasResult" style="background-color:#f5f5f5; border-radius:0px 0px 10px 10px;">
          <v-subheader>候補...</v-subheader>
          <v-list-tile
            v-for="stationInfo in stationList.slice(0, 5)"
            :key="stationInfo.stationId"
            @click="onClickStationList(stationInfo)" >
            <v-list-tile-avatar>
              <v-icon>🚉</v-icon>
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-list-tile-title v-text="stationInfo.stationName"></v-list-tile-title>
              <v-list-tile-sub-title v-text="stationInfo.railwayName"></v-list-tile-sub-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-card-text>
    </div>
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
      // urlTileMap: Leaflet.js のタイルマップのURL
      urlTileMap: "https://cyberjapandata.gsi.go.jp/xyz/std/{z}/{x}/{y}.png", // 地理院地図
      // urlTileMap: "http://{s}.tile.osm.org/{z}/{x}/{y}.png",                   // OpenStreetMap

      // zoom: Leaflet.js Map のズームスケール
      zoom: 14,
      // center: Leaflet.js Map の中心座標
      center: {
        lat: 35.680446,
        lng: 139.761801
      },
      // bounds: Leaflet.js Mapの表示範囲
      bounds: {
        // 左下の座標
        _southWest: {
          lat: 35.63532680480169,
          lng: 139.73910056054595
        },
        // 右上の座標
        _northEast: {
          lat: 35.691113860493594,
          lng: 139.79489050805572
        }
      },
      // markerList: 地図上にプロットされるマーカーのリスト
      markerList: [],
      // stationList: キーワード検索結果のリスト
      stationList: [],
      // textField: キーワード検索欄の文字列
      textField: "",
      // hasResult: キーワード検索の結果があるかどうかのフラグ
      hasResult: false
    };
  },
  methods: {
    // 現在位置にある駅舎一覧を取得する
    async getMarkersInCurrentRect() {
      try {
        let resp = await axios.get(
          `
          http://${window.location.hostname}:1323/buildings`,
          {
            params: {
              begin_latitude: this.bounds._northEast.lat,
              begin_longitude: this.bounds._northEast.lng,
              end_latitude: this.bounds._southWest.lat,
              end_longitude: this.bounds._southWest.lng
            }
          }
        );

        let markers = resp.data.map(elem => ({
          lat: elem.latitude,
          lng: elem.longitude,
          name: elem.name,
          id: elem.id
        }));

        return markers;
      } catch (error) {
        console.error("ERROR @ getMarkersInCurrentRect ");
        throw error;
      }
    },

    // キーワードから駅を検索して、駅一覧リストを取得する
    async getStationListByKeyword(keyword) {
      console.log(`Keyword: ${keyword}`);

      try {
        let resp = await axios.get(`http://${window.location.hostname}:1323/stations/suggest?keyword=${keyword}`);
        let stationList = Array();

        stationList = resp.data.map(elem => ({
          lat: elem.latitude,
          lng: elem.longitude,
          stationName: elem.name,
          railwayName: elem.railway_line_name,
          orderInRailway: elem.order_in_railway,
          stationId: elem.station_id,
          buildingId: elem.building_id
        }));

        return stationList;
      } catch (error) {
        console.error(`ERROR @ getStationListByKeyword (${keyword})`);
        throw error;
      }
    },

    // 駅 ID から駅情報を取得する
    async getStationById(stationId) {
      try {
        let resp = await axios.get(`http://${window.location.hostname}:1323/stations/${stationId}`);

        let stationInfo = resp.data.map(elem => ({
          name: elem.name,
          lat: elem.latitude,
          lng: elem.longitude,
          railwayName: elem.railway_name,
          orderInRailway: elem.order_in_railway
        }));

        return stationInfo[0];
      } catch (error) {
        console.error(`ERROR @ getStationById (${stationId})`);
        throw error;
      }
    },

    // 駅情報内にある緯度経度の位置にフォーカスする
    forcusToStation(stationInfo) {
      this.$refs.mainMap.mapObject.panTo([stationInfo.lat, stationInfo.lng]);
    },

    // キーワードに完全一致した駅にフォーカスする
    checkCompleteMatchAndForcus(keyword) {
      const matchedToKeywordCompletely = this.stationList.filter(elem => elem.stationName == keyword);

      if (matchedToKeywordCompletely.length > 0) {
        // マッチした中で1番目の駅にフォーカス
        this.forcusToStation(matchedToKeywordCompletely[0]);
      }
    },

    // キーワードに基づく駅検索
    searchStation() {
      let keyword = this.textField;

      this.getStationListByKeyword(keyword)
        .then(stationList => {
          this.stationList = stationList;
        })
        .catch(error => {
          console.error(error);
        });

      // もし完全一致する駅が存在すれば検索結果の
      // 1つ目の駅にフォーカス
      this.checkCompleteMatchAndForcus(keyword);
    },

    /*****************************************************************/
    /************************** Event Handlers ***********************/
    /*****************************************************************/

    // ツールバーの「現在地」アイコンを押したとき
    onClickMyLocationIcon() {
      this.getMarkersInCurrentRect()
        .then(markerList => {
          this.markerList = markerList;
        })
        .catch(error => {
          console.error(error);
        });
    },

    // キーワード検索結果の候補をクリックしたとき
    onClickStationList(stationInfo) {
      this.forcusToStation(stationInfo);
    },

    /**********************************/
    /*******  Map (Leaflet.js)  *******/
    /**********************************/

    // ズームスケールが変更されたとき
    onUpdateZoom(zoom) {
      this.zoom = zoom;
    },

    // 中心座標が変更されたとき
    onUpdateCenter(center) {
      this.center = center;
    },

    // 表示範囲が変更されたとき
    onUpdateBounds(bounds) {
      this.bounds = bounds;
    }
  },

  // このコンポーネントがマウントされたときに実行される処理
  mounted: function() {
    this.$nextTick(function() {
      // 初期位置・ズームの設定
      this.bounds = this.$refs.mainMap.mapObject.getBounds();
    });
  },

  // 変数の監視処理
  watch: {
    // textField: キーワード検索文字列
    textField(str) {
      // 何も入力されてなければリストを非表示
      if (isEmpty(str)) {
        this.hasResult = false;
      } else {
        this.getStationListByKeyword(this.textField)
          .then(stationList => {
            if (stationList.length >= 1) {
              this.stationList = stationList;
              this.hasResult = true;
            }
          })
          .catch(error => {
            console.log(error);
          });
      }
    }
  }
};

// 空文字列かどうかチェックする関数
function isEmpty(str) {
  return !str || /^\s*$/.test(str);
}
</script>
