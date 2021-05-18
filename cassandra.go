package main

import (
    "github.com/gocql/gocql"
    "log"
    "os"
    "strconv"
)

const (
   CASSANDRA_KEYSPACE = "keyspace"
   CASSANDRA_HOST = "localhost"
   CASSANDRA_USERNAME = "cassandra"
   CASSANDRA_PASSWORD = "cassandra"
   CASSANDRA_CQL_PORT = "9042"
   PAGE_SIZE = 10
)

type Cassandra struct {
    // get version, host, date, etc. from Makefile
}

type clusterConfig struct {
    hosts string // TODO(gmodena) should be a list of hosts
    keyspace string
    cql_port int
    authenticator gocql.Authenticator
}


func lookupEnvOrElse(key string, fallback string) string {
    val, err := os.LookupEnv(key)
    if err == false {
       val = fallback
    }
    return val
}

func getConfig() clusterConfig {
    host := lookupEnvOrElse("CASSANDRA_HOST", CASSANDRA_HOST)
    username := lookupEnvOrElse("CASSANDRA_USERNAME", CASSANDRA_USERNAME)
    password := lookupEnvOrElse("CASSANDRA_PASSWORD", CASSANDRA_PASSWORD)
    keyspace := lookupEnvOrElse("CASSANDRA_KEYSPACE", CASSANDRA_KEYSPACE)
    cql_port, err := strconv.Atoi(lookupEnvOrElse("CASSANDRA_CQL_PORT", CASSANDRA_CQL_PORT))
    if err != nil {
        log.Fatal(err)
    }
    authenticator := gocql.PasswordAuthenticator{Username: username, Password: password}

    return clusterConfig{hosts: host, keyspace: keyspace, cql_port: cql_port, authenticator: authenticator}
}

func initCluster(config clusterConfig) *gocql.ClusterConfig {
    cluster := gocql.NewCluster(config.hosts)
    cluster.Authenticator = config.authenticator
    cluster.Keyspace = config.keyspace
    cluster.Consistency = gocql.One
    cluster.Port = config.cql_port
    cluster.PageSize = PAGE_SIZE
    return cluster
}

func IterRows(query gocql.Query) chan map[string]interface{} {
    err := query.Exec()
    if err != nil {
        log.Fatal(err)
    }

    iter := query.Iter()
    c := make(chan map[string]interface{})
    go func() {
        rowValues := make(map[string]interface{})
        for iter.MapScan(rowValues) {
            c  <- rowValues
            rowValues = make(map[string]interface{})
        }
        close(c)
    }()
    return c
}

func getSession() *gocql.Session {
    config := getConfig()
    cluster := initCluster(config)
    session, err := cluster.CreateSession()
    if  err != nil {
        log.Fatal(err)
    }
    return session
}

func GetCassandraSession() *gocql.Session {
    return getSession()
}