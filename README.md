# Go-CORS

This is a module to handle CORS for Go web API or Web Server.

## Usage

```go
mux.Handle("/", goCors.Set(http.Handler))
```

> NOTE : Be aware of setting CORS `*` for `dev` mode only, in case of `prod`, BE careful it will allow any requests from any origin.

## Example

```go
import (
	"log"
	"net/http"

	goCors "github.com/mattdamon108/go-cors"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", goCors.Set(Index{}))

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Listening to... port 8080")
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

type Index struct{}

func (i Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := "This is CORS test on path: " + r.URL.Path
	w.Write([]byte(path))
}
```

# Next to do

- [ ] Add the whitelist
