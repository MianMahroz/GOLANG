FROM golang

WORKDIR D:/mahrozrepo/repo/GOLANG/microservices-using-gin/order-service/

# COPY go.mod, go.sum and download the dependencies
COPY go.* ./
RUN go mod download

# COPY All things inside the project and build
COPY . .
RUN go build -o D:/mahrozrepo/repo/GOLANG/microservices-using-gin/order-service/build/myapp .

EXPOSE 8001
ENTRYPOINT [ "D:/mahrozrepo/repo/GOLANG/microservices-using-gin/order-service/build/myapp" ]