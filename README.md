# go-parser
JSON RPC endpoint to parse passed URL

### Requirements
- Docker 1.12.3-rc1 or above
- Docker Composer 1.8.1 or above
- GNU Make

### Setup
Once you've cloned the repo run:


```
make compile
```

Compile the go app.

```
docker-compose up
```

Starts the container

```
make testv
```

Runs the tests

```
make cover
```

Run coverage report and compile into html in the reports directory


### Making JSON RPC calls

*Via CURL*
`curl --data-binary '{"id":1,"method":"url.Parse", "params":["https://github.com/vavas/go-parser"], "jsonrpc": "2.0"}' -H 'content-type:application/json;' -X POST http://172.18.0.2:8000/rpc`

**Response should looks like:**
```
{"result":{"Title":"GitHub - vavas/go-parser","Description":"Contribute to vavas/go-parser development by creating an account on GitHub."},"error":null,"id":1}
```