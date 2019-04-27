# ImageManagerAPI

A simple API to manage images. :hash: :bridge_at_night:

# Index

- [About](#about)
- [Usage](#usage)
  - [Installation](#installation)
  - [Configure](#configure)
  - [Commands](#commands)
- [Development](#development)
  - [Pre-Requisites](#pre-requisites)
  - [Development Environment](#development-environment)
  - [File Structure](#file-structure)
  - [How to](#how-to)
  - [Build](#build)
- [Contribution](#contribution)
- [Resources](#resources)
- [Gallery](#gallery)
- [License](#license)

## About

ImageManagerAPI is a simple and small example of an API for storing images on the server and adding them to the database. It only performs `Post` and `Get` so once an image is uploaded, it can't be `Deleted` or `Updated`. This API can be controlled from a  single config file and it will also log all the activities in a log file.

## Usage

#### Installation

Follow steps given below to install ImageManagerAPI in your system.

1. Download the build from [here](https://github.com/ramantehlan/ImageManagerAPI/blob/master/build/ImageManagerAPI?raw=true).
2. Download the config file from [here](https://raw.githubusercontent.com/ramantehlan/ImageManagerAPI/master/build/config.yaml).
3. Place both the downloaded file in the same folder.

#### Configure

You need to make sure that you have `MySql` server installed on your system and has the following table in any database, also make sure to add database details in the config file in next step.

```sql
CREATE TABLE images (
  id INT NOT NULL UNIQUE AUTO_INCREMENT,
  image VARCHAR(255) NOT NULL,
  thumbnail VARCHAR(255) NOT NULL,
  caption text NULL,
  PRIMARY KEY (id),
  UNIQUE (id)
);
```

Now to start our API, we need to make sure that everything is configured properly. Below is the default configurations in `config.yaml`. You can make changes according to your need.

```yaml
# API Name
name: ImageManagerAPI
# Log folder location
# Do add '/' in the end
logAddr: log/
# Img folder location
# Do add '/' in the end
imgAddr: img/
# True if running in production
# False if running in development
production: false
# Database config
database:
  addr: 127.0.0.1
  name: ImageManagerAPI
  user: root
  pass: 1234567890
# Server config
server:
  port: 8082
  # Unit: Seconds
  readTimeout: 10
  writeTimeout: 10
  # Unit: MB
  maxHeaderBytes: 1

```

#### Commands

If you have installed and configured everything, you can start the API using the build you downloaded.

```console
$ ./ImageManagerAPI
```

Now, if you don't see any errors in your console then you can use [Postman](https://www.getpostman.com/), [curl](https://curl.haxx.se/) or your browser to test that the API, further you can use this [ImageManagerAPI App](https://github.com/ramantehlan/ImageManagerAPI/tree/master/app). Following are the endpoints currently available:

```go
{
	Route{"GET", "/"},
	Route{"GET", "/v1"},
	Route{"GET", "/v1/image/all"},
	Route{"GET", "/v1/image/{id}"},
	Route{"POST", "/v1/image/upload"},
}
```

## Development

If you are interested in adding more endpoints or in general playing with this project, you can read the following guidelines for the development.

#### Pre-Requisites

To develop and build this project, you need to have the following items installed on your system.

1. **Go**
```
# For Debian based
sudo apt-get install golang-go
```

2. **MySql server**  
```
# For Debian based
sudo apt-get install mysql-server
```

#### Development Environment

1. Fork this repo.
2. Clone the forked repo. `$ git clone https://github.com/YOUR_USERNAME/ImageManagerAPI`
3. Install all the dependencis. Run this command inside `src` folder. `$ go get ./...`

#### File Structure

```console
.
├── .gitignore
├── LICENSE
├── README.md
├── app
│   ├── index.html
│   ├── ss.png
│   ├── ss2.png
│   └── style.css
├── build
│   ├── config.yaml
│   └── ImageManagerAPI
└── src
    ├── config.go
    ├── config.yaml
    ├── create.go
    ├── database.go
    ├── logger.go
    ├── read.go
    ├── router.go
    ├── server.go
    ├── structure.go
    └── util.go
```

 No | File/Folder | Details
 ---|-------------|--------
  1 | ./app/... | This folder is to store the front end to test the API.
  2 | ./build/... | Contains the production build with a config file.
  3 | ./src/... | Go files are stored here.
  4 | config.go | Load the config file in the program.
  5 | config.yaml | config file for development
  6 | create.go | Handle all the `POST` requests. To store images in database and server.
  7 | database.go | Database connection is handled here.
  8 | logger.go | Log the activities on a file.
  9 | read.go | Handle all the `GET` requests.
  10 | router.go | Routes are handled and assigned.
  11 | server.go | Entry point to the API and server configurations.
  12 | structure.go | Store all the `struct` used in the program.
  13 | util.go | Other important functions are defined here.

#### How to

**Add Route**

If you need to add more routes, just add a `Route{Name, Method, Path, Handler}` object to routes variable. For example:

```go
Route{"NewRoute", "POST", "/v1/image/{key}", CreateImage}
```

**Log**

To log about different levels, you should use the function `Log(level, message)`. Log has 7 levels, which are `trace`, `debug`, `info`, `warn`, `error`, `fatal`, `panic`. For example:

```go
Log("info", "Endpoint Hit: ReadSingleImage")
```

**Connect To Database**

To access database anywhere, you should use the global variable `DB`. For example:

```go
DB.Query("SELECT * FROM images")
```

#### Build

The default build is for an `amd64` architecture, but if you want it for any other architecture, you can make your own build by using the command `$ go build` in `src` folder.


## Contribution

 Your contributions are always welcome and appreciated. Following are the things you can do to contribute to this project.

1. **Report a bug** <br>
If you think you have encountered a bug, and I should know about it, feel free to report it [here](https://github.com/ramantehlan/ImageManagerAPI/issues/new) and I will take care of it.

2. **Request a feature** <br>
You can also request for a feature [here](https://github.com/ramantehlan/ImageManagerAPI/issues/new), and if it will viable, it will be picked for development.  

3. **Create a pull request** <br>
It can't get better then this, your pull request will be really appreciated by the community. You can get started by picking up any open issues from [here](https://github.com/ramantehlan/ImageManagerAPI/issues) and make a pull request.

> If you are new to open-source, make sure to check read more about it [here](https://www.digitalocean.com/community/tutorial_series/an-introduction-to-open-source) and learn more about creating a pull request [here](https://www.digitalocean.com/community/tutorials/how-to-create-a-pull-request-on-github).

## Resources

- [MySql-Go-Driver](http://go-database-sql.org/accessing.html)
- [pkg:sirupsen/logrus](https://github.com/sirupsen/logrus)
- [pkg:spf13/viper](https://github.com/spf13/viper)
- [pkg:json-iterator/go](https://github.com/json-iterator/go)

## Gallery

![APP Frontend](https://raw.githubusercontent.com/ramantehlan/ImageManagerAPI/master/app/ss.png)
![API](https://raw.githubusercontent.com/ramantehlan/ImageManagerAPI/master/app/ss2.png)

## License

Copyright 2019 Raman Tehlan

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and
associated documentation files (the "Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the
following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial
portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT
LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN
NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
