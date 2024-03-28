
const updateButton = (button_prefix) => {
    fetch('/' + button_prefix + '/queue_count', {
        method: 'post',
    })
        .then(res => res.json())
        .then(data => {

            document.getElementById(button_prefix + "_button_text").innerHTML = "items: " + data.count;

        })

}

function updateButtons() {
    updateButton("lesson")
    updateButton("review")
    updateButton("lesson/learning")
    updateButton("review/learning")
}


updateButtons()
const interval = setInterval(function () {
    updateButtons()
}, 1000 * 60 * 5);

//  clearInterval(interval); // thanks @Luca D'Amico
