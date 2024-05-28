FROM golang:latest as build
LABEL authors="darleet"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o build/server ./cmd/server/main.go

FROM alpine

COPY --from=build /app/build/* /opt/

ENTRYPOINT ["/opt/server"]