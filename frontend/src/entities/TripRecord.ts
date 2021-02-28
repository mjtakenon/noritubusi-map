export interface RecordValue {
  id: number
  name: string
}

export interface TripRecord {
  stationFrom: RecordValue
  stationTo: RecordValue
  railway: RecordValue
}
