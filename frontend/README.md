# frontend

## 開発環境の起動

```bash
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open
```

## 開発用ビルド

```bash
npm run build
```

apisvr や関連するミドルウェアの開発用コンテナも起動します。


## プロダクションビルド

```bash
npm run preview
```

> To deploy your app, you may need to install an [adapter](https://kit.svelte.dev/docs/adapters) for your target environment.


## gRPC クライアント生成

```bash
npm run grpc:gen
```

## テスト

```bash
npm run test
```

単体テストのみ

```bash
npm run test:unit
```

integrationテスト(E2Eテスト)のみ

```bash
npm run test:integration
```

## 主な技術

- [SvelteKit](https://kit.svelte.jp/)
- [flowbite-svelte](https://flowbite-svelte.com/)
- [Tailwind CSS](https://tailwindcss.com)
- [素のHTMLの要素](https://developer.mozilla.org/ja/docs/Web/HTML/Element)
- [Playwright](https://playwright.dev/)
