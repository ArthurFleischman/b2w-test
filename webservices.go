package main

import (
	"b2w-test/planet"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func getPlanets(c *gin.Context) {
	if planetList, err := planet.GetPlanets(_db); err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "could not fetch planets")
	} else {
		log.Printf("%s received planet list", c.ClientIP())
		c.JSON(http.StatusOK, planetList)
	}
}
func insertPlanet(c *gin.Context) {
	newPlanet := planet.Planet{}
	c.BindJSON(&newPlanet)
	//set bson id object
	newPlanet.SetID()
	//set number of apearences
	newPlanet.SetApearences()

	log.Printf("trying to insert planet %s\n", newPlanet.GETIDFormated())
	//try to insert planet
	if err := planet.InsertPlanet(newPlanet, _db); err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "could not insert planet")
	} else {
		responseMessage := fmt.Sprintf("planet id: %v inserted successfully", newPlanet.GETIDFormated())
		log.Printf(responseMessage)
		c.String(http.StatusCreated, responseMessage)
	}
}

//StartWebService : start API service
func startWebService() error {
	r := gin.New()
	//GET
	r.GET("/planets", getPlanets)
	//POST
	creationPath := r.Group("/new")
	creationPath.POST("/planet", insertPlanet)
	//REMOVE

	return r.Run()
}

//#######################################
//############## DB #####################
//#######################################
//StartDBConnection start db client
func startDBConnection() error {
	client, err := mgo.Dial(url)
	if err != nil {
		return err
	}
	_db = client
	return nil
}

//#######################################
//############## DETAILS ################
//#######################################
func printDetails() {
	//put here all printing specifications
	log.Println("system started")
}
