# 学习笔记

1. errors.Error(), return e.s
2. errors.New 友好描述，指明包名
3. errors.New 为什么返回指针？每次调用errors.New都是返回一个新对象，会比较内存地址是否一致，保证每次返回的error内容不同
4. recovery 野生goroutine，在基础库准备一个降级或者兜底函数，里边recovery
5. you should only handle error once
6. 对外暴露一个方法，function内部判定，把具体的error隐藏
7. kit库不应该wrap
8. wrap需要交给应用层处理

# 作业
1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
2. 答: 应该wrap这个error，抛给上层, 交给对外之前的顶层业务处理