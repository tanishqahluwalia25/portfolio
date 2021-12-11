const skills = document.getElementsByClassName("skill");

for (let i = 0; i < skills.length; ++i) {
  const el = skills[i].firstElementChild;

  const checkHeighAnim = () => {
    let dist = el.getBoundingClientRect().top - window.innerHeight;
    if (dist < 10) {
      el.classList.add(`${skills[i].id}-skill`);
      window.removeEventListener("scroll", checkHeighAnim);
    }
  };

  window.addEventListener("scroll", checkHeighAnim);
}

function submit() {
  alert("hi");
  const user = {
    name: "tanishq",
    message: "haujks",
    email: "ssujds",
    subject: "sdnjkls",
  };
  const options = {
    method: "POST",
    body: JSON.stringify(user),
    headers: {
      "Content-Type": "application/json",
    },
  };
  url = "http://localhost:4000/query";
  // send POST request
  fetch(url, options)
    .then((res) => {
      console.log(res.json());
    })
    .then((res) => console.log(res));
}
