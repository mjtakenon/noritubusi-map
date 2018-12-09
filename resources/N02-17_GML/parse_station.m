clear;
close all;

DOMnode = xmlread('Station_transformed.xml');

Stations = DOMnode.getElementsByTagName('ksj:Station');

newStations = struct( ...
             'locationID',            [],...
             'railwayType',         [],...
             'serviceProviderType', [],...
             'railwayLineName',     [],...
             'operationCompany',    [],...
             'stationName',         [],...
             'railroadSection',     []);

for n = 0:Stations.getLength-1
    Station = Stations.item(n);
    
    location = Station.getElementsByTagName('ksj:location').item(0).getFirstChild.getData;
    railwayType = Station.getElementsByTagName('ksj:stationName').item(0).getFirstChild.getData;
    serviceProviderType = Station.getElementsByTagName('ksj:serviceProviderType').item(0).getFirstChild.getData;
    railwayLineName = Station.getElementsByTagName('ksj:railwayLineName').item(0).getFirstChild.getData;
    operationCompany = Station.getElementsByTagName('ksj:operationCompany').item(0).getFirstChild.getData;
    stationName = Station.getElementsByTagName('ksj:stationName').item(0).getFirstChild.getData;
    if Station.getElementsByTagName('ksj:railroadSection').getLength == 1
        railroadSection = Station.getElementsByTagName('ksj:railroadSection').item(0).getFirstChild.getData;
        newStations(n+1).railroadSection = railroadSection;
    end
    newStations(n+1).locationID = location;
    newStations(n+1).railwayType = railwayType;
    newStations(n+1).serviceProviderType = serviceProviderType;
    newStations(n+1).railwayLineName = railwayLineName;
    newStations(n+1).operationCompany = operationCompany;
    newStations(n+1).stationName = stationName;
    
    
%     thisList = thisListItem.getElementsByTagName('ksj:stationName');
%     thisElement = thisList.item(0);
%     thisElement
end

