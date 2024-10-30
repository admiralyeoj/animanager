# AniManager

AniManager is a Go-based service that imports the airing schedule from AniList and stores it in a database. Every minute, it checks the airing schedule to see if any scheduled anime has aired, and if so, posts an update to BlueSky.

## Table of Contents

- [AniManager](#animanager)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Usage](#usage)
  - [Commands](#commands)

## Features

- **Airing Schedule Import:** Fetches the anime airing schedule from AniList and stores it in a database.
- **BlueSky Integration:** Posts updates to BlueSky when a scheduled anime has aired.
- **Automatic Schedule Check:** Every minute, AniManager checks for aired anime and handles postings accordingly.

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/admiralyeoj/animanager.git
   cd animanager
   ```

2. **Install dependencies:*** 
    Run the following command to install Go dependencies.
    ```
    go mod tidy
    ```

3. **Setup your database:**  
   Ensure your PostgreSQL database is running. If you're using Docker, you can create and start your database by running the following command in the root of your project directory:

   ```bash
   docker-compose up -d
   ```

    Run the following command to create the tables needed
    ```bash
    make migrations
    ```

## Configuration

To configure the application, you need to set up environment variables in a `.env` file in the root of your project. Here's how to do it:

1. **Create a `.env` file:**
   - Copy the `.env.example` file (if it exists) to create your `.env` file, or create a new `.env` file manually.
   - You can use the following command:
     ```bash
     cp .env.example .env
     ```

2. **Edit the `.env` file:**
   Open the `.env` file in a text editor and configure the following variables according to your setup:

   ```dotenv
    APP_NAME="AniManager"
    APP_PORT=8888
    APP_ENV=local

   # PostgreSQL Database Configuration
   DB_USER=your_username        # Your PostgreSQL username
   DB_PASSWORD=your_password    # Your PostgreSQL password
   DB_HOST=localhost             # Database host (usually localhost)
   DB_PORT=5432                  # Default PostgreSQL port
   DB_NAME=your_database_name    # Name of your database
   DB_URL="postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable"

   # Bluesky
   # add your bluesky handle and app password from your account
   BLUESKY_HOST="https://bsky.social"
   BLUESKY_HANDLE="HandleGoesHere"
   BLUESKY_APP_PASSWORD="PasswordGoesHere"
   ```

## Usage 
 **Usage:**  
   After starting the application, it will automatically import the airing schedule from AniList and post updates to BlueSky every minute based on the scheduled anime.

## Commands

For detailed information about the Makefile commands, refer to the [Makefile Commands](MAKEFILE_COMMANDS.md).

**Stopping the application:**  
   To stop the application, you can simply interrupt the process by pressing `Ctrl + C` in the terminal where it's running.
