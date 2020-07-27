# GOSS
List of golang utils collections that share across my projects
Beucause I'm to tired to copy from one to another

## Shutdown Hook

```go
package main

import (
    "github.com/hienduyph/goos/utils/shutdowns"
)

func main() {
    ctx := shutdowns.NewCtx()
    // shared this ctx across your application.
}

```

