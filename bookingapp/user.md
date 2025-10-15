# ユーザー関連 API の `curl` 例

## 共通情報
- ベース URL: `http://13.208.158.221`
- すべてのリクエストで `Content-Type: application/json` ヘッダを付与

## ユーザー登録
新規ユーザーを登録するには `POST /register` を叩きます。`name` と `email` は必須、`date_of_birth` は `YYYY-MM-DD` 形式です。

```bash
curl -i -X POST http://13.208.158.221/register \
  -H 'Content-Type: application/json' \
  -d '{
        "name": "山田太郎",
        "email": "taro.yamada@example.com",
        "phone_number": "090-1234-5678",
        "address": "東京都千代田区1-1-1",
        "date_of_birth": "1990-04-01"
      }'
```

成功すると `HTTP/1.1 201 Created` とともに `{"id":"<生成されたUUID>"}` が返り、メールアドレスが既に存在する場合は `409 Conflict` になります。

## 登録済みユーザー取得
ユーザー ID を指定して `GET /users/{id}` を叩きます。上の登録レスポンスで返った ID を使って確認できます。

```bash
USER_ID="取得したユーザーID"
curl -i http://13.208.158.221/users/${USER_ID}
```

存在しない ID を指定すると `404 Not Found`、フォーマットが不正な場合は `400 Bad Request` が返ります。
