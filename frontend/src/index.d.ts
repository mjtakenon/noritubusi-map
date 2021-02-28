import Vue from "vue"
import { accessor } from "./vuex"

declare module "vue/types/vue" {
  interface Vue {
    $accessor: typeof accessor
  }
}
