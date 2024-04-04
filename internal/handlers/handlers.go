package handlers

import (
	"github.com/google/uuid"
	"github.com/kelmik91/Y-URL-Shortening-Service/internal/config"
	"github.com/kelmik91/Y-URL-Shortening-Service/internal/storage"
	"io"
	"net/http"
	"strings"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {

	// Если пришел POST
	if r.Method == http.MethodPost {

		//получаем тело запроса
		body, _ := io.ReadAll(r.Body)

		//проверяем что тело не пустое
		if string(body) == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//Создаем ID для ссылки.
		//Думаю лучше использовать хеш-функцию и проверку на существование ссылки в хранилище.
		//Пока сделал более простой вариант.
		uid := uuid.New().String()
		storage.StorageUrl[uid] = string(body)

		//формируем ответ
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(*config.BaseUrlRes + uid))
		return
	}

	// Если пришел GET
	if r.Method == http.MethodGet {

		//Получаем ID из урла
		id := r.URL.Path

		//проверяем запрос на пустой ID
		if id == "/" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//проверяем есть ли ID в хранилище
		id = strings.ReplaceAll(id, "/", "")
		if _, ok := storage.StorageUrl[id]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//формируем ответ
		w.Header().Add("Location", storage.StorageUrl[id])
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	//если не GET или POST отправляем статус 400
	w.WriteHeader(http.StatusBadRequest)
	return
}
