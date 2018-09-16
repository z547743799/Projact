package mod

import "Project/xiaoyao/mine_test/data"

type (
	IDBMod interface {
		GetMinerConfig() map[int64]*data.TMineCaveConfig
	}
	IMine interface {
		GetMinerConfig()
	}
)
