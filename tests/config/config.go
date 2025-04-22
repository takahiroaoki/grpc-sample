package config

import "time"

type ClientConfig struct {
	targetAddress    string
	connectTimeout   time.Duration
	connectRetryWait time.Duration
	requestTimeout   time.Duration
}

func (cc ClientConfig) TargetAddress() string {
	return cc.targetAddress
}

func (cc ClientConfig) ConnectTimeout() time.Duration {
	return cc.connectTimeout
}

func (cc ClientConfig) ConnectRetryWait() time.Duration {
	return cc.connectRetryWait
}

func (cc ClientConfig) RequestTimeout() time.Duration {
	return cc.requestTimeout
}

func NewClientConfig() ClientConfig {
	return ClientConfig{
		targetAddress:    env.targetAddress,
		connectTimeout:   3 * time.Second,
		connectRetryWait: 3 * time.Second,
		requestTimeout:   30 * time.Second,
	}
}
