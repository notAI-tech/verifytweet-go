FROM golang:1.14.6-alpine
RUN apk update && apk add build-base imagemagick-dev file tesseract-ocr-dev git && mkdir -p /home/app
ENV GO111MODULE=on
ENV CGO_CFLAGS_ALLOW=-Xpreprocessor
WORKDIR /go/src/github.com/notAI-tech/verifytweet-go
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags "-s -w" -o /home/app/bin cmd/*.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates tzdata imagemagick tesseract-ocr
COPY --from=0 /home/app/bin /home
WORKDIR /home

EXPOSE 80
CMD ["sh", "-c","./bin"]
