FROM golang:1.22
WORKDIR /golang_htmx

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /golang_htmx

CMD ["/golang_htmx"]


