package main

import (
	"context"
	"log"
	"time"

	"go.temporal.io/sdk/client"
	"github.com/H4rish4nk4r/github-runner/workflows"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("cannot create Temporal client", err)
	}
	defer c.Close()

	workflowID := "vcluster-workflow-" + time.Now().Format("20060102150405")

	params := workflows.VClusterParams{
		GitHubToken: "ghp_xxx", // Store securely!
		RepoOwner:   "your-org",
		RepoName:    "your-repo",
		Branch:      "main",
		ClusterName: "demo-cluster",
		Namespace:   "vcluster-ns",
	}

	we, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "vcluster-task-queue",
	}, workflows.VClusterWorkflow, params)
	if err != nil {
		log.Fatalln("unable to execute workflow", err)
	}

	log.Println("Started workflow:", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
