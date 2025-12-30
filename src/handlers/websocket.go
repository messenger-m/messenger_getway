package handlers

import (
	redis_handler "api-gateway/src/redisHandler"
	"log"

	"github.com/gofiber/websocket/v2"
	"github.com/redis/go-redis/v9"
)

type Message struct {
	Text string `json:"text"`
}

func WebSocketHandler(c *websocket.Conn) {
	log.Println("WebSocket клиент подключён")

	for {
		var msg Message
		if err := c.ReadJSON(&msg); err != nil {
			log.Println("Клиент отключился:", err)
			break
		}

		log.Println("Получено:", msg.Text)

		_, err := redis_handler.Client.XAdd(redis_handler.Ctx, &redis.XAddArgs{
			Stream: "events:in",
			Values: map[string]interface{}{
				"type": "message.send",
				"text": msg.Text,
				// позже добавишь user_id, chat_id
			},
		}).Result()

		if err != nil {
			log.Println("Redis error:", err)
			break
		}
	}
}
