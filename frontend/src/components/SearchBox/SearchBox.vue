<template>
  <div v-bind:style="style.searchBox">
    <v-toolbar dense>
      <v-btn icon>
        <v-icon>fa-bars</v-icon>
      </v-btn>

      <v-text-field
        hide-details
        append-icon="search"
        single-line
        v-bind:style="style.textField"
      ></v-text-field>

      <!-- <v-btn icon>
        <v-icon>my_location</v-icon>
      </v-btn> -->
    </v-toolbar>
  </div>
</template>

<script>
// 各 Vue コンポーネントのインポート

export default {
  data() {
    return {
      style: {
        searchBox: {
          // 適度にmarginを取ってfloatingさせる
          position: "absolute",
          padding: "10px",
        },
        textField: {
          // 上に謎のスペースがあるから消す
          padding: "0px",
        },
      },
    }
  },
}
</script>
