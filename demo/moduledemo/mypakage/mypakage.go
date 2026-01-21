package mypakage

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

// 一、基本数据类型
func New() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
	fmt.Println("基本数据类型结束")
}

// 练习题
// 1· 编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。
func Dayin() {
	var zhengxing int = 18
	var fudian float64 = 3.14
	var buer bool = true
	var zifuchuan string = "沙河小王子"
	fmt.Printf("值：%v，类型：%T\n", zhengxing, zhengxing)
	fmt.Printf("值：%v，类型：%T\n", fudian, fudian)
	fmt.Printf("值：%v，类型：%T\n", buer, buer)
	fmt.Printf("值：%v，类型：%T\n", zifuchuan, zifuchuan)
	fmt.Println("基本数据类型练习1结束")

}

// 2· 编写代码统计出字符串"hello沙河小王子"中汉字的数量。
func Shahewangzi() {
	s1 := "hello沙河小王子"
	count := 0
	runes := []rune(s1)

	for _, value := range runes {
		if unicode.Is(unicode.Scripts["Han"], value) {
			count++
		}
	}
	fmt.Printf("字符串'%s'中有%d个汉字\n", s1, count)
	fmt.Println("基本数据类型练习2结束")

}

// 二、运算符（算术运算符·关系运算符·逻辑运算符·位运算符·赋值运算符）
// 算术运算符
func Suanshu() {
	var a, b = 1, 2
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	fmt.Println("算术运算符结束")
}

// 关系运算符
func Guanxi() {
	var a, b = 10, 20
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println("关系运算符结束")
}

// 逻辑运算符
func Luoji() {
	var a, b = true, false
	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!a)
	fmt.Println("逻辑运算符结束")

}

// 位运算符（位运算符对整数在内存中的二进制位进行操作。）
func Wei() {
	var a, b = 10, 20
	fmt.Println(a & b)
	fmt.Println(a | b)
	//参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。（两位不一样则为1）
	fmt.Println(a ^ b)
	fmt.Println(a << 2)
	fmt.Println(a >> 2)
	fmt.Println(^a)
	fmt.Println("位运算符结束")

}

// 赋值运算符
func Fuzhi() {
	var a, b, c = 10, 20, 30
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	a += 10
	fmt.Println(a)
	b -= 10
	fmt.Println(b)
	c *= 10
	fmt.Println(c)
	c /= 10
	fmt.Println(c)
	c %= 10
	fmt.Println(c)
	c <<= 2
	fmt.Println(c)
	c >>= 2
	fmt.Println(c)
	c &= 10
	fmt.Println(c)
	c ^= 10
	fmt.Println(c)
	fmt.Println("赋值运算符结束")
	fmt.Println("运算符结束")

}

// 练习题
// 有一堆数字，如果除了一个数字以外，其他数字都出现了两次，那么如何找到出现一次的数字？
func Chuxianyici() int {
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	result := 0
	for _, value := range a {
		result ^= value
	}
	fmt.Println("只出现了一次的数字是：", result)
	fmt.Println("运算符练习结束")
	return result
}

// 三、流程控制练习题（编写代码打印9*9乘法表。）
func JiuJiu() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

// 四、数组练习题
// 求数组[1, 3, 5, 7, 8]所有元素的和
func Shuzuhe() {
	nums := [5]int{1, 3, 5, 7, 8}
	var sum int
	for _, value := range nums {
		sum += value
	}
	fmt.Println("数组和为：", sum)
	fmt.Println("数组练习1结束")

}

// 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
func Hewei8() {
	nums := [5]int{1, 3, 5, 7, 8}
	target := 8

	for _, i := range nums {
		for _, j := range nums {
			if i+j == target {
				fmt.Println("和为8的两个元素的下标分别为：", i, j)
			}
		}
	}

	fmt.Println("数组练习2结束")
}

//五、切片练习题
//请写出下面代码的输出结果。
/*
func Qiepian() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
}
*/
//答案：[0 1 2 3 4 5 6 7 8 9]

