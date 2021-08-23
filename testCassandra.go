package main

import (
    "context"
    "fmt"
    "github.com/gocql/gocql"
    "os"
)

func main() {
    // read command line arguments
    hostname := os.Args[1]
    username := os.Args[2]
    password := os.Args[3]

    // Cluster connection/session code
    cluster := gocql.NewCluster(hostname)
    cluster.Authenticator = gocql.PasswordAuthenticator{
        Username: username,
        Password: password,
    }
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

}
