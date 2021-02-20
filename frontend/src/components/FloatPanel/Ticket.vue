<template>
  <v-card>
    <v-card-text>
      <div class="wrapper font-weight-bold">
        <div class="station-name">
          <span class="text--primary">{{ stationFrom.name }}</span>
        </div>
        <v-icon class="arrow">arrow_right</v-icon>
        <div class="station-name">
          <span class="text--primary">{{ stationTo.name }}</span>
        </div>
      </div>

      <div v-if="isFilledEitherStations">
        <v-col class="d-flex" v-if="true">
          <!-- <v-select
            :items="stationFrom.lines.map(e => e.railwayName)"
            label="路線"
          ></v-select> -->
          <v-select
            item-text="railwayName"
            item-value="railwayName"
            :items="matchedRailways"
            label="路線"
            :value="railway.name"
          />
        </v-col>
        <v-alert dense outlined type="error" v-else>
          共通する路線が存在しません
        </v-alert>
      </div>

      <!-- <div>{{ railway.name }}</div> -->
      <div class="registration" v-if="isFilledBothStations">
        <v-btn primary block outlined>登録</v-btn>
      </div>
    </v-card-text>
  </v-card>
</template>

<script>
import _ from "lodash"

export default {
  name: "Ticket",
  computed: {
    stationFrom() {
      return this.$store.getters["TripRecord/stationFrom"]
    },
    stationTo() {
      return this.$store.getters["TripRecord/stationTo"]
    },
    matchedRailways() {
      if (this.isFilledBothStations) {
        console.log(this.stationFrom.lines, this.stationTo.lines)
        return _.intersectionBy(
          this.stationFrom.lines,
          this.stationTo.lines,
          "railwayName"
        )
        // const lineNamesOfStationTo = this.stationTo.lines.map(line => line.railwayName)
        // return this.stationFrom.lines.filter(line =>
        //   lineNamesOfStationToStationTo.includes(line.railwayName)
        // )
      } else {
        if (this.stationFrom.name !== "") {
          return this.stationFrom.lines
        } else if (this.stationTo.name !== "") {
          return this.stationTo.lines
        }
      }
      return Array()
    },
    railway() {
      return this.$store.getters["TripRecord/railway"]
    },
    isFilledBothStations() {
      return this.stationFrom.name !== "" && this.stationTo.name !== ""
    },
    isFilledEitherStations() {
      return this.stationFrom.name !== "" || this.stationTo.name !== ""
    },
  },
}
</script>

<style lang="scss" scoped>
.wrapper {
  display: flex;
  width: 100%;
  height: 2rem;
  flex-direction: row;
  align-content: center;
}
.station-name {
  width: 50%;
  font-size: 1.5rem;
  display: flex;
  align-items: center;
  // v-card-titleのフォントサイズが取れず
  // cardviewのサイズが変わるため決め打ちで指定
}
.arrow {
  width: 10%;
}
.registration {
  width: 100%;
  margin-top: 8px;
}
</style>
