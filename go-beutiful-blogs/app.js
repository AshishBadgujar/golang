const form = document.querySelector('form')
form.addEventListener('submit', e => {
    e.preventDefault()
    let id = "id_" + Math.random().toString(16).slice(12)
    let name = form.elements['name'].value
    let email = form.elements['email'].value
    let message = form.elements['message'].value
    let data = { id, name, email, message }
    try {
        fetch('/api/submit', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        }).then((response) => {
            console.log("data=", response.status)
            if (response.status == 200) {
                document.getElementById("container").innerText = "Thanks!"
            }
        }).catch((error) => {
            console.log(error)
        });
    } catch (error) {
        console.log(error)
    }
})
