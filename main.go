package main

import (
	"fmt"

	"github.com/elliotchance/orderedmap"
)

func main() {
	
		myFruits := orderedmap.NewOrderedMap() //This is an empty map pointing to a memory location

		//Created an Example => setting keys and values into the empty map "myFruits"
		//myFruits.Set => sets/allocates a key and value into the empty map
		myFruits.Set("orange", 4)
		myFruits.Set("mango", 2)
		myFruits.Set("grape", 2)
		myFruits.Set("apple", 1)
		myFruits.Set("pineapple", 2)
		myFruits.Set("strawberry", 2)

		//myFruits.len() => spitsOut or returns the length of map i.e the numbers of keys in map
		fmt.Println(myFruits.Len())

		// The way myFruits.Keys() work is it creates a slice of keys and index each one
		fmt.Println(myFruits.Keys()) // so this prints the keys in a slice

		// myFruits.Get returns an integer (1 or 0 == nil) and a boolean. Hence if key isPresent in map "myFruits"
		// it returns 1 and true, if key isAbsent in map "myFruits" returns nil and false
		fmt.Println(myFruits.Get("apple")) // returns 1 and true
		fmt.Println(myFruits.Get("banana")) // returns nil and false

		//myFruits.Delete => as a name implies deletes or remove keys in the map "myFruits"
		//and returns the new map "myFruits"
		myFruits.Delete("pineapple")
		fmt.Println(myFruits.Keys())

		//This is a FOR LOOP Front : iterates over the map "myFruits" and returns items from the first index or key to the last
		//Used for larger lists of items
		//INITIALIZATION => el:= myFruits.Front()
		//CONDITION => el != nil
		//INCREMENT => el = el.Next()

		for el:= myFruits.Front(); el != nil ; el = el.Next() {
			fmt.Println(el.Key, el.Value)
		}

		//This is a FOR LOOP Back : iterates over the map "myFruits" and returns items from the last index or key to the first
		for el:= myFruits.Back(); el != nil ; el = el.Prev() {
			fmt.Println(el.Key, el.Value)
		}

		// for range, iterates over the slice of keys and prints them in the order it was listed/arranged no matter how many times you
		//iterate over it => how it works is since the map "myFruits" is now "myFruits.Keys()" and "myFruits.Keys" is a slice, the items would not
		// change position because items in a slice do not swap positions.
		for _, key := range myFruits.Keys() {
			fmt.Println(key)
		}

		//This will print out the index of each key
		for value, _ := range myFruits.Keys() {
			fmt.Println(value)
		}

		// To get the values associated to the keys in the right order, this ranges over each items and prints out the values.
		for _, key := range myFruits.Keys() {
			value, _ := myFruits.Get(key)
			fmt.Println(value)
		}

}


