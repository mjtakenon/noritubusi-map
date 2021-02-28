export interface Line {
  stationId: number
  railwayName: string
  orderInRailway: number
}

export function isLine(obj: any): obj is Line {
  return (
    obj !== null &&
    typeof obj.railwayName === "string" &&
    typeof obj.stationId === "number" &&
    typeof obj.orderInRailway === "number"
  )
}
