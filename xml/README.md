# XML

## Parse

```console
go run parse/main.go ./parse/sample.xml
person
  "
  "
  name
    "
    "
    family
      " Newmarch "
    family
    "
    "
    personal
      " Jan "
    personal
    "
  "
  name
  "
  "
  email
    "
    jan@newmarch.name
  "
  email
  "
  "
  email
    "
    j.newmarch@boxhill.edu.au
  "
  email
  "
"
person
"
"
```

## Unmarshal

```console
 go run unmarshal/main.go
 Family name: " Newmarch "
 Second email address: "
     j.newmarch@boxhill.edu.au
       "
```
