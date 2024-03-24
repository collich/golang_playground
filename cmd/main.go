package main

import (
	// "fmt"
	// "log"
	// "os/exec"
	// "flag"
	// "os"
)

func main() {
	
}

/*----------CMD output----------*/
// func main() {
// 	output, err := exec.Command("docker", "ps").Output()
// 	if err != nil{
// 		log.Fatal(err)
// 	}

// 	fmt.Println(string(output))
// }

/*----------CMD programs----------*/
// func main() {
	// cmd := exec.Command("docker")
	// if err := cmd.Run(); err != nil{
	// 	log.Fatal(err)
	// }
	// fmt.Println(cmd.Run())
// }	
	
/*----------CMD flags----------*/
// func main() {
	// fmt.Println("Program name:", os.Args[0])

	// for i := 0; i < flag.NArg(); i++ {
		// 	fmt.Printf("Non Flags %d with the word: %s\n", i, flag.Arg(i))
	// }
	
	// wordPtr := flag.String("word", "Hello", "A Hello Word")
	// wordPtrShort := flag.String("w", "Hello", "A Hello Word")
	// flag.Parse()
		
	// for i, args := range os.Args[1:]{
		// 	fmt.Printf("Number %d with args %s\n", i, args)
	// }
			
	// fmt.Println("Word: ", *wordPtr)
	// fmt.Println("Word(short): ", *wordPtrShort)
	// fmt.Println("Others: ", flag.Args())
		
// }