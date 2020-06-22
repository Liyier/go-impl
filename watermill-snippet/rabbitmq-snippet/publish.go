package rabbitmq_snippet

import (
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func publish() chan int {
	pub := gochannel.NewGoChannel()
}
