package main

import (
    "context"
    "fmt"
    "github.com/gocql/gocql"
    "os"
    "strconv"
    "time"
)

func main() {
    // set default port
    var port int = 9042
    var err error

    // read command line arguments
    hostname := os.Args[1]
    username := os.Args[2]
    password := os.Args[3]

    // debug
    //fmt.Println("args == ",len(os.Args))

    if len(os.Args) > 4 {
        port,err = strconv.Atoi(os.Args[4])
    }

    // Cluster connection/session code
    cluster := gocql.NewCluster(hostname)
    cluster.Port = port

    // auth
    cluster.Authenticator = gocql.PasswordAuthenticator{
        Username: username,
        Password: password,
    }

    // force protocol version 4
    cluster.ProtoVersion = 4
    session, err := cluster.CreateSession()
    if err != nil {
    	  fmt.Println(err)
    }
    defer session.Close()
    ctx := context.Background()
    // connection established

    // define strKey to read
    var strClusterName string
    err2 := session.Query(`SELECT cluster_name FROM system.local`).WithContext(ctx).Scan(&strClusterName)
    if err2 != nil {
        fmt.Println(err)
    } else {
        fmt.Println("cluster_name:", strClusterName)
    }

    // https://stackoverflow.com/questions/17690776/how-to-add-pause-to-a-go-program
    duration := time.Duration(10)*time.Second // Pause for 10 seconds
    time.Sleep(duration)
}
