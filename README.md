# gobel-client-example
[![GitHub license](https://img.shields.io/github/license/bmf-san/gobel-client-example)](https://github.com/bmf-san/gobel-client-example/blob/master/LICENSE)
[![CircleCI](https://circleci.com/gh/bmf-san/gobel-client-example.svg?style=svg)](https://circleci.com/gh/bmf-san/gobel-client-example)

Gobel is a headless cms built with golang. 
This is an example for client application of gobel.

# gobel
- [gobel-api](https://github.com/bmf-san/gobel-api)
- [gobel-admin-client](https://github.com/bmf-san/gobel-admin-client)
- [gobel-client-example](https://github.com/bmf-san/gobel-client-example)
- [gobel-example](https://github.com/bmf-san/gobel-example)

# Requirements
- Golang1.13
- Docker Compose

# Get started
Before you start, you need to clone [gobel-api](https://github.com/bmf-san/gobel-api).

`cp .env_example .env`

`make docker-compose-build`

`make docker-compose-up` or `make docker-compose-up-d`
Add hosts to `/etc/hosts`.

```
127.0.0.1 gobel-client-example.local
```

# Contributing
We welcome your issue or pull request from everyone.
Please make sure to read the [CONTRIBUTING.md](https://github.com/bmf-san/gobel-client-example/.github/CONTRIBUTING.md).

# License
This project is licensed under the terms of the MIT license.

# Author
bmf - Software engineer.

- [github - bmf-san/bmf-san](https://github.com/bmf-san/bmf-san)
- [twitter - @bmf-san](https://twitter.com/bmf_san)
- [blog - bmf-tech](http://bmf-tech.com/)
