# Package Reflect

Package reflect di Golang adalah library yang digunakan untuk memeriksa dan memanipulasi tipe data serta nilai dari suatu variabel pada runtime. Package reflect biasanya digunakan dalam kasus pemrosesan JSON, database, atau komunikasi antar aplikasi yang menggunakan format data variative

## `reflect.TypeOf()`

Digunakan untuk mengecek tipe data yang masuk melalui interface. Berikut contohnya:

```go
package main

import (
    "fmt"
    "reflect"
)

// Cek tipe data dengan reflect.TypeOf
func checkType(value interface{}) {
    t := reflect.TypeOf(value)
    fmt.Println("Type of value:", t)
}

func main() {
    checkType(42)
    checkType("lumoshive")
}
```

output:

```bash
Type of value: int
Type of value: string
```

## `reflect.TypeOf()` with function name, `NumFiled` and `field`

Digunakan untuk mengakses Informasi Object seperti nama object, index field dan field itu sendiri. Berikut contohnya:

```go
package main

import (
    "fmt"
    "reflect"
)

// access information object
type Person struct {
    Name string
    Age int
}

func structInfo(p interface{}) {
    t := reflect.TypeOf(p)
    fmt.Println("Struct name:", t.Name())

    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fmt.Printf("Field %d: %s (Type: %s)\n", i, field.Name, field.Type)
    }
}

func main() {
    p := Person{Name: "Aris", Age: 30}
    structInfo(p)
}
```

output:

```bash
Struct name: Person
Field 0: Name (Type: string)
Field 0: Age (Type: int)
```

## `reflect.ValueOf()`

Digunakan mengecek nilai dari sebuah data interface. Berikut contohnya:

```go
package main

import (
    "fmt"
    "reflect"
)

// check value with reflect valueof
func checkValue(value interface{}) {
    v := reflect.ValueOf(value)
    fmt.Println("Value of variable:", v)
}

func main() {
    checkValue(40)
    checkValue("lumoshive academy")
}
```

output:

```bash
Value of variable: 40
Value of variable: lumoshive academy
```

## `reflect.Valueof()` with function `MethodByName`

Digunakan untuk memanggil method yang dimiliki suatu object. Berikut contohnya:

```go
package main

import (
    "fmt"
    "reflect"
)

type User struct{}

// call method
func (t User) SayHello() {
    fmt.Println("Hello from User!")
}

func callMethodByName(obj interface{}, methodName string) {
    v := reflect.ValueOf(obj)
    method := v.MethodByName(methodName)

    if method.IsValid() {
        method.Call(nil)
    } else {
        fmt.Println("Method not found")
    }
}

func main() {
    p := User{}
    callMethodByName(p, "SayHello")
    callMethodByName(p, "Greet")
}

```

output:

```bash
Hello from user!
Mthod not found
```
