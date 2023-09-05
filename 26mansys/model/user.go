package model

// 这里思考一些问题?

// 一个书城的用户应该拥有那些属性呢?

// 姓名，年龄，收藏的图书列表

type User struct {
	Username string
	Password string
	BookList []Book
}
