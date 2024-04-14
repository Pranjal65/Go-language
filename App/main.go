package main

import "fmt"

func main(){
	fmt.Println("Hello World");
	vedios:=getVideos()
	for _, v:=range vedios{
		fmt.Println()
		vedio:=fmt.Sprintf("Id of vedio %v \n Title of vedio: %v \n Description of vedio: %v \n imageurl of vedio is: %v \n url of vedio is: %v",v.ID,v.Title,v.Description,v.ImgURL,v.URL);
		fmt.Println(vedio);
	}
	saveVedios(vedios);
}