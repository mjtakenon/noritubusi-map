<template>
  <l-popup :options="{ autoPan: false, minWidth: '500px' }">
    <h2>{{ this.name }}</h2>
    <v-divider class="my-1"></v-divider>
    <v-list dense style="max-height: 200px; overflow-y: auto;" class="py-0">
      <v-list-item
        v-for="(line, idx) in this.lines"
        :key="idx"
        @click="onClickRailwayName(line.railwayName)"
        class="px-1"
      >
        <v-list-item-avatar class="my-0 mr-0">
          <!-- <H2>{{ icon }}</H2> -->
          <v-icon>train</v-icon>
        </v-list-item-avatar>
        <v-list-item-content class="py-0 px-2">
          <v-list-item-title>{{ line.railwayName }}</v-list-item-title>
        </v-list-item-content>
        <v-btn
          @click.stop="onClickRide(line)"
          @mousedown.stop=""
          @touchstart.stop=""
          color="primary"
          outlined
        >
          乗車
        </v-btn>
        <v-btn
          @click.stop="onClickGetOff(line)"
          @mousedown.stop=""
          @touchstart.stop=""
          :disabled="!isStationFromFilled"
          color="error"
          outlined
        >
          降車
        </v-btn>
      </v-list-item>
    </v-list>
  </l-popup>
</template>

<script>
import { LPopup } from "vue2-leaflet"
import { getStationsInRailway } from "../../utils/api/search.js"

export default {
  components: {
    LPopup,
  },
  props: {
    name: {
      type: String,
      required: true,
    },
    lines: {
      type: Array,
      required: true,
      validator(lines) {
        return lines.every(line => line.hasOwnProperty("railwayName"))
      },
    },
  },
  data() {
    return {}
  },
  computed: {
    isStationFromFilled() {
      const station = this.$store.getters["TripRecord/stationFrom"]
      return station.name !== ""
    },
  },
  methods: {
    onClickRide(line) {
      this.$store.commit("TripRecord/stationFrom", {
        name: this.name,
        id: line.stationId,
      })
      // TODO: IDをバックエンドから返してもらったら変える
      this.$store.commit("TripRecord/railway", {
        name: line.railwayName,
        id: 0,
      })
      this.$store.commit("SuggestList/buildings", [])
      this.setPinsOnRailway(line.railwayName)
    },
    onClickGetOff(line) {
      this.$store.commit("TripRecord/stationTo", {
        name: this.name,
        id: line.stationId,
      })
      // TODO: IDをバックエンドから返してもらったら変える
      this.$store.commit("TripRecord/railway", {
        name: line.railwayName,
        id: 0,
      })
      this.$store.commit("SuggestList/buildings", [])
    },
    companyNameClicked(val) {
      console.log(val)
    },
    onClickRailwayName(railwayName) {
      console.log(railwayName)
      this.setPinsOnRailway(railwayName)
    },
    setPinsOnRailway(railwayName) {
      getStationsInRailway(railwayName)
        .then(response => {
          const pins = response.data.map(building => ({
            latLng: [building.latitude, building.longitude],
            popup: {
              name: building.name,
              lines: building.name !== this.name ? [] : this.lines,
            },
            openPopup: building.name === this.name,
          }))

          this.$store.dispatch("Map/setPins", pins)
        })
        .catch(error => console.error(error))
    },
  },
}
</script>
