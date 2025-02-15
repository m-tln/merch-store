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
	// jwtService := service.NewJWTService("pronin")
	// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk3MzU0OTEsInVzZXJfaWQiOjZ9.3cALV0bP6aaKgRChD_dhnpkddVsEqZGfjah2uDqXtBo"
	// token, err := jwtService.ValidateToken(tokenString)

	// fmt.Println(token.Valid)
	// if err != nil {
	// 	fmt.Println("err: ", err)
	// 	return
	// }
	// if !token.Valid {
	// 	fmt.Println("err: ", err)
	// 	return
	// }

	// claims, ok := token.Claims.(jwt.MapClaims)
	// if !ok {
	// 	fmt.Println("Missing claims")
	// }
	// fmt.Println(claims)

	// userID, ok := claims["user_id"]
	// printType(userID)
	// if !ok {
	// 	fmt.Println("Missing user_id in token")
	// }
	// ctx := context.WithValue(context.Background(), KeyUserID, userID)
	// fmt.Println("userID: ", userID)
	// fmt.Println("ctx: ", ctx.Value(KeyUserID))
}
