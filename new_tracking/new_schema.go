package new_tracking

import (
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"googlemaps.github.io/maps"
)

// MapPoint represent a point on the map, each point is associated with a feature ID, which points to a feature
// MapPoint/loc will be used as 2Dsphere index
type MapPoint struct {
	Loc       maps.LatLng `json:"loc"`
	FeatureID string      `json:"featureID"`
}

// LocationInfo represent the information associated with a Geospatial point
type LocationInfo struct {
	ID          string `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Order       int    `json:"order"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Enabled     bool   `json:"enabled"`
}

type TimeStamp struct {
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type Style struct {
	Color string `json:"color,omitempty"`
	Width int    `json:"width,omitempty"`
}

// Stop indicates where a tracked object is scheduled to arrive
type Stop struct {
	ID    string       `json:"id"             bson:"id"`
	Info  LocationInfo `json:"info"`
	Style Style        `json:"style"`
}

func (App *DBCollection) GetRouteByID(id string){

}
func (App *App) RoutesHandler(w http.ResponseWriter, r *http.Request) {
	// Find all routes in database
	var routes []Route
	// Query all routes information
	err := App.RoutesInfo.Find(bson.M{}).All(&routes)
	for _, r := routes{
		r.
	}
	// Handle query errors
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Send each route to client as JSON
	WriteJSON(w, routes)
}
