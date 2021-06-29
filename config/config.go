package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

//Configuration contains configuration parsed from the configuration file
type Configuration struct {
	//var_name "type " `yaml:"paramName"`

}

//GetConf  retun config from external file
func (c *Configuration) GetConf() *Configuration {

	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatal("Can't load config file: ", err)
	}
	return c

}
