<template>
  <l-popup> 
    ID:{{ this.data.id }} <br>
    <H1>{{ this.data.name }} </H1>
    <!-- <H2> 🚊  {{ this.data.companyName }} : {{ this.data.railwayName }} </H2>  -->
    <v-list>
      <v-list-tile
        v-for="(c, idx) in this.data.companyName"
        :key="idx"
        @click="companyNameClicked(idx)">
        <v-list-tile-avatar>
          <H1>{{icon}}</H1> <!-- <v-icon :class="[item.iconClass]">{{ item.icon }}</v-icon> -->
        </v-list-tile-avatar>
        <v-list-tile-content>
          <v-list-tile-title> {{ data.railwayName[idx] }}</v-list-tile-title>
          <v-list-tile-sub-title> {{ data.companyName[idx] }} </v-list-tile-sub-title>
        </v-list-tile-content>
      </v-list-tile>
    </v-list>
  </l-popup>
</template>

<script>
import { LPopup } from "vue2-leaflet";

export default {
  components: {
    LPopup
  },
  props: {
    data: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      icon: ""
    };
  },
  methods: {
      companyNameClicked(val) {
        console.log(val);
    }
  },
  mounted: function () {
  this.$nextTick(function () {
    // アイコンの指定
    if(this.data.railwayName[0].indexOf("新幹線") !== -1) { // 新幹線
      this.icon="🚆"
    } else { // それ以外
      this.icon="🚊"
    }
  })}
};
</script>
