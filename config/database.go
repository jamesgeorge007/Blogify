package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/mgo.v2"
	yaml "gopkg.in/yaml.v2"
)

type Configurations struct {
	DATABASE string `yaml:"DATABASE"`
	SERVER   string `yaml:"SERVER"`
}

var db *mgo.Database

func (c *Configurations) Connect() *mgo.Database {
	server, database := GetConnectionCredentials()

	session, err := mgo.Dial(server)

	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(database)

	return db
}

//Get connection creds
func GetConnectionCredentials() (string, string) {

	var c *Configurations
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
