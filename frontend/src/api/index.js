// Production WebSocket endpoint
var socket = new WebSocket("wss://chatify-back-production.up.railway.app/ws");

let connect = cb => {
    console.log("Attempting Connection...");

    socket.onopen = () => {
        console.log("Successfully Connected.");
    };

    socket.onmessage = msg => {
        console.log(msg);
        cb(msg);
    };
    
    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
        socket.send("Client Closed!");
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };
};

let sendMsg = msg => {
    console.log("Sending message: ", msg);
    socket.send(msg);
}

export { connect, sendMsg };