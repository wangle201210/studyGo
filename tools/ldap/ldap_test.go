package ldap

import (
	"fmt"
	"testing"
)

func TestAddLdapUser(t *testing.T) {
	conn := InitLdap()
	defer conn.Close()
	// 添加用户
	user := User{
		username: "wanna",
		password: "123456",
	}
	err := addUser(conn, user)
	if err != nil {
		panic(err)
	}
	fmt.Println("add success...")
}

func TestFindLdapUser(t *testing.T) {
	conn := InitLdap()
	defer conn.Close()
	// 添加用户
	user := User{
		username: "wanna",
		password: "123456",
	}
	searchResult, err := findUser(conn, user)
	if err != nil {
		panic(err)
	}
	for _, entry := range searchResult.Entries {
		fmt.Println("find user: ", entry.DN)
		for _, v := range entry.Attributes {
			if v.Name == "userPassword" {
				if v.Values[0] != user.password {
					panic("密码不对")
				} else {
					fmt.Println("密码正确")
				}
			}
			fmt.Println(v.Name, v.Values)
		}
	}
	fmt.Println("find success...")
}
