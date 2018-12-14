var m_mono = new L.tileLayer('https://tile.mierune.co.jp/mierune_mono/{z}/{x}/{y}.png', {
    attribution: "Maptiles by <a href='http://mierune.co.jp/' target='_blank'>MIERUNE</a>, under CC BY. Data by <a href='http://osm.org/copyright' target='_blank'>OpenStreetMap</a> contributors, under ODbL."
});

var m_color = new L.tileLayer('https://tile.mierune.co.jp/mierune/{z}/{x}/{y}.png', {
    attribution: "Maptiles by <a href='http://mierune.co.jp/' target='_blank'>MIERUNE</a>, under CC BY. Data by <a href='http://osm.org/copyright' target='_blank'>OpenStreetMap</a> contributors, under ODbL."
});

var t_pale = new L.tileLayer('http://cyberjapandata.gsi.go.jp/xyz/pale/{z}/{x}/{y}.png', {
    attribution: "<a href='http://www.gsi.go.jp/kikakuchousei/kikakuchousei40182.html' target='_blank'>国土地理院</a>"
});

var t_ort = new L.tileLayer('http://cyberjapandata.gsi.go.jp/xyz/ort/{z}/{x}/{y}.jpg', {
    attribution: "<a href='http://www.gsi.go.jp/kikakuchousei/kikakuchousei40182.html' target='_blank'>国土地理院</a>"
});

var o_std = new L.tileLayer('http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&amp;copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
});

var o_rail = new L.TileLayer('http://{s}.tiles.openrailwaymap.org/standard/{z}/{x}/{y}.png',
{
	attribution: '<a href="https://www.openstreetmap.org/copyright">© OpenStreetMap contributors</a>, Style: <a href="http://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA 2.0</a> <a href="http://www.openrailwaymap.org/">OpenRailwayMap</a> and OpenStreetMap'
});

var map = L.map('map', {
    // default
    // center: [35.681,139.763],
    // 適当に日本が入る値
    center: [35.7,137.0],
    zoom: 7,
    zoomControl: true,
    layers: [o_std,o_rail]
});

var Map_BaseLayer = {
    "MIERUNE地図 color": m_color,
    "MIERUNE地図 mono": m_mono,
    "地理院地図 淡色": t_pale,
    "地理院地図 オルソ": t_ort,
    "OpenStreetMap 標準": o_std,
};

var Map_AddLayer = {
    "OpenRailwayMap": o_rail
}

function addStationCircles()
{
    var csv = $.csv.toArrays(data)
}

// function readCsv(data) {
//     var target = '#csv-body';
//     var csv = $.csv.toArrays(data);
//     var insert = '';
//     $(csv).each(function() {
//         if (this.length > 0) {
//             insert += '<tr>';
//             $(this).each(function() {
//                 insert += '<td>' + this + '</td>';
//             });
//             insert += '</tr>';
//         }
//     });
//     $(target).append(insert);
// }
// var csvfile = 'res/stations.csv';
// $(function(){
//     $.get(csvfile, readCsv, 'text');
// });

// example
L.circle([35.66413037753069,139.75278854370114], {radius: 100}).addTo(map);

L.control.scale({
    imperial: false,
    maxWidth: 300
}).addTo(map);

L.control.layers(
    Map_BaseLayer,
    Map_AddLayer, 
    {
        collapsed: false
    }
).addTo(map)