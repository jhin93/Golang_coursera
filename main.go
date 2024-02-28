package main

import (
	"encoding/json"
	"fmt"
)

// User 구조체 정의
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func main() {
	// JSON 문자열
	jsonData := `{
        "id": 1,
        "first_name": "John",
        "last_name": "Doe",
        "email": "john.doe@example.com"
    }`

	// User 구조체 타입의 변수 선언
	var user User

	// json.Unmarshal을 사용하여 JSON 문자열을 User 구조체로 변환
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 변환된 User 구조체 출력
	fmt.Printf("%+v\n", user)
	// 결과 : {ID:1 FirstName:John LastName:Doe Email:john.doe@example.com}
}