// 请使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序
func Paixu() {
	var a = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a[:])
	fmt.Println(a)
	fmt.Println("切片练习1结束")
}

func Qiepian_2() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("切片练习2结束")
}

//六、map练习题
//观察下面代码，写出最终的打印结果。
/*
func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
*/
//答案：{1，2，3} {1，3} {1，3}

// 写一个程序，统计一个字符串中每个单词出现的次数。比如：“how do you do"中how=1 do=2 you=1。
func WordCount() {
	s := "how do you do"
	wordcount := make(map[string]int)

	words := strings.Fields(s)
	for _, word := range words {
		wordcount[word]++
	}
	for word, count := range wordcount {
		fmt.Printf("%s:%d\n", word, count)
	}
	fmt.Println("map练习结束")
}

//defer经典案例
/*
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func main() {
	fmt.Println(f1()) //5
	fmt.Println(f2()) //6
	fmt.Println(f3()) //5
	fmt.Println(f4()) //5
}
*/
//打印：5，6，5，5

//defer面试题
/*
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
*/
//打印：A 1 2 3
// 	   A 10 2 12
//     B 10 20 30
//     AA 10 12 22
//     BB 10 12 22
//提示：defer注册要延迟执行的函数时该函数所有的参数都需要确定其值（所以calc("A", x, y)的调用会提前执行）
//注意：recover()必须搭配defer使用。
// 	   defer一定要在可能引发panic的语句之前定义。

//七、函数练习题
/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	//定义字符兑换金币规则
	rule := map[string]int{
		"e": 1,
		"i": 2,
		"o": 3,
		"u": 4,
		"E": 1,
		"I": 2,
		"O": 3,
		"U": 4,
	}

	//总共使用的金币数量
	totalUsed := 0

	//遍历用户
	for _, user := range users {
		userCoins := 0
		//遍历用户名
		for _, char := range user {
			//如果有符合的char字符就返回coinsUsed
			if coinsUsed, ok := rule[string(char)]; ok {
				userCoins += coinsUsed
			}
		}
		//把遍历的每个用户的金币数保存到map中
		distribution[user] = userCoins
		//累计所有用户的金币
		totalUsed += userCoins
	}

	fmt.Println("分配：", distribution)
	for user, coin := range distribution {
		fmt.Printf("%s: %d\n", user, coin)
	}
	return coins - totalUsed
}
func Fenjinbi() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	fmt.Println("函数练习结束")
}

//八、指针
// 指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值。
//Go语言中的指针不能进行偏移和运算，因此Go语言中的指针操作非常简单，我们只需要记住两个符号：&（取地址）和*（根据地址取值）

//		ptr := &v    // v的类型为T
// v:代表被取地址的变量，类型为T
// ptr:用于接收地址的变量，ptr的类型就为*T，称做T的指针类型。*代表指针。

// 总结： 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。

// 变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
// 对变量进行取地址（&）操作，可以获得这个变量的指针变量。
// 指针变量的值是指针地址。
// 对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。

// new是一个内置的函数，它的函数签名如下：
// new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。
// func new(Type) *Type
// 其中，
// Type表示类型，new函数只接受一个参数，这个参数是一个类型
// *Type表示类型指针，new函数返回一个指向该类型内存地址的指针。

// make也是用于内存分配的，区别于new，它只用于slice、map以及channel的内存创建，
// 而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
// make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作。
// make函数的函数签名如下：
// func make(t Type, size ...IntegerType) Type

//new与make的区别
//二者都是用来做内存分配的。
// make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
// 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。

//九、结构体
//自定义类型
//将MyInt定义为int类型
// type MyInt int
// 通过type关键字的定义，MyInt就是一种新的类型，它具有int的特性。

//类型别名（rune和byte就是类型别名）
// 类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。
// type TypeAlias = Type

//二者的区别
//类型定义
// type NewInt int

//类型别名
// type MyInt = int

// func main() {
// 	var a NewInt
// 	var b MyInt

// 	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
// 	fmt.Printf("type of b:%T\n", b) //type of b:int
// }

