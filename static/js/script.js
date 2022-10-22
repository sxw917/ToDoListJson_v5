let items = document.getElementsByTagName("li")

const toggleItemStatus = (id, item) => {
    var xhttp = new XMLHttpRequest();
    xhttp.open("POST", "/toggle", true);
    xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) { //If connection is good
            if (this.responseText === "true") {
                item.classList.add("done");
            } else if (this.responseText === "false") {
                item.classList.remove("done")
            } else {
                alert(this.responseText);
            }
        }
    };
    xhttp.send(`id=${id}`); //??????????????????????????????
    // xhttp.send("id=" + id);

    // "id=" + id `ID￥`
}

for(let i = 0; i < items.length; i++) {
    items[i].addEventListener("click", () => {
        var id = items[i].getAttribute('data-id');
        toggleItemStatus(id, items[i]);
    })
}


