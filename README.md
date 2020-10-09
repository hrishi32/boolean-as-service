# Boolean as a service

## About Project
Boolean as a service is an api which stores boolean values along with a key identified by a unique key (UUID). API provides Get, Create, Update and Delete booleans (id, key, value) with http requests. Implementation of this API is in golang.

### Requests
- id should be `uuid`
- value should be either `true` or `false`(boolean, not string)
- key should be `string`
#### POST Request to create a boolean
```
POST /
request:

{
  "value":true,
   "key": "name" // this is optional
}

response:

{
  "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "value": true,
  "key": "name"
}
```

#### GET request to access existing boolean
```
GET /:id
response:

{
  "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "value": true,
  "key": "name"
}
```

#### PATCH request to update the existing boolean
```
PATCH /:id
request:

{
  "value":false,
  "key": "new name" // this is optional
}
response:

{
  "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "value": false,
  "key": "new name"
}
```

#### DELETE request to delete the existing boolean
```
DELETE /:id
response:
HTTP 204 No Content
```

## Installation
### On Linux/Mac

```bash
$ go build main.go
$ ./main
```

### With Docker

```bash
$ docker build -t boolean_as_service
$ docker run -it -p 8000:8000 -e DB_USER='root' -e DB_PASSWORD='m' -e DB_NAME='boolean' -e DOCKER='yes' -e DB_PORT='8084' -e DB_HOST='host.docker.internal' boolean_as_service 
```

## Tests
We have endpoint tests in controller folder. Change to controller, and execute following command

```bash
$ go test
```
Or
```bash
$ go test -cover
```