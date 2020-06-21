package main

import (
  "flag"
  "fmt"
  "os"
)

const (
        InfoColor    = "\033[1;34m%s\033[0m"
        NoticeColor  = "\033[1;36m%s\033[0m"
        WarningColor = "\033[1;33m%s\033[0m"
        ErrorColor   = "\033[1;31m%s\033[0m"
        DebugColor   = "\033[0;36m%s\033[0m"
)

func help() {
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
  fmt.Println("     \tinformation.")
  fmt.Println()
  fmt.Println("Options:")
  flag.VisitAll(func(f *flag.Flag) {
    fmt.Fprintf(os.Stderr, "  --%v\t%v", f.Name, f.Usage)
  })
}

func main() {
  flag.Usage = help
  // Up subcommand
  //upCmd := flag.NewFlagSet("up", flag.ExitOnError)
  //runCmd := flag.NewFlagSet("run", flag.ExitOnError)
  //whoCmd := flag.NewFlagSet("who", flag.ExitOnError)

  // Global flags
  helpFlag := flag.Bool("help", false, "Display this message.")

  flag.Parse()

  if *helpFlag == true {
    help()
    os.Exit(0)
  }

  if len(os.Args[1:]) < 1 {
    help()
    os.Exit(1)
  }

  switch os.Args[1] {
    case "up":
      fmt.Fprintf(os.Stdout, NoticeColor, "ctrl v1.0.0\n")
      fmt.Fprintf(os.Stdout, ErrorColor, " |----------------------------------|\n")
      fmt.Fprintf(os.Stdout, ErrorColor, " | NOT IMPLEMENTED                  |\n")
      fmt.Fprintf(os.Stdout, ErrorColor, " |----------------------------------|\n")
      fmt.Fprintf(os.Stdout, ErrorColor, " | This feature hasn't been         |\n")
      fmt.Fprintf(os.Stdout, ErrorColor, " | implemented yet.                 |\n")
      fmt.Fprintf(os.Stdout, ErrorColor, " |----------------------------------|\n")

    default:
      help()
      os.Exit(1)
  }
}
