/* if (
  localStorage.getItem("color-theme") === "dark" ||
  (!("color-theme" in localStorage) &&
    window.matchMedia("(prefers-color-scheme: dark)").matches)
) {
  document.documentElement.classList.add("dark");
} else {
  document.documentElement.classList.remove("dark");
}

var themeToggleDarkIcon = document.getElementById("theme-toggle-dark-icon");
var themeToggleLightIcon = document.getElementById("theme-toggle-light-icon");

// Change the icons inside the button based on previous settings
if (
  localStorage.getItem("color-theme") === "dark" ||
  (!("color-theme" in localStorage) &&
    window.matchMedia("(prefers-color-scheme: dark)").matches)
) {
  themeToggleLightIcon.classList.remove("hidden");
} else {
  themeToggleDarkIcon.classList.remove("hidden");
}

var themeToggleBtn = document.getElementById("theme-toggle");

themeToggleBtn.addEventListener("click", function () {
  // toggle icons inside button
  themeToggleDarkIcon.classList.toggle("hidden");
  themeToggleLightIcon.classList.toggle("hidden");

  // if set via local storage previously
  if (localStorage.getItem("color-theme")) {
    if (localStorage.getItem("color-theme") === "light") {
      document.documentElement.classList.add("dark");
      localStorage.setItem("color-theme", "dark");
    } else {
      document.documentElement.classList.remove("dark");
      localStorage.setItem("color-theme", "light");
    }

    // if NOT set via local storage previously
  } else {
    if (document.documentElement.classList.contains("dark")) {
      document.documentElement.classList.remove("dark");
      localStorage.setItem("color-theme", "light");
    } else {
      document.documentElement.classList.add("dark");
      localStorage.setItem("color-theme", "dark");
    }
  }
});*/

// set the modal menu element
const $targetEl = document.getElementById("crud-modal");

// options with default values
const options = {
  placement: "bottom-right",
  backdrop: "dynamic",
  backdropClasses: "bg-gray-900/50 dark:bg-gray-900/80 fixed inset-0 z-40",
  closable: true,
  onHide: () => {
    console.log("modal is hidden");
  },
  onShow: () => {
    console.log("modal is shown");
  },
  onToggle: () => {
    console.log("modal has been toggled");
  },
};

// instance options object
const instanceOptions = {
  id: "crud-modal",
  override: true,
};

const modal = new Modal($targetEl, options, instanceOptions);

document.querySelector("#new-date").addEventListener("click", () => {
  modal.toggle();
});

////////////////////////////////////////////////////////////////////////////////////////
const daysInMonth = (month, year) => new Date(year, month, 0).getDate();

const getFirstDayOfWeek = (month, year) =>
  new Date(year, month - 1, 1).getDay();

const fetchAppointments = (year, month) => {
  return fetch(`/api/${year}/${month}`)
    .then((res) => res.json())
    .then((data) =>
      data.map((item) => ({
        date: item.date,
        hour: item.hour,
        desc: item.desc,
        id: item.id, // Add ID here
      }))
    )
    .catch((err) => {
      console.error(err);
      return [];
    });
};

