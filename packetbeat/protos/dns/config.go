package dns

import (
	"github.com/wangjia184/beats/packetbeat/config"
	"github.com/wangjia184/beats/packetbeat/protos"
)

type dnsConfig struct {
	config.ProtocolCommon `config:",inline"`
	IncludeAuthorities    bool `config:"include_authorities"`
	IncludeAdditionals    bool `config:"include_additionals"`
}

var (
	defaultConfig = dnsConfig{
		ProtocolCommon: config.ProtocolCommon{
			TransactionTimeout: protos.DefaultTransactionExpiration,
		},
	}
)
