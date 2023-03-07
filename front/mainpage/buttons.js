
const updateButton = (button_prefix) => {
    fetch('/'+button_prefix+'/queue_count', {
        method: 'post',
    })
    .then(res => res.json())
    .then(data =>{

        document.getElementById(button_prefix+"_button_text").innerHTML = "items: "+data.count;
   
    })

}

updateButton("lesson")
updateButton("review")

const interval = setInterval(function() {
    updateButton("lesson")
    updateButton("review")
  }, 1000*60*5);
 
//  clearInterval(interval); // thanks @Luca D'Amico
