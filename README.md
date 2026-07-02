# ASCII-Art-Web-Stylize

An extension of [ascii-art-web](https://acad.learn2earn.ng/git/dodumu/ascii-art-web) that transforms the plain HTTP-based ASCII art generator into a **visually appealing, responsive, and user-friendly** web experience through the addition of CSS.

## 📖 About

This project builds on the original `ascii-art-web` server, which converts user-submitted text into ASCII art using different banner styles (`standard`, `shadow`, `thinkertoy`). While the original focused purely on functionality, **ascii-art-web-stylize** focuses on the *human-computer interface* — making the site pleasant to look at, easy to navigate, and consistent across devices.

## ✨ Goals

- Make the interface **appealing, interactive, and intuitive**
- Improve **user feedback** (loading states, error messages, hover effects, etc.)
- Ensure **text readability** regardless of color choices (sufficient contrast)
- Keep the design **consistent** across all pages and components
- Make the layout **responsive** so it works well on desktop, tablet, and mobile
- Add **interactivity** without breaking core functionality

## 🛠️ Tech Stack

- **Backend:** Go (standard library only — `net/http`, `html/template`, etc.)
- **Frontend:** HTML, CSS
- **No external Go packages** — per project constraints, only the Go standard library is allowed

## 🚀 Getting Started

### Prerequisites
- [Go](https://golang.org/doc/install) installed on your machine

### Installation
```bash
git clone https://acad.learn2earn.ng/git/dodumu/ascii-art-web-stylize
cd ascii-art-web-stylize
```

### Running the server
```bash
go run main.go
```

Then open your browser and navigate to:
```
http://localhost:8080
```

## 🎨 Usage

1. Enter the text you want converted into ASCII art.
2. Choose a banner style: `standard`, `shadow`, or `thinkertoy`.
3. Submit the form and view your styled ASCII art output.

## 📁 Project Structure

```
ascii-art-web-stylize/
├── main.go
├── HomeHandler.go
├── ErrorHandler.go
├── ArtHandler.go
├── banner/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── ascii/
│   ├── LoadBanner.go
│   ├── Render.go
│   └── Generate.go
├── templates/
│   ├── base.html
│   ├── input.html
│   ├── output.html
│   └── error.html
└── README.md
```

## ✅ Status Codes

| Code | Meaning                              |
|------|---------------------------------------|
| 200  | Success — ASCII art generated         |
| 400  | Bad request — empty or invalid input  |
| 404  | Not found — invalid banner style/page |
| 500  | Internal server error                 |

## 👥 Authors

- **Ooja Omale (Grace)**


## 📜 License

This project is developed for educational purposes as part of the 01-edu curriculum.