# EasyDB-GO client
easydb.io golang client

[![CircleCI](https://circleci.com/gh/heavenlyteam/easydb-go/tree/master.svg?style=svg)](https://circleci.com/gh/heavenlyteam/easydb-go/tree/master)
[![codecov](https://codecov.io/gh/heavenlyteam/easydb-go/branch/master/graph/badge.svg)](https://codecov.io/gh/heavenlyteam/easydb-go)


### Install 

```go
import "github.com/heavenlyteam/easydb-go"
```


### Create a new instance

```go
package main

import "github.com/heavenlyteam/easydb-go"

func main() {
    db, err := easydb.Open("databasename", "token")
    if err != nil {
        panic(err)
    }
}

```

### Save a new object

```go
var (
    db *easydb.Client
    testData = "Can be any type"
)

err := db.Put("key", testData)

```

### Get object

```go
var db *easydb.Client

result, err := db.Get("key")
```

### Delete object

```go
var db *easydb.Client

result, err := db.Delete("key")
```

### All objects list

```go
var db *easydb.Client

result, err := db.List()
```