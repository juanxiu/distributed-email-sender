package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

// --- 이메일 수신 핸들러 ---

// localDomain은 시스템의 내부 도메인을 나타냅니다. 실제 환경에서는 설정 파일에서 읽어와야 합니다.
const localDomain = "mydomain.com"

// IncomingEmailRequest는 수신된 이메일을 나타냅니다.
type IncomingEmailRequest struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// ReceiveEmailHandler는 수신된 이메일을 처리합니다.
func ReceiveEmailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST 메서드만 허용됩니다.", http.StatusMethodNotAllowed)
		return
	}

	var req IncomingEmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "잘못된 요청 본문입니다.", http.StatusBadRequest)
		return
	}

	if req.From == "" || req.To == "" {
		http.Error(w, "'From'과 'To' 필드는 필수입니다.", http.StatusBadRequest)
		return
	}

	recipientDomain := getDomainFromEmail(req.To)

	if recipientDomain == localDomain {
		// TODO: 수신된 이메일을 내부 데이터베이스에 저장하는 로직을 구현해야 합니다.
		// 예: database.SaveEmail(req)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "이메일을 수신하여 내부에 저장했습니다."})
	} else {
		// TODO: 외부로 전달할 이메일을 메시지 큐에 발행하는 로직을 구현해야 합니다.
		// 예: messagequeue.PublishEmailForDelivery(req)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{"message": "외부 이메일 전달 요청을 수락했습니다."})
	}
}

// getDomainFromEmail은 이메일 주소에서 도메인 부분을 추출합니다.
func getDomainFromEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) == 2 {
		return strings.ToLower(parts[1])
	}
	return ""
}
