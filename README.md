# beacon-model
Beacon model houses all the models used in beacon.

### time.Time Type in models
Make sure all `time.Time` types are pointers in the model struct

Example:

```go
type Example struct {
  AccessDate *time.Time
  Created *time.Time
}
```

### Auto tagging with hook

We want to auto tag when w