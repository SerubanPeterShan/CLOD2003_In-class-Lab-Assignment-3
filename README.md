## Task Create API ##
This Guide explains how to use the Task Create API using the curl command to create am existing task with given by user for title ,description or status 


**API POST Method URL:**
This URL can be accessed through the POST method:
http://localhost:8092/tasks/



**Fields to add tasks:**
```json
{
    "title": "New task title",
    "description": "New task description",
    "status": "New task status"
}
```
> **_If any of the field is empty it will asign the defualt values for each field_**

**Fields with default values:**
```json
{
    "title":"Sample Task",
    "description":"This is a sample task description",
    "status":"pending"
}
```
**ID** number is Automatically incremented

**_Sample Request_**

To create a new task, use the following curl command:

```console
curl -X POST 'http://localhost:8092/tasks/' \
--header 'Content-Type: application/json' \
--data '{
           "title": "Sample Task",
           "description": "This is a sample task description",
           "status": "pending"
}'
```


## Task Update API ##
This guide explains how to use the Task Update API using the curl command to update an existing task with new values for title, description, or status.

**API PUT Method URL:**
This URL can be accessed through the PUT method:
http://localhost:8092/tasks/<task_id>

> Replace <task_id> with the ID of the task you want to update.
>
>In header add Content-Type: application/json

**Fields to be updated:**
```json
{
    "title": "New task title",
    "description": "New task description",
    "status": "New task status"
}
```
> **_If no changes are required for any of these fields, leave the corresponding field empty in the JSON payload._**

**Sample Request**
To update a task with ID 1, use the following curl command:

```console
curl -X 'http://localhost:8092/task' \
--header 'Content-Type: application/json' \
--data '{
           "title": "New Task _ v1 ",
           "description": "Poathan",
           "status": "Chama"
}'
```


