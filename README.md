# mydriveuploader

## Overview

MyDriveUploader is a powerful command-line interface (CLI) tool that streamlines the process of uploading files and folders to Google Drive. It offers seamless integration with Google Drive, enabling users to manage their file uploads with ease and efficiency.

## Features

- **Upload Folders**: Effortlessly traverse and upload entire folders to Google Drive, preserving the folder structure.
- **OAuth2 Authentication**: Securely authenticate with Google Drive using OAuth2, ensuring your data remains protected.
- **Persistent Token Storage**: Automatically save and reuse authentication tokens for future sessions, eliminating the need for repeated logins.
- **Cobra CLI Framework**: Built using the robust Cobra framework, providing intuitive command management and extensibility.

## Installation

To install MyDriveUploader, follow these steps:

1. **Clone the Repository**: Open your terminal and run the following command to clone the repository:
   ```bash
   git clone https://github.com/HaseemKhattak01/mydriveuploader
   ```
2. **Navigate to the Project Directory**: Change into the project directory:
   ```bash
   cd mydriveuploader
   ```
3. **Build the Executable**: Compile the application by running:
   ```bash
   go build
   ```

## Setting Up Google Drive Credentials

To use MyDriveUploader, you must configure your Google Drive API credentials. Follow these detailed steps:

1. **Access Google Cloud Console**: Visit the [Google Cloud Console](https://console.cloud.google.com/).
2. **Create or Select a Project**: Either create a new project or select an existing one from your dashboard.
3. **Enable Google Drive API**: Navigate to the "APIs & Services" section, search for "Google Drive API," and enable it for your project.
4. **Create OAuth 2.0 Credentials**: Go to the "Credentials" tab, click "Create Credentials," and select "OAuth client ID." Follow the prompts to configure the consent screen and create your credentials.
5. **Download `credentials.json`**: Once created, download the `credentials.json` file and save it to the root directory of your MyDriveUploader project.

**Important**: Ensure that `credentials.json` is not included in your version control system, as it contains sensitive information that should remain private.


## Setting Up Dropbox Credentials

In addition to Google Drive, MyDriveUploader also supports uploading files to Dropbox. To use this feature, you need to configure your Dropbox API credentials. Follow these steps:

1. **Access Dropbox App Console**: Visit the [Dropbox App Console](https://www.dropbox.com/developers/apps).
2. **Create a New App**: Click on "Create App" and choose the type of access you need (either "App folder" or "Full Dropbox").
3. **Generate Access Token**: Once your app is created, navigate to the "OAuth 2" section and generate an access token. This token will allow MyDriveUploader to upload files to your Dropbox account.
4. **Set Environment Variable**: Store the generated access token in an environment variable named `DROPBOX_TOKEN`. This can be done by adding the following line to your `.env` file:
   ```
   DROPBOX_TOKEN="your_generated_access_token"
   ```

**Important**: Ensure that your `.env` file is not included in your version control system, as it contains sensitive information that should remain private.

## Features

- **Upload to Dropbox**: Seamlessly upload files and folders from your local machine to Dropbox, maintaining the directory structure.
- **OAuth2 Authentication for Dropbox**: Securely authenticate with Dropbox using OAuth2, ensuring your data remains protected.

With these configurations, you can now use MyDriveUploader to manage file uploads to both Google Drive and Dropbox efficiently.

