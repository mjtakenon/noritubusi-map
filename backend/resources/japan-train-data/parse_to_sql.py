#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from attrdict import AttrDict
import json
import sys
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

    raw_data = [AttrDict(x) for x in json.load(open(path_to_json, encoding='utf-8'))]

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
    buildingsArray = np.array([list(buildings[i].values()) for i in list(buildings.keys())])
    stationsArray = np.array([list(stations[i].values()) for i in list(stations.keys())])

    # 統合処理済かどうかのフラグ
    buildingsArray = np.array([ np.append(b, False) for b in buildingsArray ])
    # ラベルに従って統合    
    deleteIdx = list()
    for b in buildingsArray :
        if b[4] :
            continue
        # 同名駅のリストを取得
        sameNameIdx = np.where(b[1] == buildingsArray[:,1])
        buildingsArray[sameNameIdx,4] = True
        # 緯度経度が近ければ統合
        # 統合する閾値(500m)
        mergeDistanceThreshold = 0.00545
        latlongs = np.array([p for p in buildingsArray[sameNameIdx,2][0]] , dtype=np.float64)
        labels = MeanShift(bandwidth = mergeDistanceThreshold).fit_predict(latlongs)
        
        for l in range(max(labels)+1):
            for d in range(np.sum(labels==l)-1):
                # 位置情報を統合(そのbuildingsを消し、stationsのbuildings_idが存在したら置き換える)
                idx = buildingsArray[sameNameIdx[0][np.where(labels==l)[0]]][d+1,0] 
                deleteIdx.append(idx)
                # stationsにbuilding_idが含まれていたら置き換える
                for a in np.where(stationsArray[:,3] == str(idx))[0]:
                    stationsArray[a,3] = str(buildingsArray[sameNameIdx[0][np.where(labels==l)[0]]][0,0])
    # あとでstationsのbuilding_idの更新?
    buildingsArray = np.delete(buildingsArray, np.array(deleteIdx)-1, 0) # idなので+1されているものをidxに変換する
    buildingsArray = np.delete(buildingsArray, 4, 1) # マージしたかのフラグは不要なので消す

    return buildingsArray, stationsArray, railways


if __name__ == '__main__':
    argc, argv = len(sys.argv), sys.argv

    if argc != 2:
        sys.stderr.write(f"usage: {argv[0]} [path to raw-data.json] \n")
        sys.exit(-1)

    else:

        buildings, stations, railways = proc(argv[1])
        # 生成した SQL クエリを 'seeds.sql' ファイルに出力

        with open('seeds.sql', 'wt') as f:
            for v in buildings:
                f.write(
                    SQL_INSERT_INTO(
                        'buildings',
                        ['id', 'name', 'latlong','connected_railways_num'],
                        [v[0], v[1], SQL_GEOMFROMTEXT(v[2]), v[3]], 
                        [0, 2, 3]
                    )
                )

            for v in railways.values():
                f.write(
                    SQL_INSERT_INTO(
                        'railways',
                        ['id', 'name'],
                        [v['id'], v['name']],
                        [0]
                    )
                )

            for v in stations:
                f.write(
                    SQL_INSERT_INTO(
                        'stations',
                        ['id', 'name', 'railway_id', 'building_id', 'num_in_railway'],
                        [v[0], v[1], v[2], v[3], v[4]],
                        [0, 2, 3, 4]
                    )
                )