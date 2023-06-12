package rabbitmq

import (
	"sync"
)

var oncePool sync.Once
var instancePool *RabbitPool
