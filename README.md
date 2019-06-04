# Go-CORS

This is a module to handle CORS for Go web API or Web Server.

## Usage

```go
import goCors "github.com/mattdamon108/go-cors"

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", goCors.Set(&GraphQL{Schema: schema}))

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Listening to... port 8080")
	if err = s.ListenAndServe(); err != nil {
		panic(err)
	}
}

type GraphQL struct {
	Schema *graphql.Schema
}

func (h *GraphQL) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// ...

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
```

# Next to do

- [ ] Add the whitelist
