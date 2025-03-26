package main

import (
    "tree-visualization/routers"
)

func main() {
    router := routers.SetupRouter()
    router.Run(":8080") // Run on port 8080
}