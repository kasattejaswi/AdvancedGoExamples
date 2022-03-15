document.onload = onLoaded();

function onLoaded() {
    var source = new EventSource("/sse/dashboard");
    source.onmessage = function(event) {
        console.log("onmessage called;");
        console.dir(event);
        var dashboard = JSON.parse(event.data)

        document.getElementById("tusers").innerHTML = dashboard["users"];
        document.getElementById("lusers").innerHTML = dashboard["users_logged_in"];
        var items = dashboard["inventory"]["items"]
        document.getElementById("bicyclesPrice").innerHTML = items["bicycle"].price;
        document.getElementById("bicyclesQuantity").innerHTML = items["bicycle"].quantity;

        document.getElementById("booksPrice").innerHTML = items["book"].price;
        document.getElementById("booksQuantity").innerHTML = items["book"].quantity;

        document.getElementById("wbPrice").innerHTML = items["wbottle"].price;
        document.getElementById("wbQuantity").innerHTML = items["wbottle"].quantity;

        document.getElementById("rcPrice").innerHTML = items["rccar"].price;
        document.getElementById("rcQuantity").innerHTML = items["rccar"].quantity;
    }
}