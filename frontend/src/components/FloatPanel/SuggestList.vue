<template>
  <div v-show="buildings.length !== 0">
    <v-list two-line class="suggest-list">
      <v-list-item-group color="primary">
        <v-list-item
          v-for="(building, i) in this.buildings"
          :key="i"
          @click="onClickSuggestedBuilding(building)"
        >
          <v-list-item-avatar class="list-icon">
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

<script lang="ts">
import Vue from "vue"

import { SuggestBuilding } from "@/entities/SuggestBuilding"

type Data = {
  isOverflowns: Array<boolean>
}

export default Vue.extend({
  props: {},
  data(): Data {
    return {
      isOverflowns: [],
    }
  },
  created() {
    this.isOverflowns = [...Array(this.buildings.length).fill(false)]
  },
  watch: {
    buildings(_, oldBuildings: Array<SuggestBuilding>) {
      if (oldBuildings.length > 0) {
        const scrollables = this.$refs.scrollables as Array<HTMLSpanElement>
        this.isOverflowns = scrollables.map((scrollable: HTMLSpanElement) => {
          const parent = scrollable.parentNode
          if (parent !== null && parent instanceof Element) {
            return scrollable.clientWidth > parent.clientWidth
          }
          return false
        })
      }
    },
  },
  // メソッド
  methods: {
    railwayNameJoinWithComma(building: SuggestBuilding) {
      return building.lines.map(line => line.railwayName).join(", ")
    },
    onClickSuggestedBuilding(building: SuggestBuilding) {
      this.$accessor.Map.setPinAndFocus(building)
    },
  },
  // 算出プロパティ
  computed: {
    buildings(): Array<SuggestBuilding> {
      return this.$accessor.SuggestList.buildings
    },
  },
})
</script>

<style lang="scss" scoped>
.suggest-list {
  width: 100%;
  max-height: 400px;
  overflow-y: auto;
}
.list-icon {
  font-size: 150%;
}
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
