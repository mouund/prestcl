> :warning: **Not meant to be used on production, need to be tested before ! USE ONLY ON TEST SHOP**: Be very careful here, it is in early stages ! Code improvments need to be done, tests... If I have time I will continue working on it, contribution is welcome !

# prestctl

Prestctl is a modern CLI for Prestashop administration.

## Installation

### Binary

Simply download the binary adapted to your platform, for Linux

```bash
$ wget https://github.com/mouund/prestctl/releases/download/v0.1.0/prestctl-linux-amd64
$ sudo mv prestctl-linux-amd64 /usr/local/bin/prestctl
$ prestctl --help
```

### Building it from source

```bash
$ git clone https://github.com/mouund/prestctl.git
$ cd prestctl
$ go build
$ mv prestctl /usr/local/bin
$ prestctl --help
```

## Usage

### General help

```bash
$ prestctl --help
```

```bash
$ prestctl --help
A simple prestashop CLI to administer your instance

Usage:
  prestctl [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  delete      Delete a ressource from the prestashop website
  describe    Describe a ressource from the prestahsop website
  get         Get a list of a ressource from the prestahsop website
  help        Help about any command

Flags:
  -h, --help   help for prestctl

Use "prestctl [command] --help" for more information about a command.
```
### API authentication

To authenticate to the prestashop API, use either --shopUrl and --token flags or create a file in 

`$HOME/.prestcl`

e.g:

```
name: <name of the shop>
shopUrl: <shop URL> (ex: http://localhost:8080)
token: <API token>
```

### Examples

#### Get resource

```bash
$ prestctl get products --shopUrl <your shop Url> --token <your API token>
```

If using `$HOME/.prestcl` 

```bash
$ prestctl get products 
```
is sufficient

```bash
$ prestctl get products

ID         Name                                     Prix      
1          Hummingbird printed t-shirt              23.900000 
2          Hummingbird printed sweater              35.900000 
16         Mountain fox notebook                    12.900000 
17         Brown bear notebook                      12.900000 
18         Hummingbird notebook                     12.900000 
6          Mug The best is yet to come              11.900000 
7          Mug The adventure begins                 11.900000 
8          Mug Today is a good day                  11.900000 
9          Mountain fox cushion                     18.900000 
10         Brown bear cushion                       18.900000 
11         Hummingbird cushion                      18.900000 
15         Pack Mug + Framed poster                 35.000000 
19         Customizable mug                         13.900000 
3          The best is yet to come' Framed poster   29.000000 
5          Today is a good day Framed poster        29.000000 
12         Mountain fox - Vector graphics           9.000000  
13         Brown bear - Vector graphics             9.000000  
14         Hummingbird - Vector graphics            9.000000
```
Check help for get to have all possibilities

```bash
$ prestctl get --help
Get a list of a ressource from the prestahsop website

Usage:
  prestctl get [flags]
  prestctl get [command]

Available Commands:
  carriers    Get a list of carriers of the prestashop instance
  categories  Get a list of categories of the prestashop instance
  customers   Get a list of customers of the prestasop instance
  orders      Get a list of orders of the prestashop instance
  products    Get a list of products from your prestashop instance

Flags:
  -h, --help   help for get

Use "prestctl get [command] --help" for more information about a command.
```
#### Delete a resource 

Use --id flag

```bash
$ prestctl delete product --id 1
Product deleted successfully

```
#### Describe a product

```bash
$ prestctl describe product --id 2
details:
  id: 2
  idmanufacturer: "1"
  idsupplier: "0"
  idcategorydefault: "5"
  new: null
  cachedefaultattribute: "9"
...
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.
The CLI is not yet meant for production please use it on test environment.
