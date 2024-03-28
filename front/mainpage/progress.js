
const updateProgress = () => {
    fetch('/user/progress', {
        method: 'post',
    })
    .then(res => res.json())
    .then(data =>{
        for (var i = 0; i < data.progress.length; i++) {
            document.getElementById("progress_"+(i+1)).innerHTML = data.progress[i];
        }
   
    })

}

updateProgress()
