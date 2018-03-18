package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var key string = "api_key=f10f942d8c6c6a04af3c3774e257daa795c10589"
var searchURL string = "http://www.giantbomb.com/api/search/?"

type ResultsGame struct {
	APIDetailURL           string      `json:"api_detail_url"`
	BoxArtImages           BoxArtImage `json:"image"`
	Deck                   string
	ExpectedReleaseDay     string       `json:"expected_release_day"`
	ExpectedReleaseMonth   string       `json:"expected_release_month"`
	ExpectedReleaseQuarter string       `json:"expected_release_quarter"`
	ExpectedReleaseYear    string       `json:"expected_release_year"`
	GameRatings            []GameRating `json:"original_game_rating"`
	Name                   string
	OriginalReleaseDate    string     `json:"original_release_date"`
	Platfoms               []Platform `json:"platforms"`
	SiteDetailURL          string     `json:"site_detail_url"`
}

type GameRating struct {
	ID   int
	Name string
}

type Platform struct {
	APIDetailURL  string `json:"api_detail_url"`
	Name          string
	SiteDetailURL string `json:"site_detail_url"`
}

type BoxArtImage struct {
	IconBoxArt  string `json:"icon_url"`
	ThumbBoxArt string `json:"thumb_url"`
	LargeboxArt string `json:"super_url"`
}

type GiantBombSearchResponse struct {
	Results []ResultsGame
}

// This function must be named ServeHTTP() and these argument methods must be in this order and *http.Request must be a reference
// This is a method of the Hello stuct
// func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	//fmt.Fprint prints to a writer object, which has been passed in with the 'w'
// 	fmt.Fprint(w, "<h1> Hello from the Go web server!<h1>")
// }

func main() {
	http.HandleFunc("/", GetInit)
	http.HandleFunc("/search", GetSearchResults)

	log.Fatal(http.ListenAndServe(":5050", nil))
}

func giantBombAPI(url string) ([]byte, error) {
	// Get request to Giant Bomb
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	// close connectiong to response to ensure no resource leaks
	defer resp.Body.Close()
	// read data from response body and return a slice of bytes
	return ioutil.ReadAll(resp.Body)

}

func search(query string) ([]ResultsGame, error) {
	var g GiantBombSearchResponse
	// fmt.Println(searchURL + key + url.QueryEscape(query))
	body, err := giantBombAPI(searchURL + key + query)
	if err != nil {
		return []ResultsGame{}, err
	}
	err = json.Unmarshal(body, &g)
	fmt.Println("From search func: ", g.Results)
	return g.Results, err

}

/////////////////////// ROUTES! ////////////////////
func GetInit(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "We runnin'.")
}

func GetSearchResults(writer http.ResponseWriter, req *http.Request) {
	var results []ResultsGame
	var err error
	fmt.Println(searchURL + key + url.QueryEscape("&format=json&query='resident evil'&resources=game"))
	fmt.Println(searchURL + key + "&format=json&query=%27resident%20evil%27&resources=game")
	results, err = search("&format=json&query=%27resident%20evil%27&resources=game")
	checkRequestError(writer, err)
	fmt.Println("From GetSearchResults: ", results)
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(results)
	checkRequestError(writer, err)
}

func checkRequestError(writer http.ResponseWriter, err error) {
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
