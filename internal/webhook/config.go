// config.go
package webhook

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func LoadConfigFromYAML(filePath string) (map[string][]HourBlock, error) {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config map[string][]HourBlock
	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		return nil, err
	}

	return config, nil
}
