# Template

## How to

```console
go run main.go
The name is jan. The age is 50.
An email is &lt;h1&gt;jan@newmarch.name&lt;/h1&gt;An email is jan.newmarch@gmail.com
An employer is Monashand the role is HonoraryAn employer is Box Hilland the role is Head of HE
```

## Custome func

```console
go run customeFunc/main.go
The name is jan.

        An email is "jan at newmarch.name"

        An email is "jan.newmarch at gmail.com"
```

## Var

```console
go run var/main.go


    Name is jan, email is jan@newmarch.name

    Name is jan, email is jan.newmarch@gmail.com
```

# if

```console
go run if/main.go | jq
{
    "Name": "jan",
      "Emails": [
          "jan@newmarch.name",
          "jan.newmarch@gmail.com"
      ]
  }
}
```

# Sequence

```console
go run sequence/main.go
a, b, c, d, e, f
a in black, b in white, c in red, d in black, e in white, f in red
```
