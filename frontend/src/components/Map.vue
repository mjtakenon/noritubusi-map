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
    <!-- 路線登録時に左から出てくるメニュー -->
    <v-slide-x-transition>
      <v-card
        light
        height="100%"
        width="340px"
        style="position: absolute; top:120px;"
        transition="slide-x-transition"
        v-show="showInputDetailsModal"
      >
      <!-- 路線から駅を選択するリスト -->
      <v-list subheader v-show="isRideStationFixed" >
        <v-subheader v-text="rideStationTextFieldModel"></v-subheader>
        <v-list-group
          v-for="l in rideStation.lines"
          :key="l.railway_name" 
          @click="onClickRailwayList(l)"
          no-action
          :prepend-icon="'train'">
          <template v-slot:activator>
              <v-list-tile>
                <v-list-tile-content>
                  <v-list-tile-title> {{ l.railway_name }} </v-list-tile-title>
                </v-list-tile-content>
              </v-list-tile>
          </template>
          <v-list-tile 
          v-for="(ll, idx) in suggestedDropStationList"
          :key="idx"
          @click="onClickSuggestedDropStation(ll)">
            <v-list-tile-content>
              <v-list-tile-title> {{ ll }} </v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list-group>
      </v-list>
    </v-card>
    </v-slide-x-transition>
    <v-slide-x-transition>
      <v-card
        light
        height="120px"
        width="340px"
        style="position: absolute; top:0px; background-color:#2196f3"
        transition="slide-x-transition"
        v-show="showInputDetailsModal"
      ></v-card>
    </v-slide-x-transition>
    <!-- floatingっぽく見せるためのpadding -->
    <div class="pa-2">
      <!-- 検索フィールドの幅を指定 -->
      <v-card-text style="width: 320px; position: relative;">
        <!-- 検索フィールド(乗車駅) -->
        <v-toolbar absolute height="50px" v-bind:style="rideStationToolbarStyle" v-bind:flat="flatToolbar">
          <v-toolbar-side-icon style="color:#FFFFFF"></v-toolbar-side-icon>
          <v-text-field clearable single-line dark autofocus
            label="乗車駅を入力"
            tabindex="1"
            v-model="rideStationTextFieldModel"
            @keyup.enter="searchStation"
            @click:clear="rideStationTextFieldCleared"
          ></v-text-field>
          <v-btn icon style="color:#FFFFFF" @click="searchStation">
            <v-icon>search</v-icon>
          </v-btn>
        </v-toolbar>
      </v-card-text>
      <v-card-text style="width: 320px; position: relative;">
        <!-- 検索フィールド(降車駅) -->
        <v-toolbar v-show="showDropStationTextField" absolute height="50px" v-bind:style="dropStationToolbarStyle" v-bind:flat="flatToolbar">
          <v-btn icon style="color:#FFFFFF" @click="swapTextField">
            <v-icon>swap_vert</v-icon>
          </v-btn>
          <v-text-field clearable single-line dark
            label="降車駅を入力"
            tabindex="2"
            v-model="dropStationTextFieldModel"
            @keyup.enter="searchStation"
            @click:clear="dropStationTextFieldCleared"
          ></v-text-field>
          <v-btn icon style="color:#FFFFFF">
            <v-icon>search</v-icon>
          </v-btn>
        </v-toolbar>
      </v-card-text>
      <!-- サジェストのリスト -->
      <v-card-text v-bind:style="suggestListStyle">
        <v-slide-y-transition>
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
        </v-slide-y-transition>
      </v-card-text>
      <!-- 左下のFloatingActionButton -->
      <v-btn absolute dark fab bottom right color="pink"
        style="margin-bottom:75px;"
        @click="onClickGetCurrentPosition"
      >
        <v-icon>my_location</v-icon>
      </v-btn>
      <v-btn absolute dark fab bottom right color="blue"
        style="margin-bottom:150px;"
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
      // rideStation: 乗車駅
      rideStation: [],
      // drioStation: 降車駅
      dropStation: [],
      // suggestedDropStationList: 乗車駅に接続している駅のリスト
      suggestedDropStationList: [],
      // rideStationTextFieldModel, dropStationTextFieldModel: キーワード検索欄の文字列
      rideStationTextFieldModel: "",
      dropStationTextFieldModel: "",
      // flatToolbar: 検索バーのデザインをflatにするか
      flatToolbar: false,
      // hasResult: キーワード検索の結果があるかどうかのフラグ
      hasResult: false,
      // showDropStationTextField: 降車駅を入力するフィールドを表示するかのフラグ
      showDropStationTextField: false,
      // showSuggestList: キーワード検索の結果リストを表示するかのフラグ
      showSuggestList: false,
      // showInputDetailsModal: 駅登録の詳細を入力できるよう左から出てくるモーダル
      showInputDetailsModal: false,
      // isRideStationFixed, isDropStationFixed: それぞれのテキストフィールドが確定しているか
      isRideStationFixed: false,
      isDropStationFixed: false,
      // 表示によってテキストフィールドの角を丸めるためのスタイル指定
      rideStationToolbarStyle: {
        borderRadius: '10px',
        backgroundColor: '#2196f3'
      },
      dropStationToolbarStyle: {
        marginTop: '18px',
        borderRadius: '10px',
        backgroundColor: '#2196f3'
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
        // let resp = await axios.get(`http://${window.location.hostname}:1323/stations/suggest?keyword=${keyword}`);
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
    // 戻り値: true=1つ以上の駅が見つかった場合, false=見つからなかった場合
    checkCompleteMatchAndForcus(keyword) {
      const matchedToKeywordCompletely = this.stationList.filter(elem => elem.name == keyword);
      if (matchedToKeywordCompletely.length > 0) {
        // マッチした中で1番目の駅にフォーカス
        this.forcusToStation(matchedToKeywordCompletely[0]);
        this.markerList = [matchedToKeywordCompletely[0]];
        return matchedToKeywordCompletely[0];
      }
      return false;
    },

    // キーワードに基づく駅検索
    searchStation() {
      let keyword = "";
      if (!this.isRideStationFixed) {
        keyword = this.rideStationTextFieldModel;
      } else {
        keyword = this.dropStationTextFieldModel;
      }

      this.getStationListByKeyword(keyword)
        .then(stationList => {
          this.stationList = stationList;
        })
        .catch(error => {
          console.error(error);
        });
      // もし完全一致する駅が存在すれば検索結果の
      // 1つ目の駅にフォーカス
      var result = this.checkCompleteMatchAndForcus(keyword);
      if (!result=== false) {
        // 乗車駅の編集
        if(!this.isRideStationFixed) { 
          this.rideStationFix(result);
        } else { // 降車駅 
          this.dropStationFix(result);
        }
      }
    },
    rideStationFix(stationInfo) {
        this.showDropStationTextField = true;
        this.suggestListStyle.position = 'absolute'
        this.suggestListStyle.top = '120px'
        this.suggestListStyle.width = '372px'
        this.suggestListStyle.paddingTop = '0px'
        this.dropStationToolbarStyle.borderRadius = "0px 0px 10px 10px";
        this.isRideStationFixed = true;
        this.showInputDetailsModal = true;
        this.flatToolbar = true;
        this.showSuggestList = false;
        this.rideStation = stationInfo;
    },
    rideStationUnfix() {
      this.isRideStationFixed = false;
      if (isEmpty(this.dropStationTextFieldModel)) {
        this.showDropStationTextField = false;
        this.suggestListStyle.position = 'relative'
        this.suggestListStyle.top = '-50px'
        this.suggestListStyle.width = '352px'
        this.suggestListStyle.paddingTop = '30px'
        this.dropStationToolbarStyle.borderRadius = "0px";
        this.showInputDetailsModal = false;
        this.flatToolbar = false;
        this.showSuggestList = false;
        this.rideStation = [];
      }
    },
    dropStationFix(stationInfo) {
      this.isDropStationFixed = true;
      this.dropStationToolbarStyle.borderRadius = "0px 0px 10px 10px";
      this.showSuggestList = false;
      this.dropStation = stationInfo;
    },
    dropStationUnfix() {
      this.isDropStationFixed = false;
      this.dropStationToolbarStyle.borderRadius = "0px";
      this.showSuggestList = false;
      if (!this.isRideStationFixed) {
        this.rideStationUnfix();
      }
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
      this.stationList = [stationInfo];
      // 乗車駅の入力
      if (!this.isRideStationFixed) { 
        this.rideStationTextFieldModel = stationInfo.name;
        this.rideStationFix(stationInfo);
      } // 降車駅の入力
      else {
        this.dropStationTextFieldModel = stationInfo.name;
        this.dropStationFix(stationInfo);
      }
    },

    // 路線名をクリックしたとき
    onClickRailwayList(railwayInfo) {
      // 実装?
      // this.getStationListByRailwayName(railwayInfo.railway_name)
      this.suggestedDropStationList = [ "駅1", "駅2", "駅3" ];
      console.log(railwayInfo);
    },

    // サジェストされた降車駅をクリックしたとき
    onClickSuggestedDropStation(stationInfo) {
      console.log(stationInfo)
    },

    // 現在地ボタンをクリックしたとき
    onClickGetCurrentPosition() {
      navigator.geolocation.getCurrentPosition(this.getCurrentPositionCompleted);
    },

    // 位置情報の取得が完了したとき
    getCurrentPositionCompleted(pos) {

      this.$refs.mainMap.mapObject.panTo([pos.coords.latitude, pos.coords.longitude]);
      // this.$refs.mainMap.mapObject.setView(new L.LatLng(pos.coords.latitude,  pos.coords.longitude), this.zoom);
      // 移動後の場所にピンを立てようとしても移動前の場所になってしまう
      // this.onClickMyLocationIcon();
    },

    // テキストフィールドのクリアボタンを押したとき
    rideStationTextFieldCleared() {
      this.rideStationUnfix();
    },
    dropStationTextFieldCleared() {
      this.dropStationUnfix();
    },

    // swapボタンを押したとき
    swapTextField() {
      // フィールドの内容を交換
      this.dropStationTextFieldModel = [this.rideStationTextFieldModel, this.rideStationTextFieldModel = this.dropStationTextFieldModel][0];
      this.isDropStationFixed = [this.isRideStationFixed, this.isRideStationFixed = this.isDropStationFixed][0];
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
    },
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
    // rideStationTextFieldModel: キーワード検索文字列
    rideStationTextFieldModel(str) {
      // 入力確定後に変更があり、それが候補と違えば確定を解除
      if (this.isRideStationFixed) {
        if (str != this.stationList[0].name) {
          this.rideStationUnfix();
        } else {
          this.showSuggestList = true;
        }
      }
      // 何も入力されてなければリストを非表示にする
      if (isEmpty(str)) {
        this.hasResult = false;
        this.showSuggestList = false;
      } else {
        this.getBuildingListByKeyword(this.rideStationTextFieldModel)
          .then(stationList => {
            if (stationList.length >= 1) {
              this.stationList = stationList;
              this.hasResult = true;
              if(this.isRideStationFixed) {
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
    dropStationTextFieldModel(str) {
      // 入力確定後に変更があり、それが候補と違えば確定を解除
      if(this.isDropStationFixed && str != this.stationList[0].name) {
        this.dropStationUnfix();
      }
      // 何も入力されてなければリストを非表示
      if (isEmpty(str)) {
        this.hasResult = false;
        this.showSuggestList = false;
      } else {
        this.getBuildingListByKeyword(this.dropStationTextFieldModel)
          .then(stationList => {
            if (stationList.length >= 1) {
              this.stationList = stationList;
              this.hasResult = true;
              if (this.isDropStationFixed) {
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
        this.rideStationToolbarStyle.borderRadius = "10px 10px 0px 0px";
      } else {
        this.rideStationToolbarStyle.borderRadius = "10px";
      }
      // 下にサジェストが表示されている
      if (this.showSuggestList) {
        this.dropStationToolbarStyle.borderRadius = "0px";
      }
    },
    showDropStationTextField() {
      if (this.hasResult | this.showDropStationTextField) { 
        this.rideStationToolbarStyle.borderRadius = "10px 10px 0px 0px";
      } else {
        this.rideStationToolbarStyle.borderRadius = "10px";
      }
    }
  }
};

// 空文字列かどうかチェックする関数
function isEmpty(str) {
  return !str || /^\s*$/.test(str);
}
</script> 