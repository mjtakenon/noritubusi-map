<template>
  <div>
    <v-list
      two-line
      style="max-height: 400px; max-width: 300px; overflow-y: auto;"
    >
      <v-list-item-group color="primary">
        <v-list-item
          v-for="(building, i) in this.buildings"
          :key="i"
          @click="onClickSuggestedBuilding(building)"
        >
          <v-list-item-avatar style="font-size: 150%">
            <v-icon>train</v-icon>
          </v-list-item-avatar>
          <v-list-item-content>
            <v-list-item-title> {{ building.name }} </v-list-item-title>
            <v-list-item-subtitle class="scrollable">
              <span :class="{ scroll: isOverflowns[i] }" ref="scrollables">{{
                railwayNameJoinWithComma(building)
              }}</span>
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </v-list-item-group>
    </v-list>
  </div>
</template>

<script>
import { suggest } from "../../utils/api/search.js"

export default {
  props: {
    buildings: {
      type: Array,
      required: true,
    },
  },
  data() {
    return {
      isOverflowns: [],
    }
  },
  created() {
    this.isOverflowns = [...Array(this.buildings.length).fill(false)]
  },
  watch: {
    buildings(newBuildings, oldBuildings) {
      if (oldBuildings.length > 0) {
        this.isOverflowns = this.$refs.scrollables.map(scrollable => {
          const parent = scrollable.parentNode
          return scrollable.clientWidth > parent.clientWidth
        })
      }
    },
  },
  // メソッド
  methods: {
    railwayNameJoinWithComma(building) {
      return building.lines.map(line => line.railway_name).join(", ")
    },
    setPin(building) {
      this.$store.dispatch("Map/setPins", [
        {
          latLng: [building.latitude, building.longitude],
          popup: {
            name: building.name,
            lines: building.lines,
          },
          autoOpenPopup: true,
        },
      ])
      return
    },
    onClickSuggestedBuilding(building) {
      this.setPin(building)
      return
    },
  },
  // 算出プロパティ
  computed: {},
}
</script>

<style lang="scss" scoped>
.scrollable {
  overflow: hidden;
  & > span {
    display: inline-block;
  }
}

span.scroll {
  padding-left: 100%;
  white-space: nowrap;
  animation: scroll 15s linear infinite;
}

@keyframes scroll {
  0% {
    transform: translateX(0);
  }
  100% {
    transform: translateX(-100%);
  }
}
</style>
