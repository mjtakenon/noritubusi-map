function parsed_railroadSections = parse_railroadSection(filename)
DOMnode = xmlread(filename);

RailRoadSections = DOMnode.getElementsByTagName('ksj:RailroadSection');

parsed_railroadSections = struct( ...
             'railRoadSectionID',           [],...
             'locationID',          [],...
             'railwayType',         [],...
             'serviceProviderType', [],...
             'railwayLineName',     [],...
             'operationCompany',    [],...
             'stationID',         []);

for n = 0:RailRoadSections.getLength-1
    RailRoadSection = RailRoadSections.item(n);
    
    railRoadSectionID = RailRoadSection.getElementsByTagName('ksj:RailroadSectionID').item(0).getFirstChild.getData;
    locationID = RailRoadSection.getElementsByTagName('ksj:location').item(0).getFirstChild.getData;
    railwayType = RailRoadSection.getElementsByTagName('ksj:railwayType').item(0).getFirstChild.getData;
    serviceProviderType = RailRoadSection.getElementsByTagName('ksj:serviceProviderType').item(0).getFirstChild.getData;
    railwayLineName = RailRoadSection.getElementsByTagName('ksj:railwayLineName').item(0).getFirstChild.getData;
    operationCompany = RailRoadSection.getElementsByTagName('ksj:operationCompany').item(0).getFirstChild.getData;
    if RailRoadSection.getElementsByTagName('ksj:station').getLength == 1
        stationID = RailRoadSection.getElementsByTagName('ksj:station').item(0).getFirstChild.getData;    
        parsed_railroadSections(n+1).stationID = strrep(string(stationID),'#','');
    end
    
    parsed_railroadSections(n+1).railRoadSectionID = strrep(string(railRoadSectionID),'#','');
    parsed_railroadSections(n+1).locationID = string(locationID);
    parsed_railroadSections(n+1).railwayType = int32(str2double(railwayType));
    parsed_railroadSections(n+1).serviceProviderType = int32(str2double(serviceProviderType));
    parsed_railroadSections(n+1).railwayLineName = string(railwayLineName);
    parsed_railroadSections(n+1).operationCompany = string(operationCompany);
end

end