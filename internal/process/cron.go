package process

import (
	"context"
	"errors"
	"log"
)

// An example of worker

type worker struct {
	// usecases etc.
}

func NewWorker() *worker {
	return &worker{}
}

func (w *worker) Run(ctx context.Context) error {
	log.Println("worker running...")
	// do something. only return error if worker is failed to start
	return errors.New("unimplemented")
}
