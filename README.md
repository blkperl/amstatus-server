# Amstatus Server

## API Usage

### Amstatus

Resource: /amstatus
Method: GET

#### Curl Example
```
curl -i http://127.0.0.1:9000/amstatus
```

Response

```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 07 Sep 2014 20:56:01 GMT
Content-Length: 118

{
 "host": "filer.example.org",
  "disk": "/path/to/disk",
  "dumping": "302m",
  "dumped": "66556m"
}
```
