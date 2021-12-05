package main

import "github.com/google/wire"

/*
在Go的最佳实践中，更推荐返回具体实现而不是接口。
*/

// UserService 用户服务.
type UserService struct {
	userRepo UserRepository // <-- UserService依赖UserRepository接口
}

type User struct {
	ID   int
	Name string
}

// UserRepository 存放User对象的数据仓库接口,比如可以是 mysql,restful api .
type UserRepository interface {
	// GetUserByID 根据ID获取User, 如果找不到User返回对应错误信息
	GetUserByID(id int) (*User, error)
}

// NewUserService *UserService构造函数
func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// mockUserRepo 模拟一个UserRepository实现
type mockUserRepo struct {
	foo string
	bar int
}

// GetUserByID UserRepository接口实现
func (u *mockUserRepo) GetUserByID(id int) (*User, error) {
	return &User{}, nil
}

// NewMockUserRepo *mockUserRepo构造函数
func NewMockUserRepo(foo string, bar int) *mockUserRepo {
	return &mockUserRepo{
		foo: foo,
		bar: bar,
	}
}

// MockUserRepoSet 将 *mockUserRepo与UserRepository绑定
var MockUserRepoSet = wire.NewSet(NewMockUserRepo, wire.Bind(new(UserRepository), new(*mockUserRepo)))