// 结构体的定义
// type 类型名 struct {
//     字段名 字段类型
//     字段名 字段类型
//     …
// }
// 其中：
// 类型名：标识自定义结构体的名称，在同一个包内不能重复。
// 字段名：表示结构体字段名。结构体中的字段名必须唯一。
// 字段类型：表示结构体字段的具体类型。

// 结构体实例化
// var 结构体实例 结构体类型
// type person struct {
// 	name string
// 	age  int
// }

// func main() {
// 	var jxh person
// 	jxh.name = "将绣花"
// 	jxh.age = 20
// fmt.Println("这个人叫：", jxh.name)
// fmt.Println("这个人的年龄是：", jxh.age)
// }

// 匿名结构体
// var user struct{Name string; Age int}
//     user.Name = "小王子"
//     user.Age = 18
//     fmt.Printf("%#v\n", user)

// 创建指针类型结构体
// var p2 = new(person)
// fmt.Printf("%T\n", p2)     //*main.person
// fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", age:0}
// Go语言中支持对结构体指针直接使用.来访问结构体的成员
// var p2 = new(person)
// p2.name = "小王子"
// p2.age = 28
// fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"小王子", age:28}

// 取结构体的地址实例化
//使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
// p3 := &person{}
// fmt.Printf("%T\n", p3)     //*main.person
// fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
// p3.name = "七米"
// p3.age = 30
// p3.city = "成都"
// fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"七米", city:"成都", age:30}

// p3.name = "七米"其实在底层是(*p3).name = "七米"，这是Go语言帮我们实现的语法糖。

// 结构体初始化
//有两种方式：使用键值对初始化、使用值的列表初始化
//① 键对应结构体的字段，值对应该字段的初始值 ； 也可以对结构体指针进行键值对初始化 ； 当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。
//以下分别是三个例子
// p5 := person{
// 	name: "小王子",
// 	city: "北京",
// 	age:  18,
// }
// fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"小王子", city:"北京", age:18}
// p6 := &person{
// 	name: "小王子",
// 	city: "北京",
// 	age:  18,
// }
// fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"小王子", city:"北京", age:18}
// p7 := &person{
// 	city: "北京",
// }
// fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"", city:"北京", age:0}
//② 初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值
// p8 := &person{
// 	"沙河娜扎",
// 	"北京",
// 	28,
// }
// fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"沙河娜扎", city:"北京", age:28}
//注意：必须初始化结构体的所有字段。初始值的填充顺序必须与字段在结构体中的声明顺序一致。该方式不能和键值初始化方式混用。

//结构体占用一块连续的内存
// type test struct {
// 	a int8
// 	b int8
// 	c int8
// 	d int8
// }
// n := test{
// 	1, 2, 3, 4,
// }
// fmt.Printf("n.a %p\n", &n.a)
// fmt.Printf("n.b %p\n", &n.b)
// fmt.Printf("n.c %p\n", &n.c)
// fmt.Printf("n.d %p\n", &n.d)
//输出
// n.a 0xc0000a0060
// n.b 0xc0000a0061
// n.c 0xc0000a0062
// n.d 0xc0000a0063

//空结构体不占用空间

//方法和接收者
// Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的this或者 self。

// 方法的定义格式如下：
// func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
//     函数体
// }
// 其中，
// 1、接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，而不是self、this之类的命名。
// 例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
// 2、接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
// 3、方法名、参数列表、返回参数：具体格式与函数定义相同。
//函数和方法最本质的区别是方法名前有接收者，但是函数没有。

//指针类型的接收者
//指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。这种方式就十分接近于其他语言中面向对象中的this或者self。
// SetAge 设置p的年龄
// 使用指针接收者
// func (p *Person) SetAge(newAge int8) {
// 	p.age = newAge
// }
// func main() {
// 	p1 := NewPerson("小王子", 25)
// 	fmt.Println(p1.age) // 25
// 	p1.SetAge(30)
// 	fmt.Println(p1.age) // 30
// }
//值类型的接收者
//当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。
// SetAge2 设置p的年龄
// 使用值接收者
// func (p Person) SetAge2(newAge int8) {
// 	p.age = newAge
// }
// func main() {
// 	p1 := NewPerson("小王子", 25)
// 	p1.Dream()
// 	fmt.Println(p1.age) // 25
// 	p1.SetAge2(30) // (*p1).SetAge2(30)
// 	fmt.Println(p1.age) // 25
// }

