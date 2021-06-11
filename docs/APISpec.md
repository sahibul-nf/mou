# API Spec

## Create New User (Register)

Request 🔥

- Method : POST
- Endpoint : `/api/v1/users`
- Header :
  - Accept : application/json
- Params : 
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
    "token": "blabablabla"
  }
}
```

\
\</>

## User Sessions (Login)

Request 🔥

- Method : POST
- Endpoint : `/api/v1/sessions`
- Header :
  - Accept : application/json
- Params : 
- Body :

```json
{
  "email": "string",
  "password": "string"
}
```

Response 🚀

```json
{
  "meta": {
    "message": "Successfuly loggedin",
    "code": 200,
    "status": "success"
  },
  "data": {
    "id": 4,
    "name": "Arini Qisty Adila",
    "occuputaion": "Bussiness Syariah",
    "email": "ariniqa@gmail.com",
    "token": "blablabal"
  }
}
```

\
\</>

## User Upload Avatar

Request 🔥

- Method : POST
- Endpoint : `/api/v1/avatars`
- Header :
  - Accept : application/json
  - Authorization : Bearer tokentoken
- Params : 
- Body : FORM

```json
{

}
```

Response 🚀

```json
{
  "meta": {
    "message": "Avatar successfuly uploaded",
    "code": 200,
    "status": "success"
  },
  "data": {
    "is_uploaded": true
  }
}
```

\
\</>

## User Email Checkers

Request 🔥

- Method : POST
- Endpoint : `/api/v1/avatars`
- Header :
  - Accept : application/json
  - Authorization : Bearer tokentoken
- Params : 
- Body : 

```json
{
  "email": "ariniqaa@gmail.com"
}
```

Response 🚀

```json
{
  "meta": {
    "message": "Email is available",
    "code": 200,
    "status": "success"
  },
  "data": {
    "is_available": true
  }
}
```

\
\</>

## Get List of Campaigns

Request 🔥

- Method : GET
- Endpoint : `/api/v1/campaigns`
- Header :
  - Accept : application/json
- Params : 
  - user_id
- Body : 

```json
{

}
```

Response 🚀

```json
{
  "meta": {
    "message": "Successfuly get list of campaigns",
    "code": 200,
    "status": "success"
  },
  "data": [
    {
      "id": 1,
      "name": "Moyu Invesment",
      "short_description": "blablabla",
      "image_url": "campaign-images/moyu.png",
      "goal_amount": 1000000,
      "current_amount": 20000,
      "slug": "moyu",
      "user_id": 1
    },
    {
      "id": 2,
      "name": "Cloviel",
      "short_description": "balbal",
      "image_url": "",
      "goal_amount": 20000000,
      "current_amount": 20000,
      "slug": "cloviel",
      "user_id": 2
    }
  ]
}
```
