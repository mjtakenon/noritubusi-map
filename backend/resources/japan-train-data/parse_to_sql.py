#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from attrdict import AttrDict
import json
import sys
from os import path
from collections import Counter, defaultdict
import numpy as np
from sklearn.cluster import MeanShift


"""
Key-Value データから INSERT INTO 文を生成する

Parameters
----------
TBL_NAME : str
    INSERT INTO 文の対象となるテーブル名

KEYS : list[str]
    INSERT INTO 文で挿入されるカラム名
    KEYS に含まれるカラム名は TBL_NAME で指定したテーブルに
    存在している必要がある

VALUES : list[str]
    INSERT INTO 文で挿入されるレコード値
    VALUES の要素順は KEYS の要素順に対応している必要がある

IDX_NOT_QUOTED : list[int]
    VALUES 句において、クォーテーションを付けない要素
    のインデックス番号を指定する。デフォルト値は [] 
    SQL では int 型や関数の出力結果を格納する場合は、
    クォーテーションがあるとエラーとなる。

Returns
-------
query : str
    生成した INSERT INTO 文
"""


def SQL_INSERT_INTO(TBL_NAME, KEYS, VALUES, IDX_NOT_QUOTED=[]):
    str_keys = ', '.join([f"`{x}`" for x in KEYS])
    str_values = ', '.join(
        [f"'{x}'" if i not in IDX_NOT_QUOTED else str(x) for (i, x) in enumerate(VALUES)])

    return f"INSERT INTO `{TBL_NAME}` ({str_keys}) VALUES ({str_values});\n"


"""
緯度, 経度の配列から MySQL の GeomFromText 関数表記を得る

Parameters
----------
latlong : list[float, float]
    緯度, 経度の配列

Returns
-------
query : str
    GeomFromText 関数表記の文字列
"""


def SQL_GEOMFROMTEXT(latlong):
    return "GeomFromText('POINT({0:.8f} {1:.8f})')".format(*latlong)


