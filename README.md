# bookstore_users-api
Users API

Bookstore Microservice architecture
Bookstore Microservice architecture
Designed and developed the bookstore microservice architecture.
Integrated SQL & NoSQL databases, search engines (ElasticSearch)

Microservice Breakdown:

1. bookstore_items-api [link: https://github.com/sharkx018/bookstore_items-api]
 a. Developed REST APIs for creating and getting the items
 b. Integrated the ElasticSearch 
 c. Added the security layer


2. bookstore_users-api [link: https://github.com/sharkx018/bookstore_users-api/tree/main]
 a. Users API
 b. Integrated SQL database
 c. Developed REST APIs to perform CRUD operation on users


3. bookstore_oauth-api [link: https://github.com/sharkx018/bookstore_oauth-api]
 a. Service to generate and validate the access token
 b. Integrated the Cassandra DB


4. bookstore_oauth-go [link: https://github.com/sharkx018/bookstore_oauth-go/tree/master]
 a. Auth Library to authenticate the request
 b. Calls the bookstore_oauth-api internally


5. bookstore_utils-go [link: https://github.com/sharkx018/bookstore_utils-go]
 a. Go utils shared across our entire micro services
 b. Added logger and improved error library
