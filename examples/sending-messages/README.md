# Sending and Receiving Messages

Basic example showing how to publish and subscribe for messages.

```shell
$ ./sending-messages.bin
2022/01/02 16:29:38
GET /v1/admin/ready HTTP/1.1
Host: 127.0.0.1:4000
User-Agent: OpenAPI-Generator/1.0.0/go
Accept: application/json
Accept-Encoding: gzip


2022/01/02 16:29:38
HTTP/2.0 200 OK
Content-Length: 16
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:29:38 GMT

{"success":true}
2022/01/02 16:29:38
GET /v1/admin/ready HTTP/1.1
Host: 127.0.0.1:4000
User-Agent: OpenAPI-Generator/1.0.0/go
Accept: application/json
Accept-Encoding: gzip


2022/01/02 16:29:38
HTTP/2.0 200 OK
Content-Length: 16
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:29:38 GMT

{"success":true}
2022/01/02 16:29:38
POST /v1/admin/stream HTTP/1.1
Host: 127.0.0.1:4000
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 102
Accept: application/json
Content-Type: application/json
Accept-Encoding: gzip

{"max_age":600000000000,"name":"665f64b2-5df0-4c9f-aae5-5a6cb9afd725","subjects":["subj.1","subj.2"]}

2022/01/02 16:29:38
HTTP/2.0 200 OK
Content-Length: 16
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:29:38 GMT
Httpmq-Request-Id: 765049e0-898a-4a27-80a1-26a24286625f

{"success":true}
2022/01/02 16:29:38
POST /v1/admin/stream/665f64b2-5df0-4c9f-aae5-5a6cb9afd725/consumer HTTP/1.1
Host: 127.0.0.1:4000
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 105
Accept: application/json
Content-Type: application/json
Accept-Encoding: gzip

{"filter_subject":"subj.*","max_inflight":1,"mode":"push","name":"5cb52b29-0f38-4c21-af33-0a882affc692"}

2022/01/02 16:29:38
HTTP/2.0 200 OK
Content-Length: 16
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:29:38 GMT
Httpmq-Request-Id: 11a85e2c-4d0b-48bf-8013-86b9fed5c6c6

{"success":true}
2022/01/02 16:29:38
POST /v1/data/subject/subj.2 HTTP/1.1
Host: 127.0.0.1:4001
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 16
Accept: application/json
Content-Type: text/plain
Accept-Encoding: gzip

SGVsbG8gd29ybGQ=
2022/01/02 16:29:38
HTTP/2.0 200 OK
Content-Length: 16
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:29:38 GMT
Httpmq-Request-Id: 440e218f-abe9-4708-82af-b4557dce4ae4

{"success":true}
Read: 'Hello world'
2022/01/02 16:29:38
POST /v1/data/stream/665f64b2-5df0-4c9f-aae5-5a6cb9afd725/consumer/5cb52b29-0f38-4c21-af33-0a882affc692/ack HTTP/1.1
Host: 127.0.0.1:4001
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 26
Accept: application/json
Content-Type: application/json
Accept-Encoding: gzip

{"consumer":1,"stream":1}

2022/01/02 16:29:38
HTTP/2.0 200 OK
Content-Length: 16
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:29:38 GMT
Httpmq-Request-Id: 9e7ae637-9f1f-41bd-af8c-98cd117357b4

{"success":true}
Push subscription complete. Request ID
Subscription errors: context canceled
2022/01/02 16:29:38
DELETE /v1/admin/stream/665f64b2-5df0-4c9f-aae5-5a6cb9afd725 HTTP/1.1
Host: 127.0.0.1:4000
User-Agent: OpenAPI-Generator/1.0.0/go
Accept: application/json
Accept-Encoding: gzip


2022/01/02 16:29:38
HTTP/2.0 200 OK
Content-Length: 16
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:29:38 GMT
Httpmq-Request-Id: 8eb0d37b-fd2f-4a40-91b5-d8e21e44dc57

{"success":true}
```
