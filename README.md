# GitHub Profile Viewer

A web application that displays GitHub user profiles with detailed information, repositories, and visitor tracking.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Installation](#installation)
- [Documentation](#documentation)
- [GitHub OAuth Integration](#github-oauth-integration)

## Overview

GitHub Profile Viewer is a web application that allows users to view GitHub profiles, including: 
- Profile Picture
- Full Name
- Email
- Number of followers and following
- The first 6 repositories
- Login using their GitHub account via OAuth 

The application also tracks the total number of visitors and displays the 3 most recent visitors.


By default, if a user is not logged in, the page will redirect to `/octocat`.

## Features

- Display GitHub user profiles
- OAuth login using GitHub account
- Show profile picture, full name, email, followers, and following
- Display the first 6 repositories of a user
- Total number of visitors tracking
- Show the 3 most recent visitors
- Responsive UI based on Figma design
- Redirect to `/octocat` for non-logged-in users

## Tech Stack

- Frontend:
  - Next.js 13.0.6
  - Tailwind CSS
  - React Context
  - Custom Hooks

- Backend:
  - Golang 1.21
  - Router-chi for routing
  - SQLite3 as the database
  - GitHub OAuth integration

## Installation

### Backend
1. Clone the repository:
   ```bash
   git clone https://github.com/B0urG3ois/pixel8labs-test.git

2. Install Golang 1.21:
   ```bash
   https://go.dev/doc/install

3. Run application:
   ```bash
   cd /backend
   go run cmd/app/main.go

4. Application Start:
   ```bash
   Application start on port 8000
   http://localhost:8000/api/*


### Frontend
1. Clone the repository:
   ```bash
   git clone https://github.com/B0urG3ois/pixel8labs-test.git

2. Install Node.js and NPM:
   ```bash
   https://docs.npmjs.com/downloading-and-installing-node-js-and-npm

3. Run application:
   ```bash
   cd /frontend
   npm install
   npm run dev

4. Application Start:
   ```bash
   Application start on port 3000, and open your browser
   http://localhost:3000


## Documentation

- Postman Collection
  ```bash
  https://api.postman.com/collections/12053329-8d8ac234-9f8d-4305-a8b7-3566f438328d?access_key=PMAT-01HN79YPPTZBJ7T8FS6BCBS3WY

- Endpoints Collection
  - `/api/login`: Initiate the OAuth flow to link with a GitHub account.
  - `/api/logout`: Log the user out.
  - `/api/callback`: Manage code exchange against the front-callback layer's code.
  - `/api/v1/user`: Retrieve the user information details.
  - `/api/v1/repositories`: Retrieve the user's six most recent repositories.

- Figma Design
  ```bash
  https://www.figma.com/file/fLiLQfjSF6X7pEfHli2Lwh/Fullstack-Engineer-Test-Case?type=design&node-id=0%3A1&mode=design&t=RfULQB2MF956TxTT-1
  
- Backend Public URL
  ```bash
  https://pixel8labs-test-be31bc5e9b1e.herokuapp.com
  
- Frontend Public URL
  ```bash
  https://pixel8labs-assignment.vercel.app/


## GitHub OAuth Integration

Go-based wrappers abstract all interactions with the GitHub API in the backend. Efficient data retrieval from the backend by the frontend eliminates the requirement for additional client libraries.

----------------------------------------------------------------

The documentation aims to offer a comprehensive insight into the GitHub OAuth application, encompassing its components and instructions for local setup and execution. For a more in-depth understanding of the code structure and implementation, consult the source code and comments within the relevant repository files.
