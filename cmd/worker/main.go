package main

import (
	"github.com/TomasCruz/users/internal/infra/worker"
)

func main() {
	w := worker.WorkerApp{}
	w.Start()
}
