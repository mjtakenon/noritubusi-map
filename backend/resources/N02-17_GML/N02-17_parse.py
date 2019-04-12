#!/usr/bin/env python3

import xml.etree.ElementTree as ET
import os
import sys
import re
import json
import numpy as np
import numpy.matlib
import itertools

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


def SQL_INSERT_INTO(TBL_NAME, KEYS, VALUES, IDX_NOT_QUOTED=[]):
    str_keys = ', '.join([f"`{x}`" for x in KEYS])
    try:
        str_values = ', '.join(
            [f"'{x}'" if i not in IDX_NOT_QUOTED else x for (i, x) in enumerate(VALUES)])
    except TypeError as e:
        import pdb
        pdb.set_trace()

    return f"INSERT INTO `{TBL_NAME}` ({str_keys}) VALUES ({str_values});\n"


if __name__ == '__main__':
    argc, argv = len(sys.argv), sys.argv

    if argc != 2:
        sys.stderr.write(f"usage: {argv[0]} [xml-file-path] \n")
        sys.exit(-1)

    # --------------------#
    # --- Main Routine ---#
    # --------------------#
    path_xml_file = argv[1]

    sys.stderr.write("Pre-processing ... \n")

    # Get Namespace Map
    namespace_map = dict()
    with open(path_xml_file) as f:
        for line in f.readlines():
            m = re.match(r'\s*xmlns:(.*)="(.*)".*', line)
            if m is not None:
                namespace_map.update({m.group(1): m.group(2)})

    # XML Parse
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
        # curves.update(
        #     {
        #         attr_id: {
        #             'posList': pos_list
        #         }
        #     }
        # )

    # 2. RailroadSection (路線情報) のパース
    # パース結果は dict 型の変数 railroad_sections に格納される
    # railroad_sections は Key に 「gml:id 属性値」、
    # Value を「路線情報を格納する dict」とする
    # 路線情報は、ksj:RailloadSection 要素直下の各 XML 要素について、
    # 属性名: 値 を Key-Value ペアで保持するものとする
    # 以下に railload_sections の凡例を示す
    # railload_sections = {
    #   'gml:id 属性値': {
    #     'railwayType': '鉄道区分コード (普通鉄道, 軌道, モノレールなど)',
    #     'serviceProviderType': '事業者種別コード (JR, 新幹線, 私鉄など)',
    #     'railwayLineName': '路線名',
    #     'operationCompany': '運営会社名',
    #     'location': '座標リスト を参照するためのID'
    #   }
    # }
    xml_railroad_sections = root_elem.findall(
        'ksj:RailroadSection', namespaces=namespace_map)
    railroad_sections = dict()
    for elem in xml_railroad_sections:
        # gml:id 属性値の取得
        attr_id = get_id_attr(elem, namespace_map)
        # 路線情報を格納する dict を初期化
        railroad_sections[attr_id] = dict()
        # ksj:RailroadSection 内の各要素を順にたどる
        for child in elem:
            # 要素が値を持つ場合
            if child.text is not None:
                # 要素名を Key, 要素値を Value として
                key = re.sub(f"{{{namespace_map['ksj']}}}", "", child.tag)
                val = child.text
                # 路線情報の dict を更新する
                railroad_sections[attr_id][key] = val

        # 'location' を設定
        railroad_sections[attr_id]['location'] = get_link(
            elem, 'ksj:location', namespace_map)

    # 3. Station (駅情報) のパース
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
    xml_stations = root_elem.findall('ksj:Station', namespaces=namespace_map)
    stations = dict()
    for elem in xml_stations:
        # gml:id 属性値の取得
        attr_id = get_id_attr(elem, namespace_map)
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

        # 'location', 'railroadSection' を設定
        stations[attr_id]['location'] = get_link(
            elem, 'ksj:location', namespace_map)
        stations[attr_id]['railroadSection'] = get_link(
            elem, 'ksj:railroadSection', namespace_map)

    # 4. 各データセット間の統合処理
    # 「座標リスト」「路線情報」「駅情報」間の情報をIDをもとに紐づけて、1つのデータセットに
    # 統合する処理を行う

    # 駅単位のデータにしたいため、駅情報 (stations) に対し for-each を行い、統合する
    for station in stations.values():
        # (1) 座標リストの取得
        # 駅情報内の 'location' 値が null 出ない場合
        if station['location'] != "null":
            # 'location' 値を curves から対応する座標リストで更新する
            station['location'] = curves[station['location']]

            # 重心計算
            # 座標リストは駅の端と端の2点間の座標を持っている(？)ため、そこから重心を
            # 計算し、駅の座標値とする
            # 行列演算を行いたいため、 Python の list から numpy.ndarray に変換し、
            # 'location' 値を重心値で更新する
            arr = np.array(station['location'], dtype=np.float64)
            station['location'] = arr.mean(axis=0).tolist()

    import pdb
    pdb.set_trace()
