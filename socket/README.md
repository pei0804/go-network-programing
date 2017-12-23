# Network Programing

## IPアドレス

```console
go run ip/main.go 192.168.2.1
The address is  192.168.2.1
```

## サブネットマスク

```console
go run mask/main.go 192.168.100.1
Address is  192.168.100.1  Default mask length is  32  Leading ones count is  24  Mask is(hex)  ffffff00  Network is  192.168.100.0
```

## 名前解決

```console
go run resolve/main.go google.co.jp
Resolved address is  216.58.221.163
```

## nslookup的な

```console
go run lookupHost/main.go google.com
172.217.26.110
2404:6800:400a:809::200e
```

## ウェルノウンポート

`/etc/services` にリスト化されている

```console
go run lookupPort/main.go tcp http
Service port  80
```

## TCPアドレス解析

```console
go run resolveTCPAddr/main.go tcp google.com:80
Resolved address is  216.58.196.238:80
```

## TCP コネクション

```console
go run tcpSockets/main.go HEAD www.google.com:80
HTTP/1.0 302 Found
Cache-Control: private
Content-Type: text/html; charset=UTF-8
Referrer-Policy: no-referrer
Location: http://www.google.co.jp/?gfe_rd=cr&dcr=0&ei=C-A8WtipFceQ8QeOnpjoDg
Content-Length: 271
Date: Fri, 22 Dec 2017 10:35:55 GMT
```

```console
go run tcpSockets/main.go GET google.com:80
HTTP/1.0 302 Found
Cache-Control: private
Content-Type: text/html; charset=UTF-8
Referrer-Policy: no-referrer
Location: http://www.google.co.jp/?gfe_rd=cr&dcr=0&ei=hOE8WtLQIs2Q8QeX0bLoCw
Content-Length: 271
Date: Fri, 22 Dec 2017 10:42:12 GMT

<HTML><HEAD><meta http-equiv="content-type" content="text/html;charset=utf-8">
<TITLE>302 Moved</TITLE></HEAD><BODY>
<H1>302 Moved</H1>
The document has moved
<A HREF="http://www.google.co.jp/?gfe_rd=cr&amp;dcr=0&amp;ei=hOE8WtLQIs2Q8QeX0bLoCw">here</A>.
</BODY></HTML>
```

# daytime 時刻サービス

```console
go run daytimeServer/main.go

telnet localhost 1200
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
2017-12-22 19:54:56.994908505 +0900 JSTConnection closed by foreign host.
```

## マルチスレッドサーバー

```console
go run multiThreadedServer/main.go
```

```console
telnet localhost 1201
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
a
a
a
a
test
test
```

```console
telnet localhost 1201
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
te
te
dsa
dsa
adwadad
adwadad
```

## UDP Daytime

```console
go run udpDaytimeServer/main.go
```

```console
go run udpDaytimeClient/main.go 127.0.0.1:1200
2017-12-23 13:13:46.158747639 +0900 JST
```
