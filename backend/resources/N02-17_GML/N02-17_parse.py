#!/usr/bin/env python3

import xml.etree.ElementTree as ET
import os, sys, re, json
import numpy as np

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
  str_values = ', '.join([ f"'{x}'" if i not in IDX_NOT_QUOTED else x for (i, x) in enumerate(VALUES) ])
  
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
        railways.append( (s['operationCompany'], s['railwayLineName']) )
        buildings.append( [s['stationName'], np.array([float(s['location']['center'][0]),float(s['location']['center'][1])])] )

      buildings = np.array(buildings)
      
      # 駅: 同名かつユークリッド距離が近い駅(0.5km以内)を統合して出力
      mergeDistanceThreshold = 0.00545

      # 全ての駅でループ(実装)
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

      buildings = buildings.tolist()

      for b in buildings:
        # print(SQL_INSERT_INTO('buildings', ['name', 'latlong'] , [b[0], f"GeomFromText('POINT({' '.join(map(str,b[1]))})')"] ))
        fp.write(
          SQL_INSERT_INTO('buildings', ['name', 'latlong'] , [b[0], f"GeomFromText('POINT({' '.join(map(str,b[1]))})')"] )
        )

      # 路線名の出力
      railways = list(set(railways))
      for r in railways:
        # print(SQL_INSERT_INTO('railways', ['operation_company', 'railway_line_name'], r))
        fp.write(
          SQL_INSERT_INTO('railways', ['operation_company', 'railway_line_name'], r)
        )

      # TODO: 駅の出力 

      # for s in stations.values():
      #   kv = {
      #     'station_name': s['stationName'], 
      #     'center_latlong': f"GeomFromText('POINT({' '.join(s['location']['center'])})')", 
      #     'operation_company': s['operationCompany'], 
      #     'service_provider_type': s['serviceProviderType'], 
      #     'railway_line_name': s['railwayLineName'], 
      #     'railway_type': s['railwayType']
      #   }
      # # for r in railways.values():
      # #   print(SQL_INSERT_INTO('railways', r.keys(), r.values()))
      #   fp.write(
      #     SQL_INSERT_INTO('stations', kv.keys(), kv.values(), IDX_NOT_QUOTED=[1, 3, 5])
      #   )
