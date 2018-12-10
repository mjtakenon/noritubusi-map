function parsedCurves = parse_curve(filename)

DOMnode = xmlread(filename);

Curves = DOMnode.getElementsByTagName('gml:Curve');

parsedCurves = struct( ...
             'curveID', [],...
             'posList', []);

for n = 0:Curves.getLength-1
    Curve = Curves.item(n);

    curveID = Curve.getElementsByTagName('gml:CurveId').item(0).getFirstChild.getData;
    parsedCurves(n+1).curveID = char(curveID);

    posList = Curve.getElementsByTagName('gml:segments').item(0).getElementsByTagName('gml:LineStringSegment').item(0).getElementsByTagName('gml:posList').item(0).getFirstChild.getData;
    posList = strtrim(string(split(posList,'\n')));
    idx = find(cellfun('isempty',posList));
    posList(idx) = [];
    parsedCurves(n+1).posList = double(split(posList));
end

end