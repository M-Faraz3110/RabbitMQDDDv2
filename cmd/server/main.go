package main

import (
	"rabbitmqdddv2/pkg/app"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	app.Start()
	wg.Wait()
}
