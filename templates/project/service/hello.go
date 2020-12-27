package service

import (
	"context"

	"github.com/influx6/sabuhp"
	"github.com/influx6/sabuhp/actions"
)

func helloWorldAction(ctx context.Context, job actions.Job) {
	_ = job.Transport.SendToAll(sabuhp.NewMessage(job.Msg.FromAddr, job.To, []byte("hello world")), -1)
}

var HelloWorldRequest = actions.WorkerRequest{
	ActionName:  "helloWorld",
	PubSubTopic: "hello-world",
	WorkerCreator: func(config actions.WorkerConfig) *actions.WorkerGroup {
		config.Action = actions.ActionFunc(helloWorldAction)
		return actions.NewWorkGroup(config)
	},
}
