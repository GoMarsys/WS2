package main

import (
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// r.Body  => return the Body IO Object
	//
	// return http Error in the response writer with the given status code
	// http.Error(w, err.Error(), http.StatusInternalServerError)
	io.WriteString(w, "OK")
}

func main() {

	// ServeMux is an HTTP request multiplexer.
	// It matches the URL of each incoming request against a list of registered
	// patterns and calls the handler for the pattern that
	// most closely matches the URL.
	//
	// Patterns name fixed, rooted paths, like "/favicon.ico",
	// or rooted subtrees, like "/images/" (note the trailing slash).
	// Longer patterns take precedence over shorter ones, so that
	// if there are handlers registered for both "/images/"
	// and "/images/thumbnails/", the latter handler will be
	// called for paths beginning "/images/thumbnails/" and the
	// former will receive requests for any other paths in the
	// "/images/" subtree.
	//
	// Note that since a pattern ending in a slash names a rooted subtree,
	// the pattern "/" matches all paths not matched by other registered
	// patterns, not just the URL with Path == "/".
	//
	// If a subtree has been registered and a request is received naming the
	// subtree root without its trailing slash, ServeMux redirects that
	// request to the subtree root (adding the trailing slash). This behavior can
	// be overridden with a separate registration for the path without
	// the trailing slash. For example, registering "/images/" causes ServeMux
	// to redirect a request for "/images" to "/images/", unless "/images" has
	// been registered separately.
	//
	// Patterns may optionally begin with a host name, restricting matches to
	// URLs on that host only. Host-specific patterns take precedence over
	// general patterns, so that a handler might register for the two patterns
	// "/codesearch" and "codesearch.google.com/" without also taking over
	// requests for "http://www.google.com/".
	//
	// ServeMux also takes care of sanitizing the URL request path,
	// redirecting any request containing . or .. elements or repeated slashes
	// to an equivalent, cleaner URL.
	mux := http.NewServeMux()

	// HandleFunc registers the handler function for the given pattern.
	mux.HandleFunc("/", handler)

	log.Println("starting server")

	// ListenAndServe listens on the TCP network address addr
	// and then calls Serve with handler to handle requests
	// on incoming connections.
	// Accepted connections are configured to enable TCP keep-alives.
	// Handler is typically nil, in which case the DefaultServeMux is
	// used.
	//
	// A trivial example server is:
	//
	//	package main
	//
	//	import (
	//		"io"
	//		"net/http"
	//		"log"
	//	)
	//
	//	// hello world, the web server
	//	func HelloServer(w http.ResponseWriter, req *http.Request) {
	//		io.WriteString(w, "hello, world!\n")
	//	}
	//
	//	func main() {
	//		http.HandleFunc("/hello", HelloServer)
	//		log.Fatal(http.ListenAndServe(":12345", nil))
	//	}
	//
	// ListenAndServe always returns a non-nil error.
	log.Fatal(http.ListenAndServe(":8080", mux))
}