const generateCalendar = (month, year, appointments) => {
  const agenda = document.getElementById("agenda");
  agenda.innerHTML = "";

  const daysInPrevMonth = daysInMonth(month - 1, year);
  const daysInCurrentMonth = daysInMonth(month, year);
  const firstDayOfWeek = getFirstDayOfWeek(month, year);

  let dayCount = 1;
  let prevMonthDayCount = daysInPrevMonth - firstDayOfWeek + 1;
  let nextMonthDayCount = 1;

  const today = new Date();
  const todayString = `${today.getFullYear()}-${(today.getMonth() + 1)
    .toString()
    .padStart(2, "0")}-${today.getDate().toString().padStart(2, "0")}`;

  for (let i = 0; i < firstDayOfWeek; i++) {
    const dayElement = document.createElement("div");
    dayElement.classList.add("day", "empty");
    dayElement.textContent = prevMonthDayCount++;
    agenda.appendChild(dayElement);
  }

  for (let i = 0; i < daysInCurrentMonth; i++) {
    const dayElement = document.createElement("div");
    dayElement.classList.add("day");

    const dateString = `${year}-${month.toString().padStart(2, "0")}-${dayCount
      .toString()
      .padStart(2, "0")}`;
    dayElement.id = dateString; // Add ID to the day element

    if (dateString === todayString) {
      dayElement.classList.add("today"); // Highlight today's date with a red background
    }

    const dateSpan = document.createElement("span");
    dateSpan.classList.add("date");
    dateSpan.textContent = dayCount;
    dayElement.appendChild(dateSpan);

    const appointmentsContainer = document.createElement("div");
    appointmentsContainer.classList.add("appointments");

    const dayAppointments = appointments.filter(
      (app) => app.date === dateString
    );

    if (dayAppointments.length > 1) {
      const countElement = document.createElement("div");
      countElement.classList.add("appointment-count");
      countElement.textContent = dayAppointments.length;
      appointmentsContainer.appendChild(countElement);
    } else if (dayAppointments.length === 1) {
      const dot = document.createElement("div");
      dot.classList.add("appointment-dot");
      appointmentsContainer.appendChild(dot);
    }

    dayElement.appendChild(appointmentsContainer);
    agenda.appendChild(dayElement);

    dayElement.addEventListener("mouseenter", (event) => {
      showPopup(event, dayAppointments);
    });

    dayElement.addEventListener("mouseleave", hidePopup);

    dayCount++;
  }

  const totalCells = 42;
  const daysToAdd = totalCells - (firstDayOfWeek + daysInCurrentMonth);

  for (let i = 0; i < daysToAdd; i++) {
    const dayElement = document.createElement("div");
    dayElement.classList.add("day", "empty");
    dayElement.textContent = nextMonthDayCount++;
    agenda.appendChild(dayElement);
  }
};

const showPopup = (event, appointments) => {
  const popup = document.getElementById("popup");
  popup.innerHTML = ""; // Clear previous content

  appointments.forEach((appointment) => {
    const appElement = document.createElement("div");
    appElement.classList.add("appointment");
    appElement.id = `appointment-${appointment.id}`; // Add ID to the appointment element
    appElement.textContent = `${appointment.hour} - ${appointment.desc}`;
    popup.appendChild(appElement);
  });

  const rect = event.target.getBoundingClientRect();
  popup.style.top = `${rect.bottom + window.scrollY}px`;
  popup.style.left = `${rect.left + window.scrollX}px`;
  popup.classList.add("visible");
};

const hidePopup = () => {
  const popup = document.getElementById("popup");
  popup.classList.remove("visible");
};

const updateCalendar = async (year, month) => {
  const appointments = await fetchAppointments(year, month);
  generateCalendar(month, year, appointments);
  const monthNames = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
  ];
  document.getElementById("current-month").textContent = `${
    monthNames[month - 1]
  } ${year}`;
};

let currentYear = new Date().getFullYear();
let currentMonth = new Date().getMonth() + 1;

document.getElementById("prev-month").addEventListener("click", () => {
  currentMonth--;
  if (currentMonth < 1) {
    currentMonth = 12;
    currentYear--;
  }
  updateCalendar(currentYear, currentMonth);
});

document.getElementById("next-month").addEventListener("click", () => {
  currentMonth++;
  if (currentMonth > 12) {
    currentMonth = 1;
    currentYear++;
  }
  updateCalendar(currentYear, currentMonth);
});

updateCalendar(currentYear, currentMonth);

$(document).ready(function () {
  $("#date-form").on("submit", function (event) {
    event.preventDefault();

    $.ajax({
      type: "POST",
      url: "/api/new",
      data: $(this).serialize(),
      success: function (response) {
        alert("Form submitted successfully: " + response);
        $("#date-form")[0].reset();

        updateCalendar(currentYear, currentMonth);
      },
      error: function (xhr, status, error) {
        alert("An error occurred: " + error);
        console.log(xhr.responseText);
        console.log(error);
      },
    });
  });
});

document.getElementById("popup").addEventListener("mouseenter", () => {
  document.getElementById("popup").classList.add("visible");
});

document.getElementById("popup").addEventListener("mouseleave", hidePopup);
