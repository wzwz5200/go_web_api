package main

import (
	"fmt"
	"hellow/comm"

	"github.com/gin-gonic/gin"
)
type Hero struct {
	Name string
	ad int
}


func (this Hero) Getname() {
  fmt.Println("name",this.Name)

}


func (this Hero) Setname(newName string) {
 
	this.Name = newName


}

func changeAV(p *int){

	 *p = 10
}

func ping(c *gin.Context){

		c.JSON(200, gin.H{
			"message": "pong",
		})
}

func main() {
	// c :=fool("w",3)

	
	// fmt.Println("c=" ,c)
    // var a int = 1
	// changeAV(&a)
	// fmt.Println("a=", a)
	
    comm.GetPG()
	fmt.Println("W")


	r := gin.Default()
	
	r = CollectRoute(r)

	r.GET("/ping", ping)
	


	r.Run()
 
}