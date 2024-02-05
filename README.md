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

## GorillaMux

```go
import (
    "github.com/gorilla/mux"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
 vars := mux.Vars(r)
 w.WriteHeader(http.StatusOK)
 fmt.Fprintf(w, "ID: %v\n", vars["id"])
}

func ArticleCategoryHandler(w http.ResponseWriter, r *http.Request) {
 vars := mux.Vars(r)
 w.WriteHeader(http.StatusOK)
 fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func main() {
 r := mux.NewRouter()

 // r.HandleFunc("/articles/{category}/", ArticleCategoryHandler)
 r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler).Name("article")

 //SubRouting => sub-routing and prefixing path( r.PathPrefix("/articles"))
 s := r.PathPrefix("/articles").Subrouter()

 // can contain(articles/{books}, articles/tech, articles/tech/123)
 s.HandleFunc("/{category}/", ArticleCategoryHandler)


 //Custom Path => HandleFunc chaining is possible (HandlerFunc and not HandleFunc)
r.HandleFunc("/articles/{cate}/").HandlerFunc(ArticleCategoryHandler)

 // Reverse mapping(Registered URLs)
 url, err := r.Get("article").URL("category", "books", "id", "123")
 fmt.Printf("%v, %v\n", url, err)



 srv := &http.Server{
  Handler: r,
  Addr:    "127.0.0.1:8000",
  // standard pratices to enforce timeout
  WriteTimeout: 15 * time.Second,
  ReadTimeout:  15 * time.Second,
 }

 log.Fatal(srv.ListenAndServe())

}

```

> PathPrefix (should start with) use case is when serving static files from static folders

```go
r := mux.NewRouter()
r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",http.FileServer(http.Dir("/tmp/static"))))

```
