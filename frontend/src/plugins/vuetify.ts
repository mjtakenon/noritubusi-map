import Vue from "vue"
import Vuetify from "vuetify/lib/framework"

Vue.use(Vuetify)

export default new Vuetify({
  icons: {
    iconfont: "md",
  },
  themes: {
    light: {
      primary: "#378640",
    },
  },
})
