### Prerequisites

- Go        : go1.15.6
- Bee       : v2.0.2 
- Beego     : v2.0.1

### Set Environment Variable

```
export STARCOIN_ES_URL=<url>
export STARCOIN_ES_USER=<username>
export STARCOIN_ES_PWD=<password>
export STARCOIN_ES_PREFIX=<prefix>
```

### Init & Run

```
git clone git@github.com:starcoinorg/starcoin-explorer-api.git
cd starcoin-explorer-api
go install
bee run -gendoc=true
```

### Docs

- publish swagger folder

    `bee generate docs`

- online restful api doc and test

    `bee run -gendoc=true` 
    
    `bee run -downdoc=true -gendoc=true`  - only for the fist time 
    
    > http://127.0.0.1:8080/swagger/ 
    
    > auto register routers: routers/commentsRouter_*.go 
    
    > only available while runmode = dev (conf/app.conf)


### Test
- goconvey
    - http://127.0.0.1:8080/ - web UI
    - http://127.0.0.1:8080/reports/ - view test coverage

- go test
    - with cache
        `go test ./tests`
    - without cache
        `go test ./tests -v` 

- Goland
    - open tests/*.go, press `Ctrl + Shift + R`


### Benchmarks
```
cd tests
go test -cpu 1 -run  Benchmark_ -bench=.
```
 