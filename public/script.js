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

/*fetch("/api")
  .then(function (res) {
    return res.json();
  })
  .then((data) => {
    console.log(data);
    const dates = document.getElementById("dates");
    data.forEach((date) => {
      const li = document.createElement("li");
      li.innerHTML = `
  <a href="#" class="block max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700">
    <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">${date.date}</h5>
    <p class="font-normal text-gray-700 dark:text-gray-400">${date.hour} ${date.desc}</p>
  </a>
      `;
      dates.appendChild(li);
    });
  })
  .catch((err) => {
    console.log(err);
  });
*/

// Function to get the number of days in a month
const daysInMonth = (month, year) => new Date(year, month, 0).getDate();

// Function to get the day of the week of the first day of the month
const getFirstDayOfWeek = (month, year) =>
  new Date(year, month - 1, 1).getDay();

// Fetch appointments from the server
const fetchAppointments = (year, month) => {
  return fetch(`/api/${year}/${month}`)
    .then((res) => res.json())
    .then((data) =>
      data.map((item) => ({
        date: item.date,
        hour: item.hour,
        desc: item.desc,
      }))
    )
    .catch((err) => {
      console.error(err);
      return [];
    });
};

// Generate the calendar for the specified month and year
const generateCalendar = (month, year, appointments) => {
  const agenda = document.getElementById("agenda");
  agenda.innerHTML = ""; // Clear previous content

  const daysInPrevMonth = daysInMonth(month - 1, year);
  const daysInCurrentMonth = daysInMonth(month, year);
  const firstDayOfWeek = getFirstDayOfWeek(month, year);

  let dayCount = 1;
  let prevMonthDayCount = daysInPrevMonth - firstDayOfWeek + 1;
  let nextMonthDayCount = 1;

  // Generate days for the previous month
  for (let i = 0; i < firstDayOfWeek; i++) {
    const dayElement = document.createElement("div");
    dayElement.classList.add("day", "empty");
    dayElement.textContent = prevMonthDayCount++;
    agenda.appendChild(dayElement);
  }

  // Generate days for the current month
  for (let i = 0; i < daysInCurrentMonth; i++) {
    const dayElement = document.createElement("div");
    dayElement.classList.add("day");

    const dateSpan = document.createElement("span");
    dateSpan.classList.add("date");
    dateSpan.textContent = dayCount;
    dayElement.appendChild(dateSpan);

    const appointmentsContainer = document.createElement("div");
    appointmentsContainer.classList.add("appointments");

    const dateString = `${year}-${month.toString().padStart(2, "0")}-${dayCount
      .toString()
      .padStart(2, "0")}`;
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

    dayCount++;
  }

  // Calculate how many days to add to complete 6 weeks (42 days total)
  const totalCells = 42;
  const daysToAdd = totalCells - (firstDayOfWeek + daysInCurrentMonth);

  // Generate days for the next month to fill the week
  for (let i = 0; i < daysToAdd; i++) {
    const dayElement = document.createElement("div");
    dayElement.classList.add("day", "empty");
    dayElement.textContent = nextMonthDayCount++;
    agenda.appendChild(dayElement);
  }
};

// Update the calendar with new appointments
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
let currentMonth = new Date().getMonth() + 1; // getMonth is zero-based

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

// Initialize the calendar
updateCalendar(currentYear, currentMonth);

// Handle form submission
$(document).ready(function () {
  $("#date-form").on("submit", function (event) {
    event.preventDefault(); // Prevent the form from submitting the traditional way

    $.ajax({
      type: "POST",
      url: "/api/new", // URL to submit the form data
      data: $(this).serialize(), // Serialize the form data
      success: function (response) {
        alert("Form submitted successfully: " + response); // Handle the response from the server
        $("#date-form")[0].reset(); // Optionally reset the form fields

        // Refresh the appointments for the current month
        updateCalendar(currentYear, currentMonth);
      },
      error: function (xhr, status, error) {
        alert("An error occurred: " + error); // Handle any errors
        console.log(xhr.responseText);
        console.log(error);
      },
    });
  });
});
