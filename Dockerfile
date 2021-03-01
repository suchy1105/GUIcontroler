FROM alpine

RUN apk add git go
#RUN bash
RUN go get -u github.com/go-chi/chi
RUN go get gopkg.in/yaml.v2

RUN mkdir /app
ADD . /app
WORKDIR /app
## Our project will now successfully build with the necessary go libraries included.
RUN go build -o main .
CMD ["/app/main"]