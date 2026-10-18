package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// レスポンス用のデータ構造体 (TypeScriptの type / interface に相当)
type Poll struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	VotesCount int    `json:"votes_count"`
}

func SetupRouter() *gin.Engine {
	// Gin のルーター（ルーターインスタンス）を初期化
	r := gin.Default()

	// CORS設定 (Next.js開発サーバー http://localhost:3000 からのアクセスを許可)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// ヘルスチェック用エンドポイント (動作確認用)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Go API server is running",
		})
	})

	// アンケート向けのダミーデータを返すAPI
	r.GET("/api/polls", func(c *gin.Context) {
		polls := []Poll{
			{ID: "1", Title: "好きなプログラミング言語は？", VotesCount: 42},
			{ID: "2", Title: "好きなエディタは？", VotesCount: 18},
		}
		c.JSON(http.StatusOK, gin.H{"polls": polls})
	})

	// 新規投票を作成する POST エンドポイント
	r.POST("/api/polls", func(c *gin.Context) {
		var newPoll Poll

		if err := c.ShouldBindJSON(&newPoll); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		randomId := rand.Intn(10000) + 1
		newPoll.ID = strconv.Itoa(randomId)
		newPoll.VotesCount = 0

		c.JSON(http.StatusCreated, gin.H{"poll": newPoll})
	})

	return r
}

func main() {
	r := SetupRouter()
	// サーバーの起動 (ポート 8080 で待機)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
