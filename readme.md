### Prerequisite
#### To initialize a go project there must be a go.mod file
```
go mod init github.com/ab70/fiber-restapi
```

### To add a package 
```
go get github.com/gofiber/fiber/v2
```


### To run the project
```
go run server.go
```

#### For hot reload:
```
go get -u github.com/air-verse/air

#export this
export PATH=$(go env GOPATH)/bin:$PATH

#reload .zshc file:
source ~/.zshrc

#which air and then put that in:
nano ~/.zshrc
export PATH=$PATH:/home/ab/go/bin


```

### To build for binary use this:
```
# this will generate a birany in the current directory
 go build -o server server.go
```