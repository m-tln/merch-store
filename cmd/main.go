package main

import "merch-store/internal/app"

type contextKeyUserID string

const KeyUserID contextKeyUserID = "userID"

// func printType(v any) {
// 	t := reflect.TypeOf(v)
// 	fmt.Printf("Type: %v, Value: %v\n", t, v)
// }

func main() {
	app.Init()
	
}
