FROM alpine:latest
WORKDIR /app

COPY ./*.env .

COPY gomonagent .

CMD ["./gomonagent"]
