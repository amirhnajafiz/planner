function updateDb(item) {
    let input = document.getElementById(item)
    let newItem = input.value
    fetch(`/update?olditem=${item}&newitem=${newItem}`, {method: "PUT"}).then(res =>{
        if (res.status === 200){
            alert("Database updated")
            window.location.pathname = "/"
        }
    })
}

function removeFromDb(item) {
    fetch(`/delete?item=${item}`, {method: "Delete"}).then(res =>{
        if (res.status === 200){
            window.location.pathname = "/"
        }
    })
}
