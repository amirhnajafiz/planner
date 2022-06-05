// this method updates an item
function updateDb(item) {
    let input = document.getElementById(item)
    let newItem = input.value

    let url = `/update?olditem=${item}&newitem=${newItem}`

    fetch(url, {method: "PUT"})
        .then(res => {
            if (res.status === 200) {
                alert("Item removed")

                window.location.pathname = "/"
            }
        })
        .catch(e => console.error(e))
}

// this method removes an item
function removeFromDb(item) {
    let url = `/delete?item=${item}`

    fetch(url, {method: "Delete"})
        .then(res => {
            if (res.status === 200) {
                window.location.pathname = "/"
            }
        })
        .catch(e => console.error(e))
}
