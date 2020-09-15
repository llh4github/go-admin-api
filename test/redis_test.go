package test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

// Person 测试结构体与redis中读写操作的工具人
type Person struct {
	Name     string
	Hegith   int
	Birthday time.Time
}

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       3,
	})
}

// 测试 初始化redis 连接
func TestInitRedis(t *testing.T) {

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
	if err != nil {
		t.Fatal("redis 连接失败！ ", err)
	}
}

// 测试 字符串值存取
func TestStr(t *testing.T) {
	set, err := rdb.Set(ctx, "hello", "world", time.Hour*2).Result()
	if err != nil {
		t.Log("字符串值存入失败！", err)
	}
	t.Log(set)
	hello, err := rdb.Get(ctx, "hello").Result()
	if err != nil {
		t.Log("字符串值取出失败！", err)
	}
	t.Log(hello)
}

// 测试 列表值的存取
func TestList(t *testing.T) {
	l1, err := rdb.LPush(ctx, "l1", "apple", "banana").Result()
	if err != nil {
		t.Log("列表值存入失败！", err)
	}
	t.Log(l1)
	ll1, _ := rdb.LLen(ctx, "l1").Result()
	t.Log(ll1)
	ex, err := rdb.Expire(ctx, "l1", time.Hour*2).Result()
	if err != nil {
		t.Log("列表设定过期时间失败！", err)
	}
	t.Log(ex)
}

// 测试
func TestStruct(t *testing.T) {
	p := Person{
		Name:     "Tom",
		Hegith:   123,
		Birthday: time.Now(),
	}
	hashed, _ := json.Marshal(p)
	s1, err := rdb.Set(ctx, "s1", hashed, time.Hour).Result()
	if err != nil {
		t.Log("列表设定过期时间失败！", err)
	}
	t.Log(s1)
	s1h := rdb.Get(ctx, "s1").Val()
	var p2 Person
	json.Unmarshal([]byte(s1h), &p2)
	t.Log(p2)
}
