# ASCII Art Web

This project is a web-based application that converts text into ASCII art using different banner styles.

## Features

-   Web interface to input text and select a banner style.
-   Supports multiple banner styles: standard, shadow, and thinkertoy.
-   Renders ASCII art in the browser.

## Getting Started

### Prerequisites

-   Go (Golang) installed on your system.

### Installation

1.  Clone the repository:
    ```bash
    git clone https://github.com/your-username/ascii-art-web.git
    ```
2.  Navigate to the project directory:
    ```bash
    cd ascii-art-web
    ```

### Running the application

1.  Run the Go application:
    ```bash
    go run main.go
    ```
2.  Open your web browser and go to `http://localhost:8080`

## Usage

1.  Enter the text you want to convert in the text area.
2.  Select a banner style from the dropdown menu.
3.  Click the "Submit" button to generate the ASCII art.

## Project Structure

-   `main.go`: The main application file containing the web server and logic for handling requests and generating ASCII art.
-   `banners/`: Directory containing the banner files (`standard.txt`, `shadow.txt`, `thinkertoy.txt`).
-   `static/`: Directory containing static assets like CSS and HTML.
-   `go.mod`: Go module file.
-   `main_test.go`: Test file for the main application.
