(async () => {
    try {
        let res = await fetch('/api')
        let data = await res.json()
        console.log("data=", data)
        if (data) {
            let div = document.getElementById('todos')
            div.innerHTML = ''
            data.map(item => {
                div.innerHTML += `<div class="mt1">
                <a class="delete" id="${item.id}">&#x2715</a>
                <p>${item.todo}</p>
                </div>`
            })
        }
    } catch (error) {
        console.log(error)
    }
})();

const form = document.querySelector('form')
form.addEventListener('submit', async e => {
    e.preventDefault()
    let id = "id_" + Math.random().toString(16).slice(12)
    let todo = form.elements['todo'].value
    let data = { id, todo }
    try {
        await fetch('/api/submit', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })

        let div = document.getElementById('todos')
        div.innerHTML += `<div class="mt1">
        <a class="delete" id="${data.id}">&#x2715</a>
        <p>${data.todo}</p>
    </div>`
        document.querySelector('input').value = ''
    } catch (error) {
        console.log(error)
    }
})

document.body.addEventListener("click", async function (e) {
    if (e.target.classList.contains("delete")) {
        if (e.target.id) {
            try {
                await fetch(`/api/delete?id=${e.target.id}`, {
                    method: 'DELETE',
                })
                e.target.parentNode.remove();
            } catch (error) {
                console.log(error)
            }
        }
    }
})