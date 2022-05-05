package models

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var Pool *redis.Pool

func init() {
	//redisHost, _ := beego.AppConfig.String("redisHost")
	//log.Fatal(redisHost)
	//redisPort, _ := beego.AppConfig.String("redisPort")
	//redisPass, _ := beego.AppConfig.String("redisPass")
	//redisDB, _ := beego.AppConfig.Int("redisDB")
	//redisAddr := beego.AppConfig.String("redisAddr")
	//redisPort, _ := beego.AppConfig.String("redisPort")
	//var conn = beego.AppConfig.String("redisHost") + ":" + beego.AppConfig.String("redisPort")

	Pool = &redis.Pool{ //实例化一个连接池
		MaxIdle:     16,  //最初的连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Wait:        true,
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			//dial, err2 := redis.Dial("tcp", "127.0.0.1:6379")
			dial, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				log.Fatal(err)
			}
			return dial, err
		},
	}
	c := Pool.Get()
	defer c.Close()
}

func Lock(key string, expireTime int) int {
	client := Pool.Get()
	n, _ := redis.Int(client.Do("SETNX", key, expireTime))
	if n != 1 { //没有获取到锁
		// 判断锁是否过期
		lockTime, _ := redis.Int64(client.Do("GET", key))
		unix := time.Now().Unix()
		// 锁已过期，删除锁，重新获取
		if unix > lockTime {
			Unlock(key)
			n, _ = redis.Int(client.Do("SETNX", key, unix+int64(expireTime)))
		}
	}
	return n
}

func Unlock(key string) {
	client := Pool.Get()
	client.Do("DEL", key)
}
