<template>
  <div class="float-panel">
    <Ticket />
    <SearchBox />
  </div>
</template>

<script lang="ts">
import Vue from "vue"

import SearchBox from "@/components/FloatPanel/SearchBox.vue"
import Ticket from "@/components/FloatPanel/Ticket.vue"

export default Vue.extend({
  name: "FloatPanel",
  components: {
    SearchBox,
    Ticket,
  },
})
</script>

<style lang="scss" scoped>
.float-panel {
  position: absolute;
  padding: 10px;
  display: flex;
  flex-direction: column;
  width: 400px;
  & > * {
    margin-bottom: 16px;
  }
  &:last-child {
    margin-bottom: 0;
  }
}
</style>
