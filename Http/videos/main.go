package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main(){
     // this is used to handle request and response over the server
	 // firt parameter is the  route to the page and second is function to call which takes two inputs responsewriter and request
	http.HandleFunc("/home",Hello);
	//get  API
	http.HandleFunc("/video",HandleGetVideos);
	//post  API
	http.HandleFunc("/update",HandleUpdateVideos);

	//http.ListenAndServe() starts an HTTP server with given address and handler
	//handler is usually il which means use defaultServeMux
	//st time of port no make use of colon always
	http.ListenAndServe(":8000",nil);

}
func Hello(w http.ResponseWriter, r *http.Request){
	// all the headers are present int the Header section of the httprequest we can iterate over it.
	for header_key,values := range r.Header{
		fmt.Printf("\nheader key: %v and values is:%v",header_key,values);
	}

	// to write back to the header of request Add()function is provided 
	w.Header().Add("Token","1234567890")
	// write function takes stream/slice of bytes to write response on webpage
	w.Write([]byte("Welcome to The Go language Tutorial"));
}
func HandleGetVideos(w http.ResponseWriter, r *http.Request){
	videos:=getVideos();
	videoBytes, err:=json.Marshal((videos));
	if(err!=nil){
		panic(err);
	}
	w.Write(videoBytes);
}
func HandleUpdateVideos(w http.ResponseWriter, r *http.Request){
	if(r.Method=="POST"){
            body,err:=io.ReadAll(r.Body);
			if(err!=nil){
				panic(err);
			}
			var videos []video;
			err=json.Unmarshal(body,&videos);
            if(err!=nil){
				w.WriteHeader(400);
				//Fprintf formats according to a format specifier and writes to w. It returns the number of bytes written and any write error encountered.
				fmt.Fprintf(w,"Bad request");
			}
			saveVedios(videos);
	}else{
		w.WriteHeader(405)
		fmt.Fprintf(w,"Method not supported");

	}
}