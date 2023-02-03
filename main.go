package main

import (
	"fmt"
	"sync"
)

func main() {
	c1 := make(chan string)
	var wg sync.WaitGroup

	//Appel et affichage de la 1er API
	for port := 3000; port <= 4000; port++ {
		wg.Add(1)
		go firstAPI(port, &wg, c1)
	}
	string := <-c1
	key := string[19:]
	fmt.Println(key)

	//Appel et affichage de la 2eme API
	string2 := secondAPI(key)
	key2 := string2[20:]
	fmt.Println(key2)

	//Download du fichier manuellement ...

	//Appel et affichage de la 3eme API
	thirdAPI()

}
