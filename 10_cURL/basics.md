# CURL

Make request from command line, and can be used on scripts.

 * http get:
```
curl http://example.com/uwu
```

 * http post with params and header:
```
curl -X POST http://example.com/uwu -d 'param1=value1&param2=value2' -H "Content-Type: application/json"
```

 * http get filtered with jq:
```
curl https://jsonplaceholder.typicode.com/users/1 | jq .username
```

 * http get filter a whole field with jq:
```
curl https://jsonplaceholder.typicode.com/users | jq '.[].username'
```

 * http get filter multiple fields with jq:
```
curl https://jsonplaceholder.typicode.com/users/1 | jq '.name, .email'
```