//Q: 什么时候使用指针类型接收者
//A：1、需要修改接收者中的值
// 	 2、接收者是拷贝代价比较大的大对象
// 	 3、保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

// 任意类型添加方法
// 在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。
// MyInt 将int定义为自定义MyInt类型
// type MyInt int
// // SayHello 为MyInt添加一个SayHello的方法
// func (m MyInt) SayHello() {
// 	fmt.Println("Hello, 我是一个int。")
// }
// func main() {
// 	var m1 MyInt
// 	m1.SayHello() //Hello, 我是一个int。
// 	m1 = 100
// 	fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt
// }
//注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。

// 结构体的匿名字段
// 结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。
// Person 结构体Person类型
// type Person struct {
// 	string
// 	int
// }

// func main() {
// 	p1 := Person{
// 		"小王子",
// 		18,
// 	}
// 	fmt.Printf("%#v\n", p1)        //main.Person{string:"北京", int:18}
// 	fmt.Println(p1.string, p1.int) //北京 18
// }
//注意事项： 这里匿名字段的说法并不代表没有字段名，而是默认会采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。

// 嵌套结构体
//一个结构体中可以嵌套包含另一个结构体或结构体指针
//Address 地址结构体
// type Address struct {
// 	Province string
// 	City     string
// }

// //User 用户结构体
// type User struct {
// 	Name    string
// 	Gender  string
// 	Address Address
// }

// func main() {
// 	user1 := User{
// 		Name:   "小王子",
// 		Gender: "男",
// 		Address: Address{
// 			Province: "山东",
// 			City:     "威海",
// 		},
// 	}
// 	fmt.Printf("user1=%#v\n", user1)//user1=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}
// }

// 嵌套匿名字段
// 上面user结构体中嵌套的Address结构体也可以采用匿名字段的方式，例如：
// Address 地址结构体
// type Address struct {
// 	Province string
// 	City     string
// }

// // User 用户结构体
// type User struct {
// 	Name    string
// 	Gender  string
// 	Address //匿名字段
// }

// func main() {
// 	var user2 User
// 	user2.Name = "小王子"
// 	user2.Gender = "男"
// 	user2.Address.Province = "山东"    // 匿名字段默认使用类型名作为字段名
// 	user2.City = "威海"                // 匿名字段可以省略
// 	fmt.Printf("user2=%#v\n", user2) //user2=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}
// }
//当访问结构体成员时会先在结构体中查找该字段，找不到再去嵌套的匿名字段中查找。

//嵌套结构体的字段名冲突
//嵌套结构体内部可能存在相同的字段名。在这种情况下为了避免歧义需要通过指定具体的内嵌结构体字段名。
//Address 地址结构体
// type Address struct {
// 	Province   string
// 	City       string
// 	CreateTime string
// }

// //Email 邮箱结构体
// type Email struct {
// 	Account    string
// 	CreateTime string
// }

// //User 用户结构体
// type User struct {
// 	Name   string
// 	Gender string
// 	Address
// 	Email
// }

// func main() {
// 	var user3 User
// 	user3.Name = "沙河娜扎"
// 	user3.Gender = "男"
// 	// user3.CreateTime = "2019" //ambiguous selector user3.CreateTime
// 	user3.Address.CreateTime = "2000" //指定Address结构体中的CreateTime
// 	user3.Email.CreateTime = "2000"   //指定Email结构体中的CreateTime
// }

//结构体的“继承”
// Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。
/*
// Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！
}
*/

// 结构体字段的可见性
// 结构体中字段大写开头表示可以公开访问，小写表示私有（仅在定义当前结构体的包中可以访问）

