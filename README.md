# Welcome to Teaset
![Version](https://img.shields.io/badge/version-0.0.1-blue.svg?cacheSeconds=2592000)
[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)



### [Homepage](https://tomkys144.github.io/teaset/)

## Roadmap

- [ ] fully automated setup
- [ ] Displays fully autonomous environment, which is user-friendly for non-programmeres in which users can:
	- [ ] Create issues
	 - [ ] Communicate with developers
	- [ ] Create feature requests
	- [ ] Upload files linked to the issue
- [ ] receives emails and telegram messages and creates issues from them

## Instalation

```sh
go install github.com/tomkys144/teaset
```
now you can choose two options:

### A) Interactive setup

```sh
./$GOPATH/bin/teaset setup --interactive
```
here you get all options asked in command linked

### B) Automated setup

```sh
wget https://raw.githubusercontent.com/tomkys144/teaset/dev/.env.example -O /path/to/.env.example
```
now edit values in `/path/to/.env.example`
```sh
teaset setup -s /path/to/.env.example -d /path/to/prefered/installation/directory
```

## Author

**Tomáš Kysela**

* Website: [tkysela.cz](https://tkysela.cz)
* Twitter: [@tomkys144](https://twitter.com/tomkys144)
* Github: [@tomkys144](https://github.com/tomkys144)
* LinkedIn: [@tomáš-kysela-7078aa1b8](https://linkedin.com/in/tomáš-kysela-7078aa1b8)


## License

Copyright © 2021 [Tomáš Kysela](https://github.com/tomkys144).

This project is [BSD 3 Clause](https://spdx.org/licenses/BSD-3-Clause.html) licensed.
