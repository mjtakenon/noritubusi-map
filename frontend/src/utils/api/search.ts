import { Station, isStation } from "@/entities/Station"
import { SuggestBuilding, isSuggestBuilding } from "@/entities/SuggestBuilding"
import axiosBase from "axios"
const axios = axiosBase.create({
  baseURL: "http://localhost:1323",
  headers: {
    "Content-Type": "application/json",
    "X-Requested-With": "XMLHttpRequest",
  },
  responseType: "json",
})

export async function searchBuildingsByKeyword(
  keyword: string
): Promise<Array<SuggestBuilding>> {
  // TODO: リクエスト間隔を空けるような実装に (例: 100ms ごと)

  return await axios
    .get("/buildings/suggest", {
      params: { keyword },
      headers: { "Content-Type": "application/x-www-form-urlencoded" },
    })
    .then(response => {
      // 200 OK なら成功なのでレスポンスを返す それ以外なら例外を投げる
      if (response.status === 200) {
        const responseData = response.data as Array<any>
        if (responseData.every(elem => isSuggestBuilding(elem))) {
          return responseData as Array<SuggestBuilding>
        } else {
          console.error("Invalid response data structure")
          console.debug(response.data)
          throw response
        }
      } else {
        throw response
      }
    })
    .catch(error => {
      throw error.response
    })
}

export async function getStationsInRailway(
  railwayName: string
): Promise<Array<Station>> {
  // TODO: リクエスト間隔を空けるような実装に (例: 100ms ごと)

  return await axios
    .get(`/railways/${railwayName}`)
    .then(response => {
      // 200 OK なら成功なのでレスポンスを返す それ以外なら例外を投げる
      if (response.status === 200) {
        const responseData = response.data as Array<any>
        if (responseData.every(elem => isStation(elem))) {
          return responseData as Array<Station>
        } else {
          console.error("Invalid response data structure")
          throw response
        }
      } else {
        throw response
      }
    })
    .catch(error => {
      throw error.response
    })
}
