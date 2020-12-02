## 学习笔记

我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

1. dao 层是和底层交互的，所以要 Wrap
2. 中间层传递错误，可以用 WithMessage 保存信息，也可以使用 withStack 仅用来保存堆栈信息
3. 最上层使用 error 判等操作来屏蔽底层的 SQL 报错（业务无需关心底层的数据实现），仅返回 404 Not Found
