FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY go.mod go.sum ./
RUN go mod download

# copy other code
COPY . .

RUN swag init

# Compile application
RUN go build -o main .

# port
EXPOSE 8080

# run app
CMD ["./main"]