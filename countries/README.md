# Countries

Demonstrate simple POST GET and DELETE operations

The curl demo:
```
curl -i -d '{"Code":"FR","Name":"France"}' http://127.0.0.1:8080/countries
curl -i -d '{"Code":"US","Name":"United States"}' http://127.0.0.1:8080/countries
curl -i http://127.0.0.1:8080/countries/FR
curl -i http://127.0.0.1:8080/countries/US
curl -i http://127.0.0.1:8080/countries
curl -i -X DELETE http://127.0.0.1:8080/countries/FR
curl -i http://127.0.0.1:8080/countries
curl -i -X DELETE http://127.0.0.1:8080/countries/US
curl -i http://127.0.0.1:8080/countries
```
