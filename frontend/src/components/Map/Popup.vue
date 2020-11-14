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
          @click.stop="onClick"
          @mousedown.stop=""
          @touchstart.stop=""
          color="primary"
          outlined
        >
          乗車
        </v-btn>
        <v-btn
          @click.stop="onClick"
          @mousedown.stop=""
          @touchstart.stop=""
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
import { railways } from "../../utils/api/search.js"

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
  methods: {
    onClick() {
      console.log(`onClick`)
    },
    companyNameClicked(val) {
      console.log(val)
    },
    onClickRailwayName(railwayName) {
      console.log(railwayName)

      railways(railwayName)
        .then(response => {
          // here
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
