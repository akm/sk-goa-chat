## Stage type

- production
- staging
- local
- github

Stage type のインスタンスとして Stage が複数個存在し、リポジトリにそれらの違いをコミットしなければならない場合もありえます。
例えば staging に対して staging1, staging2 それぞれで設定が異なる場合が考えられます。
逆に local のインスタンスとして各開発者の環境はそれぞれ別物ですが、それらの違いをコミットする必要がなければ local1, local2 
というような Stage を登録する必要はありませんし、通常はそのように行うべきではありません。
ただし Stage type として local_windows と  local_mac のような区別を行った方が良い場合もあります。

## ENV_TYPE

(正確に言えば process cluster type みたいな名前のはずですが分かりやすい名前じゃないので ENV_TYPE にしました)

- server
- dev
- ui/test:integration
- ui/test:unit
- apisvr/test

ほとんど名前が説明してくれていると思います。
ui/test:integration は SvelteKit が標準でサポートするテストで、Playwrightを使ったテストです。
これは一般的には E2Eテストと呼ばれるものです。
staging/production では通常、ドメイン名やロードバランサーなどをを含めたテストになりますが、実行するテストのシナリオ等は
基本的に同じものを使用します。

## Ports

ENV_TYPE            | STAGE_TYPE               | apisvr HTTP | apisvr gRPC | ui HTTP  | mysql    | firebase authentication
--------------------|--------------------------|-------------|-------------|----------|----------|------------------------
server              | staging,production       | 8000        | 8080        | 4173     | 3306     | ?
dev                 | local                    | 8000        | 8080        | 5173     | 3306     | 9099
ui/test:integration | staging,production       | 8000        | 8080        | 4173     | 3306     | ?
ui/test:integration | local,github             | 8001        | 8081        | 4173     | 3307     | 9090
ui/test:unit        | local,github             | -           | -           | -        | -        | -
apisvr/test         | local,github             | -           | -           | -        | 3311     | 9091

staging や production に対する ui/test:integration では、起動済みのサーバーに対してアクセスするので、新たにサーバーを起動することはありません。
