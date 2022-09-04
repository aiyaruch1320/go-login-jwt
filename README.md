# GO Login & JWT
## Features
- Create user with shell script
- JWT
- Authentication username and password
- Data into mongodb
- Docker compose

## Tech
- <a href="https://go.dev/dl/"> Go </a>
- <a href="https://www.docker.com/products/docker-desktop/"> Docker </a>
- <a href="https://www.postman.com/downloads/"> Postman </a>
- Code Editor

## Try
- before start you should install the program from above
- then use command -> ```git clone git@github.com:aiyaruch1320/go-login-jwt.git```
- run docker and use command in terminal -> ```docker compose up -d```
- use Postman test api


#### Health
```
GET -> http://localhost:8000/health
```
Response should be return:
```
OK
```
#### Create user
```
POST -> http://localhost:8000/api/register
```
Body:
```
{
    "username": "admin",
    "password": "12345678",
    "role": "admin" or "user"
}
```
Response when username already exist:
```
{
    "message": "username already exist"
}
```
Response when create user success:
```
{
    "message": "created user"
}
```
Response when role wrong:
```
{
    "message": "role must be admin or user"
}
```
#### Login
```
POST -> http://localhost:8000/login
```
Body:
```
{
    "username": "admin",
    "password": "12345678",
}
```
Response should be return when success:
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjYzMDg1OTNkYTNmYzJhZjc3ZTQzNGRkYSIsInVzZXJuYW1lIjoidXNlciIsInJvbGUiOiJ1c2VyIiwiY2xpZW50X2lkIjoiIiwic3ViIjoiIiwiYXV0aF90aW1lIjowLCJsZHAiOiIiLCJpYXQiOjAsInNjb3BlIjpudWxsLCJhbXIiOm51bGx9.b3PFwJGcP2GJXcwXOdotZqXwIWb-Mi13nC1fw3ysVwM"
}
```
Response when username or password wrong:
```
{
    "message": "invalid username or password"
}
```

#### Try User
```
GET -> http://localhost:8000/api/user
```
Response should be return when success:
```
admin&user can see
```
Response when has no _JWT_ :
```
{
    "message": "invalid or expired jwt"
}
```

#### Try Admin
```
GET -> http://localhost:8000/api/admin
```
Response should be return when success:
```
only admin can see
```
Response when _Role_ not admin:
```
{
    "message": "you are not authorized to access this resource"
}
```
Response when has no _JWT_ :
```
{
    "message": "invalid or expired jwt"
}
```
