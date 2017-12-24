# Data Serialaisation

## ASN.1

```console
go run asn1/main.go
After marshal/unmarshal:  13
```

## ASN.1 Daytime

ASN.1エンコードされた文字列をやり取りする

```console
go run asn1DaytimeServer/main.go
```

```console
go run asn1DaytimeClient/main.go 127.0.0.1:1200
After marshal/unmarshal:  2017-12-24 21:45:30 +0900 JST
```

## JSON

```console
cd json && go run main.go && cat person.json | jq
{
  "Name": {
    "Family": "Newmarch",
    "Personal": "Jan"
  },
  "Email": [
    {
      "Kind": "home",
      "Address": "jan@newmarch.name"
    },
    {
      "Kind": "work",
      "Address": "j.newmarch@boxhill.edu.au"
    }
  ]
}
```

## JSON Echo

```console
go run jsonEchoServer/main.go
```


```console
go run jsonEchoClient/main.go 127.0.0.1:1200
Jan Newmarch
home: jan@newmarch.name
work: j.newmarch@boxhill.edu.au
Jan Newmarch
home: jan@newmarch.name
work: j.newmarch@boxhill.edu.au
Jan Newmarch
home: jan@newmarch.name
work: j.newmarch@boxhill.edu.au
Jan Newmarch
home: jan@newmarch.name
work: j.newmarch@boxhill.edu.au
Jan Newmarch
home: jan@newmarch.name
work: j.newmarch@boxhill.edu.au
Jan Newmarch
home: jan@newmarch.name
work: j.newmarch@boxhill.edu.au
Jan Newmarch
home: jan@newmarch.name
work: j.newmarch@boxhill.edu.au
Jan Newmarch
home: jan@newmarch.name
work: j.newmarch@boxhill.edu.au
Jan Newmarch
home: jan@newmarch.name
work: j.newmarch@boxhill.edu.au
Jan Newmarch
home: jan@newmarch.name
work: j.newmarch@boxhill.edu.au
```

## gob

```console
cd gobOnMemory && go run ../gob/main.go && go run main.go
Person Jan Newmarch
home: jan@newmarch.name
work: j.newmarch@boxhill.edu.au
```

## Base64

バイナリデータを文字列として符号化する  

```console
go run base64/main.go
AQIDBAUGBwg=
123456780000
```
