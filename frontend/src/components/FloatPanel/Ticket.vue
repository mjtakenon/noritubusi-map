<template>
  <v-card>
    <v-card-text>
      <div class="wrapper font-weight-bold">
        <div class="station">
          <span class="text--primary">{{ stationFrom.name }}</span>
        </div>
        <v-icon class="arrow">arrow_right</v-icon>
        <div class="station">
          <span class="text--primary">{{ stationTo.name }}</span>
        </div>
      </div>
      <div class="registration" v-if="true">
        <v-btn primary block outlined>登録</v-btn>
      </div>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: "Ticket",
  computed: {
    stationFrom() {
      return this.$store.getters["TripRecord/stationFrom"]
    },
    stationTo() {
      return this.$store.getters["TripRecord/stationTo"]
    },
    isFilledBothStations() {
      return this.stationFrom.name !== "" && this.stationTo.name !== ""
    },
  },
}
</script>

<style lang="scss" scoped>
.wrapper {
  display: flex;
  width: 100%;
  height: 2rem;
  font-size: 1.5rem;
  flex-direction: row;
  align-content: center;
}
.station {
  width: 50%;
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
