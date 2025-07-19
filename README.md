# envy

envy is a minimal go package that provides type-safe access to environment variables. envy avoids using `reflect` and instead uses generics for type assertion. Struct parsing is not included as it would require using reflect and the same functionality can be achieved with envy.Env() combined with a package such as `github.com/mitchellh/mapstructure`. 

## Examples
Getting all environment variabls:
```go
envs := envy.Env() // map[string]string
```
Getting a single environment variable:
```go
val, set, err = envy.Get[uint16]("PORT", envy.WithFallback(8080))
// if set:
//  val = 80 (uint16)
//  set = true
//  err = nil (errors happen only if there is a strcon parse error or the type is not supported )
// else:
//  val = 8080
//  set = false
//  err = nil
```