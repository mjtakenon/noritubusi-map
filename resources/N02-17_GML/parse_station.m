function parsedStations = parse_station(filename)

DOMnode = xmlread(filename);

Stations = DOMnode.getElementsByTagName('ksj:Station');

parsedStations = struct( ...
             'stationID',           [],...
             'locationID',          [],...
             'railwayType',         [],...
             'serviceProviderType', [],...
             'railwayLineName',     [],...
             'operationCompany',    [],...
             'stationName',         [],...
             'railroadSectionID',   [],...
             'weightPos',           [],...
             'posList',             []);

for n = 0:Stations.getLength-1
    Station = Stations.item(n);
    
    stationID =  Station.getElementsByTagName('gml:id').item(0).getFirstChild.getData;
    location = Station.getElementsByTagName('ksj:location').item(0).getFirstChild.getData;
    railwayType = Station.getElementsByTagName('ksj:railwayType').item(0).getFirstChild.getData;
    serviceProviderType = Station.getElementsByTagName('ksj:serviceProviderType').item(0).getFirstChild.getData;
    railwayLineName = Station.getElementsByTagName('ksj:railwayLineName').item(0).getFirstChild.getData;
    operationCompany = Station.getElementsByTagName('ksj:operationCompany').item(0).getFirstChild.getData;
    stationName = Station.getElementsByTagName('ksj:stationName').item(0).getFirstChild.getData;
    if Station.getElementsByTagName('ksj:railroadSection').getLength == 1
        railroadSection = Station.getElementsByTagName('ksj:railroadSection').item(0).getFirstChild.getData;
        parsedStations(n+1).railroadSectionID = strrep(char(railroadSection),'#','');
    end
    
    parsedStations(n+1).stationID = char(stationID);
    parsedStations(n+1).locationID = strrep(char(location),'#','');
    parsedStations(n+1).railwayType = int32(str2double(railwayType));
    parsedStations(n+1).serviceProviderType = int32(str2double(serviceProviderType));
    parsedStations(n+1).railwayLineName = char(railwayLineName);
    parsedStations(n+1).operationCompany = char(operationCompany);
    parsedStations(n+1).stationName = char(stationName);
end

end