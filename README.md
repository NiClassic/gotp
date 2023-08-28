# gotp

gotp is a small cli tool that helps you generate TOTP-Tokens, typically used to secure your account for different services (2FA).

## Installation

To use this tool, you have to have [Go](https://go.dev/doc/install) installed. After that, clone the repository and build it using:

```bash
go build .
```

## Usage

```bash
gotp -s [step] -d [digits] <SECRET_KEY_FROM_SERVICE>
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)