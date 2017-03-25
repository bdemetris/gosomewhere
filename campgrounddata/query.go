package campgrounddata

import (
  "encoding/xml"
  "net/http"
  "fmt"
)

func QueryCampground(amenityID string) (ResultSet, error) {
  apikey := LoadConfig("./campgrounddata/config.json")

    url := fmt.Sprintf("http://api.amp.active.com/camping/campgrounds?pstate=CA&amenity=%s&api_key=%s", amenityID, apikey.Key)

    fmt.Println(url)

    resp, err := http.Get(url)

    // responseData,err := ioutil.ReadAll(resp.Body)
    // if err != nil {
    //   log.Fatal(err)
    // }
    //
    // responseString := string(responseData)
    //
    // fmt.Println(responseString)

    if err != nil {
        return ResultSet{}, err
    }

    defer resp.Body.Close()

    var d ResultSet

    if err := xml.NewDecoder(resp.Body).Decode(&d); err != nil {
        return ResultSet{}, err
    }

    return d, nil
}

type ResultSet struct {
  	XMLName   xml.Name   `xml:"resultset"`
	AmenityID int `xml:"amenity,attr"`
	Count	int `xml:"count,attr"`
	State	string `xml:"pstate,attr"`
	Type	string	`xml:"resultType,attr"`
  	Results []Result `xml:"result"`
}

type Result struct {
	XMLName		xml.Name	`xml:"result"`
	FacilityID 	int 		`xml:"facilityID,attr"`
	FacilityName 	string 		`xml:"facilityName,attr"`
	FacilityPhoto 	string 		`xml:"faciltyPhoto,attr"`
	Latitude 	float64 	`xml:"latitude,attr"`
	Longitude 	float64 	`xml:"longitude,attr"`
	State 		string		`xml:"state,attr"`
}
