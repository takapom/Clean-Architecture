# Booking App API

Go 言語と GORM を用いて宿泊予約を管理するシンプルな API サービスです。Clean Architecture を意識したレイヤ分割を採用し、ドメインロジックとインフラ依存コードを分離しています。

## 主な特徴
- MySQL（またはメモリリポジトリ）を利用したプラン・予約データ管理
- 予約作成時のバリデーション（宿泊日、人数、プラン存在チェック）
- RESTful なエンドポイント（プラン検索、予約 CRUD の一部）
- サーバ起動時の自動マイグレーションと初期データ投入

## ディレクトリ構成
```
.
├── cmd/api            # エントリーポイント（HTTP サーバ）
├── internal
│   ├── domain         # ドメインエンティティ & リポジトリインターフェース
│   ├── usecase        # ユースケース（アプリケーションサービス）
│   ├── interface/http # HTTP ハンドラ層
│   └── infrastructure # DB 接続・リポジトリ実装（MySQL / メモリ）
├── go.mod, go.sum
└── docker-compose.yml # MySQL 起動用定義
```

## アーキテクチャ概要
- `internal/domain/entity` にプラン (`Plan`) と予約 (`Reservation`) のドメインモデルを定義。`Reservation.Nights()` などのビジネスロジックをエンティティに寄せています。
- `internal/domain/repository` ではユースケースが依存するポート（インターフェース）を宣言。
- `internal/usecase/reservation_uc.go` はプラン検索や予約作成のアプリケーションロジックを担当し、入力バリデーションと料金計算を行います。
- `internal/interface/http` が HTTP リクエストを受け、ユースケースを呼び出して JSON を返却します。
- `internal/infrastructure` で具体的なアダプタを実装。`repository/mysql` は GORM を利用した永続化、`memory` はインメモリ実装です。

## 依存関係
- Go 1.24 以上
- MySQL 8 系（Docker Compose で起動可）
- ライブラリ
  - `gorm.io/gorm`
  - `gorm.io/driver/mysql`

## 環境変数
`cmd/api/main.go` では以下の環境変数を読み込み、未設定の場合はデフォルト値を使用します。

| 変数名   | デフォルト値  | 説明                 |
|----------|----------------|----------------------|
| `DB_HOST`| `127.0.0.1`    | MySQL ホスト         |
| `DB_PORT`| `3306`         | MySQL ポート         |
| `DB_USER`| `root`         | 接続ユーザー         |
| `DB_PASS`| `password`     | 接続パスワード       |
| `DB_NAME`| `booking`      | 使用するデータベース |

## 起動方法
1. 依存環境を用意
   ```bash
   docker compose up -d mysql
   ```
   > `docker-compose.yml` は root パスワード `password`、DB 名 `booking` で MySQL 8.0 を立ち上げます。

2. API サーバを起動
   ```bash
   go run ./cmd/api
   ```
   起動すると `:8080` で HTTP サーバが待ち受けます。

### マイグレーションとシード
サーバ起動時に以下が自動で実行されます。
- GORM の `AutoMigrate` による `plans` / `reservations` テーブル生成
- `plans` テーブルが空の場合、初期プラン 3 件を投入
  - 例: `ID=100, Name="富士プレミアム", Price=12000`

## ドメインロジック
- 予約作成 (`ReservationUsecase.Create`)
  - チェックイン < チェックアウト、人数 >= 1 を検証
  - 指定プランの存在確認（未存在時は `ErrPlanNotFound`）
  - 宿泊数（`Reservation.Nights()`）と人数、プラン単価から合計金額を算出
  - リポジトリ経由で保存し、生成された ID を返却
- 予約参照 (`Get`, `List`) とプラン検索 (`SearchPlans`) もユースケースを経由

## HTTP API
| メソッド | パス                | 説明                           |
|----------|---------------------|--------------------------------|
| `POST`   | `/reservations`     | 予約を新規作成                 |
| `GET`    | `/reservations`     | 予約一覧を取得                 |
| `GET`    | `/reservations/{id}`| 予約詳細を取得                 |
| `GET`    | `/plans`            | キーワードでプランを検索       |

### リクエスト/レスポンス例
**予約作成**
```bash
curl -X POST http://localhost:8080/reservations \
  -H 'Content-Type: application/json' \
  -d '{
        "plan_id": 100,
        "number": 2,
        "checkin": "2025-10-12",
        "checkout": "2025-10-14"
      }'
```
レスポンス
```json
{ "id": 1 }
```

**予約一覧**
```bash
curl http://localhost:8080/reservations
```
レスポンス（例）
```json
[
  {
    "id": 1,
    "plan_id": 100,
    "number": 2,
    "checkin": "2025-10-12",
    "checkout": "2025-10-14",
    "total": 48000,
    "nights": 2
  }
]
```

**プラン検索**
```bash
curl "http://localhost:8080/plans?keyword=富士"
```
レスポンス（例）
```json
[
  {
    "id": 100,
    "name": "富士プレミアム",
    "keyword": "富士 山 静岡",
    "price": 12000
  }
]
```

### エラーレスポンス
- リクエスト JSON / 日付フォーマット不正: `400 Bad Request`
- 宿泊日逆転・人数不足: `400 Bad Request`
- 存在しないプラン指定: `404 Not Found`
- その他予期しないエラー: `500 Internal Server Error`

## テストや拡張のヒント
- インメモリリポジトリ（`internal/infrastructure/memory`）を利用してユニットテストを書けます。
- バリデーション強化（例: 最大人数、予約重複チェック）や、キャンセル API 追加などの拡張が容易です。
- HTTP レイヤは `net/http` 標準ライブラリのままなので、Echo や Chi などに置き換える場合もユースケース層は流用可能です。

## ライセンス
本リポジトリにライセンス表記がない場合は、利用前に作成者へ確認してください。
