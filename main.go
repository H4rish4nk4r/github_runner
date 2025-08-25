package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"github.com/H4rish4nk4r/github-runner/worflows"
	"github.com/H4rish4nk4r/github-runner/worflows"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, "vcluster-task-queue", worker.Options{})
	w.RegisterWorkflow(VClusterWorkflow)
	w.RegisterActivity(TriggerGitHubWorkflow)

	log.Println("Starting worker...")
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("worker failed", err)
	}
}
