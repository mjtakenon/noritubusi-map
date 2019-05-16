<style scoped>
</style>

<template>
  <v-list>
    <v-list-tile>
      <v-list-tile-avatar>
        <v-icon x-large>portrait</v-icon>
      </v-list-tile-avatar>
      <v-list-tile-content>
        <v-list-tile-title v-if="isLoggedIn">{{ userInfo.name }}</v-list-tile-title>
        <v-list-tile-title v-else>未ログイン</v-list-tile-title>
      </v-list-tile-content>
    </v-list-tile>
  </v-list>
</template>

<script>
export default {
  components: {},
  props: {
    isLoggedIn: {
      type: Boolean,
      required: true
    },
    userInfo: {
      type: Object,
      required: true
    }
  },
  data() {
    return {};
  }
};
</script>
