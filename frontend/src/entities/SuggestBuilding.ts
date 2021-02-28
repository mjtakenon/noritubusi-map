import { Line, isLine } from "./Line"

export interface SuggestBuilding {
  buildingId: number
  name: string
  latitude: number
  longitude: number
  lines: Array<Line>
}

export function isSuggestBuilding(obj: any): obj is SuggestBuilding {
  console.log(obj)
  return (
    obj !== null &&
    typeof obj === "object" &&
    typeof obj.buildingId === "number" &&
    typeof obj.name === "string" &&
    typeof obj.latitude === "number" &&
    typeof obj.longitude === "number" &&
    Array.isArray(obj.lines) &&
    (obj.lines as Array<Line>).every(line => isLine(line))
  )
}
