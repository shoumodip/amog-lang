# Amog Lang
A parody language based on [Amogus](https://www.innersloth.com/games/among-us/).

## Quick Start
```console
$ go build amog.go
$ ./amog sample.amog
```

## Syntax
A sequence of lines of the form `<ACTION> <TARGET>`.

### `imposter NAME`
```conf
imposter a # var a = 0
```

### `teammate NAME`
```conf
imposter a # const a = 0
```

### `yell NAME`
```conf
yell a # println(a)
```

### `sus NAME`
```conf
sus a # a = a - 1
```

### `trust NAME`
```conf
trust a # a = a + 1
```

### `eject NAME`
```conf
eject a # delete a
```
