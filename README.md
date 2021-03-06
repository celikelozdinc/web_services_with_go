
## Install mysql driver
```bash
go get -u -v github.com/go-sql-driver/mysql
```

## Install node.js
[Installation directives](https://muhammetkucuk.com/install-node-js-from-linux-tar-gz-file/)
[Installation](https://github.com/nodesource/distributions/blob/master/README.md)

## Install angular CLI for building project && install dependencies
```bash
cd inventory-mgmt
npm install -g @angular/cli
npm install
```


## Run Application

```bash
ng serve --open
```


## Testing Basic HTTP Handlers

### bar endpoint (*`GET`*)
![bar](./img/bar.png)

### baz endpoint (*`GET`*)
![baz](./img/baz.png)

### products endpoint (*`GET`*)
![products](./img/products.png)

### product/<id> endpoint (*`GET`*)
![path parameters](./img/url_parameters.png)
