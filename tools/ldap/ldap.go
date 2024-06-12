package ldap

import (
	"crypto/tls"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

// ldap：未加密
// ldaps：加密
var ldapURL = "ldap://127.0.0.1:389"
var objectClass = "person"
var baseDN = "dc=mycompany,dc=com"

type LdapConfig struct {
	Addr             string
	BindUserDn       string
	BindUserPassword string
	BaseDn           string
	LoginName        string
	ObjectClass      []string
}

type User struct {
	username    string
	password    string
	telephone   string
	emailSuffix string
	snUsername  string
	uid         string
	gid         string
}

func loginBind(config *LdapConfig) (*ldap.Conn, error) {
	l, err := ldap.DialURL(ldapURL, ldap.DialWithTLSConfig(&tls.Config{InsecureSkipVerify: true}))

	if err != nil {
		panic(err)
		return nil, err
	}
	_, err = l.SimpleBind(&ldap.SimpleBindRequest{
		Username: config.BindUserDn,
		Password: config.BindUserPassword,
	})
	if err != nil {
		fmt.Println("ldap password is error: ", ldap.LDAPResultInvalidCredentials)
		return nil, err
	}
	fmt.Println("bind success...")
	return l, nil
}

// 创建用户
func addUser(conn *ldap.Conn, user User) error {
	// 添加用户
	addRequest := ldap.NewAddRequest(fmt.Sprintf("cn=%s,%s", user.username, baseDN), nil)
	addRequest.Attribute("objectClass", []string{objectClass})
	// addRequest.Attribute("ou", []string{"QA Group"})
	// addRequest.Attribute("o", []string{"My Company"})
	// addRequest.Attribute("cn", []string{"41234123"})
	addRequest.Attribute("sn", []string{"xx2"})
	// addRequest.Attribute("uid", []string{"10001"})
	addRequest.Attribute("userPassword", []string{user.password})
	err := conn.Add(addRequest)
	if err != nil {
		fmt.Println("add user error: ", err)
		return err
	}
	return nil
}

// 查询用户
func findUser(conn *ldap.Conn, user User) (*ldap.SearchResult, error) {
	// 多个条件：(&(cn=wangmazi)(ou=QA))
	filter := fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(user.username))
	request := ldap.NewSearchRequest(fmt.Sprintf("%s", baseDN),
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filter,
		[]string{"userPassword"},
		nil,
	)
	searchResult, err := conn.Search(request)
	if err != nil {
		fmt.Println("search user error: ", err)
		return nil, err
	}
	return searchResult, nil
}

// 删除用户
func deleteUser(conn *ldap.Conn, config *LdapConfig, user User) error {
	dn := fmt.Sprintf("cn=%s,ou=QA,%s", user.username, config.BaseDn)
	fmt.Printf("del dn %v", dn)
	delRequest := ldap.NewDelRequest(dn, nil)
	err := conn.Del(delRequest)
	if err != nil {
		fmt.Printf("Failed to delete user %s: %v\n", dn, err)
		return err
	}
	fmt.Printf("User %s successfully deleted.\n", dn)
	return nil
}

func InitLdap() *ldap.Conn {
	// Ldap Config(用于校验后续的操作，包括查询用户是否存在、添加、删除等)
	config := new(LdapConfig)
	config.Addr = ldapURL
	config.BaseDn = baseDN
	config.BindUserDn = "cn=admin," + baseDN
	config.LoginName = "admin"
	config.BindUserPassword = "admin"
	// 客户不配置username，我们需要根据配置的ObjectClass查询出对应的用户。
	// 因为如果用户配置的是cn，那么可能会查询出一些组织、其他设备等，所以为了将Ldap第三方用户纳管过来，我们需要添加ObjectClass
	config.ObjectClass = []string{objectClass}

	// 与建立ldap服务建立连接（方便后续查询新增删除项）
	conn, err := loginBind(config)
	if err != nil {
		panic(err)
	}
	return conn
}

func RunLdap() {
	conn := InitLdap()
	defer conn.Close()
	TestAddUser(conn)
}

// TestAddUser 测试添加用户
func TestAddUser(conn *ldap.Conn) {
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

// TestFindUser 测试查询用户
func TestFindUser(conn *ldap.Conn, config *LdapConfig) {
	user := &User{
		username: "wangmazi",
	}
	searchResult, err := findUser(conn, *user)
	if err != nil {
		panic(err)
	}
	for _, entry := range searchResult.Entries {
		fmt.Println("find user: ", entry.DN)
		for _, v := range entry.Attributes {
			fmt.Println(v.Name, v.Values)
		}
	}
	return
}

func TestDeleteUser(conn *ldap.Conn, config *LdapConfig) {
	user := User{
		username: "wangmazi",
	}
	err := deleteUser(conn, config, user)
	if err != nil {
		panic(err)
	}

}
