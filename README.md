# Bookstore Microservice architecture
## Integrated SQL & NoSQL databases, search engines (ElasticSearch)

## Microservice Breakdown:

1. ### bookstore_items-api [link: https://github.com/sharkx018/bookstore_items-api]
 * Developed REST APIs for creating and getting the items
 * Integrated the ElasticSearch 
 * Added the security layer

2. ### bookstore_users-api [link: https://github.com/sharkx018/bookstore_users-api/tree/main]
 * Users API
 * Integrated SQL database
 * Developed REST APIs to perform CRUD operation on users

3. ### bookstore_oauth-api [link: https://github.com/sharkx018/bookstore_oauth-api]
 * Service to generate and validate the access token
 * Integrated the Cassandra DB

4. ### bookstore_oauth-go [link: https://github.com/sharkx018/bookstore_oauth-go/tree/master]
 * Auth Library to authenticate the request
 * Calls the bookstore_oauth-api internally

5. ### bookstore_utils-go [link: https://github.com/sharkx018/bookstore_utils-go]
 * Go utils shared across our entire micro services
 * Added logger and improved error library

# bookstore_users-api
Users API
