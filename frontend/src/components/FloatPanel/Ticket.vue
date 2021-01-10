<template>
  <v-card>
    <v-card-title>
      <div class="wrapper">
        <div class="station">
          <span>{{ stationFrom.name }}</span>
        </div>
        <v-icon class="arrow">arrow_right</v-icon>
        <div class="station">
          <span>{{ stationTo.name }}</span>
        </div>
      </div>
    </v-card-title>
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
  },
}
</script>

<style lang="scss" scoped>
.wrapper {
  display: flex;
  width: 100%;
  flex-direction: row;
  align-content: center;
}
.station {
  width: 50%;
  display: flex;
  align-items: center;
}
.arrow {
  width: 10%;
}
</style>
