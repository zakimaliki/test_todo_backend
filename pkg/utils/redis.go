package utils

import (
	"context"
	"encoding/json"
	"strconv"
	"time"
	"todo-api/pkg/config"
	"todo-api/pkg/models"
)

// Fungsi untuk menyimpan data ke cache Redis
func CacheTask(ctx context.Context, id int, task *models.Task) {
	cacheKey := "task:" + strconv.Itoa(id)
	jsonTask, _ := json.Marshal(task)
	config.RedisClient.Set(ctx, cacheKey, jsonTask, 5*time.Minute)
}

// Fungsi untuk menghapus data dari cache Redis
func DeleteCacheTask(ctx context.Context, id int) {
	cacheKey := "task:" + strconv.Itoa(id)
	config.RedisClient.Del(ctx, cacheKey)
}
