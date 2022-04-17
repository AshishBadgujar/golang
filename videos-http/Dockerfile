FROM golang:1.15-alpine as build
WORKDIR /videos
COPY ./videos/* .
RUN go build -o videos

FROM alpine as runtime
COPY --from=build /videos /
CMD ./videos