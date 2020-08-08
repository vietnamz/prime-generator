# Prime generator Application

This is a golang structure was cloned from https://github.com/golang-standards/project-layout. It is important to have well structured at the beginning of the project, otherwise we will  suffer from painful like improper import, etc... 
With Go 1.14 [`Go Modules`](https://github.com/golang/go/wiki/Modules) are finally ready for production. We should use unless we have a special reason for that.
# My current environment is
* MacOS Catalina 10.15.1
* Clang version 11.0.0
* Go 1.14.6

# First day.

Originally, Build a simple CLI application. Consider using cobra [`Cobra`](https://github.com/spf13/cobra) a very powerful library to build cli.
secondly, Logging is critical part. Should pay attention first to make sure, We can track our mistake during development. Consider using [`Logrus`](https://github.com/sirupsen/logrus)

#### Code structure and design is based on [`Go moby`] project (https://github.com/moby/moby)
## result for first day.
spent 12h to work on, and now I was able to call the api to 127.0.0.1/ping. return "OK"

## Issue.
* spend a lot of time to resolve the issue related to. lesson leart is really careful with dependency. Need to learn more about this area.

```bash
go: found github.com/Sirupsen/logrus in github.com/Sirupsen/logrus v1.6.0
go: github.com/docker/docker/pkg/term imports
        github.com/docker/docker/pkg/term/windows imports
        github.com/Sirupsen/logrus: github.com/Sirupsen/logrus@v1.6.0: parsing go.mod:
        module declares its path as: github.com/sirupsen/logrus
                but was required as: github.com/Sirupsen/logrus
```

https://github.com/sensu/sensu-go/issues/1261





## How to run for the first day.
```bash
go build github.com/vietnamz/project-layout/cmd/prime_cal
```

## How to run.
```mysql based
./prime_cal
```

# Note: This app only support for unix platform. Window has never tested out.