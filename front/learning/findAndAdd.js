
const form = document.getElementById('wf-form-serch-Form');
const dialog = document.querySelector("dialog");
const dialogCloseButton = dialog.querySelector("#cancelBtn");
const dialogConfirmButton = dialog.querySelector("#confirmBtn");
const radios = document.getElementsByClassName("w-form-formradioinput w-radio-input")

function checkOrCross(bool) {
  if (bool) {
    return '✓'
  } else {
    return "𐄂"
  }
}
function createNewRow(table, json, funtionOnClick) {
  characters = json["data"]["characters"]
  pReading = findPrimary(json["data"]["readings"])["reading"]
  pMening = findPrimary(json["data"]["meanings"])["meaning"]
  isLearning = json["data"]["is_learning"]
  isBurned= json["data"]["is_burned"]
  // isBurned = findPrimary(json["data"]["meanings"])["meaning"]

  var row = table.insertRow(1);

  // Insert new cells (<td> elements) at the 1st and 2nd position of the "new" <tr> element:
  var cell1 = row.insertCell(0);
  var cell2 = row.insertCell(1);
  var cell3 = row.insertCell(2);
  var cell4 = row.insertCell(3);
  var cell5 = row.insertCell(4);
  cell1.innerHTML = '<a target="_blank">' + characters + '</a>'
  cell2.innerHTML = '<a target="_blank">' + pReading + '</a>'
  cell3.innerHTML = '<a target="_blank">' + pMening + '</a>'

  cell4.innerHTML = '<a target="_blank">' + checkOrCross(isLearning) + '</a>'
  cell5.innerHTML = '<a target="_blank">' + checkOrCross(isBurned) + '</a>'

  row.addEventListener("click", () => {
    dialog.querySelector("p").innerHTML = characters + " - " + pReading + " - " + pMening
    funtionOnClick()
  });
}



function findPrimary(xs) {
  for (const x of xs) {
    if (x["primary"] == true) {
      return x
    }
  }
}


dialogConfirmButton.addEventListener("click", () => {

  fetch("/learning/add", {
    method: "POST",
    body: JSON.stringify({
      id: Number(document.getElementById("selectedItemId").value)
    }),
    headers: {
      "Content-type": "application/json; charset=UTF-8"
    }
  })

  // TODFO
  dialog.close();
});
dialogCloseButton.addEventListener("click", () => {

  dialog.close("test")
});

function showItem(json) {


  table = document.getElementById("voc-table")
  createNewRow(table, json, function () {
    document.getElementById("selectedItemId").value = json["id"];
    dialog.showModal();
  });
}

function getResults() {
  fetch("/learning/find/submit", {
    method: "POST",
    body: JSON.stringify({
      type: form.elements["type"].value,
      searchReading: form.elements["checkbox-2"].checked,
      searchWriting: form.elements["checkbox-3"].checked,
      searchText: form.elements["inText"].value
    }),
    headers: {
      "Content-type": "application/json; charset=UTF-8"
    }
  })
    .then((response) => response.json())
    .then((json) =>
      json.forEach(showItem)
    );

}

form.addEventListener("submit", function (event) {
  // stop form submission
  event.preventDefault();

  getResults();
});

for (const radio of radios) {
  radio.onclick = (e) => {
    var checkboxes = document.getElementsByClassName("checkbox-field -v")
    if (radios[0].checked) {
      for (const box of checkboxes) {
        box.style.visibility = "visible"
      }
    } else {
      for (const box of checkboxes) {
        box.style.visibility = "hidden"
      }
    }
  }
}