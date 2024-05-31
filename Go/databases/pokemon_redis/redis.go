package redis

import (
	t "Go/time_completion"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type db interface {
	Delete(ctx context.Context, key string)
	Get(ctx context.Context, key string) ([]byte, bool)
	Set(ctx context.Context, key string, value []byte)
	Len(ctx context.Context) int
}

type RdisModel struct {
	DB *redis.Client
}

func NewConnection(ctx context.Context) (rm RdisModel, err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	rm = RdisModel{
		DB: rdb,
	}

	for i := 0; i < 3; i++ {
		_, err = rm.DB.Ping(ctx).Result()
		if err == nil {
			break
		}
		time.Sleep(time.Duration(i) * time.Second)
	}
	return

}

func (r RdisModel) Get(ctx context.Context, key string) ([]byte, bool){
	defer t.FunctionTimer(r.Get)()
	i, err := r.DB.Get(ctx, key).Bytes()
	if err != nil {
		println(err.Error())
		return nil, false
	}
	return i, true
}

func (r RdisModel)Delete(ctx context.Context, key string) {
	defer t.FunctionTimer(r.Delete)()
	res, err := r.DB.Del(ctx, key).Result()
	if err != nil {
		println(err.Error())
		return
	}
	if res == 0{
		println("value not found")
	}

}

func (r RdisModel)Set(ctx context.Context, key string, value []byte){
	defer t.FunctionTimer(r.Set)()
	_, err := r.DB.Set(ctx, key, value, 0).Result()
	if err != nil {
		println(err.Error())
		return
	}
}

func (r RdisModel)Len(ctx context.Context,) int {
	defer t.FunctionTimer(r.Len)()
	result, err := r.DB.DBSize(ctx).Result()
	if err != nil {
		println(err.Error())
		return 0
	}
	return int(result)
}
