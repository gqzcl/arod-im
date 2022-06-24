(function(win) {
    const rawHeaderLen = 16;
    const packetOffset = 0;
    const headerOffset = 4;
    const verOffset = 6;
    const opOffset = 8;
    const seqOffset = 12;

    var Client = function(options) {
        var MAX_CONNECT_TIMES = 10;
        var DELAY = 15000;
        this.options = options || {};
        this.createConnect(MAX_CONNECT_TIMES, DELAY);
    }

    var appendMsg = function(text) {
        var span = document.createElement("SPAN");
        var text = document.createTextNode(text);
        span.appendChild(text);
        document.getElementById("box").appendChild(span);
    }

    Client.prototype.createConnect = function(max, delay) {
        var self = this;
        if (max === 0) {
            return;
        }
        connect();

        var textDecoder = new TextDecoder();
        var textEncoder = new TextEncoder();
        var heartbeatInterval;
        function connect() {
            var ws = new WebSocket('ws://127.0.0.1:7700/');
            //var ws = new WebSocket('ws://127.0.0.1:3102/sub');
            ws.binaryType = 'arraybuffer';
            ws.onopen = function() {
                auth();
            }

            ws.onmessage = function(evt) {
                var data = evt.data;
                var dataView = new DataView(data, 0);
                const obj=JSON.parse(data)
                appendMsg(obj);
                console.log(obj);

                switch(op) {
                    case 8:
                        // auth reply ok
                        document.getElementById("status").innerHTML = "<color style='color:green'>ok<color>";
                        appendMsg("receive: auth reply");
                        // send a heartbeat to server
                        heartbeat();
                        heartbeatInterval = setInterval(heartbeat, 30 * 1000);
                        break;
                    case 3:
                        // receive a heartbeat from server
                        console.log("receive: heartbeat");
                        appendMsg("receive: heartbeat reply");
                        break;
                    case 9:
                        // batch message
                        for (var offset=rawHeaderLen; offset<data.byteLength; offset+=packetLen) {
                            // parse
                            var packetLen = dataView.getInt32(offset);
                            var headerLen = dataView.getInt16(offset+headerOffset);
                            var ver = dataView.getInt16(offset+verOffset);
                            var op = dataView.getInt32(offset+opOffset);
                            var seq = dataView.getInt32(offset+seqOffset);
                            var msgBody = textDecoder.decode(data.slice(offset+headerLen, offset+packetLen));
                            // callback
                            messageReceived(ver, msgBody);
                            appendMsg("receive: ver=" + ver + " op=" + op + " seq=" + seq + " message=" + msgBody);
                        }
                        break;
                    default:
                        var msgBody = textDecoder.decode(data.slice(headerLen, packetLen));
                        messageReceived(ver, msgBody);
                        appendMsg("receive: ver=" + ver + " op=" + op + " seq=" + seq + " message=" + msgBody);
                        break
                }
            }

            ws.onclose = function() {
                if (heartbeatInterval) clearInterval(heartbeatInterval);
                setTimeout(reConnect, delay);

                document.getElementById("status").innerHTML =  "<color style='color:red'>failed<color>";
            }

            function heartbeat() {
                var headerBuf = new ArrayBuffer(rawHeaderLen);
                var headerView = new DataView(headerBuf, 0);
                headerView.setInt32(packetOffset, rawHeaderLen);
                headerView.setInt16(headerOffset, rawHeaderLen);
                headerView.setInt16(verOffset, 1);
                headerView.setInt32(opOffset, 2);
                headerView.setInt32(seqOffset, 1);
                ws.send(headerBuf);
                console.log("send: heartbeat");
                appendMsg("send: heartbeat");
            }

            function auth() {
                var token = '{"uid":"100001", "group_id":"group@1000"}'
                var textEncoder = new TextEncoder();
                var bodyBuf = textEncoder.encode(token);
                ws.send(bodyBuf)
                ws.onmessage = function(evt) {
                    console.log(evt.data);
                }

                appendMsg("send: auth token: " + token);
            }

            function messageReceived(ver, body) {
                var notify = self.options.notify;
                if(notify) notify(body);
                console.log("messageReceived:", "ver=" + ver, "body=" + body);
            }

            function mergeArrayBuffer(ab1, ab2) {
                var u81 = new Uint8Array(ab1),
                    u82 = new Uint8Array(ab2),
                    res = new Uint8Array(ab1.byteLength + ab2.byteLength);
                res.set(u81, 0);
                res.set(u82, ab1.byteLength);
                return res.buffer;
            }

            function char2ab(str) {
                var buf = new ArrayBuffer(str.length);
                var bufView = new Uint8Array(buf);
                for (var i=0; i<str.length; i++) {
                    bufView[i] = str[i];
                }
                return buf;
            }

        }

        function reConnect() {
            self.createConnect(--max, delay * 2);
        }
    }

    win['MyClient'] = Client;
})(window);

// const {io}=require("socket.io-client");

// const socket = io("ws://127.0.0.1:7700/");
// const form = document.querySelector('form');
// const input = document.querySelector('input');
// const ul = document.querySelector('ul');

// form.addEventListener('submit',e=>{
//     e.preventDefault();
//     if (input.value){
//         socket.emit('chat message',input.value);
//         input.value='';
//     }
// });
// socket.on('chat message',msg=>{
//     const li = document.createElement('li');
//     li.textContent = msg;
//     ul.appendChild(li);
// });

// /**
//  * Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
//  * Use of this source code is governed by a MIT style
//  */

// (function(win) {
//     var Client = function(options) {
//         var DELAY = 15000;
//         this.options = options || {};
//         this.createConnect(DELAY);
//     }
//     Client.prototype.createConnect = function(delay) {
//         connect();
//         function connect() {
//             var ws = new WebSocket('ws://127.0.0.1:7700/');
//             ws.binaryType = 'arraybuffer';
//             ws.onopen = function() {
//                 auth();
//             }

//             ws.onmessage = function(evt) {
//                 console.log("receiveHeader: packetLen=");

//                 switch(op) {
//                     case 8:
//                         // auth reply ok
//                         document.getElementById("status").innerHTML = "<color style='color:green'>ok<color>";
//                         //appendMsg("receive: auth reply");
//                         // send a heartbeat to server
//                         //heartbeat();
//                         //heartbeatInterval = setInterval(heartbeat, 30 * 1000);
//                         break;
//                     case 3:
//                         // receive a heartbeat from server
//                         console.log("receive: heartbeat");
//                         //appendMsg("receive: heartbeat reply");
//                         break;
//                     default:
//                         var msgBody = textDecoder.decode(data.slice(headerLen, packetLen));
//                         //messageReceived(ver, msgBody);
//                         //appendMsg("receive: ver=" + ver + " op=" + op + " seq=" + seq + " message=" + msgBody);
//                         break
//                 }
//             }
//         }
//         function auth(){
//             var token = '{"uid":"123", "group_id":"group@1000"}'
//             var textEncoder = new TextEncoder();
//             var bodyBuf = textEncoder.encode(token);
//             ws.send(bodyBuf)
//             ws.onmessage = function(evt) {
//                 console.log(evt.data);
//             }
//         }
//     }
//     win['MyClient'] = Client;
// })(window);