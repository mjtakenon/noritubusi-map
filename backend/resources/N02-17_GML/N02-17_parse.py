#!/usr/bin/env python3

import os
import sys
import re
import json
import itertools
from collections import Counter
import xml.etree.ElementTree as ET
import numpy as np
from sklearn.cluster import MeanShift

"""
要素から "gml:id" 属性の値を取り出す

Parameters
----------
xml_elem : xml.etree.Element
    属性値を取り出す対象となる XML 要素

namespace_map : dict
    XML 名前空間を解決するための連想配列
    Key に「名前空間」、Value に「解決される値」が
    セットされていることを想定する


Returns
-------
attr_val : str
    取り出した "gml:id" 属性値
    該当する属性が存在しなかった場合、"null" を返す
"""
def get_id_attr(xml_elem, namespace_map):
    attr_val = elem.get(f"{{{namespace_map['gml']}}}id")
    return attr_val if attr_val is not None else "null"


"""
要素直下にある指定した要素の持つ "xlink:href" 属性の値を取り出す

Parameters
----------
xml_elem : xml.etree.Element
    "xlink:href" 属性を持つ要素を持つ親 XML 要素

str_selector : str
    xml_elem 内にある "xlink:href" 属性をもつ XML 要素名

namespace_map : dict
    XML 名前空間を解決するための連想配列
    Key に「名前空間」、Value に「解決される値」が
    セットされていることを想定する

Returns
-------
attr_val : str
    取り出した "xlink:href" 属性値
    該当する属性が存在しなかった場合、"null" を返す
"""
def get_link(xml_elem, str_selector, namespace_map):
    elem = xml_elem.find(str_selector, namespaces=namespace_map)
    if elem is not None:
        if elem.get(f"{{{namespace_map['xlink']}}}href") is not None:
            return elem.get(f"{{{namespace_map['xlink']}}}href").lstrip('#')
    return "null"


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



