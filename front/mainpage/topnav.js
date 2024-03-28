
const logoutBtn = document.querySelector('#logout_button');

logoutBtn.addEventListener('click', () => {
    alertBox("loging out")
    setTimeout(function () {
        window.location.href = "/user/pleas_log_out_because_js_cant";
     }, 1000);
})


const alertBox = (data) => {
    const alertContainer = document.querySelector('.alert-box');
    const alertMsg = document.querySelector('.alert');
    alertMsg.innerHTML = data;
    alertContainer.style.background = 'red'

    alertContainer.style.top = `5%`;
    setTimeout(() => {
        alertContainer.style.top = null;
    }, 5000);
}
