# go-starwars
StarWars ascii animation written in Go.

# Support OS
- Mac
- Linux
- Windows

# Requirements
- Go 1.11.4~  
- github.com/gobuffalo/packr

# Installtion
Use Go

```
git clone https://github.com/skanehira/go-starwars.git
cd go-starwars
GO111MODULE=on packr install
```

Use Docker
```
docker run --rm -it skanehira/go-starwars
```

# Usage
```
# run 
$ go-starwars

# options
$ go-starwars -h
Usage of go-starwars:
  -s int
        speed, up speed by making it smaller (default 40)
```

# Credits
- Original art work : Simon Jansen (http://www.asciimation.co.nz/)
- Inspiration: Martin W. Kirst (ascii-telnet-server)
- Stream ASCII movies over HTTP: (ascii-tv)
