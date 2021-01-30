package planet

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/peterhellberg/swapi"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type queryPlanet struct {
	PlanetList []swapi.Planet `json:"results"`
}

//Planet represents the plannet data struct
type Planet struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Climate     string        `json:"climate" Gbson:"climate"`
	Terrain     string        `json:"terrain" bson:"terrain"`
	Appearances int           `json:"appearances" bson:"appearances"`
}

//SetID sets bson id into model
func (p *Planet) SetID() {
	p.ID = bson.NewObjectId()
}

func fetchDataFromAPI(name string) ([]string, error) {
	queryString := fmt.Sprintf("https://swapi.dev/api/planets/?search=%s", name)
	log.Println("trying to fetch: " + queryString)
	if response, err := http.Get(queryString); err != nil {
		return nil, err
	} else if response.StatusCode == http.StatusOK {
		gotPlanet := queryPlanet{}
		json.NewDecoder(response.Body).Decode(&gotPlanet)
		return gotPlanet.PlanetList[0].FilmURLs, nil

	} else {
		return nil, fmt.Errorf("could not fetch planet info")
	}
}

//SetApearences fetch the number of apearences
func (p *Planet) SetApearences() error {
	if moviesShown, err := fetchDataFromAPI(p.Name); err != nil {
		return err
	} else {
		p.Appearances = len(moviesShown)
		return nil
	}
}

//GETIDFormated get formated id
func (p *Planet) GETIDFormated() string {
	return strings.Split(p.ID.String(), "\"")[1]
}

//GetPlanets fetch all planets data
func GetPlanets(db *mgo.Session) ([]Planet, error) {
	log.Println("trying to fecth planets")
	planetList := []Planet{}
	if err := db.DB("b2w").C("planets").Find(nil).All(&planetList); err != nil {
		return nil, err
	} else {
		return planetList, nil
	}
}

//InsertPlanet Create a planet in database
func InsertPlanet(newPlanet Planet, db *mgo.Session) error {
	if err := db.DB("b2w").C("planets").Insert(&newPlanet); err != nil {
		return err
	} else {
		return nil
	}
}
