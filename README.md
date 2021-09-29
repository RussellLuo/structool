# structool

A codec for Go structs with support for chainable encoding/decoding hooks.


## Features

1. Provide a uniform codec by combining [mapstructure][1] and [structs][2].
2. Make encoding/decoding hooks chainable.


## Installation

```bash
$ go get -u github.com/RussellLuo/structool
```


## Why?!

1. Why to use `structs`
   
    `mapstructure` has limited support for decoding structs into maps ([issues/166][3] and [issues/249][4]).

2. Why to make a fork of `fatih/structs`

    [fatih/structs][5] has been archived, but it does not support encoding hooks yet.

3. Why chainable hooks may be useful

    Both `mapstructure` and `structs` support hooks in the form of a single function, and this forces us to couple various conversions together.
   
    Chainable hooks (much like HTTP middlewares), on the other hand, promote separation of concerns, which makes individual hook reusable and composable.


## Documentation

Check out the [Godoc][6].


## License

[MIT](LICENSE)


[1]: https://github.com/mitchellh/mapstructure
[2]: https://github.com/RussellLuo/structs
[3]: https://github.com/mitchellh/mapstructure/issues/166
[4]: https://github.com/mitchellh/mapstructure/issues/249
[5]: https://github.com/fatih/structs
[6]: https://pkg.go.dev/github.com/RussellLuo/structool
