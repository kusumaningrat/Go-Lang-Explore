# Simple Todo List App with Slice in Go lang

This is simple todo list system that built with go lang and utilize slice as temporary database.
This documentation is purely to practice and gain understanding basic data structure especially in Go Lang.

## Features / Functions

### Formated Response

```
{
    "message": string,
    "data": any
}
```

#

### Create Todo

Request

```

- Create() ->
  {
    "title": string,
    "description": string,
    "completed": bool (default false)
  }
```

Response

```

{
    "message": "Todo created successfully",
    "data": {
        "id": 1,
        "title": "Todo1",
        "description": "Description Todo1",
        "completed": false
    }
}

```

#

### Get All Todos

Response

```
- GetAll() ->
  {
    "message": "Users loaded successfully",
    "data": [
        {
            "id": 1,
            "title": "Todo1",
            "description": "Description Todo1",
            "completed": false
        },
        {
            "id": 2,
            "title": "Todo2",
            "description": "Description Todo2",
            "completed": true
        }
    ]
  }

```

#

### Get Single Todo

Pass id's of todo

Request param (id)

Response

```
- GetOne(id) ->
    {
        "message": "User loaded successfully",
        "data": {
            "id": 1,
            "title": "Todo1",
            "description": "Description Todo1",
            "completed": false
        },
    }
```

#

### Update Todo

Pass id's of todo

Request Body

```
- Update(id) ->
    {
        "title": "Todo1 Update",
        "description": "Description Todo1 Update",
        "completed": true
    },
```

Response

```
{
    "message": "User updated successfully",
    "data": {
        "id": 1,
        "title": "Todo1 Update",
        "description": "Description Todo1 Update",
        "completed": true
    },
}
```

#

### Delete Todo

Pass id's of todo

Requst param(id)

Response

```

- Delete(id) ->
{
    "message": "User deleted successfully",
    "data": null
}

```

#
