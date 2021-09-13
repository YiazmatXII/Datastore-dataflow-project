package config

import (
	"context"

	"cloud.google.com/go/datastore"
	"cloud.google.com/go/pubsub"
)

const (
	PubsubTopicExec = "Exec"
	WebhookKind     = "info"
	ComputationKind = "result"
	ProjectId       = "projectID_here"
)

var (
	DatastoreClient *datastore.Client
	PubsubClient    *pubsub.Client
	TopicExec       *pubsub.Topic
)

func init() {
	var err error

	DatastoreClient, err = datastore.NewClient(context.Background(), ProjectId)
	if err != nil {
		panic(err)
	}
	PubsubClient, err = pubsub.NewClient(context.Background(), ProjectId)
	if err != nil {
		panic(err)
	}
	TopicExec = PubsubClient.Topic(PubsubTopicExec)
}
