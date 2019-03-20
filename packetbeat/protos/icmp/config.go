package icmp

import (
	"time"

	"github.com/wangjia184/beats/packetbeat/protos"
)

type icmpConfig struct {
	SendRequest        bool          `config:"send_request"`
	SendResponse       bool          `config:"send_response"`
	TransactionTimeout time.Duration `config:"transaction_timeout"`
}

var (
	defaultConfig = icmpConfig{
		TransactionTimeout: protos.DefaultTransactionExpiration,
	}
)
