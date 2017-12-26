# Web Scokets

WebSocketがどのように起動されるかは、ユーザエージェントが「Webソケットに切り替える」という特別なHTTP要求を送信することです。 HTTPリクエストの基礎となるTCP接続は開いたままですが、HTTPレスポンスを取得してソケットを閉じる代わりに、ユーザーエージェントとサーバーの両方がWebソケットプロトコルを使用するように切り替えます。

## Echo

```console
go run server/main.go
Echoing
Sending to client: Hello 0
Received back from client: Hello 0
Sending to client: Hello 1
Received back from client: Hello 1
Sending to client: Hello 2
Received back from client: Hello 2
Sending to client: Hello 3
Received back from client: Hello 3
Sending to client: Hello 4
Received back from client: Hello 4
Sending to client: Hello 5
Received back from client: Hello 5
Sending to client: Hello 6
Received back from client: Hello 6
Sending to client: Hello 7
Received back from client: Hello 7
Sending to client: Hello 8
Received back from client: Hello 8
Sending to client: Hello 9
Received back from client: Hello 9
```

```console
go run client/main.go ws://localhost:12345
Received from server: Hello 0
Received from server: Hello 1
Received from server: Hello 2
Received from server: Hello 3
Received from server: Hello 4
Received from server: Hello 5
Received from server: Hello 6
Received from server: Hello 7
Received from server: Hello 8
Received from server: Hello 9
```

## JSON

```console
go run jsonClient/main.go ws://localhost:12345
```

```console
go run jsonServer/main.go
Name: Jan
An email: ja@newmarch.name
An email: jan.newmarch@gmail.com
```

- [XML](http://tumregels.github.io/Network-Programming-with-Go/websockets/the_codec_type.html)
- [TLS](http://tumregels.github.io/Network-Programming-with-Go/websockets/web_sockets_over_tls.html)
