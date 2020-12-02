# 学习笔记

1. errors.Error(), return e.s
2. errors.New 友好描述，指明包名
3. errors.New 为什么返回指针？每次调用errors.New都是返回一个新对象，会比较内存地址是否一致，保证每次返回的error内容不同
4. recovery 野生goroutine，在基础库准备一个降级或者兜底函数，里边recovery
5. you should only handle error once
6. 对外暴露一个方法，function内部判定，把具体的error隐藏
7. kit库不应该wrap
8. wrap需要交给应用层处理
9. 

# 作业
