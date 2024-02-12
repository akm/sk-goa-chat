# frontend/src/lib

## ディレクトリ

大前提としてディレクトリ構造は [SvelteKitのルール](https://kit.svelte.jp/docs/modules#$lib-$lib-server) に従います。

| ディレクトリ名 | 実行場所        | 内容                  |
| -------------- | --------------- | --------------------- |
| firebase       | ブラウザ, uisvr | firebase に関するもの |
| models         | ブラウザ, uisvr | frontend上でのモデル  |
| server         | uisvr           | サーバー上で動く処理  |

## APIクライアント

frontend でのAPIクライアントには以下の2種類あります。

- SSR(Server-Side Rendering) で使われる gRPC クライアント
- ブラウザ上で動作するHTTPのOpenAPIクライアント

### gRPCクライアント

gRPC クライアントは uisvr 上で動作する API クライアントです。
[SvelteKitのルール](https://kit.svelte.jp/docs/modules#$lib-$lib-server) に従い [frontend/src/lib/server](./frontend/src/lib/server/) 以下に定義する必要があります。

`npm run grpc:gen` で生成できます。
gRPC クライアントはリソース毎に生成されるので、リソースが追加する場合には frontend/Makefiles/grpc.mk に定義されている `GRPC_RESOURCES` に対象を追加してください。

https://github.com/akm/sk-goa-chat/blob/a82020596c74a94389c72ef39b4fc57d26717587/frontend/Makefiles/grpc.mk#L4

### HTTPクライアント (あるいは OpenAPIクライアント)

HTTPクライアント はブラウザ上で動作する API クライアントです。
現時点では HTTP クライアントを生成するタスクは用意されていませんが、近い将来準備する予定です。
