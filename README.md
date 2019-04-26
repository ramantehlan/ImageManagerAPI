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
  - [How to work](#how-to-work)
  - [Build](#build)
- [Contribution](#contribution)
- [Resources](#resources)
- [Gallery](#gallery)
- [License](#license)

## About
  - Long introduction
  - Status
## Usage
#### Installation
What and how to install
#### Configure
What and how to configure

#### Commands

What and how to use the command

## Development

#### Pre-Requisites

#### Development Environment

sudo apt-get install mysql-server
http://go-database-sql.org/accessing.html

CREATE TABLE images (
  id INT NOT NULL UNIQUE AUTO_INCREMENT,
  image VARCHAR(255) NOT NULL,
  thumbnail VARCHAR(255) NOT NULL,
  caption text NULL,
  PRIMARY KEY (id),
  UNIQUE (id)
);

INSERT INTO images VALUES(null,'img/1.png','img/thumb-1.png',"This is a sample image");

#### File Structure

```console
.
├── .gitignore
├── LICENSE
├── README.md
├── app
│   ├── index.html
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

#### How to work

#### Build

## Contribution

 Your contributions are always welcome and appreciated. Following are the things you can do to contribute to this project.

1. **Report a bug** <br>
If you think you have encountered a bug, and I should know about it, feel free to report it [here]() and I will take care of it.

2. **Request a feature** <br>
You can also request for a feature [here](), and if it will viable, it will be picked for development.  

3. **Create a pull request** <br>
It can't get better then this, your pull request will be really appreciated by the community. You can get started by picking up any open issues from [here]() and make a pull request.

> If you are new to open-source, make sure to check read more about it [here](https://www.digitalocean.com/community/tutorial_series/an-introduction-to-open-source) and learn more about creating a pull request [here](https://www.digitalocean.com/community/tutorials/how-to-create-a-pull-request-on-github).

## Resources

- [pkg:sirupsen/logrus](https://github.com/sirupsen/logrus)
- [pkg:spf13/viper](https://github.com/spf13/viper)
- [pkg:json-iterator/go](https://github.com/json-iterator/go)

## Gallery
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
