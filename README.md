# Prime generator Application

## Ubuntu 18.04 
This is a golang structure was cloned from https://github.com/golang-standards/project-layout. It is important to have well structured at the beginning of the project, otherwise we will  suffer from painful like improper import, etc... 
With Go 1.14 [`Go Modules`](https://github.com/golang/go/wiki/Modules) are finally ready for production. We should use unless we have a special reason for that.
# My current environment is
* MacOS Catalina 10.15.1
* Clang version 11.0.0
* Go 1.14.6

Originally, Build a simple CLI application. Consider using cobra [`Cobra`](https://github.com/spf13/cobra) a very powerful library to build cli.
secondly, Logging is critical part. Should pay attention first to make sure, We can track our mistake during development. Consider using [`Logrus`](https://github.com/sirupsen/logrus)

#### Code structure and design is based on [`Go moby`] project (https://github.com/moby/moby)


## How to build locally
```bash
go build github.com/vietnamz/project-layout/cmd/prime_cal
```

## How to run.
```mysql based
./prime_cal
```