# Manage Streams

Basic example showing how to manage streams with the client.

```shell
$ ./manage-streams.bin
2022/01/02 16:35:12
GET /v1/admin/ready HTTP/1.1
Host: 127.0.0.1:4000
User-Agent: OpenAPI-Generator/1.0.0/go
Accept: application/json
Accept-Encoding: gzip


2022/01/02 16:35:12
HTTP/2.0 200 OK
Content-Length: 16
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:35:12 GMT

{"success":true}
2022/01/02 16:35:12
POST /v1/admin/stream HTTP/1.1
Host: 127.0.0.1:4000
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 102
Accept: application/json
Content-Type: application/json
Accept-Encoding: gzip

{"max_age":600000000000,"name":"fea6a70a-a4c7-4a3a-a204-ccd9c60e3c6e","subjects":["subj.1","subj.2"]}

2022/01/02 16:35:12
HTTP/2.0 200 OK
Content-Length: 16
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:35:12 GMT
Httpmq-Request-Id: 25e4dfa7-2ad5-4faf-9e39-8e24b1cafbeb

{"success":true}
2022/01/02 16:35:12
GET /v1/admin/stream/fea6a70a-a4c7-4a3a-a204-ccd9c60e3c6e HTTP/1.1
Host: 127.0.0.1:4000
User-Agent: OpenAPI-Generator/1.0.0/go
Accept: application/json
Accept-Encoding: gzip


2022/01/02 16:35:12
HTTP/2.0 200 OK
Content-Length: 419
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:35:12 GMT
Httpmq-Request-Id: 3b5f7d0d-9aa7-4e67-a2da-e324acd776bb

{"success":true,"stream":{"config":{"name":"fea6a70a-a4c7-4a3a-a204-ccd9c60e3c6e","subjects":["subj.1","subj.2"],"max_consumers":-1,"max_msgs":-1,"max_bytes":-1,"max_age":600000000000,"max_msgs_per_subject":-1,"max_msg_size":-1},"created":"2022-01-03T00:35:12.802740214Z","state":{"messages":0,"bytes":0,"first_seq":0,"first_ts":"0001-01-01T00:00:00Z","last_seq":0,"last_ts":"0001-01-01T00:00:00Z","consumer_count":0}}}
{
  "config": {
    "max_age": 600000000000,
    "max_bytes": -1,
    "max_consumers": -1,
    "max_msg_size": -1,
    "max_msgs": -1,
    "max_msgs_per_subject": -1,
    "name": "fea6a70a-a4c7-4a3a-a204-ccd9c60e3c6e",
    "subjects": [
      "subj.1",
      "subj.2"
    ]
  },
  "created": "2022-01-03T00:35:12.802740214Z",
  "state": {
    "bytes": 0,
    "consumer_count": 0,
    "first_seq": 0,
    "first_ts": "0001-01-01T00:00:00Z",
    "last_seq": 0,
    "last_ts": "0001-01-01T00:00:00Z",
    "messages": 0
  }
}
Request ID 3b5f7d0d-9aa7-4e67-a2da-e324acd776bb
2022/01/02 16:35:12
PUT /v1/admin/stream/fea6a70a-a4c7-4a3a-a204-ccd9c60e3c6e/subject HTTP/1.1
Host: 127.0.0.1:4000
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 33
Accept: application/json
Content-Type: application/json
Accept-Encoding: gzip

{"subjects":["subj.2","subj.3"]}

2022/01/02 16:35:12
HTTP/2.0 200 OK
Content-Length: 16
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:35:12 GMT
Httpmq-Request-Id: 6b5eda7b-f480-4191-9179-7d415ffcc84b

{"success":true}
2022/01/02 16:35:12
GET /v1/admin/stream/fea6a70a-a4c7-4a3a-a204-ccd9c60e3c6e HTTP/1.1
Host: 127.0.0.1:4000
User-Agent: OpenAPI-Generator/1.0.0/go
Accept: application/json
Accept-Encoding: gzip


2022/01/02 16:35:12
HTTP/2.0 200 OK
Content-Length: 419
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:35:12 GMT
Httpmq-Request-Id: bf703503-dc6b-4347-be11-c8e783ec002e

{"success":true,"stream":{"config":{"name":"fea6a70a-a4c7-4a3a-a204-ccd9c60e3c6e","subjects":["subj.2","subj.3"],"max_consumers":-1,"max_msgs":-1,"max_bytes":-1,"max_age":600000000000,"max_msgs_per_subject":-1,"max_msg_size":-1},"created":"2022-01-03T00:35:12.802740214Z","state":{"messages":0,"bytes":0,"first_seq":0,"first_ts":"0001-01-01T00:00:00Z","last_seq":0,"last_ts":"0001-01-01T00:00:00Z","consumer_count":0}}}
{
  "config": {
    "max_age": 600000000000,
    "max_bytes": -1,
    "max_consumers": -1,
    "max_msg_size": -1,
    "max_msgs": -1,
    "max_msgs_per_subject": -1,
    "name": "fea6a70a-a4c7-4a3a-a204-ccd9c60e3c6e",
    "subjects": [
      "subj.2",
      "subj.3"
    ]
  },
  "created": "2022-01-03T00:35:12.802740214Z",
  "state": {
    "bytes": 0,
    "consumer_count": 0,
    "first_seq": 0,
    "first_ts": "0001-01-01T00:00:00Z",
    "last_seq": 0,
    "last_ts": "0001-01-01T00:00:00Z",
    "messages": 0
  }
}
Request ID bf703503-dc6b-4347-be11-c8e783ec002e
2022/01/02 16:35:12
PUT /v1/admin/stream/fea6a70a-a4c7-4a3a-a204-ccd9c60e3c6e/limit HTTP/1.1
Host: 127.0.0.1:4000
User-Agent: OpenAPI-Generator/1.0.0/go
Content-Length: 25
Accept: application/json
Content-Type: application/json
Accept-Encoding: gzip

{"max_age":900000000000}

2022/01/02 16:35:12
HTTP/2.0 200 OK
Content-Length: 16
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:35:12 GMT
Httpmq-Request-Id: 3868bd3c-af22-466c-8e02-ab2af716d1c0

{"success":true}
2022/01/02 16:35:12
GET /v1/admin/stream/fea6a70a-a4c7-4a3a-a204-ccd9c60e3c6e HTTP/1.1
Host: 127.0.0.1:4000
User-Agent: OpenAPI-Generator/1.0.0/go
Accept: application/json
Accept-Encoding: gzip


2022/01/02 16:35:12
HTTP/2.0 200 OK
Content-Length: 419
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:35:12 GMT
Httpmq-Request-Id: 3dbb955d-2cde-4b72-b9ee-108a2480a43b

{"success":true,"stream":{"config":{"name":"fea6a70a-a4c7-4a3a-a204-ccd9c60e3c6e","subjects":["subj.2","subj.3"],"max_consumers":-1,"max_msgs":-1,"max_bytes":-1,"max_age":900000000000,"max_msgs_per_subject":-1,"max_msg_size":-1},"created":"2022-01-03T00:35:12.802740214Z","state":{"messages":0,"bytes":0,"first_seq":0,"first_ts":"0001-01-01T00:00:00Z","last_seq":0,"last_ts":"0001-01-01T00:00:00Z","consumer_count":0}}}
{
  "config": {
    "max_age": 900000000000,
    "max_bytes": -1,
    "max_consumers": -1,
    "max_msg_size": -1,
    "max_msgs": -1,
    "max_msgs_per_subject": -1,
    "name": "fea6a70a-a4c7-4a3a-a204-ccd9c60e3c6e",
    "subjects": [
      "subj.2",
      "subj.3"
    ]
  },
  "created": "2022-01-03T00:35:12.802740214Z",
  "state": {
    "bytes": 0,
    "consumer_count": 0,
    "first_seq": 0,
    "first_ts": "0001-01-01T00:00:00Z",
    "last_seq": 0,
    "last_ts": "0001-01-01T00:00:00Z",
    "messages": 0
  }
}
Request ID 3dbb955d-2cde-4b72-b9ee-108a2480a43b
2022/01/02 16:35:12
DELETE /v1/admin/stream/fea6a70a-a4c7-4a3a-a204-ccd9c60e3c6e HTTP/1.1
Host: 127.0.0.1:4000
User-Agent: OpenAPI-Generator/1.0.0/go
Accept: application/json
Accept-Encoding: gzip


2022/01/02 16:35:12
HTTP/2.0 200 OK
Content-Length: 16
Content-Type: application/json
Date: Mon, 03 Jan 2022 00:35:12 GMT
Httpmq-Request-Id: 89abafef-52ef-44e1-b9e8-c0d1691456c4

{"success":true}
```
