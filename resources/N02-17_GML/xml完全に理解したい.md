# 宿敵N02-17.xml

* gml:Curve
  * 曲線(直線)の緯度経度データ(配列)
    * 駅と区間の両方が含まれる
    * 駅も1区間として記録されている
      * [区間-区間-区間]ではなく、[駅-区間-駅-区間-駅]という記録のされ方
    * 駅がターミナルの場合複数の線で表される
  * LineStringSegment.xmlにトリミングした
  
* ksj:RailLoadSection
  * 路線データの情報
    * 路線名や区分など
  * ksj:locationにgml:Curveへのリンク
  * RailroadSection.xmlにトリミングした


* ksj:Station
  * 駅の情報
    * 駅名や会社名等
  * 1駅につき1路線のため重複がある
  * ksj:locationにgml:Curveへのリンクがある
  * ksj:railroadSectionにもリンクが
  * Station.xmlにトリミングした

## 重要でない部分

* Dataset
  * タグの仕様書へのリンク
* desctiption
  * 実質ファイル名
* boundedBy
  * 路線図全体の情報
    * lowerCorner:左下?
    * upperCorner:右上?

## TODO

* とりあえず地図上に重ねて表示させたい

* ksj:Station
  * 今回の乗り潰しマップを作成するにあたって、曲線である必要はなさそう(駅=点でいい)
  * gml:Curveから緯度経度を持ってきて重心を出し、それを配列としたい
  * 元の直線も(とりあえず)残しておきたい

* ksj:RailroadSection
  * 区間は曲線(複数の直線)で表示させたい
    * かっこわるいので(通信量やばかったらのちのち検討したい)
  * 区間の次の駅を検索するには緯度経度で調べるしかなさそうなので、1路線としてのテーブルを作成した方が後々楽そう

* gml:Curve
  * これもIDごとのテーブルにしたい
  * このままだと(Accessでimportすると)1つのテーブルになってしまうので変形が必要