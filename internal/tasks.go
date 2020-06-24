package internal

import (
  "log"
  "os/exec"
  "strings"
)

type Task struct {
  Type string
  Action string `yaml:",omitempty"`
  Script []string `yaml:",omitempty"`
}

func ValidateTask(task Task) (bool, string) {
  if task.Type != "shell" && task.Script != nil {
    log.Fatal("Parameter 'script' can only be used with the 'shell' type.")
  }

  return true, "Valid"
}

func RunTask(task Task) {
  switch task.Type {
    case "shell":
      shellTask(task)

    default:
      return
  }
}

func shellTask(task Task) {
  for _, cmd := range task.Script {
    split := strings.Fields(cmd)
    command := exec.Command(split[0], strings.Join(split[1:], " "))
    err := command.Run()
    if err != nil {
      log.Fatal("Something went wrong.\n" + err.Error())
    }
  }
}
