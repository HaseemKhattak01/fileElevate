# MyDriveUploader

## Overview

MyDriveUploader is a command-line interface (CLI) tool designed to simplify the process of uploading files and folders to Google Drive. It provides a seamless integration with Google Drive, allowing users to manage their file uploads efficiently.

## Features

- **Upload Folders**: Traverse and upload entire folders to Google Drive.
- **OAuth2 Authentication**: Securely authenticate with Google Drive using OAuth2.
- **Persistent Token Storage**: Save and reuse authentication tokens for future sessions.
- **Cobra CLI Framework**: Built using the Cobra framework for easy command management.

## Installation

To install MyDriveUploader, clone the repository and build the executable:

   ## Setting Up Google Drive Credentials

   To use this application, you need to set up your Google Drive API credentials:

   1. Go to the [Google Cloud Console](https://console.cloud.google.com/).
   2. Create a new project or select an existing one.
   3. Enable the Google Drive API for your project.
   4. Create OAuth 2.0 credentials and download the `credentials.json` file.
   5. Place the `credentials.json` file in the root directory of this project.

   **Note**: Do not commit `credentials.json` to version control. It contains sensitive information.