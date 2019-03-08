# Windowsユーザーのための環境構築

## 必要環境
  * Windows 10 Pro or Education より上位
  * Windows Subsystem on Linux


## Windows 機能の有効化と無効化
* **Hyper-V** にチェックを付ける
  * もしすべての項目にチェックが両方とも入らなかったら BIOS をから CPU-Virtualization (VT-d、またはそれと同等のもの) を Enable にする
* **Windows Subsystem on Linux** にもチェックを付ける

## Dockerの導入

* Docker for Windows をインストール
	* https://hub.docker.com/editions/community/docker-ce-desktop-windows
* ESET を使ってるときは Firewall の設定をする
	* ポート 445 が拒否になってるので探してチェックを外す
	* https://qiita.com/toro_ponz/items/167250639211f264b43c
* Docker が起動するので右クリック → Settings から Expose daemon on... のチェックを付ける
* Shared Drives のチェックもつけとく

## リポジトリをクローン

* WSL を起動
```sh
# git, git-lfs のインストール
$ sudo apt install git git-lfs

# クローン & Git LFS ファイルのプル
$ git clone https://github.com/mjtakenon/noritubusi-map
$ git lfs pull
```

## 環境構築 (バックエンド)

```sh
# docker-compose のインストール
$ sudo apt install docker-compose
# DOCKER_HOST 環境変数の設定
$ echo "export DOCKER_HOST='tcp://0.0.0.0:2375'" > ~/.bashrc

# リポジトリの 'backend' ディレクトリに移動
$ cd noritubusi-map/backend

# 既存の docker イメージがある場合は削除 (初回の環境構築時はおそらく不要)
$ docker-compose down --rmi all --volumes --remove-orphans
# docker イメージのビルドと立ち上げ
$ docker-compose build
$ docker-compose up -d && docker-compose logs -f
```
* `http://localhost:1323/` にアクセスして `Hello World!` と表示されれば OK
* ただしこのままだとCORSエラーが出るので仮にCORS-Allowプラグインを導入する
  * https://chrome.google.com/webstore/detail/allow-control-allow-origi/nlfbmbojpeacfghkpbjhddihlkkiljbi

## 環境構築 (フロントエンド)

* 別のWSLを起動

```sh
# node.js 用バージョン管理ツール 'Nodebrew' のインストール
$ curl -L git.io/nodebrew | perl - setup
# Nodebrew のパスを通す
$ echo "export PATH=$HOME/.nodebrew/current/bin:$PATH" > ~/.bashrc

# node.js (v.10.15.3) を Nodebrew 経由でインストール
$ nodebrew install-binary v10.15.3
# node.js が入っているか確認
$ node -v
# 'v.10.15.3' と表示されれば OK
# node.js 用パッケージ管理ツール 'yarn' のインストール
$ npm install --global yarn

# リポジトリの 'backend' ディレクトリに移動
$ cd noritubusi-map/frontend

# yarn による依存パッケージインストール
$ yarn

# 開発用フロントエンドサーバー立ち上げ
$ yarn serve
```
* `http://localhost:8080` にアクセスしてマップが表示されればOK
