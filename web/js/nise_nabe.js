$(function() {
    "use strict";
    var
        li = [
            {url: "http://d.hatena.ne.jp/nise_nabe"},
            {url: "http://twitter.com/nise_nabe"},
            {url: "http://github.com/nise-nabe"}
        ],
        template = "<li><a href=\"{{url}}\">{{url}}</li>",
        content = $("#content");
    li.forEach(function(elem) {
        $("#links").append($.mustache(template, elem));
    });
});
