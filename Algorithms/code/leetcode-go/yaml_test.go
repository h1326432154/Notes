package leetcode

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v2"
)

type Yaml struct {
	Redis []*RedisOpt `yaml:"redis"`
}

type RedisOpt struct {
	Host string `yaml:"Host"`
	PWD  string `yaml:"Password"`
}

func TestYaml(t *testing.T) {
	f := filepath.Join("./example.yaml")
	yamlFile, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	c := Yaml{}
	if err := yaml.Unmarshal(yamlFile, &c); err != nil {
		panic(err)
	}
	t.Log(c)
}
