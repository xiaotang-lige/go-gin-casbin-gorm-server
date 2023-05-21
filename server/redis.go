package server

import (
	"context"
	"encoding/json"
	"log"
	"messageServe/linker"
	"messageServe/model"
	"time"
)

func redisLPush(ctx context.Context, mes *model.Message) error {
	ll, err := json.Marshal(mes)
	if err != nil {
		log.Println(err)
	}
	return linker.Db.RedisDb.LPush(ctx, mes.Target, ll).Err()
}
func redisBRPop(ctx context.Context, time time.Duration, userId string) ([]string, error) {
	v, err := linker.Db.RedisDb.BRPop(ctx, time, userId).Result()
	return v, err
}
