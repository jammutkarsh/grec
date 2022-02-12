// click.innerHTML = "Click me";
// click.onclick = function () {
//   copy("data");
// };
// let dc = document.body.appendChild(click);

var func = () => {
  let tableBody = document.getElementById("data");
  var td, btn;
  tableBody.appendChild((td = document.createElement("td")));
  td.appendChild((btn = document.createElement("button")));

  btn.innerHTML = "Click Me";
  btn.onclick = function () {
    copy("data");
  };
  td.appendChild(btn);
};
fetch("/.hidden/lists.json").then((res) => {
  res.json().then((data) => {
    // let myJSON = data;
    // myArray = JSON.parse(myJSON);
    console.log(data);
    if (data.length > 0) {
      var temp = "";

      for (let i = 0; i < data.length; i++) {
        temp += "<tr>";
        temp += "<td>" + data[i] + "</td>";
        temp += "<td>" + func() + "</td></tr>";
      }
      document.getElementById("data").innerHTML = temp;
    }
  });
});

function copy(element_id) {
  var aux = document.createElement("div");
  aux.setAttribute("contentEditable", true);
  aux.innerHTML = document.getElementById(element_id).innerHTML;
  aux.setAttribute("onfocus", "document.execCommand('selectAll',false,null)");
  document.body.appendChild(aux);
  aux.focus();
  document.execCommand("copy");
  document.body.removeChild(aux);
}
