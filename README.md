# simple-api-using-go
A Go Lang based API that will be performing operations like Updation, Creation, and Reading the data

This Project is about making a lite API that will be helping in codes.
So, its a library API that will store some books, 
Here we will be able to check in a book, check out a book, update the books, delete the books and get a book by its ID.

We need the gin package
> go get github.com/gin-gonic/gin

Steps: 
1. //Firstly we will be having a struct type for storing our data and then we will be adding few things to get this serialized and converted in JSON, so the API can take JSON version of struct.

2. // using the capital case allows the entities to be viewed as public entities, while in process of serialization we want them to be within scope so, we are converting the caps into lowercase in JSON format, if we don't use the caps everytime we return we get empty objects, because JSON will not be able to read every field name out of this objects

3. We create our own slice that would be acting as our data for now, later on we can connect it with some databases

4. We created a router using gin, that allows us to handle different routes and endpoints. Like oour route is /books, that operates on the localhost:8080, and it will call the function getBooks, that has gin.context, that is all the information about the request generated, and allows to return a response, and it will return a JSON, which is properly indented

// Now we have created a endpoint from where we are able to get the data, now we will be creating a POST method go that we can update or add something to the existing data

5. We are creating another method that would be taking up, the c *gin.context(), and would be initiating up with the query parameters available

> we need to bind the JSON which was the part of data payload of this request to this book object

What we done is that we created a "book" object newBook, and then we are trying to Bind the data of JSON to the new book by passing its pointer, if err then return, otherwise, we are appending the newBook to slice of books, and we are returning in with statusCode
and in main, we get a POST method with the same route

For running the command: 
curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"

// GETTING A BOOK BY ITS ID
 we will be creating a function that would be taking id in strings, and would be returning a book function and error if the book is not found

It iterates over the our data of books and if the data is found out it return the data at that particular ID, otherwise returns with error statuscode, 
and an another function has the parameter that has data of the file, would be calling upon the previous function, if successfully returned, it return the data if not return the erros
path would be specified as
'./books/:id'
and connection would be established by
curl localhost:8080/books/:id

//Checking the book and checking out the book
There will be two function serving each one of the operation, and they will increase the quantity of book when checked in and decrease it when checked out and return error if 0

> Check out function:
	firstly,we will be trying to get the id of the book, and if the ID is not valid, it would be returning a bad request and message
	Secondly, id is found,then we will be calling getBookbyID function to find out the book, if not found then same bad request and message
	Next, we would be checking the quantity if its 0 then error, otherwise quantity-- and returning the status code with book
Route: router.PATCH("/checkout", checkOutBook)
Command: curl localhost:8080/checkout?id=2 --request "PATCH"
