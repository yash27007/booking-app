# 🎟️ Ticket Booking Console App (Go)

A simple console-based ticket booking application written in Go.  
This app allows users to book tickets for a conference with input validation, user-friendly prompts, and modular structure.

---

## 🛠 Features

- Book tickets for a conference via CLI
- Validates user input (name, email, ticket count)
- Displays confirmation and remaining tickets
- Clean project structure using packages (`main` and `helper`)
- Uses Go features like structs, slices, functions, and packages

---

## 📁 Project Structure

```

booking-app/
├── go.mod                # Go module file
├── main.go               # Main program logic
└── helper/
└── helper.go         # Helper functions (greetings, validation)

````

---

## 🚀 Getting Started

### Prerequisites

- Go 1.24.0 or higher installed

### Running the App

```bash
git clone https://github.com/yash27007/booking-app.git
cd booking-app
go run main.go
````

Follow the prompts to book your ticket!

---

## 📦 Modules & Packages

* **main.go**: Contains the core logic – input, booking, and managing state.
* **helper/helper.go**: Includes:

  * `GreetUsers()` – Displays a welcome message.
  * `ValidateUserInput()` – Validates name, email, and ticket count.

---

## 🧪 Example Run

```
Welcome to Go Conference booking application!
We have a total of 50 tickets and 50 are still available.
Get your tickets here to attend!

Enter your first name: John
Enter your last name: Doe
Enter your email address: john.doe@example.com
Enter number of tickets: 2
Thank you John Doe for booking 2 tickets. A confirmation email has been sent to john.doe@example.com.
48 tickets remaining for Go Conference.
```

---

## 📚 Concepts Used

* Structs
* Functions
* Packages
* Slices
* Input/output (`fmt.Scan`)
* String operations (`strings.Contains`)
* Basic validation

---

## 🤝 Contributing

Pull requests and suggestions are welcome!
Feel free to fork the repository and improve it (e.g., add concurrency, store bookings in a file, etc.)

---

## 📃 License

This project is open-source and free to use under the [MIT License](LICENSE).


