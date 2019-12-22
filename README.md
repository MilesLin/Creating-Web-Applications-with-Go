This is the course pratice from [Creating Web Applications with Go](https://app.pluralsight.com/library/courses/creating-web-applications-go-update/table-of-contents)

It uses postgresql in the demo, which I uploaded to docker hub https://hub.docker.com/repository/docker/mileslin/go_lab_postgresql

`docker run -d mileslin/go_lab_postgresql`
* database: lss
* user: postgres
* password: 2wsx#EDC


**Other commands:**

Gen TLS Key

`go run C:\Go\src\crypto\tls\generate_cert.go -host localhost`

snapshot profile

`go tool pprof http://localhost:8081/debug/pprof/heap`

