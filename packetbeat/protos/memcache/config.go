package memcache

import (
	"time"

	"github.com/wangjia184/beats/packetbeat/config"
	"github.com/wangjia184/beats/packetbeat/protos"
)

type memcacheConfig struct {
	config.ProtocolCommon `config:",inline"`
	MaxValues             int
	MaxBytesPerValue      int
	UDPTransactionTimeout time.Duration
	ParseUnknown          bool
}

var (
	defaultConfig = memcacheConfig{
		ProtocolCommon: config.ProtocolCommon{
			Ports:              []int{11211},
			TransactionTimeout: protos.DefaultTransactionExpiration,
		},
		UDPTransactionTimeout: protos.DefaultTransactionExpiration,
	}
)
