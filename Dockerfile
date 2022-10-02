# FROM golang
FROM alpine
# FROM golang:1.15

COPY ./kcalendar /tmp/kcalendar

WORKDIR /tmp/

RUN chmod +x kcalendar