//结构体和JSON序列化
//JSON(JavaScript Object Notation) 是一种轻量级的数据交换格式。易于人阅读和编写，同时也易于机器解析和生成。
// JSON键值对是用来保存JS对象的一种方式，键/值对组合中的键名写在前面并用双引号""包裹，使用冒号:分隔，然后紧接着值；多个键值之间使用英文,分隔。
/*
//Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string
	Students []Student
}

func main() {
	c := &Class{
		Title:    "101",
		Students: make([]Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)

	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}
*/

//结构体标签（Tag）
//Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。 Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：
//  `key1:"value1" key2:"value2"`
// 结构体tag由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。同一个结构体字段可以设置多个键值对tag，不同的键值对之间使用空格分隔。

// 注意事项： 为结构体编写Tag时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。
// 例如不要在key和value之间添加空格。
/*
//Student 学生
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func main() {
	s1 := Student{
		ID:     1,
		Gender: "男",
		name:   "沙河娜扎",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"男"}
}
*/

//结构体和方法补充知识点
//因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意。我们来看下面的例子：

/* type Person struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person) SetDreams(dreams []string) {
	p.dreams = dreams
}

func main() {
	p1 := Person{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams)  // ?
}
*/
//正确的做法是在方法中使用传入的slice的拷贝进行结构体赋值。
/*
func (p *Person) SetDreams(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}
*/
//同样的问题也存在于返回值slice和map的情况，在实际编码过程中一定要注意这个问题。

// 九、练习题 - 结构体
// 使用“面向对象”的思维方式编写一个学生信息管理系统。
// 学生有id、姓名、年龄、分数等信息
// 程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能

/*
package mypakage

import (
	"fmt"
)

// Student 学生结构体
type Student struct {
	ID    int
	Name  string
	Age   int
	Score float64
}

// PrintInfo 打印单个学生信息
func (s *Student) PrintInfo() {
	fmt.Printf("学号：%d 姓名：%s 年龄：%d 分数：%.2f\n", s.ID, s.Name, s.Age, s.Score)
}

// StudentManager 学生管理器
type StudentManager struct {
	students map[int]*Student
	nextID   int
}

// NewStudentManager 创建新的学生管理器
func NewStudentManager() *StudentManager {
	return &StudentManager{
		students: make(map[int]*Student),
		nextID:   1,
	}
}

// AddStudent 添加学生
func (sm *StudentManager) AddStudent(name string, age int, score float64) {
	student := &Student{
		ID:    sm.nextID,
		Name:  name,
		Age:   age,
		Score: score,
	}
	sm.students[sm.nextID] = student
	sm.nextID++
	fmt.Printf("成功添加学生: %s，学号: %d\n", name, student.ID)
}

// ListStudents 展示所有学生列表
func (sm *StudentManager) ListStudents() {
	if len(sm.students) == 0 {
		fmt.Println("暂无学生信息")
		return
	}

	fmt.Println("=== 学生列表 ===")
	for _, student := range sm.students {
		student.PrintInfo()
	}
}

// GetStudentByID 根据ID获取学生
func (sm *StudentManager) GetStudentByID(id int) (*Student, bool) {
	student, exists := sm.students[id]
	return student, exists
}

// UpdateStudent 更新学生信息
func (sm *StudentManager) UpdateStudent(id int, name string, age int, score float64) error {
	student, exists := sm.GetStudentByID(id)
	if !exists {
		return fmt.Errorf("学号为 %d 的学生不存在", id)
	}

	student.Name = name
	student.Age = age
	student.Score = score

	fmt.Printf("成功更新学号为 %d 的学生信息\n", id)
	return nil
}

// DeleteStudent 删除学生
func (sm *StudentManager) DeleteStudent(id int) error {
	_, exists := sm.GetStudentByID(id)
	if !exists {
		return fmt.Errorf("学号为 %d 的学生不存在", id)
	}

	delete(sm.students, id)
	fmt.Printf("成功删除学号为 %d 的学生\n", id)
	return nil
}

// 示例使用
func StudentSystemDemo() {
	manager := NewStudentManager()

	// 添加学生
	manager.AddStudent("张三", 20, 85.5)
	manager.AddStudent("李四", 19, 92.0)
	manager.AddStudent("王五", 21, 78.5)

	// 显示学生列表
	manager.ListStudents()

	// 更新学生信息
	manager.UpdateStudent(1, "张三丰", 21, 95.0)

	// 再次显示列表
	fmt.Println("\n更新后的学生列表:")
	manager.ListStudents()

	// 删除学生
	manager.DeleteStudent(2)

	// 最终列表
	fmt.Println("\n删除学生后的列表:")
	manager.ListStudents()
}
*/

