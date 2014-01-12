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
$ jc conf -n mydomain http://jenkins.mydomain.com/
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
jenkinsci - https://ci.jenkins-ci.org/
version: 1.539-SNAPSHOT (rc-11/11/2013 15:37 GMT-kohsuke)
server: Jetty(8.y.z-SNAPSHOT)
```

## License
MIT
