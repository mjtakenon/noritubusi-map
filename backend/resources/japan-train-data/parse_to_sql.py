#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from attrdict import AttrDict
import json
import sys


def proc(path_to_json):

    raw_data = [AttrDict(x) for x in json.load(open(path_to_json))]

    buildings = dict()
    railways = dict()
    stations = dict()

    railways_id, buildings_id, stations_id = 1, 1, 1

    # prefecture: 県
    for prefecture in raw_data:
        sys.stderr.write(f" Prefecture[{prefecture.name.ja}] ... \r")
        sys.stderr.flush()
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
                ref_railways_id = railways[railway.name.ja]

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
                    ref_buildings_id = buildings[latlong_station]

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
                            'railways_id': ref_railways_id,
                            'building_id': ref_buildings_id,
                            'num_in_railway': num_in_railway
                        }
                    }
                )

    return buildings, stations, railways


if __name__ == '__main__':
    argc, argv = len(sys.argv), sys.argv

    if argc != 2:
        sys.stderr.write(f"usage: {argv[0]} [path to raw-data.json] \n")
        sys.exit(-1)

    else:

        buildings, stations, railways = proc(argv[1])
