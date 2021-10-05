# go-auth-server

Simple authentication server build with golang

> Improved security via JWT-based session tokens that can only be generated using authorized service accounts

## Methods

| API | Method            | Description         |
| ----------- | --------------- | --------- |
| /api/register     | POST          |Create a new email and password user by issuing an HTTP `POST` request|
| /api/login      | POST |Sign in a user with an email and password by issuing an HTTP `POST` request|
| /api/logout  | POST           |Logout current user by issuing an HTTP `POST` request|
| /api/user  | GET           |Validate and return the user info|



#### `/api/register`

```javascript
fetch("http://localhost:3000/api/register", {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
  },
  body: JSON.stringify({
    name,
    email,
    password,
  }),
});
```



#### `/api/login`

```javascript
fetch("http://localhost:3000/api/login", {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
  },
  credentials: "include",
  body: JSON.stringify({
    email,
    password,
  }),
});
```



#### `/api/logout`

```javascript
fetch("http://localhost:3000/api/logout", {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
  },
  credentials: "include",
});
```



#### `/api/logout`

```javascript
fetch("http://localhost:3000/api/user", {
  method: "GET",
  headers: {
    "Content-Type": "application/json",
  },
  credentials: "include",
});
```

