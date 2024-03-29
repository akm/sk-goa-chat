# 開発の進め方

## 概要

- 進め方
- 作業タスク
    - 画面仮実装
    - API仮実装
    - 画面APIクライアント組み込み
    - DBマイグレーション追加
    - APIテスト実装
    - API実装
    - 画面実装
    - E2Eテスト実装

## 進め方

進め方には以下のような進め方があります。

- 画面ドリブン
- APIドリブン
- DBドリブン

これらは作業タスクの実行順番が異なるだけですが、どのような進め方をするのかによって作業にかかる手戻りは大きく変わると思います。

### 画面ドリブン

最も基本的で堅実な進め方です。画面の仕様に一つでも不明な点がある場合にはこの方法がおすすめです。

1. 画面仮実装
2. 画面レビュー
3. API仮実装
4. 画面APIクライアント組み込み
5. DBマイグレーション追加
6. APIテスト実装
7. API実装
8. 画面実装
9. E2Eテスト実装

ポイントは 画面レビュー です。仕様が決まっていない場合、誰かに本当にこの動きで良いのかを確認する必要がありますが、
言葉やワイヤーフレームだけで仕様を正確に把握できる開発者や仕様の責任者は稀です（同じ人でもそのときの状況によっては
イメージできない状況のこともあるかもしれません）。ですので、紙芝居のように実際のデータとは連携しなくても実際に動く
画面で、やろうとしていることを説明するのは非常に効果的です。

ですので画面仮実装の時点で、表示のために必要な（本来であればAPI等で取得する）データを一旦はハードコーディングしたり、
ボタンやリンクのクリックなどについては、ダミーの処理を入れて画面遷移をするように実装します（このデータはAPI仮実装
の参考になります）。

画面の仮実装が終わった時点で、ステージング環境にデプロイを行い、画面レビューができる状態にします。

APIの実装に時間がかからない場合は。API仮実装は行わない選択も可能ですが、時間かからないとは言い切れない場合には
仮実装を行った方が良いでしょう。API仮実装では、画面仮実装でハードコーディングしたデータをそのままAPIで返すように
実装します。

API仮実装ができたら、画面にAPIクライアントを組み込むことができます。

バックエンド側は必要に応じてDBマイグレーション等を行いつつ、APIテスト実装とAPI実装を行います。
フロントエンド側は、その間に、見た目の細かな調整等を含む画面実装を行います。

最後にそれらが連携して動作することを確認するために E2Eテストを実装します。


### APIドリブン

ほとんど画面ドリブンと同じ進め方ですが、画面仮実装を行わずAPI実装を先に進める方法です。
API定義を行う際にはフロントエンド側がどのように画面を操作して、どのようにAPIを使用するのかをイメージする必要がありますが、
実際に作ってみてわかることも多々ありえるので、あまりお勧めの方法ではありません。


### DBドリブン

例えば管理者しか使わない画面で、あまり見た目にこだわりがない場合に使える進め方です。
データベースのレコードを追加変更削除できるようにAPIを用意して、画面の実装を行います。

1. DBマイグレーション追加
2. API仮実装
3. APIテスト実装
4. API実装
5. 画面実装
6. E2Eテスト実装

Ruby on Rails の scaffold と呼ばれる機能を使う場合と同じような進め方です。
Ruby on Rails の場合は、マイグレーション の作成と実行を行った後は scaffold を呼び出すと
残りの実装が行われますが、そのような機能はまだないので、それぞれ手動で実装する必要があります。


## 作業タスク

### 画面仮実装, 画面実装

基本的な画面遷移を行うためには [SvelteKit](https://kit.svelte.jp/) でのルーティング等を理解する必要があります。
次に画面に配置するコンポーネント ( あるいは [素のHTMLの要素](https://developer.mozilla.org/ja/docs/Web/HTML/Element) ) として何を使うのかについては
(HTMLの基本的な理解と) [flowbite-svelte](https://flowbite-svelte.com/) のコンポーネントへの理解が必要です。
またそのレイアウトについては flowbite が使用する [Tailwind CSS](https://tailwindcss.com/) を理解する必要があります。

詳しくは [clients/uisvr/README.md](./clients/uisvr/README.md) を参照してください。

### API仮実装

APIを定義するためには [HTTP](https://developer.mozilla.org/ja/docs/Web/HTTP/Basics_of_HTTP) あるいは [gRPC](https://grpc.io/) （できれば両方）の基本的な理解が必要です。
その上で [Goa](https://goa.design/) の [DSL](https://pkg.go.dev/goa.design/goa/v3@v3.14.6/dsl) を理解する必要がありますが、この [DSL](https://pkg.go.dev/goa.design/goa/v3@v3.14.6/dsl) がどのように使われるのかは、
Goaのアプリケーションを使ってみるのが一番良いです。また [examples](https://github.com/goadesign/examples) も参考になります。
また有料ですが [Goa v3 入門](https://zenn.dev/ikawaha/books/goa-design-v3) が日本語で解説してあり、最も分かりやすいと思います。

詳しくは [servers/apisvr](./servers/apisvr/) を参照してください。


### 画面APIクライアント組み込み

Goaが生成するAPI定義から、現時点ではgRPC のクライアントのみを生成しておりますが、
OpenAPIのクライアントも生成するべきです。OpenAPI でのAPI定義については、
[Swagger UI](https://swagger.io/tools/swagger-ui/) で確認することができます。

詳しくは [clients/uisvr/src/lib](./clients/uisvr/src/lib/) を参照してください。


### DBマイグレーション追加, APIテスト実装, API実装

RDBのデータマイグレーションには [goose](https://github.com/pressly/goose) を用います。
データベースのデータの操作には [sqlboiler](https://github.com/volatiletech/sqlboiler) を用いて生成されたモデルを通じて行います。
ですので、テーブルの作成やカラムの変更を行うマイグレーションを追加した場合、モデルを生成し直す必要があります。

詳しくは [servers/biz](./servers/biz/) と [servers/apisvr](./servers/apisvr/) を参照してください。



### E2Eテスト実装

E2Eテストは clients/uisvr 内の tests/integration に [Playwright](https://playwright.dev/) を使って実装します。

詳しくは [clients/uisvr/tests/integration](./clients/uisvr/tests/integration/) を参照してください。
