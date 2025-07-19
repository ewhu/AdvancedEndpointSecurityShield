// cmd/advancedendpointsecurityshield/main.go
package main

import (
"flag"
"log"
"os"

"advancedendpointsecurityshield/internal/advancedendpointsecurityshield"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := advancedendpointsecurityshield.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
