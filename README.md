# jc [![Build Status](https://travis-ci.org/hhatto/jc.png?branch=master)](https://travis-ci.org/hhatto/jc)

jenkins cli

![jc](https://i.gyazo.com/eedd8d82131d080fd0eae3dd02ac00d8.png)

## Installation
```
$ go get github.com/hhatto/jc
```

## Usage

### configuration
```sh
$ jc conf https://ci.jenkins-ci.org/
$ jc conf -n mydomain --add http://jenkins.mydomain.com/
```

jc's config file saves in `$HOME/.config/jc`

```sh
$ cat $HOME/.config/jc
{
  "hosts": [
    {
      "name": "default",
      "hostname": "https://ci.jenkins-ci.org/"
    },
    {
      "name": "mydomain",
      "hostname": "http://jenkins.mydomain.com/"
    }
  ]
}
```

### print job status
```sh
$ jc jobs
default - http://jenkins.hexacosa.net/
 ✔  ☁  autopep8
 ✔  ☁  genzshcomp
 ✔  ☁  gruffy
 ✔  ☁  pgmagick
```

### print server status
```sh
$ jc status -n mydomain
```

![jc staus](https://i.gyazo.com/c719fdb2bf487a025a9f4f8e2a2567f3.png)

## License
MIT
