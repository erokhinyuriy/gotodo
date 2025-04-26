# Project "gotodo"

Web API for ToDo list application.

## Description

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