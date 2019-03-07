# Windowsユーザーのための環境構築

## 必要環境
  * Windows 10 Pro or Education より上位
  * Windows Subsystem on Linux


## Windows 機能の有効化と無効化
* Hyper-V にチェックを付ける
  * もしすべての項目にチェックが両方とも入らなかったらBIOSをからCPU-Virtualization(VT、またはそれと同等のもの)をEnableにする
* Windows Subsystem on Linux にもチェックを付ける

## Dockerの導入

* Docker for Windows をインストール
	* https://hub.docker.com/editions/community/docker-ce-desktop-windows
* ESET を使ってるときは Firewall の設定をする
	* ポート445が拒否になってるので探してチェックを外す
	* https://qiita.com/toro_ponz/items/167250639211f264b43c
* Dockerが起動するので右クリック → Settings から Expose daemon on...のチェックを付ける
* Shared Drivesのチェックもつけとく

## リポジトリをクローン

* WSLを起動
* `$ sudo apt install git`
* `$ sudo apt install git-lfs`
* `$ git clone https://github.com/mjtakenon/noritubusi-map`
* `$ git lfs pull`

## backendの環境構築

* `$ sudo apt install docker-compose`
* `$ cd backend`
* `$ emacs ~/.bashrc`
  * 以下を追記
  * `export DOCKER_HOST='tcp://0.0.0.0:2375'`
* `$ docker-compose down --rmi all --volumes --remove-orphans`
  * 始めてdockerを使う際は不要
* `$ docker-compose build`
* `$ docker-compose up`
* `http://localhost:1323/`にアクセスして`Hello World!`と表示されればOK
* ただしこのままだとCORSエラーが出るので仮にCORS-Allowプラグインを導入する
  * `https://chrome.google.com/webstore/detail/allow-control-allow-origi/nlfbmbojpeacfghkpbjhddihlkkiljbi`

## frontendの環境構築

* 別のWSLを起動
* `$ sudo apt install npm`
* `$ sudo apt install yarn`
* `$ curl -L git.io/nodebrew | perl - setup`
* `$ emacs ~/.bashrc`
  * 以下を追記
  * `export PATH=$HOME/.nodebrew/current/bin:$PATH`
* `$ nodebrew install-binary v10.15.3`
* `$ cd noritubusi-map/frontend`
* `$ npm install`
* `$ npm run dev`
* `http://localhost:8080`にアクセスしてマップが表示されればOK