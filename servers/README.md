# backend

## ディレクトリ

Name | 内容
---------|---------------------
[apisvr](./apisvr/) | APIサーバーの実装
[applib](./applib/) | アプリケーションレベルのライブラリ
[biz](./biz/)       | ビジネスロジック
[containers](./containers/)     | 開発に使用するコンテナとその設定群
[dbmigrations](./dbmigrations/) | RDBマイグレーション

### applib vs biz

[applib](./applib/) は Go言語の標準ライブラリのパッケージングや、外部のGoモジュールをベースにパッケージングされたライブラリ群です。

[biz](./biz/) は RDBとマッピングされるモデルなどの業務ドメインに属する定義を提供する。




## 開発環境の起動

```bash
make dev
```

### コンパイル

```bash
make build
```

### テスト

```bash
make test
```
