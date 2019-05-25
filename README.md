# beacon-model
Beacon model houses all the models used in beacon.

## Some useful and must do when using GRPC and Go
Because we're using grpc and golang and at a point we want to tranform data from grpc to golang and vice versa, we ensure the field names are the same.

DO NOT follow go-lint when naming fields of structs. Ensure that all fields are *UpperCamelCase* as possible
### time.Time Type in models
Make sure all `time.Time` types are pointers in the model struct

Example:

```go
type Example struct {
  AccessDate *time.Time
  Created *time.Time
}
```
