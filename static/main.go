package main

import(
	"fmt"
	"log"
	"net/http"

)

func formhandler(w http.ResponseWriter,r *http.Request){
	if err :=r.ParseForm(); err !=nil{
		fmt.Fprintf(w,"parseform() err:%v",err)
		return
	}
	fmt.Fprintf(w,"post request successful\n")
	name:=r.FormValue("name")
	address:=r.FormValue("address")
	fmt.Fprintf(w,"Name=%s\n",name)
	fmt.Fprintf(w,"Address=%s",address)


}

func hellohandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path !="/hello"{
	http.Error(w ,"not found",http.StatusNotFound)
	return
	}
	fmt.Fprintf(w,"hello")
}

func main(){
	fileServer := http.FileServer(http.Dir("./"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formhandler)
	http.HandleFunc("/hello",hellohandler)

	fmt.Print("server is running at port 8080\n")
	if err :=http.ListenAndServe(":8080",nil); err !=nil{
		log.Fatal(err)
	}
}