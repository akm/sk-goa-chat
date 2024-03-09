# applib

## 概要

applib は Go言語の標準ライブラリのパッケージングや、外部のGoモジュールをベースにパッケージングされた、このアプリケーション固有のライブラリ群です。

## パッケージング

ここではパッケージングとは、ディレクトリ名とパッケージ名の決定を指します。
既存のライブラリの拡張であれば、そのパッケージングに則ってパッケージングを行ってください。既存ライブラリの拡張ではなく、新たな機能については、他のパッケージと重複しないように、パッケージングしてください。

新しい機能は通常何らかの他のパッケージや言語機能を使って実装するはずですので、何を以て新しいと判断するのかは難しいかもしれませんが、アプリケーション固有のライブラリであり、外部と共有されるものではないので、関係する開発者の間で合意できれば良いです。

## 何を含めるべきではないか

- ビジネスロジック
- sqlboiler が生成するモデルに依存するもの
- goa が生成する gen ディレクトリに依存するもの

## 何を含めるべきか

- ビジネスロジックに関係しない処理
    - 例えば標準ライブラリの拡張
    - 例えば外部ライブラリの使用範囲を限定する定義
- テストに関する処理
    - フィクスチャは biz で定義すること
- 抽象的な処理
    - ジェネリクスやインターフェース等を使って対象を抽象化した処理

## パッケージの命名

### 既存のライブラリの拡張/限定

import の際に指定するパッケージのパスが `applib/(拡張対象のパッケージ)` となるように定義します。例えば [time](https://pkg.go.dev/time) パッケージ を拡張/限定するパッケージの場合は `applib/time` という名前にします

### 新しい機能の追加

関連するパッケージが既に存在する場合は、そのパッケージのディレクトリ構造に含まれるように定義します。例えば、HTTPのCORSに関する機能の追加は、HTTPに関する機能が標準ライブラリに `net/http` があるので `net/http/cors` として追加します。

関連するパッケージが存在しない場合は、新たにディレクトリを追加します。例えば `applib/collection` はジェネリクスを使ってスライスを操作する関数を定義するパッケージで、言語機能のみを利用し、何も import していません。ですので、他のパッケージの名前とは関係しない collection という名前を付けられました。

### パッケージ名に使ってはいけない語句

- common
- util

## エイリアス

既存のライブラリの拡張/限定 を行うためにエイリアスを定義してください。

例えば applib/time/aliases.go には 標準ライブラリの time パッケージで定義される型、定数、関数（変数に代入）が再定義されています。このエイリアスを追加しておくと、applib/time パッケージが import されたコードでは、エイリアスが定義された標準ライブラリの機能を使うことができます。

単に標準の機能を利用するだけならばエイリアスの追加も必要もなく、そのまま標準ライブラリを使えば良いのですが、エイリアスを使う理由は他の機能を拡張したい場合です。

例えば applib/time では Now 関数が拡張されています。Now関数はテスト時には決まった時刻を返すようなモックに差し替えることができるようになっています。しかしNow関数を使う方は time.Now と同じように使いたいものです。

ですので time を import して使っているコードがあった場合、そのimport を applib/time に変えるだけで、テスト時にはモックも使えるものになるのが理想的だと思います。エイリアスはこの際に Now 以外のtimeパッケージの提供するものを辻褄を合わせるための仕掛けだと考えると分かりやすいかもしれません。
（import を解析して applib/time ではなく time パッケージを使っていた場合には警告やエラーにするlintの機能を追加したいと思っています）


### インターフェースのエイリアス

例えばFirebaseのような外部サービスを使う場合、そのクライアントライブラリはたくさんの機能を提供することがあります。それらを使ったコードのテストを行う場合、実際に呼び出すことができない機能については、テスト用にモックを作成する必要があります。モックは使用するメソッドを持つインターフェースを定義し、[github.com/stretchr/testify/mock](https://pkg.go.dev/github.com/stretchr/testify/mock) などを利用して、そのインターフェースを実装して作成します。

この際、ライブラリが提供するクライアントが提供するメソッドすべてを実装していたらキリがありません。テストのビルドのためだけに、クライアントにメソッドが追加されたら、そのインターフェースにメソッドを追加するような仕事はしたくありません。Goでは Java などに見られる implements キーワードは存在しませんので、このような仕事はする必要がありません。

[Go の interface は構造体の利用側が定義すると言う話](https://engineering.mobalab.net/2021/10/04/where-should-interface-defined-in-go/) で紹介されている [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments#interfaces) にあるように、機能を持つ構造体を提供する側がインターフェースを定義するのではなく、その構造体を使う側が定義することを前提としています。ですので、。テストに必要なインターフェースも最小限に留めることができます。