# --------------------#
# --- Main Routine ---#
# --------------------#
if __name__ == '__main__':
    argc, argv = len(sys.argv), sys.argv

    if argc != 2:
        sys.stderr.write(f"usage: {argv[0]} [xml-file-path] \n")
        sys.exit(-1)

    path_xml_file = argv[1]

    sys.stderr.write("Generate Namespace Map ... \n")

    # XML 名前空間マップの作成
    namespace_map = dict()
    with open(path_xml_file) as f:
        for line in f.readlines():
            m = re.match(r'\s*xmlns:(.*)="(.*)".*', line)
            if m is not None:
                namespace_map.update({m.group(1): m.group(2)})

    # XML のパース
    sys.stderr.write("XML Parsing ... \n")
    tree = ET.parse(path_xml_file)
    root_elem = tree.getroot()


    # N02-17.xml は「ある駅における駅名, 路線情報, 座標リスト」がまとまっている構造ではなく、
    # 「駅名」「路線情報」「座標リスト」というくくりでまとまっているため、扱いにくい
    # そのため、gml:id 属性値をキーとする Key-Value マップ (連想配列, dict型)で
    # 駅名、路線情報、座標リストを管理して、最後に、ある駅単位で情報をまとめるような処理にする

    # -- gml:id 属性値について --
    # gml:id 属性値は同一の駅についての値でも「駅名」「路線情報」「座標リスト」で異なる値を
    # 保持しているという、やっかいな仕様になっている
    # I.    座標リスト (gml:Curve 要素) が持つ gml:id 値は 「cv_XXXX」という形式の値となっており、
    #       路線情報 (ksj:RailroadSection)、駅情報 (ksj:Stations) が持つ ksj:location 要素内の
    #       xlink:href 属性値に対応している
    # II.    gml:id 値は 「eb02_XXXXXX」という形式の値と
    #       なっており、 が持つ ksj:railroadSection 要素内の xlink:href
    #       属性値に対応している
    # III.  駅情報 (ksj:Stations) が持つ gml:id 値は 「eb03_XXXXXX」という形式の値となっており、
    #       路線情報 (ksj:RailroadSection) が持つ ksj:station 要素内の xlink:href 属性値に対応
    #       している


    # 1. Curve (座標リスト) のパース
    # パース結果は dict 型の変数 curves に格納される
    # curves は Key に 「gml:id 属性値」、 Value を「その駅の座標リスト」とする
    sys.stderr.write("Parsing -- gml:Curve ... \n")
    xml_curves = root_elem.findall('gml:Curve', namespaces=namespace_map)
    curves = dict()
    for elem in xml_curves:
        # gml:id 属性値の取得
        attr_id = get_id_attr(elem, namespace_map)
        # 座標リスト (gml:posList 要素) の取得
        pos_list = elem.find(
            'gml:segments/gml:LineStringSegment/gml:posList', namespaces=namespace_map)
        # 得られる座標値を改行文字('\n')で分割し、各行を空白文字(' ')で区切り、(緯度, 経度) の list を生成
        pos_list = [x.split(' ')
                    for x in re.split(r'\n\s*', pos_list.text) if x]

        # curves を更新
        curves[attr_id] = pos_list
        

    # 2. Station (駅情報) のパース
    # パース結果は dict 型の変数 stations に格納される
    # stations は Key に 「gml:id 属性値」、 Value を「駅情報を格納する dict」とする
    # 駅情報は、ksj:Station 要素直下の各 XML 要素について、属性名: 値 を Key-Value ペアで
    # 保持するものとする
    # 以下に stations の凡例を示す
    # stations = {
    #   'gml:id 属性値: {
    #     'railwayType': '鉄道区分コード (普通鉄道, 軌道, モノレールなど)',
    #     'serviceProviderType': '事業者種別コード (JR, 新幹線, 私鉄など)',
    #     'railwayLineName': '路線名',
    #     'operationCompany': '運営会社名',
    #     'railwayLineName: '駅名',
    #     'location': '座標リストを参照するためのID',
    #     'railroadSection': '路線情報を参照するためのID',
    #   }
    # }
    sys.stderr.write("Parsing -- ksj:Station ... \n")
    xml_stations = root_elem.findall('ksj:Station', namespaces=namespace_map)
    stations = dict()
    for elem in xml_stations:
        # gml:id 属性値の取得
        attr_id = get_id_attr(elem, namespace_map)

        # 'ksj:location' 要素または 'ksj:railroadSection' 要素を持たないものは、スキップ
        if get_link(elem, 'ksj:location', namespace_map) != "null" and get_link(
                elem, 'ksj:railroadSection', namespace_map) != "null":

            # 駅情報を格納する dict を初期化
            stations[attr_id] = dict()
            # ksj:Station 内の各要素を順にたどる
            for child in elem:
                # 要素が値を持つ場合
                if child.text is not None:
                    # 要素が値を持つ場合
                    key = re.sub(f"{{{namespace_map['ksj']}}}", "", child.tag)
                    val = child.text
                    # 駅情報の dict を更新する
                    stations[attr_id][key] = val

            # 'location' を設定
            location_id = get_link(elem, 'ksj:location', namespace_map)
            stations[attr_id]['location'] = np.array(curves[location_id], dtype=np.float64).mean(axis=0).tolist()



    # 3. 同一駅舎内の駅を統合する処理 (これ大事)

    # 駅統合の緯度, 経度の閾値 (〜450 m )
    merge_distance_threshold = 0.00545

    # (1) 駅名称に重複があるものをまとめる
    #     collection.Counter はリスト要素の重複数をカウントしてくれる
    #     これを利用して、駅名 ( 'stationName' ) の重複カウントを行う
    count = Counter([v['stationName'] for v in stations.values()])
    duplicated_names = [k for k, v in count.items() if v != 1]

    # 統合結果は dict 型の変数 merged_stations に格納する
    # Key を 「重複している駅名称」、Value を
    # 「駅名称(駅名, 路線名, 運営会社) とクラスタ番号のリスト」
    # とする
    merged_stations = dict()

    sys.stderr.write("Merge stations ... \n")

    for name in duplicated_names:
        # name に一致する駅情報を list 型の変数 dup_stations に取り出す
        dup_stations = [v for v in stations.values() if v['stationName'] == name]
        # 座標リストを numpy.ndarray に変換する
        datas = np.array( [ v['location'] for v in dup_stations ], dtype=np.float64)
        # (2) scikit-learn の MeanShift でクラスタリング
        #     引数 bandwidth は統合する際の閾値
        labels = MeanShift(
            bandwidth=merge_distance_threshold).fit_predict(datas)

        # 統合結果で merged_stations を更新
        # merged_stations[name] = [
        #     tuple(['{0}駅 ({2} {1})'.format(*a), str(b)]) for a, b in zip(dup_stations.keys(), labels)]

        merged_stations[name] = dict()
        for label_num, station in zip(labels, dup_stations):
            if label_num not in merged_stations[name]:
                merged_stations[name][label_num] = dict()
                merged_stations[name][label_num]['stations'] = list()
            merged_stations[name][label_num]['stations'].append(station)

        for v in merged_stations[name].values():
            v['center_latlng'] = np.array(
                [vv['location'] for vv in v['stations']], dtype=np.float64).mean(axis=0).tolist()

        # merged_stations[name] = {
        #     a: b for a, b in zip(labels, samples.items())
        # }

    #  4. 重複のない駅舎を統合する処理 (これ忘れてた)

    sys.stderr.write("Get single stations ... \n")

    # 複数駅舎を持たない駅名のリストを生成
    non_duplicated_names = [k for k, v in count.items() if v == 1]
    # merged_stations と同様のスキームになるように、 single_stations を生成
    single_stations = {
        k: {
            0: {
                'stations': [v for v in stations.values() if v['stationName'] == k],
                'center_latlng': [v['location'] for v in stations.values() if v['stationName'] == k][0]
            }
        } for k in non_duplicated_names
    }

    # single_stations を merged_stations に統合
    merged_stations.update(single_stations)

    # 5. SQL の生成
    # '1_schema.sql' に定義したスキーマに合わせた INSERT INTO 文を生成する
    # buildings, stations, railways 間は id によるリレーションがあるため、
    # id 値を設定しながら SQL を生成する

    sys.stderr.write ("Genrate SQL ... \n")


    # buildings, stations, railways の SQL 文を格納する
    sql_buildings, sql_stations, sql_railways = list(), list(), dict()

    # リレーション用の ID 値の初期化
    station_id, building_id, railway_id = 1, 1, 1
    

    for building_name, buildings in merged_stations.items():
        for building in buildings.values():

            # sql_buildings に現在の駅舎情報を追加
            sql_buildings.append(
                tuple([
                    building_id,
                    building_name,
                    SQL_GEOMFROMTEXT(building['center_latlng'])
                ])
            )

            # buildings が持つ駅を station としてループ
            for station in building['stations']:
                
                # railways は (路線名, 運営会社) のペアでユニーク
                # であるため、これをキーとして sql_railways に
                # 路線情報を格納する
                key_sql_railways = tuple([ 
                    station['railwayLineName'],
                    station['operationCompany'] 
                ])

                # 現在見ている駅の路線情報が sql_railways に含まれていない場合 
                if key_sql_railways not in sql_railways:
                    
                    # sql_railways に現在の路線情報を追加
                    sql_railways[key_sql_railways] = tuple([
                        railway_id,
                        station['railwayLineName'],
                        int(station['railwayType']),
                        station['operationCompany'],
                        int(station['serviceProviderType']),
                    ])
                    
                    # 今格納した路線情報のID値 (railway_id) を
                    # station_railway_id にセットする
                    # station_railway_id は sql_stations への
                    # 格納時に利用される
                    station_railway_id = railway_id

                    # railway_id をインクリメント
                    railway_id += 1

                # すでに路線情報が格納されている場合
                else:
                    
                    # 格納されている路線情報の ID 値を取り出す
                    station_railway_id = sql_railways[key_sql_railways][0]

                # sql_stations に現在の駅情報を追加
                sql_stations.append(
                    tuple([
                        station_id,
                        station['stationName'],
                        building_id,
                        station_railway_id
                    ])
                )

                # station_id をインクリメント
                station_id += 1

            # building_id をインクリメント
            building_id += 1

    sys.stderr.write('Write SQL as "seeds.sql" ... \n')


    # 生成した SQL クエリを 'seeds.sql' ファイルに出力
    with open('seeds.sql', 'wt') as f:

        for query in sql_buildings:
            f.write(
                SQL_INSERT_INTO(
                    'buildings',
                    ['id', 'name', 'latlong'],
                    query,
                    [0, 2]
                )
            )

        for query in sql_railways.values():
            f.write(
                SQL_INSERT_INTO(
                    'railways',
                    ['id', 'name', 'type', 'company_name', 'service_provider_type'],
                    query,
                    [0, 2, 4]
                )
            )

        for query in sql_stations:
            f.write(
                SQL_INSERT_INTO(
                    'stations',
                    ['id', 'name', 'building_id', 'railway_id'],
                    query,
                    [0, 2, 3]
                )
            )
