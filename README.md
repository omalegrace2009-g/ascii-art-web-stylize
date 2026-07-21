# 🎨 ASCII Art Web Stylize

A responsive web application built with **Go** that converts plain text into ASCII art using different banner styles. This project enhances the original **ASCII Art Web** by adding a clean, modern, and responsive user interface.

---

## 📚 Table of Contents

* [Overview](#-overview)
* [Features](#-features)
* [Demo](#-demo)
* [Technology Stack](#-technology-stack)
* [Project Structure](#-project-structure)
* [Getting Started](#-getting-started)
* [Usage](#-usage)
* [Application Flow](#-application-flow)
* [HTTP Status Codes](#-http-status-codes)
* [Future Improvements](#-future-improvements)
* [Author](#-author)
* [License](#-license)

---

## 📖 Overview

ASCII Art Web Stylize is a Go web application that generates ASCII art from user input. It extends the original project with CSS styling, responsive layouts, and an improved user experience while keeping the original rendering functionality.

---

## ✨ Features

* Generate ASCII art from text
* Three banner styles:

  * Standard
  * Shadow
  * Thinkertoy
* Responsive interface
* Custom error pages
* Go template rendering
* Static CSS support
* Uses only Go's standard library

---

## 🖥️ Demo

Run the server and visit:

```text id="xyq1pg"
http://localhost:8080
```

---

## 🛠 Technology Stack

### Backend

* Go
* `net/http`
* `html/template`

### Frontend

* HTML5
* CSS3

### Tools

* Git
* GitHub

---

## 📁 Project Structure

```text id="4njlwm"
ascii-art-web-stylize/
│
├── ascii/
│   ├── Generate.go
│   ├── LoadBanner.go
│   └── Render.go
│
├── banner/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
│
├── static/
│   └── style.css
│
├── templates/
│   ├── base.html
│   ├── input.html
│   ├── output.html
│   └── error.html
│
├── ArtHandler.go
├── ErrorHandler.go
├── HomeHandler.go
├── main.go
├── go.mod
└── README.md
```

---

## 🚀 Getting Started

### Prerequisites

* Go 1.24 or later
* Git

### Clone the Repository

```bash id="9cb6t3"
git clone https://github.com/OmaleGrace/ascii-art-web-stylize.git
cd ascii-art-web-stylize
```

### Run the Application

```bash id="k1s5xa"
go run .
```

Open:

```text id="z9o3iv"
http://localhost:8080
```

---

## 🎨 Usage

1. Enter your text.
2. Select a banner style.
3. Click **Generate**.
4. View the generated ASCII art.

---

## 🔄 Application Flow

```text id="xxh8fw"
User Input
    ↓
HTTP Request
    ↓
Go Handler
    ↓
Generate ASCII Art
    ↓
Render Template
    ↓
Display Result
```

---

## 🌐 HTTP Status Codes

| Code    | Description              |
| ------- | ------------------------ |
| **200** | Success                  |
| **400** | Invalid input            |
| **404** | Page or banner not found |
| **500** | Internal server error    |

---

## 🎯 Future Improvements

* More banner styles
* Dark mode
* Copy-to-clipboard
* Download ASCII art
* Docker support
* Automated testing

---

## 👨‍💻 Author

**Ooja Omale (Grace)**

GitHub: https://github.com/OmaleGrace

---

## 📜 License

Developed for educational purposes as part of the **01-edu** curriculum.
