#Subject Registration Portal

Overview
Subject Registration Portal is a simple HTTP REST API built using the GoFr framework. It provides an interface for managing student subject registration requests, including adding, viewing, updating, and deleting requests. This application demonstrates CRUD operations with a MySQL database backend.

Features
- Add Registration: Students can make requests for enrolling in subjects.
- View Registrations: View a list of students waiting for approval.
- Update Registration: Update the status of a registration request (approve or reject).
- Delete Registration: Remove a request from the database once it is approved.

Prerequisites
- Go 
- Docker
- Postman for API testing

Testing the API with Postman
1. Add Registration
Method: POST
URL: http://localhost:8000/registrations
Body: {"name": "John Doe", "rollNo": 123, "subjectName": "Mathematics"}

2. View Registrations
Method: GET
URL: http://localhost:8000/registrations

3. Update Registration
Method: PUT
URL: http://localhost:8000/registrations/{id}
Body: {"status": "approved"}

4. Delete Registration
Method: DELETE
URL: http://localhost:8000/registrations/status/approved