//十、包
//我们可以根据自己的需要创建自定义包。一个包可以简单理解为一个存放.go文件的文件夹。
// 该文件夹下面的所有.go文件都要在非注释的第一行添加如下声明，声明该文件归属的包。
//package packagename
// 其中：
// package：声明包的关键字
// packagename：包名，可以不与文件夹的名称一致，不能包含 - 符号，最好与其实现的功能相对应。
//另外需要注意一个文件夹下面直接包含的文件只能归属一个包，同一个包的文件不能在多个文件夹下。
// 包名为main的包是应用程序的入口包，这种包编译后会得到一个可执行文件，而编译不包含main包的源代码则不会得到可执行文件。

//标识符可见性
//在同一个包内部声明的标识符都位于同一个命名空间下，在不同的包内部声明的标识符就属于不同的命名空间。
// 想要在包的外部使用包内部的标识符就需要添加包名前缀，例如fmt.Println("Hello world!")，就是指调用fmt包中的Println函数。
// 如果想让一个包中的标识符（如变量、常量、类型、函数等）能被外部的包使用，那么标识符必须是对外可见的（public）。
// 在Go语言中是通过标识符的首字母大/小写来控制标识符的对外可见（public）/不可见（private）的。在一个包内部只有首字母大写的标识符才是对外可见的。
/*
package demo

import "fmt"

// 包级别标识符的可见性

// num 定义一个全局整型变量
// 首字母小写，对外不可见(只能在当前包内使用)
var num = 100

// Mode 定义一个常量
// 首字母大写，对外可见(可在其它包中使用)
const Mode = 1

// person 定义一个代表人的结构体
// 首字母小写，对外不可见(只能在当前包内使用)
type person struct {
	name string
	Age  int
}

// Add 返回两个整数和的函数
// 首字母大写，对外可见(可在其它包中使用)
func Add(x, y int) int {
	return x + y
}

// sayHi 打招呼的函数
// 首字母小写，对外不可见(只能在当前包内使用)
func sayHi() {
	var myName = "七米" // 函数局部变量，只能在当前函数内使用
	fmt.Println(myName)
}
*/
//同样的规则也适用于结构体，结构体中可导出字段的字段名称必须首字母大写。
/*
type Student struct {
	Name  string // 可在包外访问的方法
	class string // 仅限包内访问的字段
}
*/

//包的引入
//要在当前包中使用另外一个包的内容就需要使用import关键字引入这个包，并且import语句通常放在文件的开头，package声明语句的下方。完整的引入声明语句格式如下:
//import importname "path/to/package"
// 其中：
// importname：引入的包名，通常都省略。默认值为引入包的包名。
// path/to/package：引入包的路径名称，必须使用双引号包裹起来。
// Go语言中禁止循环导入包。
//同时引入多个包时，多个包之间用逗号隔开。
/*
import "fmt"
import "net/http"
import "os"
*/
//批量引入的方式
/*
import (
    "fmt"
  	"net/http"
    "os"
)
*/
// 当引入的多个包中存在相同的包名或者想自行为某个引入的包设置一个新包名时，都需要通过importname指定一个在当前文件中使用的新包名。例如，在引入fmt包时为其指定一个新包名f。
//import f "fmt"
//如果引入一个包的时候为其设置了一个特殊_作为包名，那么这个包的引入方式就称为匿名引入。一个包被匿名引入的目的主要是为了加载这个包，从而使得这个包中的资源得以初始化。 被匿名引入的包中的init函数将被执行并且仅执行一遍。
//import _ "github.com/go-sql-driver/mysql"
// 匿名引入的包与其他方式导入的包一样都会被编译到可执行文件中。
// 需要注意的是，Go语言中不允许引入包却不在代码中使用这个包的内容，如果引入了未使用的包则会触发编译错误。

