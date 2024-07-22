# Go Photo Tagging App

This is a Go web application that allows users to upload photos and automatically tag them using Cloudinary's API. The application is built using the Gin framework.

## Features

- Upload photos
- Automatic tagging with Cloudinary
- Display uploaded photos and tags

## Getting Started

### Prerequisites

- Go 1.16+
- Cloudinary account

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/go-photo-tagging-app.git
    cd go-photo-tagging-app
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Create a `.env` file in the root directory and add your Cloudinary credentials:
    ```plaintext
    CLOUDINARY_CLOUD_NAME=your_cloud_name
    CLOUDINARY_API_KEY=your_api_key
    CLOUDINARY_API_SECRET=your_api_secret
    ```

### Running the Application

1. Start the server:
    ```bash
    go run main.go
    ```

2. Open your browser and go to `http://localhost:8080`

### Usage

- Upload a photo using the form on the homepage.
- View the uploaded photo and its tags on the results page.

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.