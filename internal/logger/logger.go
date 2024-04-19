package logger

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

var Sugar zap.SugaredLogger

func init() {
	// создаём предустановленный регистратор zap
	logger, err := zap.NewDevelopment()
	if err != nil {
		// вызываем панику, если ошибка
		panic(err)
	}
	defer logger.Sync()

	// делаем регистратор SugaredLogger
	Sugar = *logger.Sugar()
}

type (
	// Берём структуру для хранения сведений об ответе
	responseData struct {
		status int
		size   int
	}

	// LoggingResponseWriter добавляем реализацию http.ResponseWriter
	LoggingResponseWriter struct {
		http.ResponseWriter // встраиваем оригинальный http.ResponseWriter
		responseData        *responseData
	}
)

func (r *LoggingResponseWriter) Write(b []byte) (int, error) {
	// записываем ответ, используя оригинальный http.ResponseWriter
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size // захватываем размер
	return size, err
}

func (r *LoggingResponseWriter) WriteHeader(statusCode int) {
	// записываем код статуса, используя оригинальный http.ResponseWriter
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode // захватываем код статуса
}

// WithLoggingFunc добавляет дополнительный код для регистрации сведений о запросе
// и возвращает новый http.HandlerFunc.
func WithLoggingFunc(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lw := LoggingResponseWriter{
			ResponseWriter: w, // встраиваем оригинальный http.ResponseWriter
			responseData: &responseData{
				status: 0,
				size:   0,
			},
		}
		// точка, где выполняется Handler
		h.ServeHTTP(&lw, r) // обслуживание оригинального запроса

		duration := time.Since(start)

		// отправляем сведения о запросе в zap
		Sugar.Infoln(
			"uri", r.RequestURI,
			"method", r.Method,
			"status", lw.responseData.status, // получаем перехваченный код статуса ответа
			"duration", duration,
			"size", lw.responseData.size, // получаем перехваченный размер ответа
		)

	}
}
