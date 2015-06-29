package main

import (
    "encoding/json"
    "io/ioutil"
)

type Location struct {
    UUID      string
    Latitude  string
    Longitude string
    Message   string
}

func ToJson(locations []Location) []byte {
    b, _ := json.Marshal(locations)
    return b
}

func LoadLocations(filename string) []Location {
    locs := []Location{}
    in, err := ioutil.ReadFile(filename)
    if err == nil {
        json.Unmarshal(in, &locs)
    }
    return locs
}

func SaveLocations(filename string, locations []Location) {
    ioutil.WriteFile(filename, ToJson(locations), 0644)
}

func IndexOf(locations []Location, UUID string) int {
    for i, l := range locations {
        if l.UUID == UUID {
            return i
       }
    }
    return -1
}

func GetByUUID(locations []Location, UUID string) (Location, bool) {
    var location Location
    index := IndexOf(locations, UUID)
    if index > -1 {
        return locations[index], true
    }
    return location, false
}

func AppendOrReplace(locations []Location, location Location) []Location {
    index := IndexOf(locations, location.UUID)
    if index > -1 {
        locations[index] = location
    } else {
        locations = append(locations, location)
    }
    return locations
}

func DeleteByUUID(locations []Location, UUID string) []Location {
    i := IndexOf(locations, UUID)
    if i > -1 {
        locations = append(locations[:i], locations[i+1:]...)
    }
    return locations
}
