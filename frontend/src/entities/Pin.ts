import { LatLng } from "./Common"
import { isLine, Line } from "./Line"

export interface Popup {
  stationName: string
  lines: Array<Line>
}

export interface Pin {
  latLng: LatLng
  popup: Popup
  openPopup: boolean
  focusPin: boolean
}

export function isPopup(obj: any): obj is Popup {
  return (
    obj !== null &&
    typeof obj.name == "string" &&
    Array.isArray(obj.lines) &&
    (obj.lines as Array<Line>).every(line => isLine(line))
  )
}
