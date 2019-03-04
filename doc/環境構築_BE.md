# 環境構築 (バックエンド)

```sh
$ cd backend
```

## Requirements

- Docker

## Setup

#### 路線データの流し込み ( `N02-17.xml` )

- Requirements: Python 3.x, numpy

```bash
# シード用 SQL ファイルの生成
$ cd resouces/N02-17_GML
$ ./N02-17_parse.py N02-17.xml
# シードファイル を コンテナ立ち上げ時に流し込むため、initdb.d ディレクトリに移動
$ mv N02-17_seeds.sql ../../config/mysql/initdb.d/2_seeds.sql

# backend ディレクトリに戻る
$ cd ../../
# 既存イメージ, Volume の削除
# 過去に docker-compose で環境構築を行っていない場合、不要？
$ docker-compose down --rmi all --volumes --remove-orphans
# コンテナのビルド
$ docker-compose build
# コンテナ立ち上げ
$ docker-compose up -d
# CLI 上からデータ流し込みの確認
# root パスワードは docker-compose.yml を確認
docker-compose exec mysql mysql -u root -p
Enter password: ...

mysql> USE noritubusi_map;
mysql> SELECT `id`, `station_name`, X(center_latlong) AS `lat`, Y(center_latlong) AS `long`, `operation_company`, `service_provider_type`, `railway_line_name`, `railway_type` FROM stations';
```

### Docker イメージのビルド

```sh
$ docker-compose build
```

## Run

```sh
$ docker-compose up -d 
```

## Tips -- Docker

- エイリアスを貼っておくと楽？
```sh
$ alias d='docker'
$ alias dc='docker-compose'
```

- docker-compose: 起動と終了
```sh
# 標準起動の場合(標準出力がロギングのために奪われる)
$ docker-compose up 
# デーモン起動の場合
$ docker-compose up -d

# 終了
$ docker-compose down
```

- docker-compose: ログ表示
```sh
# 全体ログ表示
$ docker-compose logs
# 全体ログ表示(実行後に生成されたログも表示する)
$ docker-compose logs -f
# 特定のコンテナに限定してログ表示(例: app コンテナ)
$ docker-compse -f app
```

- docker-compose: 標準入力にアタッチ
  - デバッグ時のインタプリタ入力にアタッチする際に使用するかも
```sh
$ docker attach $(docker-compose ps -q )
```

- docker イメージの削除
```
# ビルド失敗時に生成された一時イメージの削除
$ docker rmi --force $(docker images -q --filter "dangling=true")

# 全イメージの削除
$ docker-compose down --rmi all --volumes --remove-orphans
```

