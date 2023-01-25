# govcr-sample

<h3> Setup Golang Project with Go Modules </h3>

* Create a new Project directory
* Open Terminal and `cd <project directory>`
* Run `go mod init <module name>`. For this project, we used `go mod init github.com/ksindhwani/govcr-sample`. This will create `go.mod` file.
* Now creata a main file `main.go`.
* Run `go get <dependency url>`. This will download the dependency in `GOPATH` directory. For this project we used `go get github.com/gorilla/mux`.
* Import this dependency in main.go file via `import "github.com/gorilla/mux"`.
* Run `go mod vendor`. This will create `go.sum` , `vendor folder` and `modules.txt` file.
* You may get error in `main.go` file that dependency not found . Run `go mod vendor` to download in vendor folder.

<h3> Setup </h3>

* Fill the following fields in `main.go` file 
  * <b> Api url /b> - API url to hit 
  * <b> Request Method </b> - Http Request Method 
  * <b> basic auth username </b> - Basic Auth Username - We use Basic Authentication in the API.
  * <b> <basic auth password </b> - Basic Auth Password
  
* Run the command `go run main.go`. This will call actual api first time and then create a Cassette `MyCassette1.json` and then for subsequent call they will use that Cassette.
