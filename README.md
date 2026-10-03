# nice-vote-be

リアルタイム投票・集計サービス「nice-vote」のバックエンド API サーバーです。  
Go言語とGinフレームワークを使用し、高速で軽量なWeb APIおよびWebSocket通信を提供します。

## 🛠 使用技術 (Tech Stack)

* **Language:** Go
* **Framework:** Gin (`github.com/gin-gonic/gin`)
* **Database:** PostgreSQL (開発環境: Docker / 本番環境: Supabase)
* **Realtime:** Gorilla WebSocket (`github.com/gorilla/websocket`)
* **Deployment:** Render

## 🚀 ローカル開発環境の起動方法

### 前提条件 (Prerequisites)
* Go 1.22 以上

### セットアップ順序

1. **リポジトリのクローン**
   ```bash
   git clone https://github.com/TakehiroTsukagoshi/nice-vote-be.git
   cd nice-vote-be
   ```

2. **依存関係のインストール**
   ```bash
   go mod download
   ```

3. **サーバーの起動**
   ```bash
   go run main.go
   ```
   * 起動後、`http://localhost:8080/health` にアクセスして動作確認ができます。

## 📡 エンドポイント一覧 (API Specifications)

| Method | Endpoint | Description | Auth |
|---|---|---|---|
| GET | `/health` | サーバーのヘルスチェック | 不要 |
| GET | `/api/polls` | アンケート一覧の取得 | 不要 |
| POST | `/api/polls` | 新規アンケートの作成 | 不要 |
| GET | `/ws/polls/:id` | リアルタイム投票結果受信用 WebSocket | 不要 |
