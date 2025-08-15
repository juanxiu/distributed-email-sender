package api

import (
	"net/http"
)

// SetupRoutes는 모든 API 라우트를 설정합니다.
func SetupRoutes() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/api/receive-email", ReceiveEmailHandler)

	// 여기에 다른 라우트를 추가할 수 있습니다.

	return router
}
