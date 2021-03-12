'use strict';

// Documentation at https://www.eclipse.org/paho/files/jsdoc/Paho.MQTT.Client.html
var mqtt;
const host = "127.0.0.1";
const port = 9001

function onConnect() {
    console.log("Successfully connected to mqtt broker");
}

function onConnectionLost(responseObject) {
    if (responseObject.errorCode !== 0) {
        console.log("onConnectionLost:" + responseObject.errorMessage);
    }
}

function onMessageArrived(message) {
    console.log("onMessageArrived:" + message.payloadString);
    handleMessage(message.payloadString)
}

function publish(topic, message) {
    mqtt.send(topic, message, 1, false)
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