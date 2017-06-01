package conf

import (
	"io/ioutil"
	"github.com/go-yaml/yaml"
	"log"
)

func fetchConfig() (interface{}, error) {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var configuration map[string]interface{}
	err = yaml.Unmarshal(data, &configuration)
	if err != nil {
		return nil, err
	}

	return configuration, nil
}

func Get(param string) interface{}{
	config, err := fetchConfig()
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return config.(map[string]interface{})[param]
}
