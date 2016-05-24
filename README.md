# curry

curry is a command line tool for querying currency rates, which are provided by [OpenExchangeRates](https://openexchangerates.org/).


## Disclaimer

Like OpenExchangeRates service, exchange rates are provided for informational purposes only and do not constitute financial advice of any kind. Although every attempt is made to ensure quality, no guarantees are made of accuracy, validity, availability, or fitness for any purpose.


## Install

```
$ go get -u github.com/cwchentw/curry
```

This app is developed with Go version 1.6 and tested on both GNU/Linux and Mac.


## Usage

You need an app id from OpenExchangeRates to use this application. They provide a free plan with 1,000 queries per month, or buy an enhanced plan for more advanced features.

First, initialize basic config by answering some questions:

```
$ curry init
```

Then, update related data:

```
$ curry update
```

List available currencies:

```
$ curry list
```

Query the amount of selected currencies:

```
$ curry from EUR to USD 30
```

The data will be automatically updated after 24 hours.

You may set your base currency:

```
$ curry favor USD
```

Later, query the amount to your base currency:

```
$ curry from JPY 1000
```

You may view the usage of this app with `go doc curry`.


## Copyright

Copyright 2016 Michael Chen. All rights reserved.

The app is licensed under GPLv3.
