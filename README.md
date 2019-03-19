# Lumen

Lumen is a CLI that allows you tou generate API with [base-api](https://github.com/adrien3d/base-api) framework, generating models, controllers, router and store.


### Prerequisites

Define `export PATH=$PATH:$GOPATH/bin`

### Installing

`git clone https://github.com/adrien3d/lumen.git` where you want to have it.

`cd lumen`

`go install`

## Usage

* Generate a first model with `lumen model`
* For every generation, you have to set the correct namespace, for example `github.com/adrien3d/base-api`

### Auto mode
* Select entities for each selected model and generate controller, store and router with `lumen generate`

### Manual mode
* Then, generate the corresponding controller with `lumen controller`
* Then, generate the corresponding store with `lumen store`
* Finally, generate the router with `lumen router`

## Built With

* [cobra](github.com/spf13/cobra) - A Commander for modern Go CLI interactions.
* [ishell](github.com/abiosoft/ishell) - Library for creating interactive cli applications.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
* **Adrien Chapelet** - *Initial work* - [IoThings](https://github.com/adrien3d)

