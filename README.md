# service-template-go

This branch contains a demo endpoint built atop `service-scaffold-golang` and `https://github.com/eevans/servicelib-golang`.

Depends on
```
go get github.com/eevans/servicelib-golang/logger
go get "github.com/google/uuid"
```

# Scaffold endpoints

* Health status is available at `/healtz`
```
curl http://localhost:8000/healthz        
{"version":"365c332","build_date":"1621532979","build_host":"wmf2799","go_version":"go1.15.4"}
```

# Example endpoint
`imagematching.go` implements an endpoint that fetch results from the `imagerec` keyspace 
provided at [https://github.com/gmodena/wmf-streaming-imagematching](https://github.com/gmodena/wmf-streaming-imagematching).

Clone the repo and start a dockerized Cassandra cluster as described in the project README.md.
```bash
git clone https://github.com/gmodena/wmf-streaming-imagematching.git
make cassandra
```

The endpoint depends on the `gocql` package:
```
go get github.com/gocql/gocql
```

Then execute the service with
```
make run
```

And query with:
```
$ curl "http://localhost:8000/predict?wiki=enwiki&page_id=10003732"
{"prediction":[{"image_id":"Genova_12-8-05_040.jpg"}],"model_version":"1a"}
```

### Docker Quickstart

Generated a Dockerfile for service variant with `blubber .pipeline/blubber.yaml <variant> > Dockerfile`,
and build using regular Docker tools.


For example, build and run a `development` variant of a service with:
```
blubber .pipeline/blubber.yaml development > Dockerfile
docker build -t service-scaffold-golang .
docker run -p 8000:8000  service-scaffold-golang
```

