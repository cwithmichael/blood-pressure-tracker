FROM golang:1.14
WORKDIR /code
COPY ./server/ /code
RUN go mod tidy
RUN go build
CMD ./blood-pressure-tracker
