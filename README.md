# 乗り潰しマップ自分用メモ

* 鉄路乗り潰すオタク用ツール

## 仕様

* ブラウザで動く
* これらの情報を表示
  * 地図
  * 路線図
    * 乗り潰したところは色を変える
    * JRと私鉄でも変える
* 機能
  * 乗り潰し路線の登録/削除
  * 地図上で選択できるように

## やること

* 地図と路線図の表示
  * 地図はとりあえずgoogle map api
  * 路線図をDBに格納して動的に読み込みたい
    * 通信量が ...

* 詳細は issue を見て

## 環境構築

- バックエンド   → [doc/環境構築_BE.md](doc/環境構築_BE.md)
- フロントエンド → [doc/環境構築_FE.md](doc/環境構築_FE.md)

- ~~`create-dev-environment` ブランチにて展開中 ...~~
  + `create-db-datasets` にて作業中 ... (2019/01/07 15:14)

- Git LFSについて
  + `config/mysql/init.d/2_seeds.sql`をGitLFS管理にしました。
  + このブランチを取り込む前にGitLFSを導入してください
    + https://blog.amedama.jp/entry/2017/11/19/091626

## 参考リンク

* 路線図の仕様
  (http://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-N02-v2_3.html)
* 路線図データをDBに登録
  (https://qiita.com/mima_ita/items/243922b3d5178f0315e0)
