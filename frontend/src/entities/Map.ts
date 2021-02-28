import { LatLng } from "./Common"

export interface Bounds {
  _southWest: LatLng
  _northEast: LatLng
}

export interface MapOptions {
  zoom: number
  center: LatLng
  bounds: Bounds
}
