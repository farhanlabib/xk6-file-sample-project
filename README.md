# xk6-file-sample-project

This project uses xk6, an extension system for k6, to supercharge your
performance testing scripts. With xk6, you can extend the core
functionality of k6 by adding custom modules. This project specifically
uses the file extension of xk6 to store data and use it for different
scripts or chaining APIs with that data.

## Prerequisites

-   Go
-   Git
-   xk6

## Installation

1.  Install Go from the official Go website at <https://go.dev/dl/>
2.  Add the following lines to your `~/.zshrc` file:

``` bash
export GOROOT="/usr/local/go"
export GOPATH="$HOME/Documents/yourProjectName”
export PATH="$HOME/Documents/yourProjectName/bin:$PATH"
```

3.  Reload the settings with the following command:

``` bash
source $HOME/.zshrc
```

4.  Install xk6 with the following command:

``` bash
go install go.k6.io/xk6/cmd/xk6@latest
```

5.  Build xk6 with the file extension:

``` bash
xk6 build v0.46.0 --with github.com/avitalique/xk6-file@latest
```

## Usage

The file extension allows us to use data stored in files for different
scripts or chaining APIs with that data. Here's a basic script to
illustrate its usage:

![image](https://github.com/farhanlabib/xk6-file-sample-project/assets/22010826/76467d29-68fd-4720-9b1b-8d741104ffda)



## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

