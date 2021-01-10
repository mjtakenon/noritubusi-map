<template>
  <div class="float-panel">
    <Ticket />
    <SearchBox />
  </div>
</template>

<script>
import SearchBox from "./SearchBox.vue"
import Ticket from "./Ticket.vue"

export default {
  name: "FloatPanel",
  components: {
    SearchBox,
    Ticket,
  },
}
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
