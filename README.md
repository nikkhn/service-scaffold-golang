# service-template-go

## Cassandra connector

`cassandra.go` implements a simple wrapper around the `gocql` Cassandra wrapper.
```
go get github.com/gocql/gocql
```

### Configuration
TODO: storing a password in env var is a bad idea. Maybe move secrets to a config file?

The connection to a Cassandra cluster can be configured by setting the following env variables:
* `CASSANDRA_KEYSPACE` the default keyspace to use.
* `CASSANDRA_HOST` the host to connect to (TODO: this should be multiple hosts).
* `CASSANDRA_USERNAME` for password based authentication.
* `CASSANDRA_PASSWORD` for password based authentication.
* `CASSANDRA_CQL_PORT` = 9042


### API (WIP)
Public methods. WIP.
  * `GetCassandraSession() gocql.Session` Init a Cassandra connection, and returns a `Session`.
  * `IterRows() chan map[string]interface{}` returns a generator over a paginated result set.

### Example
`imagerec_example.go` shows how to fetch results from the `imagerec` keyspace 
provided at [https://github.com/gmodena/wmf-streaming-imagematching](https://github.com/gmodena/wmf-streaming-imagematching).

Clone the repo and start a dockerized Cassandra cluster as described in the project README.md.
```bash
git clone https://github.com/gmodena/wmf-streaming-imagematching.git
make cassandra
```
Port `9042` (CQL protocol) is mapped from the container to the host at 127.0.0.1:9042

Execute the example with
```
go build imagerec_example.go cassandra.go 
CASSANDRA_KEYSPACE=imagerec ./imagerec_example
```