//init初始化函数
//在每一个Go源文件中，都可以定义任意个如下格式的特殊函数：
/*
func init(){
  // ...
}
*/
//这种特殊的函数不接收任何参数也没有任何返回值，我们也不能在代码中主动调用它。当程序启动的时候，init函数会按照它们声明的顺序自动执行。

/* 一个包的初始化过程是按照代码中引入的顺序来进行的，所有在该包中声明的init函数都将被串行调用并且仅调用执行一次。
每一个包初始化的时候都是先执行依赖的包中声明的init函数再执行当前包中声明的init函数。
确保在程序的main函数开始执行时所有的依赖包都已初始化完成。 */

//每一个包的初始化是先从初始化包级别变量开始的。例如从下面的示例中我们就可以看出包级别变量的初始化会先于init初始化函数。
/*
package main

import "fmt"

var x int8 = 10

const pi = 3.14

func init() {
  fmt.Println("x:", x)
  fmt.Println("pi:", pi)
  sayHi()
}

func sayHi() {
	fmt.Println("Hello World!")
}

func main() {
	fmt.Println("你好，世界！")
}
*/
//输出结果
// x: 10
// pi: 3.14
// Hello World!
// 你好，世界！

//go module 网址：https://www.liwenzhou.com/posts/Go/package/
//相关命令
/*
go mod init	初始化项目依赖，生成go.mod文件
go mod download	根据go.mod文件下载依赖
go mod tidy	比对项目文件中引入的依赖与go.mod进行比对
go mod graph	输出依赖关系图
go mod edit	编辑go.mod文件
go mod vendor	将项目的所有依赖导出至vendor目录
go mod verify	检验一个依赖包是否被篡改过
go mod why	解释为什么需要某个依赖
*/

//十、练习题 - 包
//编写一个calc包实现加减乘除四个功能函数，在snow这个包中引入calc包并调用其加减乘除四个函数实现数学运算。
//现在假设我们已经编写了calc包，并且已经将其保存在 moduledemo 目录下。
// calc/calc.go
/*
package calc

import "errors"

// Add 实现加法运算
func Add(a, b float64) float64 {
	return a + b
}

// Subtract 实现减法运算
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply 实现乘法运算
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide 实现除法运算
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}
*/
//创建snow包并且调用
// snow/snow.go
/*
package snow

import (
	"fmt"
	"moduledemo/calc" // 根据您的项目结构调整导入路径
)

// CalculateExample 演示如何使用calc包进行数学运算
func CalculateExample() {
	fmt.Println("开始演示calc包的功能：")

	// 加法运算
	resultAdd := calc.Add(10, 5)
	fmt.Printf("10 + 5 = %.2f\n", resultAdd)

	// 减法运算
	resultSub := calc.Subtract(10, 5)
	fmt.Printf("10 - 5 = %.2f\n", resultSub)

	// 乘法运算
	resultMul := calc.Multiply(10, 5)
	fmt.Printf("10 * 5 = %.2f\n", resultMul)

	// 除法运算
	resultDiv, err := calc.Divide(10, 5)
	if err != nil {
		fmt.Printf("除法运算错误: %v\n", err)
	} else {
		fmt.Printf("10 / 5 = %.2f\n", resultDiv)
	}

	// 测试除零错误
	_, err = calc.Divide(10, 0)
	if err != nil {
		fmt.Printf("测试除零运算: %v\n", err)
	}
}
*/
//因地制宜，我们直接在主程序 main.go 根据模块路径导入snow包，然后调用 CalculateExample 函数。如有差异可尝试命令 go mod tidy

// 十一、接口
// 每个接口类型由任意个方法签名组成，接口的定义格式如下：
/*
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
*/
// 其中：
// 接口类型名：Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有关闭操作的接口叫closer等。接口名最好要能突出该接口的类型含义。
// 方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
// 参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。
