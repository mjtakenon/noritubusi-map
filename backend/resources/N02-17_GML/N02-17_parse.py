#!/usr/bin/env python3

import xml.etree.ElementTree as ET
import os, sys, re, json
import numpy as np
import numpy.matlib

def get_id(xml_elem):
  ret_val = elem.get(f"{{{ns['gml']}}}id")
  return ret_val if ret_val is not None else "null"

def get_link(xml_elem, str_selector):
  elem = xml_elem.find(str_selector, namespaces=ns)
  if elem is not None:
    if elem.get(f"{{{ns['xlink']}}}href") is not None:
      return elem.get(f"{{{ns['xlink']}}}href").lstrip('#')
  return "null"

def SQL_INSERT_INTO(TBL_NAME, KEYS, VALUES, IDX_NOT_QUOTED=[]):
  str_keys   = ', '.join([ f"`{x}`" for x in KEYS])
  try:
    str_values = ', '.join([ f"'{x}'" if i not in IDX_NOT_QUOTED else x for (i, x) in enumerate(VALUES) ])
  except TypeError as e:
    import pdb; pdb.set_trace()

  return f"INSERT INTO `{TBL_NAME}` ({str_keys}) VALUES ({str_values});\n"

if __name__ == '__main__':
  argc, argv = len(sys.argv), sys.argv

  if argc != 2:
    sys.stderr.write(f"usage: {argv[0]} [xml-file-path] \n")
    sys.exit(-1)
  else:

    sys.stderr.write("Pre-processing ... \n")

    # Get Namespace Map
    ns = dict()
    with open(argv[1]) as f:
      lines = f.readlines()
      for l in lines:
        m = re.match(r'\s*xmlns:(.*)="(.*)".*', l)
        if m is not None:
          ns.update( { m.group(1): m.group(2) } )

    # XML Parse
    tree = ET.parse(argv[1])
    root = tree.getroot()

    # gml:Curve
    xml_curves = root.findall('gml:Curve', namespaces=ns)
    curves = dict()
    for elem in xml_curves:
      attr_id = get_id(elem)
      pos_list = elem.find('gml:segments/gml:LineStringSegment/gml:posList', namespaces=ns)
      pos_list = [ x.split(' ') for x in re.split(r'\n\s*', pos_list.text) if x ]

      curves.update(
        {
          attr_id: {
            'posList': pos_list
          }
        }
      )

    # ksj:RailroadSection
    xml_railroad_sections = root.findall('ksj:RailroadSection', namespaces=ns)
    railroad_sections = dict()
    for elem in xml_railroad_sections:
      attr_id = get_id(elem)
      rs = dict()
      for child in elem:
        if child.text is not None:
          key = re.sub(f"{{{ns['ksj']}}}", "", child.tag)
          val = child.text
          rs.update({key: val})
      rs['location'] = get_link(elem, 'ksj:location')

      railroad_sections.update(
        { attr_id: rs }
      )

    # ksj:Station
    xml_stations = root.findall('ksj:Station', namespaces=ns)
    stations = dict()
    for elem in xml_stations:
      attr_id = get_id(elem)
      st = dict()
      for child in elem:
        if child.text is not None:
          key = re.sub(f"{{{ns['ksj']}}}", "", child.tag)
          val = child.text
          st.update({key: val})
      st['location'] = get_link(elem, 'ksj:location')
      st['railroadSection'] = get_link(elem, 'ksj:railroadSection')

      stations.update(
        { attr_id: st }
      )


    # Merge Datasets
    for station in stations.values():
      if station['location'] != "null":
        station.update({'location': curves[station['location']]})
      # 重心計算
      arr = np.array(station['location']['posList'], dtype=np.float64)
      station['location']['center'] = [ f"{p:.8f}" for p in arr.mean(axis=0) ]

    # Print as JSON
    # sys.stderr.write("Print as JSON ... \n")
    # with open((os.path.splitext(argv[1])[0] + ".json"), "wt") as fp:
    #   json.dump(stations, fp, ensure_ascii=False, indent=2, sort_keys=True, separators=(',', ': '))


    # Generate SQL Seeds
    sys.stderr.write("Print as SQL ... \n")

    # Print as SQL
    with open((os.path.splitext(argv[1])[0] + "_seeds.sql"), "wt") as fp:
      # 1. companies
      railways = list()
      buildings = list()
      for s in stations.values():
        railways.append( (s['railwayLineName'], s['railwayType'], s['operationCompany'], s['serviceProviderType']) )
        buildings.append( [s['stationName'], np.array([float(s['location']['center'][0]),float(s['location']['center'][1])])] )
        
      buildings = np.array(buildings)

      # 駅: 同名かつユークリッド距離が近い駅(0.5km以内)を統合して出力
      mergeDistanceThreshold = 0.00545

      # 全ての駅でループ
      for b in range(len(buildings)):
        if b >= len(buildings):
          break
        # 駅名で検索、リスト作成
        station = buildings[buildings[:,0] ==  buildings[b,0],:]
        # 統合されたら削除
        isMerged = [False] * len(station)
        for n in range(len(station)):
          for nn in range(len(station)):
            # もし統合されていたら被統合対象とせずスキップ
            if isMerged[n] or isMerged[nn] or n == nn:
              continue
            # 距離が一定以下だったら位置情報を追加し統合フラグをたてる 統合した駅の座標を見ていない為改善が必要
            distance = np.linalg.norm(station[n][1][0:1] - station[nn][1][0:1])
            if distance < mergeDistanceThreshold:
              station[n][1] = np.append(station[n][1],station[nn][1])
              isMerged[nn] = True
        for n in range(len(station)):
          if not isMerged[n]: # 被統合
            # 位置情報の再計算
            station[n][1] = [sum(station[n][1][0::2]) / len(station[n][1][0::2]), sum(station[n][1][1::2]) / len(station[n][1][1::2])]
        # 親リストに反映
        # 位置情報の更新
        buildings[np.where(buildings[:,0] == buildings[b,0])[0][np.logical_not(isMerged)],:][0][1] = station[np.logical_not(isMerged),1]
        # 重複駅の削除
        buildings = np.delete(buildings,np.where(buildings[:,0] == buildings[b,0])[0][isMerged],0)
      
      # 路線: 重複を削除
      # railways = np.array(list(set(railways)))
      railways = np.array(list(dict.fromkeys(railways)))

      # 駅: 対応する路線と建物インデックスの計算
      for s in stations.values():
        s['railway_id'] = np.where((railways[:,0] == s['railwayLineName']) & (railways[:,2] == s['operationCompany']))[0][0]
        # 駅名が完全一致している緯度経度リストを取得
        sameNameList = buildings[np.where(buildings[:,0] == s['stationName'])]
        # 上のリストの中から最も近い建物IDのインデックスを取得
        nearestBuildingIdx = np.linalg.norm(sameNameList[:,1].tolist() - np.matlib.repmat(np.array([float(s['location']['center'][0]),float(s['location']['center'][1])]),len(sameNameList),1),None,1).argmin()
        # print(sameNameList)
        # print(nearestBuildingIdx)
        # そのインデックスに対応する建物IDを取得
        s['building_id'] = np.where(buildings[:,0] == s['stationName'])[0][nearestBuildingIdx]
        # print(s['railway_id'])
        # print(s['building_id'])
        # import pdb; pdb.set_trace()
        # np.whereでの位置情報の検索が上手くいかなかったためこの形で実装したので後で修正したい



      # 建物の出力
      buildings = buildings.tolist()
      for n, b in enumerate(buildings):
        # print(SQL_INSERT_INTO('buildings', ['name', 'latlong'] , [b[0], f"GeomFromText('POINT({' '.join(map(str,b[1]))})')"] ))
        kv = {
          'id': str(n),
          'name': b[0],
          'latlong': f"GeomFromText('POINT({' '.join(map(str,b[1]))})')"
        }
        fp.write(
          SQL_INSERT_INTO('buildings', ['id', 'name', 'latlong'] , [str(n), b[0], f"GeomFromText('POINT({' '.join(map(str,b[1]))})')"],IDX_NOT_QUOTED=[0, 2])
          # SQL_INSERT_INTO('buildings', kv.keys(), kv.values(), IDX_NOT_QUOTED=[0, 2])
        )

      # 路線名の出力
      railways = list(railways)
      for n, r in enumerate(railways):
        # kv = {
        #   'id': n,
        #   'name': r[0],
        #   'type': r[1],
        #   'company_name': r[2],
        #   'service_provider_type': r[3],
        # }
        fp.write(
          SQL_INSERT_INTO('railways', ['id', 'name', 'type','company_name','service_provider_type'], [str(n), r[0], str(r[1]), r[2], str(r[3])], IDX_NOT_QUOTED=[0, 2, 4])
          # SQL_INSERT_INTO('railways', kv.keys(), kv.values(), IDX_NOT_QUOTED=[0, 2, 4])
        )

      # 駅の出力
      for n, s in enumerate(stations.values()):
        kv = {
          'id': str(n),
          'name': s['stationName'],
          'building_id': str(s['building_id']),
          'railway_id': str(s['railway_id']),
        }
        fp.write(
          # SQL_INSERT_INTO('stations', kv.keys(), kv.values(), IDX_NOT_QUOTED=[0, 2, 3])
          SQL_INSERT_INTO('stations', ['id', 'name','building_id','railway_id'], [str(n), s['stationName'],str(s['building_id']),str(s['railway_id'])] , IDX_NOT_QUOTED=[0, 2, 3])
        )

      # for s in stations.values():
      #   kv = {
      #     'station_name': s['stationName'],
      #     'center_latlong': f"GeomFromText('POINT({' '.join(s['location']['center'])})')",
      #     'operation_company': s['operationCompany'],
      #     'service_provider_type': s['serviceProviderType'],
      #     'railway_line_name': s['railwayLineName'],
      #     'railway_type': s['railwayType']
      #   }
      # for r in railways.values():
      #   print(SQL_INSERT_INTO('railways', r.keys(), r.values()))
      #   fp.write(
      #     SQL_INSERT_INTO('stations', kv.keys(), kv.values(), IDX_NOT_QUOTED=[1, 3, 5])
      #   )
