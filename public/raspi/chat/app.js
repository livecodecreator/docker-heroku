new Vue({

    el: "#app",

    self: null,
    ws: null,

    data: {
        message: "",
        messages: "",
    },

    created: function() {

        self = this
    },

    methods: {

        send: function () {

            self.ws.send(
                JSON.stringify({
                    message: this.message
                })
            );

            self.message = "";
        },

        wsConnect: function() {

            var protocol = null

            if (location.protocol == "http:") protocol = "ws:"
            if (location.protocol == "https:") protocol = "wss:"

            var url = [protocol, "", location.host, "public", "raspi", "message"].join("/")

            self.ws = new WebSocket(url)

            self.ws.addEventListener("open", function(event) {
                console.log(event)
            })

            self.ws.addEventListener("close", function(event) {
                console.log(event)
                self.ws = null
            })

            self.ws.addEventListener("error", function(event) {
                console.log(event)
                self.ws = null
            })

            self.ws.addEventListener("message", function(event) {
                console.log(event)
                var data = JSON.parse(event.data);
                self.messages += data.message + "<br/>";
            })
        },

        wsReConnect: function() {
            if (self.ws != null) return
            self.wsConnect()
        }
    },

    mounted: function(){
        self.wsReConnectTimer = setInterval(self.wsReConnect, 5000)
        self.wsConnect()
    }
});
