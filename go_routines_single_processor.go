package main

import (
    "fmt"
    "sync"
)

func main()
{
    runtime.GOMAXPROCS(1) //max number of processors 1
    
    var wg sync.WaitGroup //sync wait group
    wg.Add(2) //


}

