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

    // define columns to read
    var strClusterName           string
    var strBroadcastAddress      string
    var strNativeProtocolVersion string
    var strReleaseVersion        string
    var strSchemaVersion         string

    // define query string
    strCQL := "SELECT cluster_name,broadcast_address,native_protocol_version,release_version,schema_version FROM system.local"

    err2 := session.Query(strCQL).WithContext(ctx).Scan(&strClusterName,&strBroadcastAddress,&strNativeProtocolVersion,&strReleaseVersion,&strSchemaVersion)
    if err2 != nil {
        fmt.Println(err)
    } else {
        fmt.Println("cluster_name:", strClusterName)
        fmt.Println("broadcast_address:", strBroadcastAddress)
        fmt.Println("native_protocol_version:", strNativeProtocolVersion)
        fmt.Println("release_version:", strReleaseVersion)
        fmt.Println("schema_version:", strSchemaVersion)
    }

}
