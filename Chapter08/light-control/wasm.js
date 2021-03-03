'use strict';

const WASM_URL = 'wasm.wasm';
var wasm;

function OnPopState(event) {
    alert("location: " + document.location + ", state: " + JSON.stringify(event.state))
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

    window.onpopstate = OnPopState
}

init();