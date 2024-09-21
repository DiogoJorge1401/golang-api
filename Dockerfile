FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go .
COPY config config
COPY db db
COPY docs docs
COPY handler handler
COPY middleware middleware
COPY router router
COPY schemas schemas

RUN CGO_ENABLED=0 GOOS=linux go build -o /gopportunities

CMD ["/gopportunities"]