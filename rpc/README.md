# RPC

## Go RPC

- 関数はpublicでなければなりません（大文字で始まります）。
- 厳密に2つの引数を持ち、最初はクライアントから関数が受け取る値データへのポインタとなり、2つ目はクライアントに返される回答を保持するポインタになる。
- 戻り値の型がerror

value引数の型は、クライアントとサーバで同じではないことに注意してください。 サーバーでは、 Valuesを使用していましたが、クライアントではArgsを使用していました。 それは重要ではありませんgobシリアライゼーションのルールと、2つの構造のフィールドのタイプの名前が一致しているためです。 プログラミングの習慣を向上させると、名前は同じでなければならないと言います。

しかし、これはGo RPCを使用する際に起こりうるトラップを指摘しています。 クライアントの構造を変更して、たとえば、 

```go
type Values struct { C, B int } 
```

gobは問題はありません。サーバー側では、アンマーシャリングはクライアントから与えられたCの値を無視し、Aにはデフォルトのゼロ値を使用します。

Go RPCを使用するには、プログラマーがフィールド名と型の安定性を厳密に強制する必要があります。 これを行うためのバージョン管理メカニズムはありませんし、 gobに可能性のある不一致を通知するメカニズムもないことに注意してください。 

### HTTP

```console
go run HTTPServer/main.go
```

```console
go run HTTPClient/main.go localhost
Arith: 17 * 8 = 136
Arith: 17 / 8 = 2 remainder 1
```

### TCP

```console
go run TCPServer/main.go localhost:1234
```

```console
go run TCPClient/main.go localhost:1234
Arith: 17 * 8 = 136
Arith: 17 / 8 = 2 remainder 1
```


[JSONRPC](http://tumregels.github.io/Network-Programming-with-Go/rpc/json.html)
