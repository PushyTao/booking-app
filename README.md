# booking-app

learning go note

## 数组和切片

array 是固定长度的一种类型，而 slice 是内存大小动态开辟的一种类型

关于循环：

![](image/README/1642843389549.png)

数组可以对指定下标的进行赋值，切片也是可以的，但是一开始为空，并且不好确定长度所以说我们可以对已经有的元素进行值的指定


```go
// bookings[0] = firstName + " " + lastName
// append添加一个元素到切片的末尾，并返回一个添加后的切片，显然可以直接赋给bookings
bookings = append(bookings, firstName+" "+lastName)
```

切片长度的确定可以使用 `len(sliceName)`来确定

![](image/README/1642926794416.png)


## 循环 & 选择

for {

}是一种无循环次数的循环，可以在cmd中使用ctrl + c停止循环

在遍历切片的过程中，可以使用for循环进行

![](image/README/1642844135603.png)

在遍历的过程当中，需要两个值进行遍历，如果命名为index, booking的话，然而并没有使用到index，所以说可以使用空白标识符"_"来标识一下，表示这里有一个东西在遍历的过程中需要用到，但是在代码逻辑上没有使用到，所以说要使用一个下划线来进行处理

因为对于一个for循环上可以加入条件，所以说关于死循环，还可以写成

for true {

}这个样子来进行处理，将意味着是一个死循环

条件较少的情况下：

if {

}或者是

if 条件1 {

}else if 条件2 {

}或者是

if 条件1 {

} else if 条件2 {

}....

else {

}

多条件适用switch语句：

![](image/README/1642846941267.png)

使用switch语句进行编程可能会简化代码

## 函数

可以使用func定义一个函数

```go
.func funcName(param1 type1,.....) return type{
}
```

因为go的函数可以返回多个值，所以说，在对于返回多个返回值的函数的返回类型声明中，可以逐个的声明返回类型，并且用圆括号包括起来，比如：

```go
func funcName(para1 para1Type,para2 para2Type,......) (returnType1, returnType2,......){
	return value1,value2......
}
```
