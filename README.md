# SSH as a microservice
An OMG service for SSH, it  is a cryptographic network protocol for operating network services securely over an unsecured network.

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)

## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```
### Test Service

* Test the service by following OMG commands

### CLI

##### SSH
```sh
<<<<<<< HEAD
$ omg run exec -a command=<COMMAND> -a username=<SERVER_USERNAME> -a password=<SERVER_PASSWORD> -e HOST=<SSH_HOST>  -e PRIVATE_KEY=<PRIVATE_KEY>
=======
$ omg run exec -a command=<COMMAND> -e username=<SERVER_USERNAME> -e password=<SERVER_PASSWORD> -e HOST=<SSH_HOST>  -e PRIVATE_KEY=<PRIVATE_KEY>
>>>>>>> 1c60e33b2d43f6f760987b9b6e4ea670575f3b76
```

## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Docker
### Build
```
docker build -t microservice-ssh .
```
### RUN
```
docker run -p 3000:3000 microservice-ssh
```
