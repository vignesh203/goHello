package handler

import (
	"encoding/json"
	"goHello/service"
	"net/http"
)

type ItemAPI struct {
	users *service.ItemService
}

func NewItemAPI() *ItemAPI {
	return &ItemAPI{
		users: service.NewItemService(),
	}
}

func OnPing(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"ok":   true,
		"pong": true,
	}

	responseBytes, _ := json.Marshal(response)
	_, _ = w.Write(responseBytes)
	return
}
