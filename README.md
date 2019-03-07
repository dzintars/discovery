Source: https://www.youtube.com/watch?v=iYrMAVSOkkA


To build minimal Docker image
https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
