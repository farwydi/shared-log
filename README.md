# shared-log
Simple shared zaplog for golang modules

## The Problem
How share log between other packages?  
Simple...
```go
package a

import "github.com/farwydi/shared-log"

func A()  {
    shared_log.Info("Hello world")
}
```
```go
package b

import "github.com/farwydi/shared-log"

func B()  {
    shared_log.Info("Hello world")
}
```

## How use
```go
package main

import (
    "github.com/farwydi/shared-log"
    "go.uber.org/zap"
)

func main()  {
    logger, _ := zap.NewProduction()
    zap.ReplaceGlobals(logger)
    
    // =======    

    shared_log.Info("Hello world")
}
```

## Trace
Passing parameters to the session
```go
package main

import (
    "github.com/farwydi/shared-log"
    "go.uber.org/zap"
    "net/http"
)

func main()  {
    logger, _ := zap.NewProduction()
    zap.ReplaceGlobals(logger)
    
    // ...
    http.HandleFunc("/", handle)
    http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request)  {
    logger := shared_log.NewTrace(zap.String("func", "base"))
    logger.Info("Log...")
    // ...
}
```