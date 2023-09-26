FROM golang:1.19.5 AS builder

WORKDIR /test_task_go_middle

COPY . .

WORKDIR /test_task_go_middle/grpc-server
RUN go build -o /test_task_go_middle/bin/grpc-server .

WORKDIR /test_task_go_middle/http-grpc-client
RUN go build -o /test_task_go_middle/bin/http-grpc-client .

WORKDIR /test_task_go_middle/test-client
RUN go build -o /test_task_go_middle/bin/test-client .

###

FROM golang:1.19.5 AS grpc-server-runtime

COPY --from=builder /test_task_go_middle/bin/grpc-server /test_task_go_middle/grpc-server
EXPOSE 50051

CMD ["/test_task_go_middle/grpc-server"]

###

FROM golang:1.19.5 AS http-grpc-client-runtime

COPY --from=builder /test_task_go_middle/bin/http-grpc-client /test_task_go_middle/http-grpc-client
EXPOSE 8080

CMD ["/test_task_go_middle/http-grpc-client"]

###

FROM golang:1.19.5 AS test-client-runtime

COPY --from=builder /test_task_go_middle/bin/test-client /test_task_go_middle/test-client

CMD ["/test_task_go_middle/test-client"]