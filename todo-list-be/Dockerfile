#build stage
FROM golang:1.20-alpine AS builder

RUN apk add --no-cache git
#migrate will be on /go/bin/migrate
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v app/http/main.go
RUN cp -r db/migrations /go/bin

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
COPY --from=builder /go/bin/migrate /migrate
COPY --from=builder /go/bin/migrations /migrations
# default values for env (can be overidden using -e)
ENV APP_DEBUG=false
ENV PORT=5555
ENV DB_STRING=mysql://root:@tcp(localhost:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local
ENV JWT_TTL=5h
ENV JWT_KEY=my-secret-key
# --- end of env
ENTRYPOINT /app
LABEL Name=dnabil/todolist-be Version=0.0.1
EXPOSE 5555
