# 環境構築 (フロントエンド)

```sh
$ cd frontend
```

- `vue-cli` で雛形作成
  - ~~`vuejs-templates/webpack` テンプレートをベースに作成~~
  - ~~REF: [Vue-cli(webpack)解剖 ーディレクトリ構成ー - Qiita](https://qiita.com/h_plum/items/86b8a6a86ac0fea8a4d1)~~

- **追記(2019/03/08)**
  - `vuejs-templates/webpack` で使用している `webpack-dev-server` に脆弱性があったため、別のテンプレートに置き換えました
  - 以前の環境を使用している場合は、再度 `yarn install` をやり直してください
  - また、いくつかコマンドが変更されています(開発用サーバー起動コマンドなど)
  ```sh
  # 念のため、'node_modules' ディレクトリを削除
  $ rm -rf node_modules
  # 依存パッケージを再インストール
  $ yarn
  # 開発用サーバーを起動
  $ yarn serve
  ```

## Requirements

- node.js ( ≥ v.10.15.x )
  - npm
  - yarn


## Setup

```sh
# 'yarn' による依存パッケージのインストール
$ yarn

# 開発用サーバーの起動
$ yarn serve
```

## Features

- Scripts
```
# 開発用サーバーの起動
$ yarn serve

# デプロイ用にビルドする
# 生成物は 'frontend/dist' ディレクトリに生成される
$ yarn build

# 構文チェック
# 修正する場合は '--fix' をつける
$ yarn lint
```
