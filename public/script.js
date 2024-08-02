fetch("/api")
  .then(function (res) {
    return res.json();
  })
  .then((data) => {
    console.log(data);
    const dates = document.getElementById("dates");
    data.forEach((date) => {
      const li = document.createElement("li");
      li.textContent = `${date.date} ${date.hourlyData.hour} ${date.hourlyData.data}`;
      dates.appendChild(li);
    });
  })
  .catch((err) => {
    console.log(err);
  });
