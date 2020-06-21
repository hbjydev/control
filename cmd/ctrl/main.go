package main

import (
  "fmt"
)

func main() {
  fmt.Println("ctrl -- A command-line development utility")
  fmt.Println()
  fmt.Println("Usage:")
  fmt.Println("  ctrl <operation> [options]")
  fmt.Println()
  fmt.Println("Operations:")
  fmt.Println("  up   -- Brings the project up using the configured")
  fmt.Println("          installation process.")
  fmt.Println("  run  -- Runs a script defined in the configuration.")
  fmt.Println("  who  -- Returns the configured maintainer(s) contact")
  fmt.Println("          information")
  fmt.Println()
  fmt.Println("Options:")
  fmt.Println("  -c,--config\tThe location of the configuration file")
}
