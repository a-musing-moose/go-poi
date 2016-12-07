package main

import (
	"github.com/gorilla/mux"
	"github.com/twinj/uuid"
	"html/template"
	"net/http"
)

const (
	locationPath = "data/locations.json"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	locs := LoadLocations(locationPath)
	t, err := template.ParseFiles("index.html")
	if err == nil {
		t.Execute(w, locs)
	}
}

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	locs := LoadLocations(locationPath)
	//locs.Load(locationPath)
	if r.Method == "POST" {
		lat := r.FormValue("lat")
		long := r.FormValue("long")
		msg := r.FormValue("msg")
		id := uuid.NewV4()
		l := Location{
			UUID:      id.String(),
			Latitude:  lat,
			Longitude: long,
			Message:   msg}
		locs = AppendOrReplace(locs, l)
		SaveLocations(locationPath, locs)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		js := ToJson(locs)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	locs := LoadLocations(locationPath)
	//locs.Load(locationPath)
	vars := mux.Vars(r)
	id := vars["uuid"]
	if r.Method == "POST" {
		lat := r.FormValue("lat")
		long := r.FormValue("long")
		msg := r.FormValue("msg")
		loc, ok := GetByUUID(locs, id)
		if ok {
			loc.Latitude = lat
			loc.Longitude = long
			loc.Message = msg
			locs = AppendOrReplace(locs, loc)
			SaveLocations(locationPath, locs)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	} else if r.Method == "DELETE" {
		locs = DeleteByUUID(locs, id)
		SaveLocations(locationPath, locs)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/locations", LocationsHandler).Methods("GET", "POST")
	r.HandleFunc("/locations/{uuid}", LocationHandler).Methods("POST", "DELETE")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
