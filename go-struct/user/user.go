package user

import (
	"errors"
	"fmt"
)

type User struct {
	Name  string
	Age   int
	Clazz string
}

// Embedding struct
type Admin struct {
	user User
	Role string
}

// Đi tìm biến quang tại địa chỉ 0x4a1a7bf2c0f0.
// Nếu muốn lấy Name: Đọc ngay tại vị trí đó (cộng 0 byte).
// Nếu muốn lấy Age: Cộng thêm 16 bytes vào địa chỉ gốc.
// Nếu muốn lấy Clazz: Cộng thêm 24 bytes (16 + 8) vào địa chỉ gốc.
func New(name string, age int, clazz string) (*User, error) {
	if name == "" || age <= 0 || clazz == "" {
		return nil, errors.New("Invalid input: Name, Age, and Clazz must be provided and valid")
	}
	return &User{
		Name:  name,
		Age:   age,
		Clazz: clazz,
	}, nil
}

func NewAdmin(name string, age int, clazz string, role string) (*Admin, error) {
	// user, err := New(name, age, clazz)
	// if err != nil {
	// 	return nil, err
	// }
	// return &Admin{
	// 	user: *user,
	// 	Role: role,
	// }, nil
	return &Admin{
		user: User{
			Name:  name,
			Age:   age,
			Clazz: clazz,
		},
		Role: role,
	}, nil
}

// Cach 2: printUser struct  - Receiver function
func (u User) PrintUser() {
	fmt.Println(u.Name)
	fmt.Println(u.Age)
	fmt.Println(u.Clazz)
}

// 1. Thử thay đổi tuổi bằng Value Receiver
func (u User) UpdateAgeValue(newAge int) {
	u.Age = newAge // Chỉ sửa trên bản sao, quang ở main KHÔNG đổi
}

// 2. Thử thay đổi tuổi bằng Pointer Receiver
func (u *User) UpdateAgePointer(newAge int) {
	u.Age = newAge // Sửa trực tiếp vào ô nhớ gốc, quang ở main SẼ đổi
}

// passing struct by value
func OutputUserDetail(user User) {
	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println(user.Clazz)
}

// Cach 1: Copy struct bằng cách gán từng trường một
func CopyUserStruct(userNeedCopy *User) User {
	var user User
	user.Name = userNeedCopy.Name
	user.Age = userNeedCopy.Age
	user.Clazz = userNeedCopy.Clazz
	return user
}
