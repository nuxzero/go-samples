package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {

	re, err := regexp.Compile("p([a-z]+)ch")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("re: %v\n", re.MatchString("peach"))

	re3 := regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(re3.FindStringSubmatch("peach punch"))

	match, err := regexp.MatchString("p([a-z]+)ch", "peach")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("match: %v\n", match)

	re2 := regexp.MustCompile(`([a-zA-Z\s]+) road`)
	fmt.Println(re2.FindString("99/1 Sukhumvit road, Bangkok"))  // Sukhumvit road
	fmt.Println(re2.FindString("99/1, Suan Phlu road, Bangkok")) // Suan Phlu road
}
