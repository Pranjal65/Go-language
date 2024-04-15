package main

import (
	"encoding/json"
	"os" // this package is used to perform operations on files
)

// struct of video which will strore the informtaion of each video
type video struct {
	ID          string `json:"Id"`  // this fields are used to map json fields to struct fields
	Title       string `json:"title"`
	Description string `json:"description"`
	ImgURL      string `json:"imgurl"`
	URL         string `json:"url"`
}

func getVideos() ([]video){
	
	filepath:="data.json"
	// variable in which we will store file data
	var videos []video;
    //Open function to open the file
	file,err:=os.ReadFile(filepath);
   if(err!=nil){
	  panic(err);
   }
   //defers its Close function with the defer keyword, and uses the json's NewDecoder function to read the file. 
  // defer file.Close();

  //Unmarshal parses the JSON-encoded data into go values and stores the result in the value pointed to by videos
  //Unmarshal return error
   err=json.Unmarshal(file,&videos);

   //marshal function is used to convert the go data into json data

   //file data is decoded s per the structs structure
 //  err = decoder.Decode(&videos);

   if(err!=nil){
	 panic(err);
   }
  // videos = append(videos,videos);

	return videos
}


func saveVedios(videos []video){
      // to convert the go structure into json data we use marshal function
	  videoBytes,err:=json.Marshal(videos)//converted into byte format

	  if(err!=nil ){
		panic(err);
	  }
	  os.WriteFile("./updated-data.json",videoBytes,0644)
}