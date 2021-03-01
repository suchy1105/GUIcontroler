FROM alpine

RUN apk add git go
#RUN bash
RUN go get -u github.com/go-chi/chi \
    gopkg.in/yaml.v2 \
    github.com/rs/zerolog \
    github.com/rs/zerolog/log


RUN mkdir /app
ADD . /app
WORKDIR /app
## Our project will now successfully build with the necessary go libraries included.
RUN go build -o main .
CMD ["/app/main"]