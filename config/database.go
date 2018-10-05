package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/mgo.v2"
	yaml "gopkg.in/yaml.v2"
)

type configurations struct {
	DATABASE string `yaml:"DATABASE"`
	SERVER   string `yaml:"SERVER"`
	PORT     int    `yaml:"PORT"`
}

var db *mgo.Database

func (c *configurations) Connect() {
	server, database := GetConnectionCredentials()

	session, err := mgo.Dial(server)

	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(database)
}

//Get connection creds
func GetConnectionCredentials() (string, string) {

	var c *configurations
	//Obtain values from yaml config
	yamlFile, err := ioutil.ReadFile("app.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	fmt.Println("DB values are", c)
	return c.SERVER, c.DATABASE

}
