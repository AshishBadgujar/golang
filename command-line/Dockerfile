FROM golang:1.15-alpine as build

WORKDIR /videos
COPY . .
RUN go build -o videos


FROM alpine as runtime 
COPY --from=build /videos .
COPY run.sh /
RUN chmod +x /run.sh
ENTRYPOINT [ "./run.sh" ]