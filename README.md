# go_stuff

## testCassandra.go
Simple test script using the gocql driver.  Takes a cluster endpoint and credentials, and returns the name of the cluster on success.

To build:

    go build testCassandra

To run:

    ./testCassandra HOSTNAME USERNAME PASSWORD

Example:

    ./testCassandra 127.0.0.1 cassandra cassandra

## testCassandraSSL.go
Simple test script using the gocql driver.  Takes a cluster endpoint, certs and credentials, and returns the name of the cluster on success.  Works with Astra DB.

    ./testCassandraSSL HOSTNAME USERNAME PASSWORD PORT CERTAUTHORITYFILE CLIENTCERT CLIENTCERTKEY

Astra DB Example:

    ./testCassandraSSL cc9blahblahd9ef0-us-central-7.db.astra.datastax.com rtFlynnLivesblahblahEwWB "xwpyKTMeZDIMC6D7blahblahblah,blahblahblah:blahDTuwdZgSkWQPTlf2CuO" 29042 /astra/ca.crt astra/cert astra/key
