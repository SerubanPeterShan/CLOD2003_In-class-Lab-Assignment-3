# CLOD2003 In Class Assignment 3 (Group A) #

Seruban Peter Shan (500235797)

## Task Management API Guide ##

This guide explains how to use the Task Management API to create, update, and retrieve tasks using the `curl` command.

## Task Create API ##

This section explains how to use the Task Create API to create a new task.

### API POST Method URL ##

Access the URL through the POST method:
<http://localhost:8092/tasks/>

### Fields to Add Tasks ##

```json
{
    "title": "New task title",
    "description": "New task description",
    "status": "New task status"
}
```

> **Note:** If any of the fields are empty, default values will be assigned.

### Fields with Default Values ###

```json
{
    "title": "Sample Task",
    "description": "This is a sample task description",
    "status": "pending"
}
```

The **ID** number is automatically incremented.

### Sample Request - Create Task ###

To create a new task, use the following `curl` command:

```sh
curl -X POST 'http://localhost:8092/tasks/' \
--header 'Content-Type: application/json' \
--data '{
           "title": "Sample Task",
           "description": "This is a sample task description",
           "status": "pending"
}'
```

## Task Get API ##

This section explains how to use the API to get all tasks or a specific task by ID.

### API GET Method URL ###

Access the URL through the GET method to **GET ALL TASKS**:
<http://localhost:8092/tasks/>

**OR**

Access the URL through the GET method to **GET TASK BY ID**:
<http://localhost:8092/tasks/(task_id)>

> Replace `<task_id>` with the ID of the task you want to retrieve.
>
> In the header, add `Content-Type: application/json`.

### Sample Request - Get all task ###

To get all task, use the following `curl` command:

```sh
curl -X GET 'http://localhost:8092/tasks/' \
--header 'Content-Type: application/json'
```

### Sample Request - Get task by id ###

To get a task by id, use the following `curl` command:

```sh
curl -X GET 'http://localhost:8092/tasks/(taskid)' \
--header 'Content-Type: application/json'
```

## Task Update API ##

This section explains how to use the Task Update API to update an existing task.

### API PUT Method URL ###

Access the URL through the PUT method:
<http://localhost:8092/tasks/(task_id)>

> Replace `<task_id>` with the ID of the task you want to update.
>
> In the header, add `Content-Type: application/json`.

### Fields to be Updated ###

```json
{
    "title": "New task title",
    "description": "New task description",
    "status": "New task status"
}
```

> **Note:** If no changes are required for any of these fields, leave the corresponding field empty in the JSON payload.

### Sample Request - Update Task ###

To update a task with ID 1, use the following `curl` command:

```sh
curl -X PUT 'http://localhost:8092/tasks/1' \
--header 'Content-Type: application/json' \
--data '{
           "title": "New Task v1",
           "description": "Updated description",
           "status": "Updated status"
}'
```
