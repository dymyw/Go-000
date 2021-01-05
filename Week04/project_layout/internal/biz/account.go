package biz

// Account
type Account struct {
	Id int
	Name string
	Address Address
}

// Address
type Address struct {
	Id int
	City string
	Area string
}

// AccountRepo
type AccountRepo interface {
	SaveAccount(*Account)
	BatchGetAccount([]Account)
}

// AccountUserCase
type AccountUserCase struct {
	// 依赖 repo
	repo AccountRepo
}

// NewAccountUserCase
func NewAccountUserCase(repo AccountRepo) *AccountUserCase {
	// 依赖注入，利用抽象、实现来解耦
	// 实现了 AccountRepo 的 repo 都可以传入，实现解耦
	return &AccountUserCase{repo: repo}
}

// Sign in
func (uc *AccountUserCase) SignIn(a *Account) {
	// 有点类似策略模式
	uc.repo.SaveAccount(a)
}
