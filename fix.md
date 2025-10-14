ユーザー会員登録

 - bookingapp/internal/domain/repository/repository.go:17 に UserRepository を追加し、ユースケースが依存するインターフェースを定義。
  - bookingapp/internal/usecase/user_uc.go:29 で登録ロジックを実装し、入力バリデーション・日付パース・重複メール検出を行って UserRepository へ委譲。
  - bookingapp/internal/infrastructure/repository/mysql/user/user_repo_mysql.go:24 に MySQL 実装を追加し、UUID 発行・DateOfBirth の NULL 対応・既存メール検索を実装。
  - bookingapp/internal/interface/http/user.go:26 にユーザーハンドラを作成し、HTTP からユースケースを呼び出して適切な HTTP ステータスを返却。
  - bookingapp/cmd/api/main.go:42 で新しいハンドラ/ユースケース/リポジトリを初期化し、POST /register を UserHandler にバインド。