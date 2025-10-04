# Go Webアプリスタータ

本プロジェクトは [Gin](https://github.com/gin-gonic/gin) フレームワークを用いて Web API を構築するためのスターターキットです。シンプルな構成とヘルスチェック、基本的な開発ツール群に加えて Next.js フロントエンド用の足場も含んでおり、小規模サービスの立ち上げを素早く行えます。

## 特徴

- `cmd/`・`internal/`・`config/` によるわかりやすいディレクトリ構成
- Graceful shutdown に対応した HTTP サーバー制御
- `/healthz` と `/api/v1/status` のヘルスチェックエンドポイント
- `Makefile` による実行・ビルド・テストなど開発タスクの自動化
- Docker Compose で Go バックエンドと Next.js フロントエンドを同時に起動可能

## ローカル開発 (ホストマシンで直接実行)

```shell
# 依存関係の取得
go mod tidy

# 開発モードで起動
make run

# テストの実行
make test
```

既定ではサーバーは `0.0.0.0:8080` で待ち受けます。環境変数でポートなどを上書きできます。

```shell
APP_ENV=production PORT=9090 make run
```

## Docker Compose を使った環境構築

1. `.env.example` を元に `.env` を作成し、必要に応じて値を調整します。
2. 初回のみ依存関係を取得するため、Go モジュールを整えます。

   ```shell
   docker compose run --rm backend sh -c "go mod tidy"
   ```

3. フロント・バックエンドをまとめて起動します。

   ```shell
   docker compose up --build
   ```

4. アクセス確認
   - バックエンド API: http://localhost:8080/healthz
   - フロントエンド: http://localhost:3000/

停止する際は `Ctrl+C` で抜けた後、必要なら `docker compose down` を実行してください。フロントエンドは `NEXT_PUBLIC_API_BASE_URL` 環境変数でバックエンドの URL を参照します (Compose では自動で `http://backend:8080` に設定しています)。

### 本番ビルド用コンテナ

- `Dockerfile.backend`: Go バイナリをビルドして軽量な Alpine 上で実行するコンテナを生成します。
- `frontend/Dockerfile`: Next.js をビルドして `next start` で配信する本番用コンテナを生成します。

それぞれ以下のようにビルドできます。

```shell
# バックエンド
docker build -f Dockerfile.backend -t sample-backend .

# フロントエンド
docker build -f frontend/Dockerfile -t sample-frontend ./frontend
```

## 採用フレームワーク (Gin)

Gin は Go 言語で広く利用されている軽量かつ高速な Web フレームワークです。

- ルーティングとミドルウェアの定義が容易で読みやすいコードを保ちやすい
- JSON ハンドリングやバリデーションなど API 開発に必要な機能を標準で備える
- `net/http` ベースのため標準ライブラリとの親和性が高く、既存資産の再利用がしやすい
- 大規模なコミュニティと豊富なサンプルがあり、情報収集が容易

本スターターでは Gin の `Logger`・`Recovery` ミドルウェアを初期設定し、ルーター層 (`internal/router`) を独立させることで将来的なルート追加やミドルウェア拡張をしやすくしています。

## ディレクトリ構成

```
.
|-- cmd/server        # アプリケーションのエントリーポイント
|-- config            # 環境変数の読み込みと設定
|-- internal/app      # アプリケーションの組み立てとライフサイクル管理
|-- internal/handler  # リクエストハンドラー群
|-- internal/router   # HTTP ルーティングの定義
|-- frontend          # Next.js フロントエンド (pages/, styles/ など)
|-- Makefile          # 開発タスクのコマンド集
|-- README.md
```

## 次のステップ

- 独自のドメインロジックを `internal/handler` 配下に追加してください。
- 必要に応じてデータベースやメッセージングなどのサービス連携を組み込みましょう。
- Docker Compose をベースに Air や Tilt などを導入すると、より快適なホットリロード環境を構築できます。
- Next.js 側からバックエンド API を呼び出す際は `NEXT_PUBLIC_API_BASE_URL` を調整し、OpenAPI などで型を共有すると保守が容易になります。
