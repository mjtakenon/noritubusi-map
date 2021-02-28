<template>
  <l-popup :options="{ autoPan: false, minWidth: 'auto' }" style="margin: 0;">
    <v-card elevation="0">
      <v-card-title class="font-weight-bold">
        {{ this.stationName }}
      </v-card-title>
      <v-card-text>
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
      </v-card-text>
    </v-card>
  </l-popup>
</template>

<script lang="ts">
import Vue, { PropType } from "vue"

import { LPopup } from "vue2-leaflet"

import { getStationsInRailway } from "../../utils/api/search"

import { Line } from "@/entities/Line"
import { Pin } from "@/entities/Pin"
import { Station } from "@/entities/Station"

export default Vue.extend({
  components: {
    LPopup,
  },
  props: {
    stationName: {
      type: String,
      required: true,
    },
    lines: {
      type: Array as PropType<Array<Line>>,
      required: true,
      default: () => [] as Array<Line>,
    },
  },
  data() {
    return {}
  },
  computed: {
    isStationFromFilled(): boolean {
      const station = this.$accessor.TripRecord.stationFrom
      return station.name !== ""
    },
  },
  methods: {
    onClickRide(line: Line) {
      this.$accessor.TripRecord.setStationFrom({
        name: this.stationName,
        id: line.stationId,
      })
      // TODO: IDをバックエンドから返してもらったら変える
      this.$accessor.TripRecord.setRailway({
        name: line.railwayName,
        id: 0,
      })
      this.$accessor.SuggestList.setBuildings([])
      this.setPinsOnRailway(line.railwayName)
    },
    onClickGetOff(line: Line) {
      this.$accessor.TripRecord.setStationTo({
        name: this.stationName,
        id: line.stationId,
      })
      // TODO: IDをバックエンドから返してもらったら変える
      this.$accessor.TripRecord.setRailway({
        name: line.railwayName,
        id: 0,
      })
      this.$accessor.SuggestList.setBuildings([])
    },
    onClickRailwayName(railwayName: string) {
      console.log(railwayName)
      this.setPinsOnRailway(railwayName)
    },
    setPinsOnRailway(railwayName: string) {
      getStationsInRailway(railwayName)
        .then(stations =>
          this.$accessor.Map.setPins(
            stations.map(
              (station: Station): Pin => ({
                latLng: {
                  lat: station.latitude,
                  lng: station.longitude,
                },
                popup: {
                  stationName: station.name,
                  lines:
                    station.name !== this.stationName
                      ? [station as Line]
                      : this.lines,
                },
                openPopup: station.railwayName === this.stationName,
                focusPin: false,
              })
            )
          )
        )
        .catch(error => console.error(error))
    },
  },
})
</script>
