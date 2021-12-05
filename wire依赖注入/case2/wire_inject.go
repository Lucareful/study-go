package case2

import "github.com/google/wire"

func InitializeUserService(foo string, bar int) *UserService {
	wire.Build(NewUserService, MockUserRepoSet) // 使用MockUserRepoSet
	return nil
}
