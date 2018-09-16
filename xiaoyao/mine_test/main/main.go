package main

import (
	"Project/xiaoyao/mine_test/mod"
	_ "Project/xiaoyao/mine_test/mod/db"
	_ "Project/xiaoyao/mine_test/mod/mine"
)

func main() {
	mod.Mine.GetMinerConfig()
}
