package maps

import "fmt"

func Map() {
	fmt.Println("This is the Map example function.")
	// to delclare the map we use the make function. the syntax is make(map[keytype]valuetype)

	var maps = make(map[string]int)
	maps["one"] = 1
	maps["two"] = 2
	maps["three"] = 3
	
	fmt.Println("map after adding values: ", maps)
	fmt.Println("Iterating through the map:")

	for key, value := range maps {
		fmt.Println(key, ":", value)
	}

	fmt.Println("Accessing a value by key:", maps["two"])

	delete(maps, "three")
	fmt.Println("map after deleting key 'three': ", maps)

	fmt.Println("Length of the map using len(): ", len(maps))
	// getting the value by key
	value, exists := maps["one"]
	if exists {
		fmt.Println("Value for key 'one':", value)
	} else {
		fmt.Println("Key 'one' does not exist in the map.")
	}


}