'use strict';

var mqtt;
const host = "127.0.0.1";
const port = 9001

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

    var payload = message.payloadString
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