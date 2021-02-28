import { isLine, Line } from "./Line"

export interface Station extends Line {
  buildingId: number
  name: string
  latitude: number
  longitude: number
}

export function isStation(obj: any): obj is Station {
  return (
    isLine(obj) &&
    typeof (obj as Station).buildingId === "number" &&
    typeof (obj as Station).latitude === "number" &&
    typeof (obj as Station).longitude === "number"
  )
}
