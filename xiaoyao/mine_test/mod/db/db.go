package db

import (
	"sync"

	"Project/xiaoyao/mine_test/mod"

	"github.com/garyburd/redigo/redis"
)

type DbMod struct {
	nameLock sync.Mutex

	gameDBPool *redis.Pool
	roleDBPool *redis.Pool

	mailLock      sync.Mutex
	arenaLogLock  []sync.Mutex
	battleLogLock []sync.Mutex
}

func init() {
	arenaLogLocks := make([]sync.Mutex, 16)
	battleLogLocks := make([]sync.Mutex, 16)
	for i := 0; i < 16; i++ {
		arenaLogLocks[i] = sync.Mutex{}
		battleLogLocks[i] = sync.Mutex{}
	}
	mod.DB = &DbMod{}
}
