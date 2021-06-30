# API Spec

## Create New User (Register)

Request 🔥

- Method : POST
- Endpoint : `/api/v1/users`
- Header :
  - Accept : application/json
- Params : None
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
- Params : None
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
- Params : None
- Body : form-data

```form-data
{
  'avatar': file
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
- Params : None
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
{}
```

Response All 🚀

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

Response All by user_id 🚀
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
      "id": 6,
      "name": "Cloviel Campaign",
      "short_description": "Vivamus magna justo.",
      "image_url": "",
      "goal_amount": 7000000,
      "current_amount": 0,
      "slug": "cloviel-campaign-1",
      "user_id": 1
    },
    {
      "id": 8,
      "name": "Campaign Keren",
      "short_description": "Vestibulum ac diam sit amet quam vehicula elementum sed sit amet dui.",
      "image_url": "",
      "goal_amount": 40000000,
      "current_amount": 0,
      "slug": "campaign-keren-1",
      "user_id": 1
    }
  ]
}
```

\
\</>

## Get Campaign Detail

Request 🔥

- Method : GET
- Endpoint : `/api/v1/campaigns/{id}`
- Header :
  - Accept : application/json
- Params :
  - id
- Body :

```json
{}
```

Response 🚀

```json
{
  "meta": {
    "message": "Successfuly get campaign detail",
    "code": 200,
    "status": "success"
  },
  "data": {
    "id": 1,
    "name": "Moyu Invesment",
    "short_description": "blablabla",
    "image_url": "campaign-images/moyu.png",
    "goal_amount": 1000000,
    "current_amount": 20000,
    "user_id": 5,
    "description": "Blabla albalb albdlbfiabf",
    "slug": "campaign-satu",
    "user": {
      "name": "Arini Qisty Adilla",
      "avatar_url": "images/arini.png"
    },
    "perks": ["Lorem ipsum dan dadn", "Lorem iand dahjbdaf"],
    "images": [
      {
        "image_url": "campaign-images/image.png",
        "is_primary": true
      },
      {
        "image_url": "campaign-images/img.png",
        "is_primary": false
      }
    ]
  }
}
```

\
\</>

## Create New Campaign

Request 🔥

- Method : POST
- Endpoint : `/api/v1/campaigns`
- Header :
  - Accept : application/json
  - Authorization : Bearer tokeneotke
- Params : None
- Body :

```json
{
    "name": "Campaign Keren",
    "short_description": "Vestibulum ac diam sit",
    "description": "Curabitur arcu erat, accumsan id...",
    "goal_amount": 40000000,
    "perks": "Quisque velit nisi, pretium ut lacinia in"
}
```

Response 🚀

```json
{
    "meta": {
        "message": "Successfuly to create new campaign detail",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 9,
        "name": "Baru Campaign Keren",
        "short_description": "Vestibulum ac diam sit amet quam vehicula elementum sed sit amet dui.",
        "image_url": "",
        "goal_amount": 455000000,
        "current_amount": 0,
        "slug": "baru-campaign-keren-1",
        "user_id": 1
    }
}
```

\
\</>

## Update Campaign

Request 🔥

- Method : PUT
- Endpoint : `/api/v1/campaigns/{id}`
- Header :
  - Accept : application/json
  - Authorization : Bearer tokeneotke
- Params :
  - id
- Body :

```json
{
    "name": "Campaign Keren",
    "short_description": "Vestibulum ac diam sit",
    "description": "Curabitur arcu erat, accumsan id...",
    "goal_amount": 40000000,
    "perks": "Quisque velit nisi, pretium ut lacinia in"
}
```

Response 🚀

```json
{
    "meta": {
        "message": "Successfuly to updated campaign",
        "code": 200,
        "status": "success"
    },
    "data": {
        "ID": 1,
        "UserID": 1,
        "Name": "No HACK Campaign Keren Udpdate",
        "ShortDescription": "Update Vestibulum ac diam sit amet quam vehicula elementum sed sit amet dui.",
        "Description": "Curabitur arcu erat, accumsan id imperdiet et, porttitor at sem. Pellentesque in ipsum id orci porta dapibus. Vestibulum ac diam sit amet quam vehicula elementum sed sit amet dui.",
        "GoalAmount": 55000000,
        "CurrentAmount": 20000,
        "BackerCount": 50,
        "Perks": "Quisque velit nisi, pretium ut lacinia in, elementum id enim",
        "Slug": "moyu",
        "CreatedAt": "2021-06-10T15:38:01+07:00",
        "UpdatedAt": "2021-06-29T20:57:49.674+07:00",
        "CampaignImages": [
            {
                "ID": 1,
                "CampaignID": 1,
                "FileName": "campaign-images/moyu.png",
                "IsPrimary": 1,
                "CreatedAt": "2021-06-10T17:01:51+07:00",
                "UpdatedAt": "2021-06-10T17:01:51+07:00"
            },
            {
                "ID": 2,
                "CampaignID": 1,
                "FileName": "campaign-images/moyu-1.png",
                "IsPrimary": 0,
                "CreatedAt": "2021-06-10T17:02:32+07:00",
                "UpdatedAt": "2021-06-10T17:02:32+07:00"
            }
        ],
        "User": {
            "ID": 1,
            "Name": "Arini QA",
            "Occupation": "Bussines Analysis",
            "Email": "arini@gmail.com",
            "PasswordHash": "$2a$04$rFH5vvg1wYuvwfHWndTBNe42CaareIqwJ4locKVYnhuOUajeCokCy",
            "AvatarFileName": "images/1-snf.png",
            "Role": "user",
            "Token": "",
            "CreatedAt": "2021-05-05T17:59:01+07:00",
            "UpdatedAt": "2021-06-10T13:39:07+07:00"
        }
    }
}
```

\
\</>

## Upload Campaign Images

Request 🔥

- Method : POST
- Endpoint : `/api/v1/campaign-images`
- Header :
  - Accept : application/json
  - Authorization : Bearer tokeneotke
- Params :
  - file
  - campaign_id
  - is_primary
- Body :

```form-data
{
    "image": file
}
```

Response 🚀

```json
{
  "meta": {
    "message": "Campaign image successfuly uploaded",
    "code": 200,
    "status": "success"
  },
  "data": {
    "is_uploaded": true
  }
}
```
