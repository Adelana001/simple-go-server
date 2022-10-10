package main 

import(
	"fmt"
	"log"
	"net/http"
)
	
func formHandler (w http.ResponseWriter , r *http.Request){
	if err := r.ParseForm() ; err != nil {
		fmt.Fprintf(w, "Parseform err: %v", err )
		return 
	}
	
	name := r.FormValue("name")
	adrress := r.FormValue("address")
	fmt.Fprintf(w,"POST request Succesful\n")
	fmt.Fprintf( w, "Name = %s\n", name )
	fmt.Fprintf(w, "Address =  %s\n", adrress)
}

func helloHandler (w http.ResponseWriter , r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w,"404 not found", http.StatusNotFound)
		return 
		} 
	if r.Method != "GET" {
		http.Error(w, "Methods is not surported" , http.StatusNotFound)
		return  
	}
	fmt.Fprintf(w,"Hello!")
}
 func main (){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Start the server at 8080\n")
	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
	}
 }
