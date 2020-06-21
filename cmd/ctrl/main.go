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
  fmt.Println("  up\tBrings the project up using the configured")
  fmt.Println("    \tinstallation process.")
  fmt.Println("  run\tRuns a script defined in the configuration.")
  fmt.Println("  who\tReturns the configured maintainer(s) contact")
  fmt.Println("     \tinformation")
  fmt.Println()
  fmt.Println("Options:")
  fmt.Println("  -c,--config\tThe location of the configuration file")
  fmt.Println("  -h,--help\tDisplays this message")
  fmt.Println("  -v,--version\tDisplays the installed Control version")
}
