# Api


```kill -9 $(lsof -t -i:8080)``` // kill a process running at port 


//install golang
```
wget https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz
sudo tar -xvf go1.12.2.linux-amd64.tar.gz
sudo mv go /usr/local
echo 'export GOROOT=/usr/local/go' >>~/.profile
echo 'export GOPATH=$HOME/go_projects' >>~/.profile
echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >>~/.profile
source ~/.profile
```
check go version

``` go version```

----------------------------------------------------------------------------------------------------------------------
Anywhere ```mkdir Api```


cd Api



```go mod init com.api.rest``` (creates go.mod shows all dependencies)



```code .```


write code with dependency



1.```go mod tidy``` (install all dependencies and creates go.sum)



2.```go mod vendor``` (shows all dependencied)


to download specific verison do ```go get``` automatically install that version and update in go mod  example -```go get github.com/spf13/cast```


now packages will also be visible like ```"../config"```   or will user ```com.api.rest/config"``` auto

for every new local or git hub package run 1 &  2 then the package will be visible
