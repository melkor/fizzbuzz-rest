#fizzbuzz-rest

FizzBuzz-rest is a simple fizz-buzz REST server. This server exposes a REST API endpoint that: Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2. Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

=))

## Installation

This module build a docker image that can be launch with docker-compose. 

You must install this module in your GOPATH.
``` shell
> go get -u github.com/melkor/fizzbuzz-rest

```

``` shell
> cd ${GOPATH}/src/github.com/melkor/fizzbuzz-rest/
> sudo docker build --tag=fizzbuzz-rest .
> sudo docker-compose up
```

## Exemples
``` shell
> curl 'localhost:8000/fizzbuzz?int1=3&int2=7&limit=25&str1=fizz&str2=fuzz'
["1","2","fizz","4","5","fizz","fuzz","8","fizz","10","11","fizz","13","fuzz","fizz","16","17","fizz","19","20","fizzfuzz","22","23","fizz","25"]
>
```

``` shell
> curl localhost:8000/mostFrequentRequest
"int1=3\u0026int2=7\u0026limit=25\u0026str1=fizz\u0026str2=fuzz"
>
```


<!--
  506  docker run -p 8000:8000 fizzbuzz-rest
  507  docker build --tag=fizzbuzz-rest .
  508  docker build --tag=fizzbuzz-rest .
  509  docker run -p 8000:8000 fizzbuzz-rest
  510  docker build --tag=fizzbuzz-rest .
  511  docker run -p 8000:8000 fizzbuzz-rest
  512  docker build --tag=fizzbuzz-rest .
  513  docker run -p 8000:8000 fizzbuzz-rest
  514  docker build --tag=fizzbuzz-rest .
  515  docker run -p 8000:8000 fizzbuzz-rest
  516  docker-compose up
  517  curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
  518  chmod +x /usr/local/bin/docker-compose
  519  docker-compose --version
  520  docker-compose up
  521  docker-compose down
-->
