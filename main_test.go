package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ヘルスチェックのエンドポイントテスト
func TestHealthEndpoint(t *testing.T) {
	router := SetupRouter()

	// テスト用のレスポンスレコーダー（結果を記録するオブジェクト）を作成
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)

	// ルーターにリクエストを流す
	router.ServeHTTP(w, req)

	// ステータスコードの検証 (200 OK)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// レスポンスボディの検証
	expectedBody := `{"message":"Go API server is running","status":"ok"}`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, w.Body.String())
	}
}

// アンケート一覧取得のエンドポイントテスト
func TestGetPollsEndpoint(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/polls", nil)

	router.ServeHTTP(w, req)

	// ステータスコードの検証
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// JSONレスポンスの構造をデコードして検証
	var response map[string][]Poll
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse JSON response: %v", err)
	}

	polls, exists := response["polls"]
	if !exists {
		t.Fatal("Expected 'polls' key in response JSON")
	}

	// 件数の検証（ダミーデータが2件あるはず）
	if len(polls) != 2 {
		t.Errorf("Expected 2 polls, got %d", len(polls))
	}

	// 1件目の内容を検証
	if polls[0].Title != "好きなプログラミング言語は？" {
		t.Errorf("Unexpected poll titile: %s", polls[0].Title)
	}
}
