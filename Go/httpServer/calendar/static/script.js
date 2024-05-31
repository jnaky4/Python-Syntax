function getCurrentFormattedDate() {
    const currentDate = new Date();
    const year = currentDate.getFullYear();
    const month = currentDate.getMonth() + 1;
    const day = currentDate.getDate();

    return `${year}-${month}-${day}`;
}

const currentFormattedDate = getCurrentFormattedDate();

// Calculate holiday dates for the current year
const currentYear = new Date().getFullYear();

function getHolidayDates(startYear, endYear) {
    const holidayDates = [];

    for (let year = startYear - 1; year <= endYear + 1; year++) {
        // New Year's Day and Independence Day have fixed dates
        const fixedHolidays = [
            new Date(year, 0, 1),    // New Year's Day (January 1st)
            new Date(year, 6, 4),    // Independence Day (July 4th)
            new Date(year, 11, 25),  // Christmas Day (December 25th)
            new Date(year, 5, 19),   // Juneteenth (June 19th)
        ];

        // Calculate Memorial Day (last Monday in May)
        const memorialDay = new Date(year, 4, 31);
        while (memorialDay.getDay() !== 1) {
            memorialDay.setDate(memorialDay.getDate() - 1);
        }

        // Calculate Labor Day (first Monday in September)
        const laborDay = new Date(year, 8, 1);
        while (laborDay.getDay() !== 1) {
            laborDay.setDate(laborDay.getDate() + 1);
        }

        // Format the dates to `${year}-${month}-${day}` and add to holidayDates
        const formattedHolidays = [
            ...fixedHolidays,
            memorialDay,
            laborDay,
        ].map(date => {
            const month = date.getMonth() + 1;
            const day = date.getDate();
            return `${year}-${month}-${day}`;
        });

        holidayDates.push(...formattedHolidays);
    }

    return holidayDates;
}

// Calculate Labor Day (first Monday in September)
function getLaborDay(year){
    const laborDay = new Date(year, 8, 1);
    while (laborDay.getDay() !== 1) {
        laborDay.setDate(laborDay.getDate() + 1);
    }
    return laborDay
}

// Calculate Memorial Day (last Monday in May)
function getMemorialDay(year){
    const memorialDay = new Date(year, 4, 31);
    while (memorialDay.getDay() !== 1) {
        memorialDay.setDate(memorialDay.getDate() - 1);
    }
    return memorialDay
}

const holidayDates = getHolidayDates(currentYear-1, currentYear+1);

// Function to style the current day cell
function styleCurrentDayCell(dayCell) {
    const cellDate = dayCell.getAttribute("id");

    if (cellDate === currentFormattedDate) {
        dayCell.style.backgroundColor = 'black';
        dayCell.style.color = 'white';
    }
}

// Function to style Sunday and Saturday cells
function styleWeekendCell(dayCell) {
    const dayOfWeek = dayCell.textContent.substr(-3);

    if (dayOfWeek === 'Sun' || dayOfWeek === 'Sat') {
        dayCell.style.backgroundColor = 'red';
        dayCell.style.color = 'white';
    }
}

// Function to style holiday cells
function styleHolidayCell(dayCell) {
    const cellDate = dayCell.getAttribute("id");

    if (holidayDates.includes(cellDate)) {
        dayCell.style.backgroundColor = 'red';
        dayCell.style.color = 'white';
    }
}

// Function to populate the calendar for a given date
function populate7DayCalendar(startDate) {
    const daysRow = document.getElementById('calendarDays');
    const textAreasRow = document.getElementById('textAreas');

    // Clear existing calendar cells
    daysRow.innerHTML = '';
    textAreasRow.innerHTML = '';

    for (let i = 0; i < 7; i++) {
        const day = new Date(startDate);
        day.setDate(startDate.getDate() + i);
        const formattedDay = day.toLocaleString('en-US', { weekday: 'short', day: 'numeric' });
        const fullyQualifiedDate = `${day.getFullYear()}-${day.getMonth() + 1}-${day.getDate()}`;

        // Create a cell for the day of the week and day with id representing the fully qualified date
        const dayCell = document.createElement('td');
        dayCell.textContent = formattedDay;
        dayCell.setAttribute("id", fullyQualifiedDate);
        dayCell.classList.add("dayCell");

        daysRow.appendChild(dayCell);
        styleCurrentDayCell(dayCell);
        styleWeekendCell(dayCell);
        styleHolidayCell(dayCell);

        // Create a cell with a textarea
        const textAreaCell = document.createElement('td');
        const textArea = document.createElement('textarea');
        textArea.style.width = '100%'; // Set a fixed width
        textArea.style.resize = 'vertical'; // Allow height to grow
        textArea.rows = 5;
        textArea.placeholder = 'Enter your notes here';
        textAreaCell.appendChild(textArea);
        textAreasRow.appendChild(textAreaCell);
    }
}

// Initialize the calendar with the current date
const currentDate = new Date();
populate7DayCalendar(currentDate);

// Event listener for the "Next" button
document.getElementById('nextButton').addEventListener('click', function () {
    currentDate.setDate(currentDate.getDate() + 7);
    populate7DayCalendar(currentDate);
});

// Event listener for the "Last" button
document.getElementById('lastButton').addEventListener('click', function () {
    currentDate.setDate(currentDate.getDate() - 7);
    populate7DayCalendar(currentDate);
});

// Add "Add Text Area" buttons and functionality
const addButtonRow = document.querySelector('.add-button-area');

for (let i = 0; i < 7; i++) {
    const addButtonCell = document.createElement('td');
    const addButton = document.createElement('button');
    addButton.textContent = 'Add Text Area';
    addButton.classList.add('add-button');

    addButton.addEventListener('click', function () {
        const textAreasRow = document.getElementById('textAreas');
        const row = textAreasRow.children[i];

        const textArea = document.createElement('textarea');
        textArea.style.width = '100%'; // Set a fixed width
        textArea.style.resize = 'vertical'; // Allow height to grow
        textArea.rows = 5;
        textArea.placeholder = 'Enter your notes here';
        row.appendChild(textArea);
    });

    addButtonCell.appendChild(addButton);
    addButtonRow.appendChild(addButtonCell);
}

// Append addButtonRow to the HTML table
const calendarTable = document.querySelector('.calendar');
calendarTable.appendChild(addButtonRow);
