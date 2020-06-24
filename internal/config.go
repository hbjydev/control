package internal

import (
  "io/ioutil"
  "log"
  "errors"

  "gopkg.in/yaml.v2"
)

type Config struct {
  Version string
  Up []Task
}

func ReadConfig() string {
  file, err := ioutil.ReadFile("ctrl.yml")
  if err != nil {
    log.Fatal(err)
  }
  return string(file)
}

func ParseConfig(rawConfig string) (Config, error) {
  config := Config{}

  err := yaml.Unmarshal([]byte(rawConfig), &config)
  if err != nil {
    return config, err
  }

  for _,task := range config.Up {
    valid, cause := ValidateTask(task)
    if valid != true {
      return config, errors.New(cause)
    }
  }

  return config, nil
}
