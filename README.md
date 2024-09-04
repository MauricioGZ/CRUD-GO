# CRUD-GO
This project shows a structured implementation of a Resfull-API using Golang as programming language and the framework Echo.

The goal is to implement a common ecomerce API, which the following databse structure.

<img src="misc/database.png"
     alt="Markdown Monster icon" />

# Architecture
The used architecture for the API is an layered architecture based on the MVC architecture, where the API interface to the user, the bussiness logic, the interaction with the database and the connection to the database are divided. 
Using the concept of dependency injection, it is possible to isolate the different layers of the whole program.

<img src="misc/architecture.png"
     alt="Markdown Monster icon"/>

# Endpoints
For the moment the following endpoints are implemented.
Note: the database contains just one user with the role admin (this user is predefined in the schema.sql).
email: admin@mail.com
password: adminpassword

## User endpoints
- Register a new user: POST HOST:PORT/user/register Request Body: {"first_name":"some first name", "last_name":"some last name", "email": "somevalid@email.com", "password":"somepassword"}
- User login: GET HOST:PORT/user/login Request Body: {"email": "somevalid@email.com", "password":"somepassword"}
## Address endpoints
- Register a new address: POST HOST:PORT/user/address/ Request Body: {"address_type":"billing" || "shipping", "address":"some address", "city":"some city", "state":"some state", "country": "some country", "zip_code":"some zip code"}. Note: This end point works only if an user has loged in.
- Get all addresses of an user; Get HOST:PORT/user/address Note: This end point works only if an user has loged in.
- Update an address: PATCH HOST:PORT/user/address/address_id Request Body: is the same from the endpoint "register a new address", but only the field, that shall be updated, has to be specified in the request body. Note: This end point works only if an user has loged in.
- Delete an address: DELETE HOST:PORT/user/address/address_id Request Body: empty. Note: This end point works only if an user has loged in.
## Categories endpoints
- Register a new category: POST HOST:PORT/categories Request Body: {"name":"some category name", "description":"some description", "parent_id": id of the parente category, if the field is empty means that the category does not have a parent category}. Note: only users with the admin role may add new categories.
- Get all categoires: GET HOST:PORT/categorie
## Products endpoints
- Register a new product: POST HOST:PORT/products Request Body: {"name":"product name", "description":"product description", "price": product price, "stock": product stock, "category_id": id of the category to which the product belongs, "image":"image url"}. Note: only users with the admin or seller role may add new products.
- Get all products: GET HOST:PORT/products.
- Get product by product_id: GET HOST:PORT/products/product_id.
- Update a product: PATCH HOST:PORT/products/product_id Request Body: is the same from the endpoint "register a new product", but only the field, that shall be updated, has to be specified in the request body. Note: only users with the admin or seller role may update a product.
## Orders endpoints
- Register a new order: POST HOST:PORT/order Request Body: [{"porduct_id":product_id, "quantity":quantity, "price":price}, ..., {"porduct_id":product_id, "quantity":quantity, "price":price}]. Note: This end point works only if an user has loged in.
- Get Orders: GET HOST:PORT/order Note: This end point works only if an user has loged in.