def proc(path_to_json):

    raw_data = [AttrDict(x) for x in json.load(
        open(path_to_json, encoding='utf-8'))]

    buildings = dict()
    railways = dict()
    stations = dict()

    railways_id, buildings_id, stations_id = 1, 1, 1

    # prefecture: 県
    for prefecture in raw_data:
        print(f" Prefecture[{prefecture.name.ja}] ... \r")
        # sys.stderr.flush()
        # railway: その県が持つ路線
        # このデータは
        # 「県」→「その県がもつ路線」→「その路線の駅とその駅にある他の路線」
        # という構造をしているため、同じ路線でも複数の都道府県を跨ぐ路線の場合
        # 同じ路線が複数回参照されることがあるため、ID処理が多少複雑になる
        for railway in prefecture.lines:

            # もし一度も railways に登録してない路線だった場合
            if railway.name.ja not in railways:
                # railways に登録
                railways.update(
                    {
                        railway.name.ja: {
                            'id': railways_id,
                            'name': railway.name.ja,
                        }
                    }
                )
                # ref_railways_id: stations で利用する railway_id の値
                ref_railways_id = railways_id
                # railways の id をインクリメント
                railways_id += 1
            # すでに一度登録している場合
            else:
                # ref_railways_id を登録されているデータから参照してセット
                ref_railways_id = railways[railway.name.ja]['id']

            # station: その県に所属する任意の路線が持つ駅(駅舎)
            # このデータは
            # 「その県がもつ路線」→「その路線の駅とその駅にある他の路線」
            # という構造をしているため、ある駅舎が複数路線を持つ場合、駅舎が
            # 複数回参照されることがあるため、ID処理が多少複雑になる
            for (num_in_railway, station) in enumerate(railway.stations, start=1):

                # stations には駅の緯度経度をキーとして登録する
                latlong_station = tuple(
                    [station.location.lat, station.location.lng])

                # もし一度も buildings に登録していない駅舎だった場合
                if latlong_station not in buildings:
                    # buildings に登録
                    buildings.update(
                        {
                            latlong_station: {
                                'id': buildings_id,
                                'name': station.name.ja,
                                'latlong': latlong_station,
                                'connected_railways_num': len(station.lines)
                            }
                        }
                    )
                    # ref_buildings_id: stations で利用する building_id の値
                    ref_buildings_id = buildings_id
                    # buildings の id をインクリメント
                    buildings_id += 1
                # すでに一度登録している場合
                else:
                    # ref_buildings_id を登録されているデータから参照してセット
                    ref_buildings_id = buildings[latlong_station]['id']
                # stations への登録
                # railway.stations はその路線の全駅のリストだが、
                # そのなかの各駅には、「その駅がある駅舎にある他の路線」について
                # の情報を持っているが、他の路線については、別の railway がフォーカス
                # されたときに登録されるので、ここでは現在の路線と一致する駅のみを登録する
                stations.update(
                    {
                        tuple([station.name.ja, railway.name.ja]): {
                            'id': stations_id,
                            'name': station.name.ja,
                            'railway_id': ref_railways_id,
                            'building_id': ref_buildings_id,
                            'num_in_railway': num_in_railway
                        }
                    }
                )
                # stationsのidをインクリメント
                stations_id += 1

    # 近い位置にある同名駅の統合
    sys.stderr.write("Merged Stations ... \n")

    # 統合する閾値(500m) : 緯度経度が近ければ統合
    mergeDistanceThreshold = 0.00545

    # building_id_map : 統合後の building_id 変換マップ
    #   Key : buildings_id
    #   Value : 変換後の building_id
    building_id_map = dict()

    sys.stderr.write("  Count building_name duplication ... \n")
    # 1. 駅名称の重複カウント
    #   collection.Counter はリスト要素の重複数をカウントしてくれる
    #   これを利用して、駅名 ( 'stationName' ) の重複カウントを行う
    count = Counter([v['name'] for v in buildings.values()])

    # 2. 重複のないものを処理
    sys.stderr.write("  Proc non-duplicated buildings ... \n")
    # non_duplicated_names : 重複のない駅名リスト
    non_duplicated_names = [k for k, v in count.items() if v == 1]

    for non_dup_name in non_duplicated_names:
        # 駅名から building_id を取得
        building_id = [v['id']
                       for v in buildings.values() if v['name'] == non_dup_name][0]

        # 重複のない駅は、building_id の変換の必要がないので
        # 同じ値を返すようにする
        building_id_map[building_id] = building_id

    # 3.  重複のあるものを処理
    sys.stderr.write("  Proc duplicated buildings ... \n")
    # duplicated_names : 重複のある駅名リスト
    duplicated_names = [k for k, v in count.items() if v != 1]

    # import ipdb
    # ipdb.set_trace()

    # 重複駅名リストを駅名ごとに処理
    for dup_name in duplicated_names:

        # dup_buildings : 同一名称の駅舎情報リスト
        #   buildings は Key に 駅舎の緯度経度、Value に駅舎情報を格納している
        #   buildings の Value にある駅名 (name) が一致するもののみ抽出
        dup_buildings = [
            v
            for v in buildings.values()
            if v['name'] == dup_name
        ]

        # latlongs : 同一名称駅舎の緯度経度のリスト
        #   dup_buildings から緯度経度 (latlong) のみを抽出
        #   scikit-learn の MeanShift が numpy.ndarray 型しか受け付けないため
        #   numpy.ndarray 型にキャスト。データ深度は np.float64 (倍精度浮動小数点)
        latlongs = np.array(
            [v['latlong'] for v in dup_buildings],
            dtype=np.float64
        )

        # scikit-learn の MeanShift でクラスタリング
        #   引数 bandwidth は統合する際の閾値
        labels = MeanShift(
            bandwidth=mergeDistanceThreshold).fit_predict(latlongs)

        # clustered_building : クラスタ番号ごとに building をまとめる
        #   Key は MeanShift のラベリング番号, Value は同じラベルに
        #   クラスタリングされた building のリスト
        #   defaultdict は dict にデフォルト値をセットして初期化する
        clustered_building = defaultdict(list)

        # zip 関数で順序を維持してループ
        #   label_num : クラスタ番号
        #   building  : 対応する駅舎
        for label_num, building in zip(labels, dup_buildings):
            # 該当するクラスタに駅舎情報を追加する
            clustered_building[label_num].append(building)

        # クラスタごとに、building_id の統一処理を行う
        #   merged_buildings : 各クラスタの駅舎集合
        for merged_buildings in clustered_building.values():

            # building_id を各クラスタの最初のものに統一する
            #   value : 各クラスタの最初の要素の building_id
            #           (変換先の値)
            value = merged_buildings[0]['id']

            # クラスタ内の全駅舎に対しループを行い
            for building_id in [v['id'] for v in merged_buildings]:
                # building_id 変換マップに値をセットする
                #   building_id : 変換元の値
                building_id_map[building_id] = value

    # buildings の重複削除
    sys.stderr.write("  Delete duplicated buildings ... \n")
    # merged_buildings : 重複を削除した buildings
    #   building_id_map の Value 集合をもとにループを回すことで
    #   変換後の building_id の駅舎のみで構成された buildings が
    #   生成できる
    merged_buildings = {
        building_id: [v for v in buildings.values() if v['id'] ==
                      building_id][0]
        for building_id in building_id_map.values()
    }

    # stations の building_id 書き換え
    sys.stderr.write("  Modify station['building_id'] ... \n")
    # stations は building_id を持つので、変換表を使って
    # 変換後の値に変更する
    for station in stations.values():
        station['building_id'] = building_id_map[station['building_id']]

    # buildings -> merged_buildings に変更して Return
    return merged_buildings, stations, railways


if __name__ == '__main__':

    # 引数処理
    argc, argv = len(sys.argv), sys.argv

    if argc != 2:
        sys.stderr.write(f"usage: {argv[0]} [path to raw-data.json] \n")
        sys.exit(-1)
    else:
        path_to_data = argv[1]

    # データ処理
    buildings, stations, railways = proc(path_to_data)

    # 生成した SQL クエリを 'seeds.sql' ファイルに出力
    with open('seeds.sql', 'wt') as f:
        # buildings
        for v in buildings.values():
            f.write(
                SQL_INSERT_INTO(
                    'buildings',
                    ['id', 'name', 'latlong', 'connected_railways_num'],
                    [v['id'], v['name'], SQL_GEOMFROMTEXT(v['latlong']),
                        v['connected_railways_num']],
                    [0, 2, 3]
                )
            )
        # railways
        for v in railways.values():
            f.write(
                SQL_INSERT_INTO(
                    'railways',
                    ['id', 'name'],
                    [v['id'], v['name']],
                    [0]
                )
            )
        # stations
        for v in stations.values():
            f.write(
                SQL_INSERT_INTO(
                    'stations',
                    ['id', 'name', 'railway_id',
                        'building_id', 'num_in_railway'],
                    [v['id'], v['name'], v['railway_id'],
                        v['building_id'], v['num_in_railway']],
                    [0, 2, 3, 4]
                )
            )
