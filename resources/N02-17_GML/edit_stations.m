% clear;
% close all;
% 
% curves = parse_curve('Curve_transformed.xml');
% 
% stations = parse_station('Station_transformed.xml');
% sections = parse_railroadSection('RailroadSection_transformed.xml');

load('data.mat');

parfor n=1:numel(stations)
    for nn=1:numel(curves)
        if stations(n).locationID ~= curves(nn).curveID
            continue;
        else
            stations(n).posList = curves(nn).posList;
            stations(n).weightPoint = sum(curves(nn).posList(:,:))./size(curves(nn).posList,1);
            break;
        end
    end
end