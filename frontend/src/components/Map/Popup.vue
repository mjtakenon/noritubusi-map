<template>
  <l-popup>
    <H1>{{ this.data.name }}</H1>
    <v-list dense>
      <v-list-item
        v-for="(l, idx) in this.data.lines"
        :key="idx"
        @click="companyNameClicked(idx)"
      >
        <v-list-item-avatar>
          <H2>{{ icon }}</H2>
        </v-list-item-avatar>
        <v-list-item-content>
          <v-list-item-title>{{
            data.lines[idx].railway_name
          }}</v-list-item-title>
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
    data: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      icon: "",
    }
  },
  methods: {
    companyNameClicked(val) {
      console.log(val)
    },
  },
  mounted: function() {
    this.$nextTick(function() {
      this.icon = "🚊"
    })
  },
}
</script>
