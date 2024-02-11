# frontend/tests/integration

## 概要

SvelteKit によって生成される このディレクトリは ブラウザを使ったテストのためのものです。
以下の２つのケースで使い方を想定しています。

- ローカル環境でのブラウザを使ったテスト
- 本番環境あるいはステージング環境でのブラウザを使った E2Eテスト

この自動テストは画面を操作して、期待した動作を行うこと（あるいは行わないこと）を確認するプログラムです。
ブラウザの操作は [Playwright](https://playwright.dev) を使って行います。

このテストのプログラムは、シナリオと [Page Object Models (POM)](https://playwright.dev/docs/pom) で構成されます。
シナリオでは、定義したPOMを用いて、ブラウザ上に表示された内容の確認と、操作を行います。

## Page Object Models (POM)

[Page Object Models](https://playwright.dev/docs/pom) は ページやその中の部品をテストしやすいように定義したオブジェクト群です。
POMは通常 [Locators](https://playwright.dev/docs/locators) を用いて実装されます。
Locator の種類としては [Other locators](https://playwright.dev/docs/other-locators#css-locator) に挙げられているものがありますが、
主に CSS locator を用います。CSS locator に対しては、セレクタとして 一般的な [CSSセレクタ](https://developer.mozilla.org/ja/docs/Web/CSS/CSS_selectors) に加え Playwright が拡張したセレクタの記法を使うことができます。


## ディレクトリ

Name | 内容
---------|---------------------
[containers](./containers/) | test/integrations をローカルで実行する際に使うコンテナ
[log](./log/)               | test/integrations をローカルで実行する際の apisvr と uisvr のログの出力先
[pom](./pom/)               | Page Object Models の定義

## キャッチアップ

- [CSSセレクタ](https://developer.mozilla.org/ja/docs/Web/CSS/CSS_selectors)
- [Playwright](https://playwright.dev)
