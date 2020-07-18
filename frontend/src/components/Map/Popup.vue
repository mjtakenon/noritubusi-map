<template>
  <l-popup>
    <h2>{{ this.name }}</h2>
    <v-divider class="my-1"></v-divider>
    <v-list dense style="max-height: 200px; overflow-y: auto;" class="py-0">
      <v-list-item
        v-for="(line, idx) in this.lines"
        :key="idx"
        @click="companyNameClicked(idx)"
        class="px-1"
      >
        <v-list-item-avatar class="my-0">
          <!-- <H2>{{ icon }}</H2> -->
          <v-icon>train</v-icon>
        </v-list-item-avatar>
        <v-list-item-content class="py-0">
          <v-list-item-title>{{line.railway_name}}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list>
  </l-popup>
</template>

<script>
import { LPopup } from "vue2-leaflet"

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
        return lines.every(
          line => line.hasOwnProperty("railway_name")
        )
      },
    },
  },
  data() {
    return {}
  },
  methods: {
    companyNameClicked(val) {
      console.log(val)
    },
  },
}
</script>
