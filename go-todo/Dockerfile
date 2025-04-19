FROM golang:1.15-alpine as build
WORKDIR /contacts
COPY ./contacts/* .
RUN go build -o contacts

FROM alpine as runtime
COPY --from=build /contacts /
CMD ./contacts