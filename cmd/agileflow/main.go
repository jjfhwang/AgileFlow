// cmd/agileflow/main.go
package main

import (
"flag"
"log"
"os"

"agileflow/internal/agileflow"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := agileflow.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
