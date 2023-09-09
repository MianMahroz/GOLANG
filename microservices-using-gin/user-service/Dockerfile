FROM golang

WORKDIR D:/mahrozrepo/repo/GOLANG/microservices-using-gin/user-service/

# COPY go.mod, go.sum and download the dependencies
COPY go.* ./
RUN go mod download

# COPY All things inside the project and build
COPY . .
RUN go build -o D:/mahrozrepo/repo/GOLANG/microservices-using-gin/user-service/build/myapp .

EXPOSE 8002
ENTRYPOINT [ "D:/mahrozrepo/repo/GOLANG/microservices-using-gin/user-service/build/myapp" ]