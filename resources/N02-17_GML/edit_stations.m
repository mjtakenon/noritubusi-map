clear;
close all;

% 各xmlを読み込み
curves = parse_curve('Curve_transformed.xml');
stations = parse_station('Station_transformed.xml');
sections = parse_railroadSection('RailroadSection_transformed.xml');

% 駅の緯度経度をcurveから探してきて重心を算出(ゴミクソ遅い)
for n=1:numel(stations)
    for nn=1:numel(curves)
        if strcmp(stations(n).locationID,curves(nn).curveID)
            stations(n).posList = curves(nn).posList;
            stations(n).weightPos = sum(curves(nn).posList(:,:))./size(curves(nn).posList,1);
            break;
        end
    end
end
clear n nn
save('data.mat');

% 書き出し用データ成形
for n=1:numel(stations)
    stations(n).weightPosY = stations(n).weightPos(1);
    stations(n).weightPosX = stations(n).weightPos(2);
end

c = reshape(struct2cell(stations),[numel(fieldnames(stations)),numel(stations)]);
c = transpose(c);
c(:,9) = c(:,11);
c(:,10) = c(:,12);
tablename = [{'stationID'},{'locationID'},{'railwayType'},{'serviceProviderType'},{'railwayLineName'},{'operationCompany'},{'stationName'},{'railroadSectionID'},{'longitude'},{'latitude'}];
c = vertcat(tablename, c(:,1:10));
xlswrite('stations.xlsx',c);
clear n c