# One-Time Secret CLI

OTS is a command-line interface (CLI) for the [onetimesecret](https://onetimesecret.com/) service.

## Installation

You will need to have Go installed, install the basic binary with the command:

```bash
go get -u github.com/olaysco/go-ots
```

Using the [onetimesecret](https://onetimesecret.com/) service requires setting up an account, create an account [here](https://onetimesecret.com/signup) if you don't have one.

Create a new file in the folder name `ots.yaml`, with the following content while replacing the username and password with your credentials.
```
username: YOUR_USERNAME
password: YOUR_PASSWORD
```

## Usage
To get the list of available commands:
```bash
ots --help
```
To get the status of the [onetimesecret](https://onetimesecret.com/) service:
```
ots status
```
To share a new secret: 
```
ots share [flags]

Flags:
      --passphrase string   Key that the recipient must know to view the secret
      --recipient string    Email address to send the secret link to
      --secret string       Secret to share
      --ttl uint            The maximum amount of time, in seconds, that the secret should survive (default 60) (default 60)
```
To retrieve a previously created secret:
```
ots retrieve [flags]

Flags:
      --passphrase string   Key that the recipient must know to view the secret
      --secret_key string   Secret key
```

## Dependencies
Dependencies are managed using  [go modules](https://golang.org/ref/mod) and require Go 1.11+ with GO111MODULE=on.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)