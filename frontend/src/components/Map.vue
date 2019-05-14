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
        <!-- 路線から降車駅を選択するリスト -->
        <v-list subheader v-if="isRideStationFixed">
          <v-subheader v-text="rideStationTextFieldModel"></v-subheader>
          <!-- 降車駅選択前 -->
          <div v-if="!isDropStationFixed">
            <v-list-group
              v-for="l in rideStation.lines"
              :key="l.railway_name" 
              @click="onClickRideRailwayList(l)"
              :prepend-icon="'train'">
              <template v-slot:activator>
                <v-list-tile>
                  <v-list-tile-content>
                    <v-list-tile-title> {{ l.railway_name }} </v-list-tile-title>
                  </v-list-tile-content>
                </v-list-tile>
              </template>
              <v-list-tile 
              v-for="(s, idx) in suggestedDropStationList"
              :key="idx"
              @click="onClickSuggestedDropStation(s,l)">
                <v-list-tile-content>
                  <v-list-tile-title> {{ s }} </v-list-tile-title>
                </v-list-tile-content>
              </v-list-tile>
            </v-list-group>
          </div>
          <!-- 降車駅選択後 -->
          <div v-else>
            <!-- 路線数が1の場合 -->
            <v-list-group :prepend-icon="'train'" v-if="rideRailway.length==1">
              <template v-slot:activator>
                <div>
                  {{ rideRailway[0].railway_name }}
                </div>
              </template>
              <v-list-tile>
                <v-list-tile-content>
                  <v-list-tile-title> hoge </v-list-tile-title>
                </v-list-tile-content>
              </v-list-tile>
            </v-list-group>
            <!-- 路線数が複数 -->
            <v-list v-else-if="rideRailway.length>=2">
              <v-list-tile 
                v-for="(r, idx) in rideRailway"
                :key="idx"
                @click="onClickUseRailwayList(r)">
                <v-list-tile-avatar>
                  <v-icon>train</v-icon>
                </v-list-tile-avatar>
                <v-list-tile-content>
                  <v-list-tile-title> {{ r.railway_name }} </v-list-tile-title>
                </v-list-tile-content>
              </v-list-tile>
            </v-list>
            <!-- 路線数がない場合 -->
            <v-list v-else>
              <v-list-tile>
                <v-list-tile-avatar>
                  <v-icon>warning</v-icon>
                </v-list-tile-avatar>
                <v-list-tile-content>
                  <v-list-tile-title> 共通の路線がありません </v-list-tile-title>
                </v-list-tile-content>
              </v-list-tile>
            </v-list>
          </div>
          <v-subheader v-if="isDropStationFixed" v-text="dropStationTextFieldModel"></v-subheader>
        </v-list>
        <v-btn absolute right color="#2196f3"
          v-show="isDropStationFixed && isRideStationFixed" 
          v-bind:disabled="!(isDropStationFixed && isRideStationFixed)" 
          v-bind:dark="isDropStationFixed && isRideStationFixed" 
          @click="onClickRegisterButton()"> 登録 </v-btn>
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
    <div class="pa-2" style="z-index:0">
      <!-- 検索フィールドの幅を指定 -->
      <v-card-text style="width: 320px; position: relative;">
        <!-- 検索フィールド(乗車駅) -->
        <v-toolbar absolute height="50px" v-bind:style="rideStationToolbarStyle" v-bind:flat="flatToolbar">
          <v-toolbar-side-icon style="color:#FFFFFF"
          @click="onClickSideIcon()"
          ></v-toolbar-side-icon>
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
      <v-btn fixed dark fab bottom right color="pink"
        style="margin-bottom:50px;"
        @click="onClickGetCurrentPosition()"
      >
        <v-icon>my_location</v-icon>
      </v-btn>
      <v-btn fixed dark fab bottom right color="blue"
        style="margin-bottom:125px;"
        @click="onClickMyLocationIcon()"
      >
        <v-icon>search</v-icon>
      </v-btn>
    </div>
    <v-slide-x-transition>
      <v-card
        height="100%"
        width="300px"
        style="position: absolute; top:0px; z-index:10; padding:10px;"
        transition="slide-x-transition"
        v-show="showSideMenu"
      >
        <v-list>
          <v-list-tile>
            <v-list-tile-avatar>
              <v-icon x-large>portrait</v-icon>
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-list-tile-title v-if="isLoggedIn"> {{ username }} </v-list-tile-title>
              <v-list-tile-title v-else> 未ログイン </v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
          <v-alert
            v-model="successAlertModel"
            type="success"
            dismissible
            transition="slide-y-transition"
          >
            {{ successAlertMessage }}
          </v-alert>

        <v-alert
          v-model="errorAlertModel"
          type="error"
          dismissible
          transition="slide-y-transition"
        >
          {{ errorAlertMessage }}
        </v-alert>
        <div class="text-xs-center" v-if="!isLoggedIn">
          <v-btn @click="showSignupMenu=true; showSigninMenu=false;"> アカウント登録 </v-btn>
          <v-btn @click="showSigninMenu=true; showSignupMenu=false;"> ログイン </v-btn>
        </div>
        <div class="text-xs-center" v-else>
          <v-btn @click="onClickSignoutButton"> ログアウト </v-btn>
        </div>
        <v-list v-if="!showSignupMenu && !showSigninMenu">
          <v-list-tile @click="showSideMenu=!showSideMenu">
            <v-list-tile-avatar>
              <v-icon>train</v-icon>
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-list-tile-title> 乗車記録をつける </v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
          <v-list-tile @click="showSideMenu=!showSideMenu">
            <v-list-tile-avatar>
              <v-icon>search</v-icon>
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-list-tile-title> 乗車記録を見る </v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
          <v-list-tile @click="showSideMenu=!showSideMenu">
            <v-list-tile-avatar>
              <v-icon>edit</v-icon>
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-list-tile-title> 乗車記録を編集する </v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
        <div v-else-if="showSignupMenu | showSigninMenu" class="pa-3">
          <v-text-field
            v-model="usernameModel"
            ref="usernameRef"
            label="ユーザー名"
            :rules="[usernameRules.required]"
            clearable
          ></v-text-field>
          <v-text-field
            v-model="passwordModel"
            ref="passwordRef"
            :append-icon="showPassword ? 'visibility' : 'visibility_off'"
            :rules="[passwordRules.required, passwordRules.min]"
            :type="showPassword ? 'text' : 'password'"
            label="パスワード"
            hint="8文字以上で入力してください。"
            counter
            clearable
            @click:append="showPassword = !showPassword"
          ></v-text-field>
          <v-text-field
            v-model="passwordConfirmModel"
            ref="passwordConfirmRef"
            :append-icon="showPasswordConfirm ? 'visibility' : 'visibility_off'"
            :rules="[passwordConfirmRules.required, passwordConfirmRules.passwordMatch]"
            :type="showPasswordConfirm ? 'text' : 'password'"
            label="パスワードの確認"
            hint="8文字以上で入力してください。"
            counter clearable
            @click:append="showPasswordConfirm = !showPasswordConfirm"
            v-if="showSignupMenu"
          ></v-text-field>
          <v-btn absolute right color="#2196f3"
          v-bind:disabled="signupFormHasError"
          v-bind:dark="!signupFormHasError"
          v-if="showSignupMenu" 
          @click="onClickSignupButton()"> 登録 </v-btn>
          <v-btn absolute right color="#2196f3"
          v-bind:disabled="signinFormHasError" 
          v-bind:dark="!signinFormHasError"
          v-if="showSigninMenu"
          @click="onClickSigninButton()"> ログイン </v-btn>
        </div>
      </v-card>
    </v-slide-x-transition>
    <v-card
      height="100%"
      width="100%"
      style="position: absolute; top:0px; z-index:9; background-color:black; opacity:0.2;"
      v-show="showSideMenu"
      @click="showSideMenu=false"
    >
    </v-card>
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
      // dropStation: 降車駅
      dropStation: [],
      // rideRailway: 利用路線
      rideRailway: [],
      // suggestedDropStationList: 乗車駅に接続している駅のリスト
      suggestedDropStationList: [],
      // rideStationTextFieldModel, dropStationTextFieldModel: キーワード検索欄の文字列
      rideStationTextFieldModel: "",
      dropStationTextFieldModel: "",
      // usernameModel,passwordModel,passwordConfirmModel: ユーザー名,パスワード,確認のパスワードの文字列
      usernameModel: "",
      passwordModel: "",
      passwordConfirmModel: "",
      // username: ユーザー名
      username: "",
      // successAlertMessage, errorAlertMessage: アラートに表示するメッセージ
      successAlertMessage: "",
      errorAlertMessage: "",
      // flatToolbar: 検索バーのデザインをflatにするか
      flatToolbar: false,
      // hasResult: キーワード検索の結果があるかどうかのフラグ
      hasResult: false,
      // showSideMenu: ログイン処理などのメニューを表示するかのフラグ
      showSideMenu: false,
      // showSignupMenu: アカウント作成メニューを表示するかのフラグ
      showSignupMenu: false,
      // showSigninMenu: ログインメニューを表示するかのフラグ
      showSigninMenu: false,
      // showPassword, showPasswordConfirm: パスワードを表示するかどうかのフラグ
      showPassword: false,
      showPasswordConfirm: false,
      // signupFromHasError, signinFormHasError: それぞれのフォームにエラーがあるかのフラグ
      signupFormHasError: true,
      signinFormHasError: true,
      // isLoggedIn: ログイン中かのフラグ
      isLoggedIn: false,
      // successAlertModel, errorAlertModel: アラートを表示するかのフラグ
      successAlertModel: false,
      errorAlertModel: false,
      // showDropStationTextField: 降車駅を入力するフィールドを表示するかのフラグ
      showDropStationTextField: false,
      // showSuggestList: キーワード検索の結果リストを表示するかのフラグ
      showSuggestList: false,
      // showInputDetailsModal: 駅登録の詳細を入力できるよう左から出てくるモーダル
      showInputDetailsModal: false,
      // isRideStationFixed, isDropStationFixed: それぞれのテキストフィールドが確定しているか
      isRideStationFixed: false,
      isDropStationFixed: false,
      // rideStationToolbarStyle,dropStationToolbarStyle: 表示によってテキストフィールドの角を丸めるためのスタイル指定
      rideStationToolbarStyle: {
        borderRadius: '10px',
        backgroundColor: '#2196f3'
      },
      dropStationToolbarStyle: {
        marginTop: '18px',
        borderRadius: '10px',
        backgroundColor: '#2196f3'
      },
      // suggestListStyle: suggestリストをフィールドに合わせて移動するためのスタイル指定
      suggestListStyle: {
        width: '352px',
        position: 'relative',
        top: '-50px',
        left: '-16px',
        paddingTop: '30px'
      },
      // usernameRules: ユーザー名の制限
      usernameRules: { 
        required: value => !!value || 'このフィールドは必須です。'
      },
      // passwordRules: パスワードの制限
      passwordRules: {
        required: value => !!value || 'このフィールドは必須です。',
        min: v => v.length >= 8 || '8文字以上で入力してください。',
      },
      // passwordConfirmRules: パスワードの制限
      passwordConfirmRules: {
        required: value => !!value || 'このフィールドは必須です。',
        passwordMatch: value => (value == this.passwordModel) || 'パスワードが一致していません。'
      }
    };
  },
  methods: {
    // アカウントを登録する
    async signupAccount(username,password) {
      var params = new URLSearchParams();
      params.append('userid', username);
      params.append('password', password);
      try {
        let resp = await axios.post(
          `http://${window.location.hostname}:1323/signup`, params);

        return resp;
      } catch (error) {
        console.error("ERROR @ signupAccount ");
        throw error;
      }
    },
    // ログイン
    async signin(username,password) {
      var params = new URLSearchParams();
      params.append('userid', username);
      params.append('password', password);
      try {
        let resp = await axios.post(
          `http://${window.location.hostname}:1323/signin`, params);

        return resp;
      } catch (error) {
        console.error("ERROR @ signin ");
        throw error;
      }
    },
    // ログアウト
    async signout() {
      var params = new URLSearchParams();
      try {
        let resp = await axios.delete(
          `http://${window.location.hostname}:1323/signin`);

        return resp;
      } catch (error) {
        console.error("ERROR @ signout ");
        throw error;
      }
    },

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

    // キーワードに完全一致、または候補が1つしかない場合に駅にフォーカスする
    // 戻り値: stationInfo=1つ以上の駅が見つかった場合, false=見つからなかった場合
    checkCompleteMatchAndForcus(keyword) {
      const matchedToKeywordCompletely = this.stationList.filter(elem => elem.name == keyword);
      if (matchedToKeywordCompletely.length > 0) {
        // マッチした中で1番目の駅にフォーカス
        this.forcusToStation(matchedToKeywordCompletely[0]);
        this.markerList = [matchedToKeywordCompletely[0]];
        return matchedToKeywordCompletely[0];
      }
      if (this.stationList.length == 1) {
        this.forcusToStation(this.stationList[0]);
        this.markerList = this.stationList;
        if(!this.isRideStationFixed) {
          this.rideStationTextFieldModel = this.stationList[0].name;
        } else {
          this.dropStationTextFieldModel = this.stationList[0].name;
        }
        return this.stationList[0];
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
          if(this.isDropStationFixed) {
            this.checkRoute();      
          }
        } else { // 降車駅 
          this.dropStationFix(result);
          this.checkRoute();      
        }
      }
    },
    // 乗車駅の確定処理
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
    // 乗車駅の確定解除
    rideStationUnfix() {
      this.isRideStationFixed = false;
      this.rideStation = [];

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
      }
    },
    // 降車駅の確定処理
    dropStationFix(stationInfo) {
      this.isDropStationFixed = true;
      this.dropStationToolbarStyle.borderRadius = "0px 0px 10px 10px";
      this.showSuggestList = false;
      this.dropStation = stationInfo;
    },
    // 降車駅の確定解除
    dropStationUnfix() {
      this.isDropStationFixed = false;
      this.dropStationToolbarStyle.borderRadius = "0px";
      this.showSuggestList = false;
      if (!this.isRideStationFixed) {
        this.rideStationUnfix();
      }
    },

    // 登録ボタンを有効にできるかフォームの入力をチェックし更新
    signupFormUpdated() {
      this.signupFormHasError = !(this.isUsernameValid(this.usernameModel) && this.isPasswordValid(this.passwordModel) && this.isPasswordConfirmValid(this.passwordModel,this.passwordConfirmModel));
    },

    signinFormUpdated() {
      this.signinFormHasError = !(this.isUsernameValid(this.usernameModel) && this.isPasswordValid(this.passwordModel));
    },

    isUsernameValid(username) {
      return !!username;
    },

    isPasswordValid(password) { 
      return !!password && (password.length>=8);
    },

    isPasswordConfirmValid(password,passwordConfirm) {
      return passwordConfirm==password;
    },


    /*****************************************************************/
    /************************** Event Handlers ***********************/
    /*****************************************************************/

    // ツールバーのメニューを開くアイコン(サイドアイコン)を押したとき
    onClickSideIcon() {
      this.showSideMenu = true;
    },

    // ユーザーアカウント登録ボタンが押されたとき
    onClickSignupButton() {
      this.signupAccount(this.usernameModel,this.passwordModel).then(resp => {
          console.log(resp);
          this.username = resp.data['userid'];
          this.isLoggedIn = true;
          this.showSignupMenu = false;
          this.successAlertModel = true;
          this.successAlertMessage = "登録しました。";
        })
        .catch(error => {
          console.error(error);
          this.errorAlertModel = true;
          this.errorAlertMessage = "登録に失敗しました。";
        });
    },

    // ログインボタンが押されたとき
    onClickSigninButton() {
      this.signin(this.usernameModel,this.passwordModel).then(resp => {
          console.log(resp);
          this.username = resp.data['userid'];
          this.isLoggedIn = true;
          this.showSigninMenu = false;
          this.successAlertModel = true;
          this.successAlertMessage = "ログインしました。";
        })
        .catch(error => {
          console.error(error);
          this.errorAlertModel = true;
          this.errorAlertMessage = "ログインに失敗しました。";
        });
    },
    // ログアウトボタンが押されたとき
    onClickSignoutButton() {
      this.signout().then(resp => {
          console.log(resp);
          this.username = "";
          this.isLoggedIn = false;
          this.successAlertModel = true;
          this.successAlertMessage = "ログアウトしました。";
        })
        .catch(error => {
          console.error(error);
          this.errorAlertModel = true;
          this.errorAlertMessage = "ログアウトに失敗しました。";
        });
    },

    // キーワード検索結果の候補をクリックしたとき
    onClickStationList(stationInfo, railwayInfo) {
      this.forcusToStation(stationInfo);
      this.markerList = [stationInfo];
      this.stationList = [stationInfo];
      // 乗車駅の入力
      if (!this.isRideStationFixed) { 
        this.rideStationTextFieldModel = stationInfo.name;
        this.rideStationFix(stationInfo);
        if (this.isDropStationFixed) {
          this.checkRoute();
        }
      } // 降車駅の入力
      else {
        this.dropStationTextFieldModel = stationInfo.name;
        this.dropStationFix(stationInfo);
        if(this.isRideStationFixed) {
          this.checkRoute();
        }
      }
    },

    // 乗車した路線名をクリックしたとき
    onClickRideRailwayList(railwayInfo) {
      console.log(railwayInfo);
      // 実装?
      // this.getStationListByRailwayName(railwayInfo.railway_name)
      this.suggestedDropStationList = [ "駅1", "駅2", "駅3" ];
    },
    
    // 後から使用した路線名をクリックしたとき
    onClickUseRailwayList(railwayInfo) {
      this.rideRailway = [railwayInfo];
    },

    // サジェストされた降車駅をクリックしたとき
    onClickSuggestedDropStation(stationInfo,railwayInfo) {
      // TODO これはダミーデータなのでエンドポイントができ次第差し替える
      stationInfo = {'id':1323, 'name': '浜松','lat': '34.703866','lng': '137.734759','lines': [{'railway_name': 'JR東海道本線(熱海～浜松)','station_id': 8517,'order_in_railway': 4},{'railway_name': 'JR東海道本線(浜松～岐阜)','station_id': 8819,'order_in_railway': 1}]};
      console.log(stationInfo);
      console.log(railwayInfo);
      // this.onClickStationList(stationInfo);
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
      this.dropStationFix(stationInfo);
      this.rideRailway = [railwayInfo];
    },

    // 乗り潰し記録の登録ボタンをクリックしたとき
    onClickRegisterButton() {
      console.log("register")
    },

    // 現在地に移動するフローティングアクションボタンをクリックしたとき
    onClickGetCurrentPosition() {
      navigator.geolocation.getCurrentPosition(this.getCurrentPositionCompleted);
    },

    // 現在地付近の駅を表示するフローティングアクションボタンをクリックしたとき
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

    // 位置情報の取得が完了したとき
    getCurrentPositionCompleted(pos) {
      this.$refs.mainMap.mapObject.panTo([pos.coords.latitude, pos.coords.longitude]);
      // this.$refs.mainMap.mapObject.setView(new L.LatLng(pos.coords.latitude,  pos.coords.longitude), this.zoom);
      // 移動後の場所にピンを立てようとしても移動前の場所になってしまう
      // this.onClickMyLocationIcon();
    },

    // テキストフィールドのクリアボタンを押したとき
    rideStationTextFieldCleared() {
      this.rideStationTextFieldModel = "";
      this.rideStationUnfix();
    },
    dropStationTextFieldCleared() {
      this.dropStationTextFieldModel = "";
      this.dropStationUnfix();
    },

    // swapボタンを押したとき
    swapTextField() {
      // フィールドの内容を交換
      this.dropStationTextFieldModel = [this.rideStationTextFieldModel, this.rideStationTextFieldModel = this.dropStationTextFieldModel][0];
      this.isDropStationFixed = [this.isRideStationFixed, this.isRideStationFixed = this.isDropStationFixed][0];
    },

    // その区間の路線数を調べる
    checkRoute() {
      console.log("ride:");
      console.log(this.rideStation);
      console.log("drop:");
      console.log(this.dropStation);
      var sameRailway = this.rideStation.lines.filter( function(d, index) {
        for (var r in this) {
          if( this[r].railway_name === d.railway_name ) return d;
        }
      }, this.dropStation.lines);
      if (sameRailway.length == 0) {
        this.rideRailway = [];
      } else {
        this.rideRailway = sameRailway;
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
    },
  },

  // このコンポーネントがマウントされたときに実行される処理
  mounted: function() {
    this.$nextTick(() => {
      // 初期位置・ズームの設定
      this.bounds = this.$refs.mainMap.mapObject.getBounds();
    });
    // 後の通信でCookieを使うようにする設定
    axios.defaults.withCredentials = true;
  },

  // 変数の監視処理
  watch: {

    // 登録/ログインフィールドの監視処理
    usernameModel() {
      this.signupFormUpdated();
      this.signinFormUpdated();
    },
    passwordModel() {
      this.signupFormUpdated();
      this.signinFormUpdated();
    },
    passwordConfirmModel() {
      this.signupFormUpdated();
    },

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