package main

import (
    "log"
    "72HW/producer/router"
)

func main() {
    r := router.SetupRouter()
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("could not run the server: %v", err)
    }
}
