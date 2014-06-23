# jc [![Build Status](https://travis-ci.org/hhatto/jc.png?branch=master)](https://travis-ci.org/hhatto/jc)

jenkins cli

![jc](https://dl.dropboxusercontent.com/u/26471561/img/jc.png)

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

![jc staus](https://dl.dropboxusercontent.com/u/26471561/img/jc_status.png)

## License
MIT
