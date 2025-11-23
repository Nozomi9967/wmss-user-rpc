package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	// MySQL 数据源配置
	Mysql struct {
		DataSource string
	}
}
