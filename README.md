# antinoLabsAssignment

Prerequisite:
1. Username, Password and Database name for respective sql server must be configured from config.json file
2. Respective database should have a table named `users_blog` with columns having properties:
   1. blogId int PRIMARY KEY auto_increment
   2. emailId varchar(255) NOT NULL
   3. blog varchar(255) NOT NULL
   4. postedOn timestamp DEFAULT CURRENT_TIMESTAMP


------------- Blogging App
This app features with the CRUD operations for blogging platform. Please, check the POSTMAN curl below for following operations:

1. To post a blog:
   curl --location 'localhost:8080/blog/post' \
   --header 'Content-Type: application/json' \
   --data-raw '{
   "EmailId": "iamnayan90@gmail.com",
   "Blog": "Welcome in gvdfgvdzvcxfd!!!"
   }'

2. To retrieve a blog by its blogId:
   curl --location 'localhost:8080/blog/getById?blogId=2'

3. To fetch all the blogs present in the table:
   curl --location 'localhost:8080/blog/getAll'

4. To update a blog by its blogId:
   curl --location --request PUT 'localhost:8080/blog/updateById' \
   --header 'Content-Type: application/json' \
   --data '{
   "BlogId":1,
   "Blog": "Welcome to USA!!"
   }'

5. To delete a blog by its blogId:
   curl --location --request DELETE 'localhost:8080/blog/deleteById?blogId=23'
