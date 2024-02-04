# Learning About Routers(Custom Mux, std ServerMux,httpRouter,GorillaMux)

this the stages before the shorturl program

## using a custom mux server

```go
package main
// std pkg

//struct for custom server mux
type CustomServerMux struct{}

 ServeHTTP implements http.Handler.
func (p *CustomServerMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

 if r.URL.Path == "/" {
  giveRandom(w, r)
  return
 }

 http.NotFound(w, r)
 return
}

func giveRandom(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintf(w, "Your random number is : %f", rand.Float64())
}

func main() {
 mux := &CustomServerMux{}
 http.ListenAndServe(":8080", mux)
}

```

## using of mux inbuild server multiplexer(router handler)

it handle the directions of all the routes in the go program

```go
func main() {
 newMux := http.NewServeMux()

 newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "%v", rand.Float64())
 })

 newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "%v", rand.Intn(100))
 })

 http.ListenAndServe(":8080", newMux)

}
```

## using (julienschmidth/httprouter)

we import and http router form julienschmidth/httprouter github and used it mux

```g
func getCommandOutput(command string, arguments ...string) string {
 // arg...(unpack)spreads the arguments of an array

 cmd := exec.Command(command, arguments...)
 var out bytes.Buffer
 var stderr bytes.Buffer
 cmd.Stdout = &out
 cmd.Stderr = &stderr

 err := cmd.Start()

 if err != nil {
  log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
 }

 err = cmd.Wait()
 if err != nil {
  log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
 }

 return out.String()
}

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
 fmt.Fprintf(w, getCommandOutput("go", "version"))
}

func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
 fmt.Fprintf(w, getCommandOutput("/bin/cat", params.ByName("name")))
}

func main() {

 router := httprouter.New()

 router.GET("/api/v1/go-version", goVersion)
 router.GET("/api/v1/show-file/:name", getFileContent)

 log.Fatal(http.ListenAndServe(":8080", router))

}
```
