# go-sapin
:christmas_tree: drawing a christmas tree in Go

[![GoDoc](https://godoc.org/github.com/moul/go-sapin?status.svg)](https://godoc.org/github.com/moul/go-sapin)
[![](https://img.shields.io/badge/appspot-sapin--as--a--service-blue.svg)](http://sapin-as-a-service.appspot.com/)

## Demo

http://sapin-as-a-service.appspot.com/

## Usage

```console
$ sapin --size=3 --balls=4 --star --emoji
```
![](https://raw.githubusercontent.com/moul/go-sapin/master/assets/sapin-size3-balls4-star-emoji.png)

---

```console
$ sapin --size=3 --balls=4 --star --color
```
![](https://raw.githubusercontent.com/moul/go-sapin/master/assets/sapin-size3-balls4-star-color.png)

---

```console
$ sapin --size=1
   *
  ***
 *****
*******
   |
```

---

```console
$ sapin --size=2
      *
     ***
    *****
   *******
    *****
   *******
  *********
 ***********
*************
     |||
     |||
```

---

```console
$ sapin --balls=4
                   *
                  ***
                 *****
                ****@**
                 *****
                *******
               *********
              ***********
             *************
              *****@**@**
             **********@**
            ********@******
           *@***************
          ****@**@***********
         *********************
           *@************@**
          *********@*********
         ***@**********@******
        ****************@******
       @************************
      @**************************
     *****************************
       *******@*****************
      ***********@***************
     *****************************
    *************************@*****
   *******************@*************
  **********@************************
 ***********************************@*
*****************@*********************
                 |||||
                 |||||
                 |||||
                 |||||
                 |||||
```

---

```console
$ sapin -h
Usage:
  sapin [OPTIONS]

Application Options:
  -s, --size= Size of the sapin

Help Options:
  -h, --help  Show this help mess
```

## Install

```console
$ go get github.com/moul/go-sapin/cmd/sapin
```

## License

MIT
