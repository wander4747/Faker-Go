# Faker GO

[![Test](https://github.com/wander4747/faker-go/actions/workflows/test.yml/badge.svg)](https://github.com/wander4747/faker-go/actions/workflows/test.yml)
[![Lint](https://github.com/wander4747/faker-go/actions/workflows/lint.yml/badge.svg)](https://github.com/wander4747/faker-go/actions/workflows/lint.yml)

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
	name := f.Person().FullName()
	fmt.Println(name)
}
```