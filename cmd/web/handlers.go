package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strconv"
)

// Define a home handler function which writes a byte slice containing
// register the home function as the handler for the "/" URL pattern.
func home(w http.ResponseWriter, r *http.Request) {
    // Check if the current request URL path exactly matches "/". If it doesn't, use
    // the http.NotFound() function to send a 404 response to the client.
    // Importantly, we then return from the handler. If we don't return the handler
    // would keep executing and also write the "Hello from SnippetBox" message.

    if r.URL.Path != "/" {
        http.NotFound(w, r)
	return
    }

    // Initialize a slice containing the paths to the two files. It's important
    // to note that the file containing our base template must be the *first*
    // file in the slice.
    files := []string{
        "./ui/html/base.tmpl.html",
	    "./ui/html/partials/nav.tmpl.html",
	    "./ui/html/pages/home.tmpl.html",
    }
    
    // Use the template.ParseFiles() function to read the template file into a
    // template set. If there's an error, we log the detailed error message and use
    // the http.Error() function to send a generic 500 Internal Server Error
    // response to the user.
    //w.Write([]byte("Hello from Snippetbox"))
    ts, err := template.ParseFiles(files...)
    if err != nil {
        log.Println(err.Error())
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	return
    }

    // Use the ExecuteTemplate() method to wtite the content of the "base"
    // template as the response body.
    err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
        log.Println(err.Error())
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

// Add a snippetView handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
    // Extract the value of the id parameter from the query string and try to
    // convert it to an integer using the strconv.Atoi() function. If it can't
    // be converted to an integer, or the value is less than 1, we return a 404
    // page not found response.
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
	return
    } 

    // Use the fmt.Fprintf() function to interpolate the id value with our response
    // and write it to the http.ResponseWriter.
    //w.Write([]byte("Display a specific snippet..."))
    fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a snippetView handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
    // Use r.Method to check whether the request is using POST or not.
    if r.Method != http.MethodPost {
        // If it's not use the w.WriteHeader() method to send a 405 status
	// code and the w.Write() method to write a "Method Not Allowed"
	// response body. We then return from the function so that the
	// subsequent code is not executed.
	w.Header().Set("Allow", http.MethodPost)
	//w.WriteHeader(405)
	//w.Write([]byte("Method Not Allowed"))
	// Use the http.Error() function to send a 405 status code and "Method Not
	// Allowed"string as the response body
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return
    }

    w.Write([]byte("Create a new snippet..."))
}

