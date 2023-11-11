package main

import "fmt"

func main() {
	// In Go language, a map is a powerful, ingenious, and versatile data structure. Golang Maps is a collection of unordered pairs of key-value. It is widely used because it provides fast lookups and values that can retrieve, update or delete with the help of keys. It is a reference to a hash table.

	languages := make(map[string]string)
	languages["js"] = "JavaScript"
	languages["go"] = "Go"
	languages["py"] = "Python"
	languages["rb"] = "Ruby"
	languages["php"] = "PHP"
	languages["c"] = "C"

	// itrteate over the map and print the key and value
	// for key, value := range languages {
	// 	fmt.Printf("for key %v value is %v\n", key, value)
	// }

	// to delete value from map or from slice
	delete(languages, "rb")

	// if we don't care about key
	for _, value := range languages {
		fmt.Printf("for key value is %v\n", value)
	}

        for _ value := range languages {
		fmt.Printf("for key value is %v\n", value)
	}
}
