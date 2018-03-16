package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Results struct {
	Deck                string
	Name                string
	OriginalReleaseDate string `json:"original_release_date"`
}

type GiantBombSearchResponse struct {
	Results []Results
}

// This function must be named ServeHTTP() and these argument methods must be in this order and *http.Request must be a reference
// This is a method of the Hello stuct
// func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	//fmt.Fprint prints to a writer object, which has been passed in with the 'w'
// 	fmt.Fprint(w, "<h1> Hello from the Go web server!<h1>")
// }

func main() {
	// 	var h Hello
	// 	var key string = "f10f942d8c6c6a04af3c3774e257daa795c10589"
	// 	var url string = "http://www.giantbomb.com/api/game/3030-4725/?api_key=f10f942d8c6c6a04af3c3774e257daa795c10589s"

	// 	// ListenAndServce takes a url and port number and the second is an instance of the Hello object...
	// 	// ... This is where the method will search and call the ServeHTTP method
	// 	err := http.ListenAndServe("localhost:5050", h)
	// 	checkError(err)

	// 	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request ) {
	// 		var result
	// 	})
	// }

	// func checkError(err error) {
	// 	if err != nil {
	// 		panic(err)
	// 	}

	results, _ := search("http://www.giantbomb.com/api/game/3030-4725/?api_key=f10f942d8c6c6a04af3c3774e257daa795c10589&format=json")
	fmt.Printf("Results : %+v", results)
	// fmt.Print(content)
	// games := gamesFromJSON(results)
	// fmt.Print(games)

}

// func search(query string) ([]) {

// }

func giantBombAPI(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	// fmt.Printf("Giantbomb response type: %T\n", resp)

	defer resp.Body.Close()
	// bytes, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	// return []byte{}, err
	// 	panic(err)
	// }

	// return string(bytes)
	return ioutil.ReadAll(resp.Body)

}

// You are getting an empty slice when []Results, but if it is not a slice you get the correct outcome.
func search(query string) ([]Results, error) {
	var g GiantBombSearchResponse
	body, err := giantBombAPI(query)
	if err != nil {
		return []Results{}, err
	}
	err = json.Unmarshal(body, &g)
	return g.Results, err

}

// func gamesFromJSON(content string) []Game {
// 	games := make([]Game, 0, 20)
// 	// use the a json decoder to decode a string object that is read by the string reader
// 	fmt.Print("From fameFromJson: \n", content)
// 	decoder := json.NewDecoder(strings.NewReader(content))
// 	// This removes the array brackets that wrap the json results
// 	// This appear to only be visible if there is more than one json object
// 	// Will this error if there is no array brackets to remove?
// 	// _, err := decoder.Token()
// 	// if err != nil {
// 	// 	// return []byte{}, err
// 	// 	panic(err)
// 	// }
// 	var game Game
// 	fmt.Print("The decode prior to decoder.More() \n", decoder)
// 	for decoder.More() {
// 		// This parses the json and pull only the fields that match the ones in my stuct
// 		err := decoder.Decode(&game)
// 		if err != nil {
// 			// return []byte{}, err
// 			panic(err)
// 		}
// 		games = append(games, game)
// 	}
// 	return games
// }

// func anotherGamesJSON(content string) []Game {

// }
