'use strict';

const WASM_URL = 'wasm.wasm';
var wasm;

var mqtt;
var reconnectTimeout = 2000;
var host = "127.0.0.1";
var port = 9001
var payload = ""

function onConnect() {
    console.log("Successfully connected to mqtt broker");

    mqtt.subscribe("weather/data");
    mqtt.subscribe("weather/alert");
}

function onConnectionLost(responseObject) {
    if (responseObject.errorCode !== 0) {
        console.log("onConnectionLost:" + responseObject.errorMessage);
    }
}

function onMessageArrived(message) {
    console.log("onMessageArrived:" + message.payloadString);

    payload = message.payloadString
    if (payload.indexOf("possible storm incoming") !== -1) {
        alertHandler(payload)
    } else {
        sensorDataHandler(payload)
    }
}

function MQTTconnect() {
    console.log("mqtt client: connecting to " + host + ":" + port);

    var cname = "weather-consumer"
    mqtt = new Paho.MQTT.Client(host, port, cname);
    var options = {
        timeout: 3,
        onSuccess: onConnect,
    };

    mqtt.onConnectionLost = onConnectionLost;
    mqtt.onMessageArrived = onMessageArrived;

    mqtt.connect(options);
}

function init() {
    const go = new Go();
    if ('instantiateStreaming' in WebAssembly) {
        WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then(function(obj) {
            wasm = obj.instance;
            go.run(wasm);
        })
    } else {
        fetch(WASM_URL).then(resp =>
            resp.arrayBuffer()
        ).then(bytes =>
            WebAssembly.instantiate(bytes, go.importObject).then(function(obj) {
                wasm = obj.instance;
                go.run(wasm);
            })
        )
    }

    MQTTconnect()
}

init();