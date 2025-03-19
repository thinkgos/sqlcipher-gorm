# GORM sqlcipher Driver

![CI](https://github.com/thinkgos/sqlcipher-gorm/workflows/CI/badge.svg)

## USAGE

```go
import (
  "github.com/thinkgos/sqlcipher"
  "gorm.io/gorm"
)

db, err := gorm.Open(sqlcipher.Open("gorm.db"), &gorm.Config{})
```

Checkout [https://gorm.io](https://gorm.io) for details.
