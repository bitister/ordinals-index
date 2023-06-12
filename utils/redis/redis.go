package redis

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
)

// redis数据超时时间
const Timeout = 3 * time.Second

//var Client *redis.ClusterClient
var Client *redis.Client

func init() {
	path := beego.AppConfig.String("redis::url")
	passwd := beego.AppConfig.String("redis::password")
	db, _ := beego.AppConfig.Int("redis::db")
	Client = redis.NewClient(&redis.Options{
		Addr:         path,
		Password:     passwd,
		MinIdleConns: 10,
		ReadTimeout:  Timeout,
		PoolSize:     1000,
		DB:           db,
	})
	err := Client.Ping().Err()
	if err != nil {
		beego.Error("Redis wrong ", err)
	}
}

//插入缓存
func RedisSet(key string, value interface{}, time time.Duration) (string, error) {
	s, e := Client.Set(key, value, time).Result()
	beego.Info(s, e, key)
	return s, e
}

//查询缓存
func RedisGet(key string) *redis.StringCmd {
	return Client.Get(key)
}

//查询缓存
func RedisDel(key string) *redis.IntCmd {
	return Client.Del(key)
}

//从列表的右边删除第一个数据，并返回删除的数据
func RedisRpop(key string) *redis.StringCmd {
	return Client.RPop(key)
}

//插入list
func RedisLpush(key string, value interface{}, time time.Duration) (int64, error) {
	s, e := Client.LPush(key, value, time).Result()
	return s, e
}

//获取list
func RedisLRange(key string, start, end int64) ([]string, error) {
	s, e := Client.LRange(key, start, end).Result()
	return s, e
}
