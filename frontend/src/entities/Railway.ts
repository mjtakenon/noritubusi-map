import { isStation, Station } from "./Station"

export interface Railway extends Station {
  railwayName: string
}

export function isRailway(obj: any): obj is Railway {
  return isStation(obj) && typeof (obj as Railway).railwayName === "string"
}
