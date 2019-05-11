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
      <TMarker v-for="marker in markerList" :key="marker.id" :data="marker"/>
    </l-map>
    <!-- floatingっぽく見せるためのpadding -->
    <div class="pa-3">
      <!-- 検索フィールドの幅を指定 -->
      <v-card-text style="width: 320px; position: relative;">
        <!-- 検索フィールド(乗車駅) -->
        <v-toolbar absolute flat height="50px" v-bind:style="rideStationTextFieldStyle">
          <v-toolbar-side-icon></v-toolbar-side-icon>
          <v-text-field
            clearable
            single-line
            label="乗車駅を入力"
            v-model="rideStationTextField"
            ref="rideStationTextFieldRef"
            @keyup.enter="searchStation"
          ></v-text-field>
          <v-btn icon>
            <v-icon>search</v-icon>
          </v-btn>
        </v-toolbar>
      </v-card-text>
      <v-card-text style="width: 320px; position: relative;">
        <!-- 検索フィールド(降車駅) -->
        <v-toolbar v-show="showDropStationTextField" absolute flat height="50px" v-bind:style="dropStationTextFieldStyle">
          <v-toolbar-side-icon></v-toolbar-side-icon>
          <v-text-field
            clearable
            single-line
            label="降車駅を入力"
            v-model="dropStationTextField"
            ref="dropStationTextFieldRef"
            @keyup.enter="searchStation"
          ></v-text-field>
          <v-btn icon>
            <v-icon>search</v-icon>
          </v-btn>
        </v-toolbar>
      </v-card-text>
      <!-- サジェストのリスト -->
      <v-card-text v-bind:style="suggestListStyle">
        <v-list subheader absolute avatar v-show="showSuggestList" style="background-color:#f5f5f5; border-radius:0px 0px 10px 10px;">
          <v-subheader>候補...</v-subheader>
          <v-list-tile
            v-for="stationInfo in stationList.slice(0, 5)"
            :key="stationInfo.stationId"
            @click="onClickStationList(stationInfo)" >
            <v-list-tile-avatar>
              <v-icon large>train</v-icon>
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-list-tile-title v-text="stationInfo.name"></v-list-tile-title>
              <v-list-tile-sub-title>
                <!-- 2路線以上だと最後の路線の後に，が入って少し見づらい -->
                <div v-if="stationInfo.lines.length>=2">
                  <span
                  v-for="lines in stationInfo.lines.slice(0, 3)"
                  :key="lines.station_id"
                  v-text="lines.railway_name + '，' "
                  ></span>
                </div>
                <div v-else>
                  <span
                  v-for="lines in stationInfo.lines"
                  :key="lines.station_id"
                  v-text="lines.railway_name"
                  ></span>
                </div>
              </v-list-tile-sub-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-card-text>
        <v-btn
          absolute
          dark
          fab
          bottom
          right
          style="margin-bottom:75px; color:gray; background-color:gray;"
          color="gray"
        >
          <v-icon>my_location</v-icon>
        </v-btn>
        <v-btn
          absolute
          dark
          fab
          bottom
          right
          style="margin-bottom:150px;"
          color="blue"
          @click="onClickMyLocationIcon"
        >
          <v-icon>search</v-icon>
        </v-btn>
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
      // rideStationTextField,dropStationTextField: キーワード検索欄の文字列
      rideStationTextField: "",
      dropStationTextField: "",
      // hasResult: キーワード検索の結果があるかどうかのフラグ
      hasResult: false,
      // showDropStationTextField: 降車駅を入力するフィールドを表示するかのフラグ
      showDropStationTextField: false,
      // showSuggestList: キーワード検索の結果リストを表示するかのフラグ
      showSuggestList: false,
      rideStationFixed: false,
      dropStationFixed: false,
      // 表示によってテキストフィールドの角を丸めるためのスタイル指定
      rideStationTextFieldStyle: {
        borderRadius: '10px'
      },
      dropStationTextFieldStyle: {
        marginTop: '18px',
        borderRadius: '10px'
      },
      // suggestリストをフィールドに合わせて移動するためのスタイル指定
      suggestListStyle: {
        width: '352px',
        position: 'relative',
        top: '-50px',
        left: '-16px',
        paddingTop: '30px'
      }

      
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
          id: elem.building_id,
          lines: elem.lines
        }));

        return markers;
      } catch (error) {
        console.error("ERROR @ getMarkersInCurrentRect ");
        throw error;
      }
    },
    
    // キーワードから駅を検索して、駅一覧リストを取得する(建物単位)
    async getBuildingListByKeyword(keyword) {
      console.log(`Keyword: ${keyword}`);

      try {
        let resp = await axios.get(`http://${window.location.hostname}:1323/buildings/suggest?keyword=${keyword}`);
        let stationList = Array();

        stationList = resp.data.map(elem => ({
          lat: elem.latitude,
          lng: elem.longitude,
          name: elem.name,
          id: elem.building_id,
          lines: elem.lines
        }));

        return stationList;
      } catch (error) {
        console.error(`ERROR @ getBuildingListByKeyword (${keyword})`);
        throw error;
      }
    },

    // キーワードから駅を検索して、駅一覧リストを取得する(駅単位)
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
      let keyword = this.rideStationTextField;

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
          console.log(this.markerList);
        })
        .catch(error => {
          console.error(error);
        });
    },

    // キーワード検索結果の候補をクリックしたとき
    onClickStationList(stationInfo) {
      this.forcusToStation(stationInfo);
      this.markerList = [stationInfo];
      this.showSuggestList = false;
      // 乗車駅の入力
      if (!this.showDropStationTextField) { 
        this.showDropStationTextField = true;
        this.rideStationTextField = stationInfo.name;
        this.suggestListStyle.top = '0px'
        this.dropStationTextFieldStyle.borderRadius = "0px 0px 10px 10px";
        this.rideStationFixed = true;
      } // 降車駅の入力
      else {
        this.dropStationTextField = stationInfo.name;
        this.dropStationFixed = true;
        this.dropStationTextFieldStyle.borderRadius = "0px 0px 10px 10px";
      }
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
    // rideStationTextField: キーワード検索文字列
    rideStationTextField(str) {
      if(this.showDropStationTextField) {
        return;
      }
      if(this.rideStationFixed) {
        this.rideStationFixed = false;
      }
      // 何も入力されてなければリストを非表示にする
      if (isEmpty(str)) {
        this.hasResult = false;
        this.showSuggestList = false;
      } else {
        this.getBuildingListByKeyword(this.rideStationTextField)
          .then(stationList => {
            if (stationList.length >= 1) {
              this.stationList = stationList;
              this.hasResult = true;
                if(this.rideStationFixed) {
                  this.showSuggestList = false;
              } else {
                this.showSuggestList = true;
              }
            }
          })
          .catch(error => {
            console.log(error);
          });
      }
    },
    dropStationTextField(str) {
      // テキストフィールドが編集されたときのFixed解除
      // if (this.dropStationFixed && this.stationList[0].name == str) {
      //   this.dropStationFixed = false;
      // }
      if(!this.showDropStationTextField) {
        return;
      }
      // 何も入力されてなければリストを非表示
      if (isEmpty(str)) {
        this.hasResult = false;
        this.showSuggestList = false;
      } else {
        this.getBuildingListByKeyword(this.dropStationTextField)
          .then(stationList => {
            if (stationList.length >= 1) {
              this.stationList = stationList;
              this.hasResult = true;
              if (this.dropStationFixed) {
                this.showSuggestList = false;
              } else {
                this.showSuggestList = true;
              }
            }
          })
          .catch(error => {
            console.log(error);
          });
      }
    },
    hasResult() {
      // リザルトに変更があった場合のデザイン
      if (this.hasResult | this.showDropStationTextField) { 
        this.rideStationTextFieldStyle.borderRadius = "10px 10px 0px 0px";
      } else {
        this.rideStationTextFieldStyle.borderRadius = "10px";
      }
      // 下にサジェストが表示されている
      if (this.showSuggestList) {
        this.dropStationTextFieldStyle.borderRadius = "0px";
      }
    },
    showDropStationTextField() {
      if (this.hasResult | this.showDropStationTextField) { 
        this.rideStationTextFieldStyle.borderRadius = "10px 10px 0px 0px";
      } else {
        this.rideStationTextFieldStyle.borderRadius = "10px";
      }
    }
  
  }
};

// 空文字列かどうかチェックする関数
function isEmpty(str) {
  return !str || /^\s*$/.test(str);
}
</script> 