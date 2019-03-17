# Lumen

Lumen is a CLI that allows you tou generate API with [base-api](https://github.com/adrien3d/base-api) framework, generating models, controllers, router and store.


### Prerequisites

Define `export PATH=$PATH:$GOPATH/bin`

### Installing

`git clone https://github.com/adrien3d/lumen.git` where you want to have it.

`cd lumen`

`go install`

## Getting Started

Generate a first model with `lumen model`

Then, generate the corresponding store with `lumen store`

Then, generate the corresponding controller with `lumen controller`

Finally, generate the router with `lumen router`

## Built With

* [cobra](github.com/spf13/cobra) - A Commander for modern Go CLI interactions.
* [ishell](github.com/abiosoft/ishell) - Library for creating interactive cli applications.

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

* **Adrien Chapelet** - *Initial work* - [IoThings](https://github.com/adrien3d)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

