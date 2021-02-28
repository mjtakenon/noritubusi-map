export interface LatLng {
  lat: number
  lng: number
}

export function isLatLng(obj: any): obj is LatLng {
  return (
    obj !== null && typeof obj.lat === "number" && typeof obj.lng === "number"
  )
}
