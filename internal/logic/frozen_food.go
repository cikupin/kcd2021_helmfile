package logic

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"reflect"

	"github.com/cikupin/kcd2021_helmfile/internal/model"
	"github.com/cikupin/kcd2021_helmfile/internal/payload"
	"github.com/go-redis/redis/v8"
)

func (l *Logic) GetFrozenFoods(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// get data from redis
	if l.config.CacheOptions.IsEnable {
		log.Println("[INFO] Trying to get frozen food data from redis....")

		res, err := l.cache.Get(ctx, l.config.CacheKeys.FrozenFoodKey).Result()
		if err != nil && err != redis.Nil {
			log.Printf("[ERROR] Error get frozen food data from redis: %s\n", err.Error())
			payload.CreateErrorResponse(w, err, "err get data from redis")
			return
		}

		if res != "" {
			data := []model.FrozenFood{}
			err = json.Unmarshal([]byte(res), &data)
			if err != nil {
				log.Printf("[ERROR] Error unmarshal frozen food data from redis: %s\n", err.Error())
				payload.CreateErrorResponse(w, err, "err unmarshal data from redis")
				return
			}

			payload.CreateSuccessResponse(w, data)
			return
		}
	}

	// get data from database
	foods, err := l.queryFrozenFoodList()
	if err != nil {
		log.Printf("[ERROR] Error get frozen food data from mysql: %s\n", err.Error())
		payload.CreateErrorResponse(w, err, "err get data from database")
		return
	}

	payload.CreateSuccessResponse(w, foods)

	// store data to cache
	if l.config.CacheOptions.IsEnable && !reflect.DeepEqual(foods, model.FrozenFood{}) {
		log.Println("[INFO] Trying to save frozen food data to redis....")

		b, err := json.Marshal(foods)
		if err != nil {
			log.Printf("[ERROR] Error marshalling frozen food data to be saved to redis: %s\n", err.Error())
			return
		}

		err = l.cache.SetEX(ctx, l.config.CacheKeys.FrozenFoodKey, b, l.config.CacheKeys.FronzenFoodTTL).Err()
		if err != nil {
			log.Printf("[ERROR] Error saving frozen food data to redis: %s\n", err.Error())
		}
	}

}

func (l *Logic) queryFrozenFoodList() (resp []model.FrozenFood, err error) {
	log.Println("[INFO] Trying to get frozen food data from mysql....")

	_, err = l.db.Select(&resp, "SELECT * FROM frozen_food_stock")
	if err != nil && err != sql.ErrNoRows {
		return resp, err
	}

	return resp, nil
}
