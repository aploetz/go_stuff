# go_stuff

## testCassandra.go
Simple test script using the gocql driver.  Takes a cluster endpoint and credentials, and returns the name of the cluster on success.

To build:
    go build testCassandra

To run:

    ./testCassandra HOSTNAME USERNAME PASSWORD

Example:

    ./testCassandra 127.0.0.1 cassandra cassandra
