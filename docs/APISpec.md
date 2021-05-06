# API Spec

## Create User

Request 🔥

- Method : POST
- Endpoint : `/api/v1/users`
- Header :
  - Accept : application/json
- Params : 
    - name
    - occupation
    - email
    - password
- Body :

```json
{
  "name": "string",
  "email": "string",
  "occupation": "string",
  "password": "string"
}
```

Response 🚀

```json
{
  "meta": {
    "message": "string",
    "code": "int",
    "status": "string"
  },
  "data": {
    "ID": 1,
    "Name": "Postman",
    "Occupation": "Superman",
    "Email": "",
    "PasswordHash": "$2a$04$NmvjRdE1SoGpZAXyQ7uipOpBnPJgLQw/cCxq9pXby2XLPI6L6SCCi",
    "AvatarFileName": "",
    "Role": "user",
    "Token": "",
    "CreatedAt": "2021-05-05T20:52:22.672+07:00",
    "UpdatedAt": "2021-05-05T20:52:22.672+07:00"
  }
}
```
