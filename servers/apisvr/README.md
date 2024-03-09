# Chat API Server

## ディレクトリ

パス     | 説明
--------|------------------
[Makefiles](./Makefiles/) | Makefile が include するターゲットを定義するファイル群
[design](./design)  | goa のAPI定義のファイル群
[services](./services/) | goaのサービスのファイル群と、cmd ディレクトリ と gen ディレクトリ


## 何をどこに実装するのか

優先順位 | 条件  | 場所
------:|---------|--------------
1 | services/gen 以下に依存するもの | [services](./services/)
2 | モデルに依存するもの | [biz](../biz/)
3 | ビジネスロジック・業務ドメインに依存するもの | [biz](../biz/)
4 | システム内の apisvr 以外の処理で使用する処理 | [applib](../applib/)
5 | 複数のシステムで使用する処理 | [applib](../applib) or 他のリポジトリ

## APIを追加するステップ

(カレントディレクトリは servers/apisvr の前提)

1. [design](./design/) に定義を追加
2. `make goa_gen` で [services/gen](./services/gen) を再生成
3. `make goa_example` で [services/cmd](./services/cmd/) を再生成し、 [services/](./services/) にサービスのファイルを生成
4. サービスを定義するファイルに Convertor を仮実装する
    - データは適当で良い
5. [services](./services/) 以下にテストのファイルを追加
    1. Convertor のテストを定義
    2. Convertor を使ったサービス自身のテストを定義
6. サービスのメソッドの内容を実装
    1. モデルを生成。詳しくは [biz](../biz/) を参照
    2. Convertor を実装
    3. サービスのメソッドを実装

## APIテスト実装

APIの自動テストには３つの種類があります。

- 単体テスト
- HTTPでの統合テスト
- gRPCでの統合テスト

どのテストもRDBに接続したり、必要に応じてFirebase等の外部サービスのモックコンテナにアクセスします。
ビジネスロジックを確認する基本的なテストは単体テストで行います。これはHTTPとgRPCのどちらにも依存しません。
HTTPでの統合テストを行うのは、URL中のクエリ文字列の扱いを確認したり、リクエストやレスポンスのボディが
本当に期待通りになるかどうかを確認する場合です。
gRPCの場合も同様ですが、gRPCの場合は Goa が生成する専用のクライアントを用いるのでHTTPのようなプロトコル
に依存したテストは少ないかもしれません。

## キャッチアップ

1. [Go言語](https://go.dev/)
2. [Goa](https://goa.design/)
    - [godoc](https://pkg.go.dev/goa.design/goa/v3)
    - [examples](https://github.com/goadesign/examples)
    - [Goa v3 入門](https://zenn.dev/ikawaha/books/goa-design-v3)
3. [testify](https://github.com/stretchr/testify)
4. [goahttpcheck](https://github.com/ikawaha/goahttpcheck)
