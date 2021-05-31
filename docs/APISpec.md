# API Spec

## Create New User (Register)

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
    "message": "Account has been registered",
    "code": 200,
    "status": "success"
  },
  "data": {
    "id": 5,
    "name": "Arini Qisty Adila",
    "occuputaion": "Bussiness Syariah",
    "email": "ariniqa@gmail.com",
    "token": "abc"
  }
}
```
