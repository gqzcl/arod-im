/**
 * Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 */

(function(win) {
    function connect() {
        //var ws = new WebSocket('ws://sh.tony.wiki:3102/sub');
        var ws = new WebSocket('ws://127.0.0.1:7700/');
    }
    function sendMsg(){
        var token = '{"uid":"123", "group_id":"group@1000"}'
        var textEncoder = new TextEncoder();
        var bodyBuf = textEncoder.encode(token);
        ws.send(bodyBuf)
        ws.onmessage = function(evt) {
            console.log(evt.data);
        }
    }
    win['MyClient'] = Client;
})(window);