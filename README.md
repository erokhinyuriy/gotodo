# Project "gotodo"

Web API service for ToDo list application.

Author: Yuriy Erokhin

## Docker

```
docker compose -f docker-compose.yml up
```

## Description

To run the service locally, use :8447 but if you run the application in a docker container, you must use :8778.

If you want to execute requests, you must be authorized. To do this, create a user (for example):

POST: http://localhost:8447/signup 

With body request:

```
{
    "username": "admin",
    "email": "admin@admin.com",
    "password": "pass1234$"
}
```

and sign in:

POST: http://localhost:8447/sighin

```
{
    "email": "admin@admin.com",
    "password": "pass1234$"
}
```

API's using Postman:

For LISTS

GetAll      GET:    http://localhost:8447/lists

CreateList  POST:   http://localhost:8447/lists

Body json:
```
{
    "name": "test1",
    "date": null
}
```

Returned id (for example):

```
"86e0dda5-f6ae-461a-8e14-d6a068c38863"
```

GetListById  GET:   http://localhost:8447/lists/86e0dda5-f6ae-461a-8e14-d6a068c38863

Returned result (for example):
```
{
    "Id": "86e0dda5-f6ae-461a-8e14-d6a068c38863",
    "Name": "test1",
    "Date": "0001-01-01T00:00:00Z",
    "Tasks": null
}
```

UpdateList  PUT:    http://localhost:8447/lists 

Body json:
```
{
    "Id": "86e0dda5-f6ae-461a-8e14-d6a068c38863",
    "Name": "test1_upd",
    "Date": "0001-01-01T00:00:00Z"
}
```

Returned result (for example):
```
"list was updated"
```

DeleteList:  DELETE: http://localhost:8447/lists/86e0dda5-f6ae-461a-8e14-d6a068c38863 

```
"list with id: 86e0dda5-f6ae-461a-8e14-d6a068c38863 was deleted"
```

For TASKS

GetAll GET: http://localhost:8447/tasks

GetTaskById  GET:   http://localhost:8447/tasks/86e0dda5-f6ae-461a-8e14-d6a068c38863

CreateTask  POST:   http://localhost:8447/tasks

UpdateTask  PUT:    http://localhost:8447/tasks

DeleteTask:  DELETE: http://localhost:8447/tasks/86e0dda5-f6ae-461a-8e14-d6a068c38863