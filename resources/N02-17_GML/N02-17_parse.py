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


if __name__ == '__main__':
  argc, argv = len(sys.argv), sys.argv
  
  if argc != 2:
    sys.stderr.write(f"usage: {argv[0]} [xml-file-path] \n")
    sys.exit(-1)
  else:

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
    with open((os.path.splitext(argv[1])[0] + ".json"), "wt") as fp:
      json.dump(stations, fp, ensure_ascii=False, indent=2, sort_keys=True, separators=(',', ': '))

