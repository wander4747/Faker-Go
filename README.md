# Faker GO

Simples lib para gerar dados falsos

## Instalação
```sh
go get  github.com/wander4747/faker-go
```

## Como usar

```go
package main

import (
	"fmt"
	"github.com/wander4747/faker-go"
)

func main() {
	f := faker.New("pt_BR")
	name := f.Name().FullName()
	fmt.Println(name)
}

```