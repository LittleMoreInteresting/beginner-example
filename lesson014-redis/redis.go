package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type ConfRedis struct {
	Addr         string
	Password     string
	Db           int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	DialTimeout  time.Duration
}

func NewRedisClient(conf *ConfRedis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Addr,
		Password:     conf.Password,
		DB:           conf.Db,
		DialTimeout:  conf.DialTimeout,
		WriteTimeout: conf.WriteTimeout,
		ReadTimeout:  conf.ReadTimeout,
	})
	if rdb == nil {
		log.Fatalf("failed opening connection to redis")
	}
	return rdb
}

func main() {
	conf := &ConfRedis{
		Addr:     ":6379",
		Password: "",
		Db:       0,
	}
	rdb := NewRedisClient(conf)
	ctx := context.Background()
	err := rdb.Set(ctx, "key1", 1, 0).Err()
	if err != nil {
		log.Fatalf("set error:%v", err)
	}
	get, err := rdb.Get(ctx, "key1").Int()
	if err == redis.Nil {
		log.Fatalf("key not found")
	}
	if err != nil {
		log.Fatalf("get error:%v", err)
	}
	fmt.Println(get)

	err = rdb.HSet(ctx, "hset", "a", 1).Err()
	if err != nil {
		log.Fatalf("hset error:%v", err)
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Printf("incr %d", i)
			err = rdb.Incr(ctx, "incr_key").Err()
			if err != nil {
				log.Fatalf("incr error:%v", err)
			}
		}(i)
	}
	wg.Wait()
	i, err := rdb.Get(ctx, "incr_key").Int()
	if err != nil {
		log.Fatalf("get error:%v", err)
	}
	fmt.Println(i)

	var incrBy = redis.NewScript(`
		local key = KEYS[1]
		local change = ARGV[1]
		
		local value = redis.call("GET", key)
		if not value then
		  value = 0
		end
		
		value = value + change
		redis.call("SET", key, value)
		
		return value
		`)
	keys := []string{"my_counter"}
	values := []interface{}{+1}
	num, err := incrBy.Run(ctx, rdb, keys, values...).Int()
	if err != nil {
		log.Fatalf("script error:%v", err)
	}
	fmt.Println(num)
}
