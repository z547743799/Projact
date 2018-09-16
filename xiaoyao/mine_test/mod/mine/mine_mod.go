package mine

import (
	"Project/xiaoyao/mine_test/data"
	"Project/xiaoyao/mine_test/mod"
	"fmt"
)

type mineMod struct {
	KV map[int64]data.TMineConfig
}

func init() {
	mod.Mine = &mineMod{}
}

func (this *mineMod) GetMinerConfig() {
	KV := mod.DB.GetMinerConfig()
	for k, v := range KV {
		fmt.Println(k)
		fmt.Println(v)
	}
}
