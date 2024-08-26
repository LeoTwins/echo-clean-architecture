# Echo Clean Architecture Template

## 環境構築
```
$ cd echo-clean-architecture
$ docker-compose up -d --build
$ docker-compose exec api /bin/bash
root@{container_id}:/go/src/app# sql-migrate up
```

## ディレクトリ構成
```
.
├── api
│   └── swagger.yml                        # API使用を定義するSwagger
├── cmd
│   └── api
│       ├── main.go                        # エントリーポイント
│       └── router
│           └── router.go                  # ルーティング
├── internal
│   ├── domain                             # ドメイン（エンティティ）層
│   │   ├── model                          # エンティティ・値オブジェクト
│   │   │   ├── account.go
│   │   │   ├── balance.go
│   │   │   ├── money.go
│   │   │   └── transaction.go
│   │   └── repository                     # リポジトリーのインターフェース
│   │       ├── account_repository.go
│   │       └── transaction_repository.go
│   ├── infrastructure                     # インフラストラクチャ層
│   │   ├── db
│   │   │   └── db.go                      # DBとの接続に関するコード
│   │   ├── repository                     # リポジトリの実装
│   │   │   ├── model                      # DBのモデル
│   │   │   │   ├── account.go
│   │   │   │   └── transaction.go
│   │   │   ├── account_repository.go
│   │   │   └── transaction_reposiotry.go
│   │   └── service　　　　　　　　　　　　    # 外部サービス(AWS等)に関するコード
│   │       └── transaction_manager.go
│   ├── interfaces　　　　　　　　　　　　     # インターフェースアダプター層
│   │   ├── handler                        # ハンドラー
│   │   │   ├── dto                        # APIリクエスト・レスポンスの構造体
│   │   │   │   ├── account.go
│   │   │   │   └── transaction.go
│   │   │   ├── account_handler.go
│   │   │   └── transaction_handler.go
│   │   ├── middleware                     # Echo Middleware
│   │   └── presenter                      # output portの実装
│   │       └── transaction_presenter.go
│   └── usecase　　　　　　　　　　 　　　　    # ユースケース層
│       ├── dto                            # input・output構造体
│       │   └── transaction.go
│       ├── interacter                     # ユースケースの実装
│       │   ├── account_interacter.go
│       │   └── transaction_interacter.go
│       └── port                           # ポート
│           ├── input                      # ユースケースのインターフェース（input port）
│           │   ├── account_input.go
│           │   └── transaction_input.go
│           └── output                     # presenterのインターフェース(output port)
│               └── transaction_output.go
├── migrations                             # sql-migrateで使用するsqlファイル
│   ├── 20240821230041-create-accounts.sql
│   └── 20240821230104-create-transactions.sql
└── pkg
    ├── apperr                             # アプリケーション固有のエラーハンドリング
    └── config                             # 環境変数を読み込むコード
        ├── config.go
        └── db_config.go

```

## 参考
Clean Architecture　達人に学ぶソフトウェアの構造と設計
<img width="1122" alt="image" src="https://github.com/user-attachments/assets/978b8f38-5f31-4acf-89f6-dcb56df86288">


## サンプルAPI
### アカウント作成
```
$ curl -X POST http://localhost:8080/open-account \
     -H "Content-Type: application/json" \
     -d '{"name": "テスト太郎", "balance": 10000}'
```

### 入金
```
$ curl -X POST http://localhost:8080/deposit \
     -H "Content-Type: application/json" \
     -d '{"id": 1, "amount": 1000}'
```

### 出金
```
$ curl -X POST http://localhost:8080/withdraw \
     -H "Content-Type: application/json" \
     -d '{"id": 1, "amount": 1000}'
```
