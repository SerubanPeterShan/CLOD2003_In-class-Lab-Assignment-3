Sample Update CURL 

1           > Task ID 
Title       > NEw title
description > New Description 
status      > New Status 



```sh
curl --location 'http://localhost:8092/updatetasks/1' \
--header 'Content-Type: application/json' \
--data '{
           "title": "New Task _ v1 ",
           "description": "Poathan",
           "status": "Chama"
}'

