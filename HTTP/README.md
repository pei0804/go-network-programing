# HTTP

## UserAgent

### Get

```console
go run userAgentGet/main.go http://google.com:80
HTTP/1.1 200 OK
Cache-Control: private, max-age=0
Content-Type: text/html; charset=Shift_JIS
Date: Tue, 26 Dec 2017 03:19:05 GMT
Expires: -1
P3p: CP="This is not a P3P policy! See g.co/p3phelp for more info."
Server: gws
Set-Cookie: 1P_JAR=2017-12-26-03; expires=Thu, 25-Jan-2018 03:19:05 GMT; path=/; domain=.google.co.jp
Set-Cookie: NID=120=o1OotndNw6JaD3Y74VWdkkFa7Ni2cBZm9p8YL6RjEQI4TBfqHc8iI5ss5Ce9It0OzBkQvaUq2b7qOVkHHJN6Uv2e2N50cr9w0fbKW7bdOYYH5R6I3dkIKc7ekLjUj0Q7; expires=Wed, 27-Jun-2018 03:19:05 GMT; path=/; domain=.google.co.jp; HttpOnly
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

Cannot handle [text/html; charset=Shift_JIS]
exit status 4
```

### Head

```console
go run userAgentHead/main.go http://google.com
200 OK
Date: [Tue, 26 Dec 2017 03:20:38 GMT]
Cache-Control: [private, max-age=0]
Content-Type: [text/html; charset=Shift_JIS]
Server: [gws]
X-Frame-Options: [SAMEORIGIN]
Vary: [Accept-Encoding]
Expires: [-1]
P3p: [CP="This is not a P3P policy! See g.co/p3phelp for more info."]
Accept-Ranges: [none]
X-Xss-Protection: [1; mode=block]
Set-Cookie: [1P_JAR=2017-12-26-03; expires=Thu, 25-Jan-2018 03:20:38 GMT; path=/; domain=.google.co.jp NID=120=HhlRAZ5rHHw8teesKhLcPJbYNaXRS_jnXAIv4p3xx3We-Jvo-s1GNASehJ9rnVaBQKdSa3-sEK56xEtS-NcV4SpHIh7NXpDQuDwyos0n3McVH9_I7grURnoWWWqcXxof; expires=Wed, 27-Jun-2018 03:20:38 GMT; path=/; domain=.google.co.jp; HttpOnly]
exit status 2
```

## Proxy Get

```console
go run main.go http://XYZ.com:8080/ http://www.google.com
GET / HTTP/1.1
Host: www.google.com
```


## Proxy Auth Get

```console
go run proxyAuthGet/main.go http://XYZ.com:8080/ http://www.google.com
GET / HTTP/1.1
Host: www.google.com
Proxy-Authorization: Basic amFubmV3bWFyY2g6bXlwYXNzd29yZA==
```

## Handle Function

```console
go run HandlerFunction/main.go
open http://localhost:8080
open http://localhost:8080/cgi-bin/printenv
```
