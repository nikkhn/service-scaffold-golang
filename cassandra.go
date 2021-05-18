package main

import (
    "fmt"
    "github.com/gocql/gocql"
    "log"
    "os"
)

const (
   CASSANDRA_KEYSPACE = "keyspace"
   CASSANDRA_HOST = "localhost"
   CASSANDRA_USERNAME = "cassandra"
   CASSANDRA_PASSWORD = "cassandra"
   PAGE_SIZE = 10
)

type Cassandra struct {
    // get version, host, date, etc. from Makefile

}

type ClusterConfig struct {
    hosts string // TODO(gmodena) should be a list of hosts
    keyspace string
    authenticator gocql.Authenticator
}

func LookupEnvOrElse(key string, fallback string) string {
    val, err := os.LookupEnv(key)
    if err == false {
       val = fallback
    }
    return val
}

func GetConfig() ClusterConfig {
    host := LookupEnvOrElse("CASSANDRA_HOST", CASSANDRA_HOST)
    username := LookupEnvOrElse("CASSANDRA_USERNAME", CASSANDRA_USERNAME)
    password := LookupEnvOrElse("CASSANDRA_PASSWORD", CASSANDRA_PASSWORD)
    keyspace := LookupEnvOrElse("CASSANDRA_KEYSPACE", CASSANDRA_KEYSPACE)

    authenticator := gocql.PasswordAuthenticator{Username: username, Password: password}

    return ClusterConfig{hosts: host, keyspace: keyspace, authenticator: authenticator}
}

func InitCluster(config ClusterConfig) *gocql.ClusterConfig {
    cluster := gocql.NewCluster(config.hosts)
    cluster.Authenticator = config.authenticator
    cluster.Keyspace = config.keyspace
    cluster.Consistency = gocql.One

    return cluster
}

func GetSession()  *gocql.Session {
    config := GetConfig()
    cluster := InitCluster(config)
    session, err := cluster.CreateSession()
    if err != nil {
        log.Fatal(err)
    }

    session.SetPageSize(PAGE_SIZE)
    return session
}

func main() {
    cassandra := GetSession()
    query := cassandra.Query("select count(*) from matches")

    fmt.Print(query.Exec())
}