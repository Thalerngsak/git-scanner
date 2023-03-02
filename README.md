# Getting Started

Install docker and run command
``docker-compose up`` to start docker and access ``http://localhost:8081/`` for Mongo Express database management dashboard  

Run command ``go run main.go`` to start application  

# API List

1. `POST /repositories`: creates a new repository  
How to test  
endpoint `http://localhost:8080/repositories'`  
request body example  
``{
   "name": "test name 3",
   "url": "https://github.com/Thalerngsak/apitest.git"
   } ``  


2. `GET /repositories`: retrieves all repositories  
   How to test  
   endpoint `http://localhost:8080/repositories'`  


3. `GET /repositories/:id`: retrieves a single repository by ID  
   How to test  
   endpoint example `http://localhost:8080/repositories/af1a94c4-8882-424f-9fb4-b3976775a4eb'`


3. `PUT /repositories/:id`: updates a repository by ID  
   How to test  
   endpoint example `http://localhost:8080/repositories/af1a94c4-8882-424f-9fb4-b3976775a4eb'`
   request body example  
   ``{
     "name": "test name 3",
     "url": "https://github.com/Thalerngsak/apitest.git"
   } ``



4. `DELETE /repositories/:id`: deletes a repository by ID  
   How to test  
   endpoint example `http://localhost:8080/repositories/af1a94c4-8882-424f-9fb4-b3976775a4eb'`  


5. `POST /scan-results`: scans a repository and creates a new scan result  
   How to test  
   endpoint example `http://localhost:8080/repositories/af1a94c4-8882-424f-9fb4-b3976775a4eb'`  
   request body example  
   ``{
      "id":"7ef1922f-7fd7-457f-bd20-377105dc95ba"
     }``

6. `GET /scan-results`: retrieves all scan results  
   How to test  
   endpoint example `http://localhost:8080/repositories/af1a94c4-8882-424f-9fb4-b3976775a4eb'`  

5. `GET /results/repository/:id`: retrieves a single scan result by ID  
   How to test  
   endpoint example `http://localhost:8080/repositories/af1a94c4-8882-424f-9fb4-b3976775a4eb'`  

# High-Level Design  
The application is designed using the Hexagonal Architecture pattern, which provides a clean separation between the business logic and the infrastructure code.
![alt text](images/hexagonal.png "title")

The database schema consists of two collections: "repositories" and "scan_results". The "repositories" collection stores the Repository entities, while the "scan_results" collection stores the ScanResult entities.

Overall, this design provides a clean separation between the different components of the application, making it easy to test, maintain, and extend.  
![alt text](images/component.png "title")

# Stack
1. Gin as the web framework  
2. MongoDB as the database for storing repository and scan result data  