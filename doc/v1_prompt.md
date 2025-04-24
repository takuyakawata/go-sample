Junieさん、以下の要素で「予約・スケジューリングシステム（Booking System）」を Go で作成してください。

---

## 1. ドメイン設計

- **会議室（Room）**
- **イベント（Event）**
- **予約枠（ReservationSlot）**

### ビジネスルール

1. **時間帯の重複チェック**
    - 同一会議室で予約枠が重複しないようにバリデーション
2. **ユーザー権限**
    - 管理者（Admin）：全ての予約を作成・承認・キャンセル可能
    - 一般ユーザー（User）：自身の予約のみ作成・申請
3. **承認フロー**
    - 一般ユーザーが予約すると「承認待ち」ステータス
    - 管理者による承認（または 2 段階承認フローをサポート）

---

## 2. アーキテクチャ／構成

- **Clean Architecture（Hexagonal Architecture）**
    - `internal/domain`：エンティティ／値オブジェクト／ドメインサービス
    - `internal/usecase`：ユースケース／アプリケーションサービス
    - `internal/interface`：HTTP（REST）／gRPC アダプタ
    - `internal/infrastructure`：PostgreSQL リポジトリ実装、Redis キャッシュ実装
- **cmd/**：サーバー起動用エントリポイント（REST と gRPC のそれぞれ）
- **pkg/**：汎用ユーティリティや共有ライブラリ

---

## 3. 技術要素

1. **REST API**
    - Gin／Echo などで予約の CRUD と承認エンドポイントを実装
2. **gRPC サービス**
    - プロトコル定義（`.proto`）で予約取得・作成・承認インターフェースを定義
3. **PostgreSQL**
    - `tsrange` 型や排他制約を使った時間帯重複チェック
    - `go-pg` や `gorm` などでマイグレーションと ORM を実装
4. **Redis キャッシュ**
    - 直近の空き状況クエリ結果をキャッシュして高速化
    - `go-redis` クライアントを利用

---

## 4. 開発フロー

1. **リポジトリ作成 & `go.mod` 初期化**
2. **ドメインモデル定義**
    - `Room`, `Event`, `ReservationSlot` のエンティティと集約ルート
3. **ユースケース実装**
    - 予約作成／重複チェック／承認ロジック
4. **インターフェース層**
    - REST：JSON エンドポイント設計 + Swagger/OpenAPI 書き起こし
    - gRPC：`.proto` → サーバー／クライアントコード生成
5. **インフラ実装**
    - PostgreSQL 接続設定、テーブル定義マイグレーション
    - Redis クライアント設定
6. **テスト**
    - テーブル駆動ユニットテスト（ドメイン・ユースケース）
    - 統合テスト（PostgreSQL / Redis を Docker Compose で起動）
7. **CI/CD & Docker**
    - `Makefile` や `docker-compose.yml` でローカル起動
    - GitHub Actions で `go fmt`・`go vet`・`go test` の自動実行

---

## 5. ドキュメント

- **README.md**：セットアップ手順／実行方法／API リファレンス
- **UML 図**：ドメインモデル図、シーケンス図（予約フロー）
- **OpenAPI/Proto**：仕様ファイルをリポジトリに含める

---

これらを踏まえて、一通りの機能が動作するサンプルプロダクトを作成してください。完成後はコードレビューとデプロイ方法（Docker Compose or Kubernetes）も一緒に確認しましょう！
