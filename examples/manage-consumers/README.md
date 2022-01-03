# Manage Consumers

Basic example showing how to manage consumers with the client.

```shell
$ ./manage-consumers.bin
2022/01/02 16:32:51
GET /v1/admin/ready HTTP/1.1
Host: 127.0.0.1:4000
User-Agent: OpenAPI-Generator/1.0.0/go
Accept: application/json
Accept-Encoding: gzip


2022/01/02 16:32:51
HTTP/2.0 200 OK
Content-Length: 16
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:32:51 GMT

{"success":true}
2022/01/02 16:32:51
POST /v1/admin/stream HTTP/1.1
Host: 127.0.0.1:4000
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 102
Accept: application/json
Content-Type: application/json
Accept-Encoding: gzip

{"max_age":600000000000,"name":"cd2f96be-1db4-4a24-baf8-5092ffaa3c99","subjects":["subj.1","subj.2"]}

2022/01/02 16:32:51
HTTP/2.0 200 OK
Content-Length: 16
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:32:51 GMT
Httpmq-Request-Id: 69aa7ede-a45a-48c4-beca-a4d5e21f2284

{"success":true}
2022/01/02 16:32:51
POST /v1/admin/stream/cd2f96be-1db4-4a24-baf8-5092ffaa3c99/consumer HTTP/1.1
Host: 127.0.0.1:4000
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 79
Accept: application/json
Content-Type: application/json
Accept-Encoding: gzip

{"max_inflight":1,"mode":"push","name":"1e86dbc2-3326-4a9d-80bf-0c2779addc8f"}

2022/01/02 16:32:51
HTTP/2.0 200 OK
Content-Length: 16
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:32:51 GMT
Httpmq-Request-Id: 70449216-3cb0-4160-bd60-cc4bbc86c9cf

{"success":true}
2022/01/02 16:32:51
GET /v1/admin/stream/cd2f96be-1db4-4a24-baf8-5092ffaa3c99/consumer/1e86dbc2-3326-4a9d-80bf-0c2779addc8f HTTP/1.1
Host: 127.0.0.1:4000
User-Agent: OpenAPI-Generator/1.0.0/go
Accept: application/json
Accept-Encoding: gzip


2022/01/02 16:32:51
HTTP/2.0 200 OK
Content-Length: 456
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:32:51 GMT
Httpmq-Request-Id: 421d82d2-8b5b-4e82-b15c-eb1c3d1316f0

{"success":true,"consumer":{"stream_name":"cd2f96be-1db4-4a24-baf8-5092ffaa3c99","name":"1e86dbc2-3326-4a9d-80bf-0c2779addc8f","created":"2022-01-03T00:32:51.295448311Z","config":{"deliver_subject":"_INBOX.PYVhfCDjtS018T3zeymwAs","max_deliver":-1,"ack_wait":30000000000,"max_ack_pending":1},"delivered":{"consumer_seq":0,"stream_seq":0},"ack_floor":{"consumer_seq":0,"stream_seq":0},"num_ack_pending":0,"num_redelivered":0,"num_waiting":0,"num_pending":0}}
{
  "ack_floor": {
    "consumer_seq": 0,
    "stream_seq": 0
  },
  "config": {
    "ack_wait": 30000000000,
    "deliver_subject": "_INBOX.PYVhfCDjtS018T3zeymwAs",
    "max_ack_pending": 1,
    "max_deliver": -1
  },
  "created": "2022-01-03T00:32:51.295448311Z",
  "delivered": {
    "consumer_seq": 0,
    "stream_seq": 0
  },
  "name": "1e86dbc2-3326-4a9d-80bf-0c2779addc8f",
  "num_ack_pending": 0,
  "num_pending": 0,
  "num_redelivered": 0,
  "num_waiting": 0,
  "stream_name": "cd2f96be-1db4-4a24-baf8-5092ffaa3c99"
}
Request ID 421d82d2-8b5b-4e82-b15c-eb1c3d1316f0
2022/01/02 16:32:51
DELETE /v1/admin/stream/cd2f96be-1db4-4a24-baf8-5092ffaa3c99 HTTP/1.1
Host: 127.0.0.1:4000
User-Agent: OpenAPI-Generator/1.0.0/go
Accept: application/json
Accept-Encoding: gzip


2022/01/02 16:32:51
HTTP/2.0 200 OK
Content-Length: 16
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:32:51 GMT
Httpmq-Request-Id: 2dd89bf8-3871-4431-93c2-a84d03872bbf

{"success":true}
```
