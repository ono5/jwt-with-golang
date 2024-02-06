# jwt-with-golang

## Usage

### Protected

```
GET http://localhost:8000/protected

Headers
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxMjAwMDRAZXhhbXBsZS5jb20iLCJpc3MiOiJjb3Vyc2UifQ.OqVWrZKgydTEds4InprvYNYggC-RZIIfq7fBSKv7qiE
```

LoginするとTokenが取れる

### Login

```
POST http://localhost:8000/login

Content-Type: application/json
{
  "email": "test120004@example.com",
  "password": "12345"
}
```

### SignUp

```
POST http://localhost:8000/signup

Content-Type: application/json
{
  "id": 6,
  "email": "test120004@example.com",
  "password": "12345"
}
```
