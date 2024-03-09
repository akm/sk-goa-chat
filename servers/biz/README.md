# biz

## ディレクトリ

パス     | 説明
--------|------------------
[Makefiles](./Makefiles/) | Makefile が include するターゲットを定義するファイル群
[models](./models/)  | sqlboiler で生成するモデル（とそのテスト）
[models_ext](./models_ext/) | モデルをカスタマイズする設定

## モデルを追加するステップ

(カレントディレクトリは servers/biz の前提)

1. DBマイグレーションを追加。詳しくは [dbmigrations](../dbmigrations/) を参照
2. `make sqlboiler_gen`
3. モデルに必要な設定を [models_ext/init.go](./models_ext/init.go) に追加
4. (TODO フィクスチャを追加するステップ)

## ビジネスロジックを追加するステップ

(TODO テスト込みのステップを追加)

## キャッチアップ

1. [Go言語](https://go.dev/)
2. [sqlboiler](https://github.com/volatiletech/sqlboiler)
3. [testify](https://github.com/stretchr/testify)
