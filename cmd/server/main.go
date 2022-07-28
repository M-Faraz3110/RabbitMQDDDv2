package main

import (
	"rabbitmqdddv2/pkg/app"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1) //Go routine to keep listening
	app.Start()
	wg.Wait()
}
