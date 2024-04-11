package handlers

import (
	"github.com/google/uuid"
	"github.com/kelmik91/Y-URL-Shortening-Service/internal/config"
	"github.com/kelmik91/Y-URL-Shortening-Service/internal/storage"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var warehouse = storage.New()

func MainHandlerGetByID(w http.ResponseWriter, r *http.Request) {
	//Получаем ID из урла
	id := r.URL.Path

	//проверяем запрос на пустой ID
	if id == "/" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//проверяем есть ли ID в хранилище
	id = strings.ReplaceAll(id, "/", "")
	warehouseURL := warehouse.Get(id)
	if warehouseURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//формируем ответ
	w.Header().Add("Location", warehouse.Get(id))
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func MainHandlerSet(w http.ResponseWriter, r *http.Request) {
	//получаем тело запроса
	body, _ := io.ReadAll(r.Body)

	_, err := url.ParseRequestURI(string(body))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//проверяем что тело не пустое
	if string(body) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Создаем ID для ссылки.
	//Думаю лучше использовать хеш-функцию и проверку на существование ссылки в хранилище.
	//Пока сделал более простой вариант.
	uid := uuid.New().String()
	warehouse.AddURL(uid, string(body))

	//формируем ответ
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(config.BaseURLRes + "/" + uid))
}
