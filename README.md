<h1 style="font-size:40px;font-weight:900;">Tables Got Your Back <img src="assets/backend.png" height="55px" width="55px"></h1>

[![Test](https://github.com/SirJager/tables/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/SirJager/tables/actions/workflows/test.yml)

Tables allows you to manage your backend database a more easier way .

## [Try Api In Postman](https://postman.com/cloudy-spaceship-448418/workspace/go/documentation/16177687-58151bd4-c7b4-402b-a5d2-c871455d2c5a)

## 🚀 What Next

- Make tests
- Add, Remove Update Primary Keys
- Add, Remove Update Foreign Keys
- Build a front end
- Join Operations

## Tech Stack

**Language:** Go

**Database:** Postgres

**Web framework:** Gin

**Authentication:** JWT, Paseto

# Endpoints

## UnAuthenticated Endpoints

| Title         | Descrption     | Method   |  Endpoint             |
|---------------|----------------|----------|-----------------------|
| Signup        | Create account | POST     | `/users/signup`       |
| Signin        | Login  account | POST     | `/users/signin`       |
| Refresh Tokne | Refresh Token  | POST     | `/users/renew-access` |

## Authenticated Endpoints

### Users Routes

| Title         | Descrption     | Method   |  Endpoint       |
|---------------|----------------|----------|-----------------|
| List Users    | List All Users | GET      | `/users`        |
| Current user  | Get  User      | GET      | `/users/me`     |
| Remove User   | Delete my acc  | DELETE   | `/users/me`     |

### Tables Routes

| Title         | Descrption     | Method   |  Endpoint                |
|---------------|----------------|----------|--------------------------|
| Create Table  | Create   Table | POST     | `/tables`                |
| List My Tables| Tables by me   | GET      | `/tables`                |
| Table Schema  | Table columns  | GET      | `/tables/{tablename}`    |
| Remove Table  | Delete my table| DELETE   | `/tables/{tablename}`    |

### Manage Columns Routes

| Title              | Descrption          | Method   |  Endpoint                         |
|--------------------|---------------------|----------|-----------------------------------|
| Add Column         | Add Column          | POST     | `/tables/{tablename}/columns`     |
| Remove Column      | Delete Column       | DELETE   | `/tables/{tablename}/columns`     |
| Add Primary key    | Add Primary Key     | POST     | `to be implemented`               |
| Delete Primary key | Delete Primary Key  | DELETE   | `to be implemented`               |
| Update Primary key | Update Primary Key  | PATCH    | `to be implemented`               |

### Manage Rows

| Title         | Descrption     | Method   |  Endpoint                      |
|---------------|----------------|----------|--------------------------------|
| Get Rows      | Get rows       | GET      | `/tables/{tablename}/rows`     |
| Insert Rows   | Insert rows    | POST     | `/tables/{tablename}/rows`     |
| Delete Rows   | Delete Rows    | DELETE   | `/tables/{tablename}/rows`     |
| Update Row    | Update Rows    | PATCH    | `to be implemented`

## Examples in  HTTP/1.1

## 1. Create User  `/users/signup`

<table>
  <tr >
    <td>Request</td>
    <td>Response</td>
  </tr>
  <tr>
    <td>

```http
POST /users/signup
Host: localhost:8080 
Content-Type: application/json 
Content-Length: 108 

{ 
  "email": "johndoe@gmail.com", 
  "username": "johndoe", 
  "password": "password", 
  "fullname": "John Doe" 
}
```

  </td>

  <td>

```json
{ 
  "id": 1, 
  "email": "johndoe@gmail.com", 
  "username": "johndoe", 
  "fullname": "John Doe", 
  "public": false, 
  "blocked": false, 
  "verified": false, 
  "updated": "2022-08-01 08:03:12.416867 +0530 IST", 
  "created": "2022-08-01 08:03:12.416867 +0530 IST"
}
```

</td>
</tr>
</table>

## 2. Create User  `/users/signin`

<table>
  <tr >
    <td>Request</td>
    <td>Response</td>
  </tr>
  <tr>
    <td>

```http
POST /users/signin HTTP/1.1
Host: localhost:8080
Authorization: Basic dXNlcjE6cGFzc3dvcmQ=
```

  </td>
  <td>

```json
{
    "access_token": "{access_token}",
    "refresh_token": "{refresh_token}",
    "session_id": "ac38a719-e382-43b4-86c0-150b40046ef4",
    "access_token_expires_at": "2033-12-27T23:45:25.707804663+05:30",
    "refresh_token_expires_at": "2022-08-02T08:45:25.707981887+05:30",
    "user": {
        "id": 1,
        "email": "johndoe@gmail.com",
        "username": "johndoe",
        "fullname": "John Doe",
        "public": false,
        "blocked": false,
        "verified": false,
        "updated": "2022-08-01 08:45:11.33764 +0530 IST",
        "created": "2022-08-01 08:45:11.33764 +0530 IST"
    }
}
```

</td>
</tr>
</table>

## 3. Renew Access Token  `/users/renew-access`

<table>
  <tr >
    <td>Request</td>
    <td>Response</td>
  </tr>
  <tr>
    <td>

```http
POST /users/renew-access HTTP/1.1
Host: localhost:8080
RefreshToken: {refresh_token}
```

  </td>

  <td>

```json
{
    "access_token": "{access_token}",
    "access_token_expires_at": "2033-12-27T23:48:20.095529481+05:30"
}
```

</td>
</tr>
</table>

## 4. Get Users  `/users`

<table>
  <tr >
    <td>Request</td>
    <td>Response</td>
  </tr>
  <tr>
    <td>

```http
GET /users HTTP/1.1
Host: localhost:8080
Authorization: Bearer {access_token}
```

  </td>

  <td>

```json
[
    {
        "id": 1,
        "email": "johndoe@gmail.com",
        "username": "johndoe",
        "fullname": "John Doe",
        "public": false,
        "blocked": false,
        "verified": false,
        "updated": "2022-08-01 08:45:11.33764 +0530 IST",
        "created": "2022-08-01 08:45:11.33764 +0530 IST"
    }
]
```

</td>
</tr>
</table>

## 5. Get My Profile  `/users/me`

<table>
  <tr >
    <td>Request</td>
    <td>Response</td>
  </tr>
  <tr>
    <td>

```http
GET /users HTTP/1.1
Host: localhost:8080
Authorization: Bearer {access_token}
```

  </td>

  <td>

```json
{
    "id": 1,
    "email": "johndoe@gmail.com",
    "username": "johndoe",
    "fullname": "John Doe",
    "public": false,
    "blocked": false,
    "verified": false,
    "updated": "2022-08-01 08:45:11.33764 +0530 IST",
    "created": "2022-08-01 08:45:11.33764 +0530 IST"
}
```

</td>
</tr>
</table>

## 6. Delete My Profile  `/users/me`

(currenlty having issues will be fixed very soon )
<table>
  <tr >
    <td>Request</td>
    <td>Response</td>
  </tr>
  <tr>
    <td>

```http
DELETE /users/me HTTP/1.1
Host: localhost:8080
Authorization: Bearer {access_token}
Content-Type: application/json
Content-Length: 1
```

  </td>

  <td>

```json
{
    "id": 1,
    "email": "johndoe@gmail.com",
    "username": "johndoe",
    "fullname": "John Doe",
    "public": false,
    "blocked": false,
    "verified": false,
    "updated": "2022-08-01 08:45:11.33764 +0530 IST",
    "created": "2022-08-01 08:45:11.33764 +0530 IST"
}
```

</td>
</tr>
</table>

## 7. Create Table  `/tables`

### Column Type

`
integer,smallint,bigint,decimal,numeric,real,'double precision',smallserial,serial,bigserial,varchar,char,character,text,timestamp,'timestamp with time zone','timestamp without time zone' date,'time with time zone',time,'time without time zone',bool,boolean,bit,'bit varying',cidr, inet,macaddr,macaddr8,json,jsonb,money,uuid
`

### Column Fields

name:- any alpha numeric, be short and simple

```
{
  "name": "name",
  "type": "serial",
  "length": 0,
  "primary": true,
  "unique": false,
  "required": false,
  "precision": 0,
  "scale": 0,
  "default": ""
}
```

<table>
  <tr >
    <td>Request</td>
    <td>Response</td>
  </tr>
  <tr>
    <td>

```http
POST /tables HTTP/1.1
Host: localhost:8080
Authorization: Bearer {access_token}
Content-Type: application/json
Content-Length: 388

{
    "table": "articles",
    "columns": [
        {
            "name": "id",
            "type": "serial",
            "primary": true
        },
        {
            "name": "title",
            "type": "varchar",
            "length": 60,
            "required": true
        },
        {
            "name": "author",
            "type": "varchar",
            "length": 60
        },
        {
            "name": "body",
            "type": "varchar",
            "required": true
        }
    ]
}
```

  </td>

  <td>

```json
{
    "id": 1,
    "name": "articles",
    "user_id": 2,
    "columns": [
        {
            "name": "id",
            "type": "serial",
            "length": 0,
            "primary": true,
            "unique": false,
            "required": false,
            "precision": 0,
            "scale": 0,
            "default": ""
        },
        {
            "name": "title",
            "type": "varchar",
            "length": 60,
            "primary": false,
            "unique": false,
            "required": true,
            "precision": 0,
            "scale": 0,
            "default": ""
        },
        {
            "name": "author",
            "type": "varchar",
            "length": 60,
            "primary": false,
            "unique": false,
            "required": false,
            "precision": 0,
            "scale": 0,
            "default": ""
        },
        {
            "name": "body",
            "type": "varchar",
            "length": 0,
            "primary": false,
            "unique": false,
            "required": true,
            "precision": 0,
            "scale": 0,
            "default": ""
        }
    ],
    "created": "2022-08-01T08:59:44.963863+05:30",
    "updated": "2022-08-01T08:59:44.963863+05:30"
}
```

</td>
</tr>
</table>

## 8. Drop Table `/tables/{tablename}`

<table>
  <tr >
    <td>Request</td>
    <td>Response</td>
  </tr>
  <tr>
    <td>

```http
DELETE /tables/articles HTTP/1.1
Host: localhost:8080
Authorization: Bearer {access_token}
```

  </td>

  <td>

```json
{
    "message": "Table 'articles' deleted"
}
```

</td>
</tr>
</table>

## 9. List All My Tables

<table>
  <tr >
    <td>Request</td>
    <td>Response</td>
  </tr>
  <tr>
    <td>

```http
GET /tables HTTP/1.1
Host: localhost:8080
Authorization: Bearer {access_token}
```

  </td>

  <td>

```json
[
    {
        "id": 2,
        "name": "articles",
        "user_id": 2,
        "columns": [
            {
                "name": "id",
                "type": "serial",
                "length": 0,
                "primary": true,
                "unique": false,
                "required": false,
                "precision": 0,
                "scale": 0,
                "default": ""
            },
            {
                "name": "title",
                "type": "varchar",
                "length": 60,
                "primary": false,
                "unique": false,
                "required": true,
                "precision": 0,
                "scale": 0,
                "default": ""
            },
            {
                "name": "author",
                "type": "varchar",
                "length": 60,
                "primary": false,
                "unique": false,
                "required": false,
                "precision": 0,
                "scale": 0,
                "default": ""
            },
            {
                "name": "body",
                "type": "varchar",
                "length": 0,
                "primary": false,
                "unique": false,
                "required": true,
                "precision": 0,
                "scale": 0,
                "default": ""
            }
        ],
        "created": "2022-08-01T09:19:02.488179+05:30",
        "updated": "2022-08-01T09:19:02.488179+05:30"
    }
]
```

</td>
</tr>
</table>

## 10. Get Table

<table>
  <tr >
    <td>Request</td>
    <td>Response</td>
  </tr>
  <tr>
    <td>

```http
GET /tables/articles HTTP/1.1
Host: localhost:8080
Authorization: Bearer {access_token}
```

  </td>

  <td>

```json
{
        "id": 2,
        "name": "articles",
        "user_id": 2,
        "columns": [
            {
                "name": "id",
                "type": "serial",
                "length": 0,
                "primary": true,
                "unique": false,
                "required": false,
                "precision": 0,
                "scale": 0,
                "default": ""
            },
            {
                "name": "title",
                "type": "varchar",
                "length": 60,
                "primary": false,
                "unique": false,
                "required": true,
                "precision": 0,
                "scale": 0,
                "default": ""
            },
            {
                "name": "author",
                "type": "varchar",
                "length": 60,
                "primary": false,
                "unique": false,
                "required": false,
                "precision": 0,
                "scale": 0,
                "default": ""
            },
            {
                "name": "body",
                "type": "varchar",
                "length": 0,
                "primary": false,
                "unique": false,
                "required": true,
                "precision": 0,
                "scale": 0,
                "default": ""
            }
        ],
        "created": "2022-08-01T09:19:02.488179+05:30",
        "updated": "2022-08-01T09:19:02.488179+05:30"
    }
```

</td>
</tr>
</table>

## 11. Add Column  `/tables/{tablename}/columns`

<table>
  <tr >
    <td>Request</td>
    <td>Response</td>
  </tr>
  <tr>
    <td>

```http
POST /tables/articles/columns HTTP/1.1
Host: localhost:8080
Authorization: Bearer {access_token}

{
    "columns": [
        {
            "name": "created",
            "type": "timestamp",
            "default": "now()"
        }
    ]
}
```

  </td>

  <td>

```json
{
    "id": 1,
    "name": "articles",
    "user_id": 2,
    "columns": [
        {
            "name": "id",
            "type": "serial",
            "length": 0,
            "primary": true,
            "unique": false,
            "required": false,
            "precision": 0,
            "scale": 0,
            "default": ""
        },
        {
            "name": "title",
            "type": "varchar",
            "length": 60,
            "primary": false,
            "unique": false,
            "required": true,
            "precision": 0,
            "scale": 0,
            "default": ""
        },
        {
            "name": "author",
            "type": "varchar",
            "length": 60,
            "primary": false,
            "unique": false,
            "required": false,
            "precision": 0,
            "scale": 0,
            "default": ""
        },
        {
            "name": "body",
            "type": "varchar",
            "length": 0,
            "primary": false,
            "unique": false,
            "required": true,
            "precision": 0,
            "scale": 0,
            "default": ""
        },
        {
            "name": "created",
            "type": "timestamp",
            "length": 0,
            "primary": false,
            "unique": false,
            "required": false,
            "precision": 0,
            "scale": 0,
            "default": "now()"
        }
    ],
    "created": "2022-08-01T08:59:44.963863+05:30",
    "updated": "2022-08-01T08:59:44.963863+05:30"
}
```

</td>
</tr>
</table>

## 12. Drop Column  `/tables/{tablename}/columns`

<table>
  <tr >
    <td>Request</td>
    <td>Response</td>
  </tr>
  <tr>
    <td>

```http
DELETE /tables/articles/columns HTTP/1.1
Host: localhost:8080
Authorization: Bearer {access_token}
Content-Type: application/json
Content-Length: 44

{
    "columns": [
        "created"
    ]
}
```

  </td>

  <td>

```json
{
    "id": 1,
    "name": "articles",
    "user_id": 2,
    "columns": [
        {
            "name": "id",
            "type": "serial",
            "length": 0,
            "primary": true,
            "unique": false,
            "required": false,
            "precision": 0,
            "scale": 0,
            "default": ""
        },
        {
            "name": "title",
            "type": "varchar",
            "length": 60,
            "primary": false,
            "unique": false,
            "required": true,
            "precision": 0,
            "scale": 0,
            "default": ""
        },
        {
            "name": "author",
            "type": "varchar",
            "length": 60,
            "primary": false,
            "unique": false,
            "required": false,
            "precision": 0,
            "scale": 0,
            "default": ""
        },
        {
            "name": "body",
            "type": "varchar",
            "length": 0,
            "primary": false,
            "unique": false,
            "required": true,
            "precision": 0,
            "scale": 0,
            "default": ""
        }
    ],
    "created": "2022-08-01T08:59:44.963863+05:30",
    "updated": "2022-08-01T08:59:44.963863+05:30"
}
```

</td>
</tr>
</table>

## 13. Insert Rows In Table  `/tables/{tablename}/rows`

### Single Row Example

```
  [
    {"k":"title", "v" : "Learn Coding"}
    {"k":"author", "v" : "John Doe"}
    {"k":"body", "v" : "article body"}
  ]
```

1. k = column
2. V = value of row for that specific column
3. if the column has required = false that can be skipped
4. you give the column where you want to store data

<table>
  <tr >
    <td>Request</td>
    <td>Response</td>
  </tr>
  <tr>
    <td>

```http
POST /tables/articles/rows HTTP/1.1
Host: localhost:8080
Authorization: Bearer {access_token}
Content-Length: 369

{
  "rows": [
    [
      {"k": "title","v": "master backend"},
      {"k": "author","v": "techschool"},
      {"k": "body","v": "Ultimate backend course"}
    ],[
      {"k": "title","v": "Learn sql"},
      {"k": "author","v": "elsa doe"},
      {"k": "body","v": "sql"}
    ],[
      {"k": "title","v": "java"},
      {"k": "author","v": "john doe"},
      {"k": "body","v": "java"}
    ]
  ]
}
```

  </td>

  <td>

```json
{
    "message": "3 rows inserted in articles"
}
```

</td>
</tr>
</table>

## 13. Get Rows from Table  `/tables/{tablename}/rows`

<table>
  <tr >
    <td>Request</td>
    <td>Response</td>
  </tr>
  <tr>
    <td>

```http
GET /tables/articles/rows HTTP/1.1
Host: localhost:8080
Authorization: Bearer {access_token}
```

  </td>

  <td>

```json
[
    {
        "author": "techschool",
        "body": "Ultimate backend course",
        "id": 1,
        "title": "master backend"
    },
    {
        "author": "elsa doe",
        "body": "sql",
        "id": 2,
        "title": "Learn sql"
    },
    {
        "author": "john doe",
        "body": "java",
        "id": 3,
        "title": "java"
    }
]
```

</td>
</tr>
</table>

## 14. Get Rows from Table  `/tables/{tablename}/rows`

### Same as above(13). But Extended

### Custom Filters

1. To Query/Fetch speicific row you must apply filters

```json
{
  "filters":{
    // values will always be in a list
    "columnName": [values...],
  }
}
```

```json
{
  "filters":{
    "anyTextColumn": ["value1 to match","value2 to match"],
    // boolean column will always have one value / either true or false
    "anyBooleanColumn": [false],
    // 
    "anyIntergerColumn": [123,42,4124],

  }
}
```

## Example

## 1

### Request

```json
{
  "filters":{
    "author":["techschool"] // Can have multiple values
  }
}
```

This is equivalent to :

`SELECT * FROM tablename  WHERE author in ('techschool')`

### Response

```json
[
    {
        "author": "techschool",
        "body": "Ultimate backend course",
        "id": 1,
        "title": "master backend"
    }
]
```

## 2

### Request

```json
{
  "filters":{
    "id":[3,1] // Can have multiple values
  }
}
```

This is equivalent to :

`SELECT * FROM tablename  WHERE id in (1)`

### Response

```json
[
    {
        "author": "techschool",
        "body": "Ultimate backend course",
        "id": 1,
        "title": "master backend"
    },
    {
        "author": "john doe",
        "body": "java",
        "id": 3,
        "title": "java"
    }
]
```

## 3

### Request

```json
{
  "filters":{
    "id":[2],
    "author":["techschool"]
  }
}
```

This is equivalent to :

`SELECT * FROM tablename  WHERE id in (2) OR author in ('techschool')`

### Response

```json
[
    {
        "author": "techschool",
        "body": "Ultimate backend course",
        "id": 1,
        "title": "master backend"
    },
    {
        "author": "elsa doe",
        "body": "sql",
        "id": 2,
        "title": "Learn sql"
    }
]
```

## 4

### Request

```json
{
  "filters":{
    "author":["ellen"],
    "&":{
      "published":[true]
    }
  }
}
```

This is equivalent to :

`SELECT * FROM tablename  WHERE author in ('ellen') AND published = true;`

### Response

```json
[
    {
        "author": "ellen",
        "body": "nextjs article body",
        "id": 5,
        "published": true,
        "title": "Learn Nextjs"
    },
    {
        "author": "ellen",
        "body": "react article body",
        "id": 6,
        "published": true,
        "title": "Learn React"
    }
]
```

## 5

### Request

```json
{
  "fields":["title"],
  "filters":{
    "author":["ellen"],
    "&":{
      "published":[true]
    }
  }
}
```

This is equivalent to :

`SELECT title FROM tablename  WHERE author in ('ellen') AND published = true;`

### Response

```json
[
    {
        "title": "Learn Nextjs"
    },
    {
        "title": "Learn React"
    }
]
```

## 6

### Request

```json
{
    "fields": [
        "title",
        "published"
    ],
    "filters": {
        "author": [
            "ellen"
        ],
        "&": {
            "published": [
                false
            ]
        }
    }
}
```

This is equivalent to :

`SELECT title FROM tablename  WHERE author in ('ellen') AND published = true;`

### Response

```json
[
    {
        "published": false,
        "title": "Learn javascript"
    }
]
```
