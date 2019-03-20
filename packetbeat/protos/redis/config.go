package redis

import (
	"github.com/wangjia184/beats/packetbeat/config"
	"github.com/wangjia184/beats/packetbeat/protos"
)

type redisConfig struct {
	config.ProtocolCommon `config:",inline"`
}

var (
	defaultConfig = redisConfig{
		ProtocolCommon: config.ProtocolCommon{
			TransactionTimeout: protos.DefaultTransactionExpiration,
		},
	}
)
