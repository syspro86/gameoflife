const go = new Go();

// WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
// 	go.run(result.instance);
// });

fetch("main.wasm").then((response) => {
    return response.arrayBuffer();
}).then((buffer) => {
    return WebAssembly.instantiate(buffer, go.importObject);
}).then((result) => {
    go.run(result.instance);

    var canvas = document.getElementById("canvas");
    initWorld(canvas.clientWidth, canvas.clientHeight);

    function render() {
        renderFrame(canvas);
        setTimeout(render, 100);
    }
    
    setTimeout(render, 100);
});
