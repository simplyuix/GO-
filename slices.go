package main

import "fmt"

func main() {

    worldTourCountries := []string{"Italy", "France", "USA", "Brazil", "India", "China"}

  
    worldTourCountries = append(worldTourCountries, "Australia")
	fmt.Println(worldTourCountries)

  
    worldTourCountries = append(worldTourCountries, "Spain")
	fmt.Println(worldTourCountries)


    worldTourCountries = append(worldTourCountries[:3], worldTourCountries[4:]...)
	fmt.Println(worldTourCountries)

 
    worldTourCountries = append(worldTourCountries[:4], worldTourCountries[5:]...)
	fmt.Println(worldTourCountries)


    worldTourCountries = append([]string{"Japan"}, worldTourCountries...)
	

    fmt.Println(worldTourCountries)
}