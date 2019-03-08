# noritubusi-map (乗りつぶしマップ)

* 鉄路乗りつぶすオタク用ツール

## Features (in the future)

* [x] ブラウザで動く
* これらの情報を表示
  [x] 地図
  * 路線図
    * 乗り潰したところは色を変える
    * JRと私鉄でも変える
* 機能
  * 乗り潰し路線の登録/削除
  * 地図上で選択できるように

## やること

* -> Go to [Project page](../../projects/1)

## Requirements

- Docker, docker-compose
- Node.js ( 10.15.x )
- Git LFS

## Setup

```sh
$ git clone https://github.com/mjtakenon/noritubusi-map
$ cd noritubusi-map
$ git lfs pull
```

- バックエンド   → [doc/環境構築_BE.md](doc/環境構築_BE.md)
- フロントエンド → [doc/環境構築_FE.md](doc/環境構築_FE.md)


## 参考リンク

* 路線図の仕様
  (http://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-N02-v2_3.html)
* 路線図データをDBに登録
  (https://qiita.com/mima_ita/items/243922b3d5178f0315e0)
