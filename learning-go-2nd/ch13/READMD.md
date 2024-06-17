# Chapter 13
API to query current time
```
GET /system/time/current
```
which will return:
```
2024-06-16T19:50:00-07:00
```


The logging will be saved in a file called "access.log" in the directory
which is the CWD of processing running the server
Example log in json format:
```json
{"time":"2024-06-16T20:05:20.562853262-07:00","level":"DEBUG","msg":"New client request received","source":"[::1]:34168"}
```


When specifying the header `Accept: application/json`, the server will return the current time
in JSON format instead of plain text
```
curl -H "Accept: application/json" http://localhost:8080/system/time/current
```