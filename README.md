# Code Sharing App Backend
This is the backend of a simple code-sharing application called NoteCode. This application will allow users to store and share coding snippets.  

Below are the requirements of the application from the challenge:
- Create a Coding Sharing App - NoteCode following given design.
- By default, users should see an HTML given snippet. See Code Guide for default HTML code.
- When users select the Share button, a new ID should be generated, and users can access the saved code with the generated ID. See Code Guide for more details.
- After code is saved and shared, Share button should be disabled until users make an edit .
- Users can choose the language and theme they want to save and share.
- The application should be responsive on all devices.
- User authentication is not required
- Deploy the solution and submit Repository URL and Demo URL.

## API Endpoints
1. POST /api/snippets: Save a new code snippet and return the generated ID.
2. GET /api/snippets/:id: Retrieve a code snippet by its ID.
3. PATCH /api/snippets/:id: Update the code snippet by its ID. 

## Live Demo
Live demo aren't available because it is not possible to deploy an API application for free.

## Environment variables
There are environment variables that is required to be provided in the .env file:
- DB_URL: the MongoDB database url.
- DB_NAME: the MongoDB database name.
- ALLOWED_ORIGINS: the allowed origins for the application that make requests to the API.
- ROUTER_PORT: the router port where the api will run.

## Installation and usage
```
git clone https://github.com/jordyf15/code-sharing-app-backend.git
cd code-sharing-app-backend
go get
go run .
```