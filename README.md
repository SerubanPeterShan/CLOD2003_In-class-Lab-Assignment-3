**Task Update API**
This guide explains how to use the Task Update API using the curl command to update an existing task with new values for title, description, or status.

**Requirements**
API URL: http://localhost:8092/updatetasks/<task_id>

Replace <task_id> with the ID of the task you want to update.
Content-Type: application/json

**Fields to be updated:**

title: New task title
description: New task description
status: New task status
If no changes are required for any of these fields, leave the corresponding field empty in the JSON payload.

**Sample Request**
To update a task with ID 1, use the following curl command:

```sh
curl --location 'http://localhost:8092/updatetasks/1' \
--header 'Content-Type: application/json' \
--data '{
           "title": "New Task _ v1 ",
           "description": "Poathan",
           "status": "Chama"
}'

