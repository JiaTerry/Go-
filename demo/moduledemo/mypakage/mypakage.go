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

// 十一、接口 *****
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

//值接收者实现接口
//使用值接收者实现接口之后，不管是结构体类型还是对应的结构体指针类型的变量都可以赋值给该接口变量。

//指针接收者实现接口
//由于Go语言中有对指针求值的语法糖，对于值接收者实现的接口，无论使用值类型还是指针类型都没有问题。但是我们并不总是能对一个值求址，所以对于指针接收者实现的接口要额外注意。

//类型和接口的关系
//一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。

//多种类型实现同一接口
//一个接口的所有方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。

//接口组合
//接口与接口之间可以通过互相嵌套形成新的接口类型
// 对于这种由多个接口类型组合形成的新接口类型，同样只需要实现新接口类型中规定的所有方法就算实现了该接口类型。
// 接口也可以作为结构体的一个字段，通过在结构体中嵌入一个接口类型，从而让该结构体类型实现了该接口类型，并且还可以改写该接口的方法。

//空接口
// 空接口是指没有定义任何方法的接口类型。因此任何类型都可以视为实现了空接口。也正是因为空接口类型的这个特性，空接口类型的变量可以存储任意类型的值。

//空接口的应用
//1、空接口作为函数的参数
//使用空接口实现可以接收任意类型的函数参数。
//2、空接口作为map的值
//使用空接口实现可以保存任意值的字典。

//接口值
//由于接口类型的值可以是任意一个实现了该接口的类型值，所以接口值除了需要记录具体值之外，还需要记录这个值属于的类型。
//也就是说接口值由“类型”和“值”组成，鉴于这两部分会根据存入值的不同而发生变化，我们称之为接口的动态类型和动态值。
//**注意：**我们不能对一个空接口值调用任何方法，否则会产生panic。
//接口值是支持相互比较的，当且仅当接口值的动态类型和动态值都相等时才相等。
//但是有一种特殊情况需要特别注意，如果接口值保存的动态类型相同，但是这个动态类型不支持互相比较（比如切片），那么对它们相互比较时就会引发panic。

//类型断言
//想要从接口值中获取到对应的实际值需要使用类型断言，其语法格式如下。
//x.(T)
// 其中：
// x：表示接口类型的变量
// T：表示断言x可能是的类型。
// 该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。
// 如果对一个接口值有多个实际类型需要判断，推荐使用switch语句来实现。
/*
由于接口类型变量能够动态存储不同类型值的特点，所以很多初学者会滥用接口类型（特别是空接口）来实现编码过程中的便捷。
只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。切记不要为了使用接口类型而增加不必要的抽象，导致不必要的运行时损耗。

在 Go 语言中接口是一个非常重要的概念和特性，使用接口类型能够实现代码的抽象和解耦，也可以隐藏某个功能的内部实现，
但是缺点就是在查看源码的时候，不太方便查找到具体实现接口的类型。

请牢记接口是一种类型，一种抽象的类型。
*/

// 小技巧： 下面的代码可以在程序编译阶段验证某一结构体是否满足特定的接口类型。
// 摘自gin框架routergroup.go
/*
type IRouter interface{ ... }

type RouterGroup struct { ... }

var _ IRouter = &RouterGroup{}  // 确保RouterGroup实现了接口IRouter
*/
// 上面的代码中也可以使用var _ IRouter = (*RouterGroup)(nil)进行验证。

//十一、练习题 - 接口
// 使用接口的方式实现一个既可以往终端写日志也可以往文件写日志的简易日志库。

//十二、Error 接口
// Go 语言中把错误当成一种特殊的值来处理，不支持其他语言中使用try/catch捕获异常的方式。
//由于 error 是一个接口类型，默认零值为nil。所以我们通常将调用函数返回的错误与nil进行比较，以此来判断函数是否返回错误。
// 注意：当我们使用fmt包打印错误时会自动调用 error 类型的 Error 方法，也就是会打印出错误的描述信息。

//创建错误 errors.New  fmt.Errorf
// func New(text string) error
// 它接收一个字符串参数返回包含该字符串的错误。
// 当我们需要传入格式化的错误描述信息时，使用fmt.Errorf是个更好fmt.Errorf("查询数据库失败，err:%v", err)
// fmt.Errorf("查询数据库失败，err:%v", err)
// 但是上面的方式会丢失原有的错误类型，只拿到错误描述的文本信息。
// 为了不丢失函数调用的错误链，使用fmt.Errorf时搭配使用特殊的格式化动词%w，可以实现基于已有的错误再包装得到一个新的错误。
// fmt.Errorf("查询数据库失败，err:%w", err)
//对于这种二次包装的错误，errors包中提供了以下三个方法。
/*
func Unwrap(err error) error                 // 获得err包含下一层错误
func Is(err, target error) bool              // 判断err是否包含target
func As(err error, target interface{}) bool  // 判断err是否为target类型
*/

//错误结构体类型
// 此外我们还可以自己定义结构体类型，实现error接口。
/*
// OpError 自定义结构体类型
type OpError struct {
	Op string
}

// Error OpError 类型实现error接口
func (e *OpError) Error() string {
	return fmt.Sprintf("无权执行%s操作", e.Op)
}
*/

// 十三、反射
// Go语言中的变量是分为两部分的:
// 类型信息：预先定义好的元信息。
// 值信息：程序运行过程中可动态变化的。

/*
反射是指在程序运行期间对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。
在运行程序时，程序无法获取自身的信息。

支持反射的语言可以在程序编译期间将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，
这样就可以在程序运行期间获取类型的反射信息，并且有能力修改它们。

Go程序在运行期间使用reflect包访问程序的反射信息。

反射就是在运行时动态的获取一个变量的类型信息和值信息。
*/

// 在Go语言的反射机制中，任何接口值都由是一个具体类型和具体类型的值两部分组成的
//在Go语言中反射的相关功能由内置的reflect包提供，任意接口值在反射中都可以理解为由reflect.Type和reflect.Value两部分组成，
//并且reflect包提供了reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Value和Type。

//TypeOf
// 在Go语言中，使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息。

// type name和type kind
// 在反射中关于类型还划分为两种：类型（Type）和种类（Kind）。
// 因为在Go语言中我们可以使用type关键字构造很多自定义类型，而种类（Kind）就是指底层的类型，但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类（Kind）。

//ValueOf
// reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的值信息。reflect.Value与原始值之间可以互相转换。
// reflect.Value类型提供的获取原始值的方法如下：
/*
方法	                    说明
Interface() interface {}	将值以 interface{} 类型返回，可以通过类型断言转换为指定类型
Int() int64	                将值以 int 类型返回，所有有符号整型均可以此方式返回
Uint() uint64	            将值以 uint 类型返回，所有无符号整型均可以此方式返回
Float() float64	            将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回
Bool() bool	                将值以 bool 类型返回
Bytes() []bytes	            将值以字节数组 []bytes 类型返回
String() string	            将值以字符串类型返回
*/

//通过反射获取值
//通过反射设置变量的值
// 想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。而反射中使用专有的Elem()方法来获取指针对应的值。

// isNil()和isValid()
// 1、isNil()
// func (v Value) IsNil() bool
// IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。
// 2、isValid()
// func (v Value) IsValid() bool
// IsValid()返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。
// IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。

//结构体反射
//任意值通过reflect.TypeOf()获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象（reflect.Type）的NumField()和Field()方法获得结构体成员的详细信息。
// reflect.Type中与获取结构体成员相关的的方法如下表所示。
/*
方法															说明
Field(i int) StructField										根据索引，返回索引对应的结构体字段的信息。
NumField() int													返回结构体成员字段数量。
FieldByName(name string) (StructField, bool)					根据给定字符串返回字符串对应的结构体字段的信息。
FieldByIndex(index []int) StructField							多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息。
FieldByNameFunc(match func(string) bool) (StructField,bool)		根据传入的匹配函数匹配需要的字段。
NumMethod() int													返回该类型的方法集中方法的数目
Method(int) Method												返回该类型方法集中的第i个方法
MethodByName(string)(Method, bool)								根据方法名返回该类型方法集中的方法
*/

//StructField类型
// StructField类型用来描述结构体中的一个字段的信息。
// StructField的定义如下：
/*
type StructField struct {
    // Name是字段的名字。PkgPath是非导出字段的包路径，对导出字段该字段为""。
    // 参见http://golang.org/ref/spec#Uniqueness_of_identifiers
    Name    string
    PkgPath string
    Type      Type      // 字段的类型
    Tag       StructTag // 字段的标签
    Offset    uintptr   // 字段在结构体中的字节偏移量
    Index     []int     // 用于Type.FieldByIndex时的索引切片
    Anonymous bool      // 是否匿名字段
}
*/

//结构体反射示例
//当我们使用反射得到一个结构体数据之后可以通过索引依次获取其字段信息，也可以通过字段名去获取指定的字段信息。

// 反射是把双刃剑
// 反射是一个强大并富有表现力的工具，能让我们写出更灵活的代码。但是反射不应该被滥用，原因有以下三个。
/*
1、基于反射的代码是极其脆弱的，反射中的类型错误会在真正运行的时候才会引发panic，那很可能是在代码写完的很长时间之后。
2、大量使用反射的代码通常难以理解。
3、反射的性能低下，基于反射实现的代码通常比正常代码运行速度慢一到两个数量级。
*/

// 十三、练习题 - 反射
// 编写代码利用反射实现一个ini文件的解析器程序。

//十四、并发 ******
//串行：顺序执行，一个接着一个
//并发：同一时间段执行多个任务，交替做多个任务
//并行：同一时刻执行多个任务，同时做多个任务
//进程（process）：程序在操作系统中的一次执行过程，系统进行资源分配和调度的一个独立单位
//线程（thread）：操作系统基于进程开启的轻量级进程，是操作系统调度执行的最小单位
//协程（coroutine）：非操作系统提供而是用用户自行创建和控制的用户态“线程”，比线程更加轻量级

//并发模型
//1、线程&锁模型
//2、Actor模型
//3、CSP模型
//4、Fork&Join模型
//Go语言中的并发程序主要是通过CSP的goroutine和channel 来实现，当然也支持使用传统的多线程共享内存的并发方式

//goroutine
//1、Goroutine 是 Go 语言支持并发的核心，在一个Go程序中同时创建成百上千个goroutine是非常普遍的，一个goroutine会以一个很小的栈开始其生命周期，一般只需要2KB。区别于操作系统线程由系统内核进行调度， goroutine 是由Go运行时（runtime）负责调度。
//2、Goroutine 是 Go 程序中最基本的并发执行单元。每一个 Go 程序都至少包含一个 goroutine——main goroutine，当 Go 程序启动时它会自动创建。
//3、在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能——goroutine，当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个 goroutine 去执行这个函数就可以了

//go关键字
//Go语言中使用 goroutine 非常简单，只需要在函数或方法调用前加上go关键字就可以创建一个 goroutine ，从而让该函数或方法在新创建的 goroutine 中执行。
//go f()  // 创建一个新的 goroutine 运行函数f
//匿名函数也支持使用go关键字创建 goroutine 去执行。
/*
go func(){
  // ...
}()
*/
//一个 goroutine 必定对应一个函数/方法，可以创建多个 goroutine 去执行相同的函数/方法。

//启动单个goroutine
//启动 goroutine 的方式非常简单，只需要在调用函数（普通函数和匿名函数）前加上一个go关键字。
//但是容易出现一个问题，主goroutine退出太快，可能导致某些函数还未执行，但是程序已经结束了。
//解决方案：1、使用Sleep，但是它时间难以精确控制，效率低度低。2、使用sync.WaitGroup，（sync是一个包）使用WaitGroup可以等待所有 goroutine 执行完毕。
//启动多个goroutine
// 关键特性：执行顺序不确定，多次运行会发现打印顺序不同，因为：
// goroutine是并发执行的；Go运行时的调度是随机的、非确定的，这是正常现象，体现了真正的并发
//总结：用go启动，用WaitGroup等待，记住main退出则全结束，执行顺序不保证。

//动态栈
//操作系统的线程一般都有固定的栈内存（通常为2MB），而Go语言的goroutine非常轻量级，一个goroutine的初始栈空间很小（一般为2KB），所以在Go语言中一次创建数万个goroutine也是可能的。
//并且goroutine的栈不是固定的，可以根据需要动态的增加或者缩小，Go的runtime会自动为goroutine分配合适的栈空间。

// goroutine调度
// 操作系统内核在调度时会挂起当前正在执行的线程并将寄存器中的内容保存到内存中，然后选出接下来要执行的线程并从内存中恢复该线程的寄存器信息，
// 然后恢复执行该线程的现场并开始执行线程。从一个线程切换到另一个线程需要完整的上下文切换。因为可能需要多次内存访问，
// 索引这个切换上下文的操作开销较大，会增加运行的cpu周期。

// 区别于操作系统内核调度操作系统线程，goroutine 的调度是Go语言运行时（runtime）层面的实现，是完全由 Go 语言本身实现的一套调度系统——go scheduler。
// 它的作用是按照一定的规则将所有的 goroutine 调度到操作系统线程上执行。

// 在经历数个版本的迭代之后，目前 Go 语言的调度器采用的是 GPM 调度模型。
//核心思想比喻：GPM调度像高效物流：每个配送站（P）管理本地包裹（G），快递员（M）优先送本站包裹，本站没货时去其他站"借"或去总仓（全局队列）取，避免任何快递员闲着。

//GOMAXPROCS
//Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个 OS 线程来同时执行 Go 代码。默认值是机器上的 CPU 核心数。
//Go语言中可以通过runtime.GOMAXPROCS函数设置当前程序并发时占用的 CPU逻辑核心数。

// 十四、练习题 - 并发
// 请写出下面程序的执行结果。
/*
for i := 0; i < 5; i++ {
	go func() {
		fmt.Println(i)
	}()
}
*/
//预测最有可能的答案是 5 5 5 5 5

//channel
// 单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。

// 虽然可以使用共享内存进行数据交换，但是共享内存在不同的 goroutine 中容易发生竞态问题。
// 为了保证数据交换的正确性，很多并发模型中必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。

// Go语言采用的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。

// 如果说 goroutine 是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制。

// Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。
// 每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

//channel类型
//var 变量名称 chan 元素类型
// 其中：
// chan：是关键字
// 元素类型：是指通道中传递元素的类型

//channel零值
//未初始化的通道类型变量其默认零值是nil

//初始化channel
//声明的通道类型变量需要使用内置的make函数初始化之后才能使用。
//make(chan 元素类型, [缓冲大小])
// 其中：
// channel的缓冲大小是可选的。

//channel操作
//通道共有发送（send）、接收(receive）和关闭（close）三种操作。而发送和接收操作都使用<-符号。
//首先定义一个通道
//ch := make(chan int)
//发送
// ch <- 10 // 把10发送到ch中
//接收
// x := <- ch // 从ch中接收值并赋值给变量x
// <-ch       // 从ch中接收值，忽略结果
//关闭
//close(ch)
// **注意：**一个通道值是可以被垃圾回收掉的。通道通常由发送方执行关闭操作，并且只有在接收方明确等待通道关闭的信号时才需要执行关闭操作。
// 它和关闭文件不一样，通常在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。

// 关闭后的通道有以下特点：
/*
1、对一个关闭的通道再发送值就会导致 panic。
2、对一个关闭的通道进行接收会一直获取值直到通道为空。
3、对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
4、关闭一个已经关闭的通道会导致 panic。
*/

//无缓冲的通道
//无缓冲的通道又称为阻塞的通道。
//deadlock表示我们程序中的 goroutine 都被挂起导致程序死锁了。
//无缓冲的通道只有在有接收方能够接收值的时候才能发送成功，否则会一直处于等待发送的阶段。
// 同理，如果对一个无缓冲通道执行接收操作时，没有任何向通道中发送值的操作那么也会导致接收操作阻塞。
// 简单来说就是无缓冲的通道必须有至少一个接收方才能发送成功。
//使用无缓冲通道进行通信将导致发送和接收的 goroutine 同步化。因此，无缓冲通道也被称为同步通道。

//有缓冲的通道
//我们可以在使用 make 函数初始化通道时，可以为其指定通道的容量
//只要通道的容量大于零，那么该通道就属于有缓冲的通道，通道的容量表示通道中最大能存放的元素数量。
// 当通道内已有元素数达到最大容量后，再向通道执行发送操作就会阻塞，除非有从通道执行接收操作。
//我们可以使用内置的len函数获取通道内元素的数量，使用cap函数获取通道的容量，虽然我们很少会这么做。

//多返回值模式
//对一个通道执行接收操作时支持使用如下多返回值模式。
//value, ok := <- ch
// 其中：
/*
value：从通道中取出的值，如果通道被关闭则返回对应类型的零值。
ok：通道ch关闭时返回 false，否则返回 true。
*/

//for range接收值
//通常我们会选择使用for range循环从通道中接收值，当通道被关闭后，会在通道内的所有值被接收完毕后会自动退出循环。
//**注意：**目前Go语言中并没有提供一个不对通道进行读取操作就能判断通道是否被关闭的方法。不能简单的通过len(ch)操作来判断通道是否被关闭。

//单向通道
//在某些场景下我们可能会将通道作为参数在多个任务函数间进行传递，通常我们会选择在不同的任务函数中对通道的使用进行限制，比如限制通道在某个函数中只能执行发送或只能执行接收操作。
//Go语言中提供了单向通道来处理这种需要限制通道只能进行某种操作的情况。
//在函数传参及任何赋值操作中全向通道（正常通道）可以转换为单向通道，但是无法反向转换。
//**注意：**对已经关闭的通道再执行 close 也会引发 panic。

// Go 通道操作说明
// 一、通道的四种状态：
/*
nil：未初始化的通道

没值：通道为空，无数据可读

有值：通道中有数据可读

满：缓冲通道已满
*/

// 二、发送操作：
/*
nil 通道 → 阻塞

没值（空缓冲） → 发送成功

有值（有数据） → 发送成功

满（缓冲满） → 阻塞
*/

// 三、接收操作：
/*
nil 通道 → 阻塞

没值（空缓冲） → 阻塞

有值（有数据） → 接收成功

满（缓冲满） → 接收成功
*/

// 四、关闭操作：
/*
nil 通道 → panic

没值（空缓冲） → 关闭成功

有值（有数据） → 关闭成功

满（缓冲满） → 关闭成功
*/

// 重要提醒：
/*
1、对 nil 通道进行发送或接收会永久阻塞
2、关闭 nil 通道会导致 panic
3、向已关闭的通道发送数据会 panic
4、关闭已关闭的通道也会 panic
*/

//select多路复用
// Select 语句具有以下特点：
/*
1、可处理一个或多个 channel 的发送/接收操作。
2、如果多个 case 同时满足，select 会随机选择一个执行。
3、对于没有 case 的 select 会一直阻塞，可用于阻塞 main 函数，防止退出。
*/

//并发安全和锁
//可能会发生数据竞争的情况
// 关键点总结
/*
1、并发不安全：多个 goroutine 无保护地访问共享资源
2、结果不可预测：每次运行可能得到不同结果
3、需要同步机制：使用锁等工具保证同一时间只有一个 goroutine 访问临界区
*/
// 解决方案：使用互斥锁（sync.Mutex）或其它同步原语保护共享资源。

//互斥锁
//互斥锁是一种常用的控制共享资源访问的方法，它能够保证同一时间只有一个 goroutine 可以访问共享资源。Go 语言中使用sync包中提供的Mutex类型来实现互斥锁。
// sync.Mutex提供了两个方法供我们使用。
/*
方法名	                     功能
func (m *Mutex) Lock()	    获取互斥锁
func (m *Mutex) Unlock()	释放互斥锁
*/
//使用互斥锁能够保证同一时间有且只有一个 goroutine 进入临界区，其他的 goroutine 则在等待锁；
// 当互斥锁释放后，等待的 goroutine 才可以获取锁进入临界区，多个 goroutine 同时等待一个锁时，唤醒的策略是随机的。

// 读写互斥锁
//互斥锁是完全互斥的，但是实际上有很多场景是读多写少的，当我们并发的去读取一个资源而不涉及资源修改的时候是没有必要加互斥锁的，这种场景下使用读写锁是更好的一种选择。读写锁在 Go 语言中使用sync包中的RWMutex类型。
// sync.RWMutex提供了以下5个方法。
/*
方法名	                               功能
func (rw *RWMutex) Lock()	          获取写锁
func (rw *RWMutex) Unlock()	          释放写锁
func (rw *RWMutex) RLock()	          获取读锁
func (rw *RWMutex) RUnlock()	      释放读锁
func (rw *RWMutex) RLocker() Locker	  返回一个实现Locker接口的读写锁
*/
//读写锁分为两种：读锁和写锁。当一个 goroutine 获取到读锁之后，其他的 goroutine 如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
// 而当一个 goroutine 获取写锁之后，其他的 goroutine 无论是获取读锁还是写锁都会等待。
//使用读写互斥锁在读多写少的场景下能够极大地提高程序的性能。不过需要注意的是如果一个程序中的读操作和写操作数量级差别不大，那么读写互斥锁的优势就发挥不出来。

//sync.WaitGroup
// 在代码中生硬的使用time.Sleep肯定是不合适的，Go语言中可以使用sync.WaitGroup来实现并发任务的同步。 sync.WaitGroup有以下几个方法：
/*
方法名	                                 功能
func (wg * WaitGroup) Add(delta int)	计数器+delta
(wg *WaitGroup) Done()	                计数器-1
(wg *WaitGroup) Wait()	                阻塞直到计数器变为0
*/
//sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。
//需要注意sync.WaitGroup是一个结构体，进行参数传递的时候要传递指针。

//sync.Once
// 在某些场景下我们需要确保某些操作即使在高并发的场景下也只会被执行一次，例如只加载一次配置文件等。
// Go语言中的sync包中提供了一个针对只执行一次场景的解决方案——sync.Once，sync.Once只有一个Do方法，其签名如下：
//func (o *Once) Do(f func())
//**注意：**如果要执行的函数f需要传递参数就需要搭配闭包来使用。

//sync.Map
//Go 语言中内置的 map 不是并发安全的
// Go语言的sync包中提供了一个开箱即用的并发安全版 map——sync.Map。
// 开箱即用表示其不用像内置的 map 一样使用 make 函数初始化就能直接使用。
// 同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法。
/*
方法名																						功能
func (m *Map) Store(key, value interface{})												存储key-value数据
func (m *Map) Load(key interface{}) (value interface{}, ok bool)						查询key对应的value
func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)		查询或存储key对应的value
func (m *Map) LoadAndDelete(key interface{}) (value interface{}, loaded bool)			查询并删除key
func (m *Map) Delete(key interface{})													删除key
func (m *Map) Range(f func(key, value interface{}) bool)								对map中的每个key-value依次调用f
*/

//原子操作
//针对整数数据类型（int32、uint32、int64、uint64）我们还可以使用原子操作来保证并发安全，通常直接使用原子操作比使用锁操作效率更高。
// Go语言中原子操作由内置的标准库sync/atomic提供。

//atomic包
/*
方法																							解释
func LoadInt32(addr *int32) (val int32)															读取操作
func LoadInt64(addr *int64) (val int64)															读取操作
func LoadUint32(addr *uint32) (val uint32)														读取操作
func LoadUint64(addr *uint64) (val uint64)														读取操作
func LoadUintptr(addr *uintptr) (val uintptr)													读取操作
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)										读取操作

func StoreInt32(addr *int32, val int32)															写入操作
func StoreInt64(addr *int64, val int64)															写入操作
func StoreUint32(addr *uint32, val uint32)														写入操作
func StoreUint64(addr *uint64, val uint64)														写入操作
func StoreUintptr(addr *uintptr, val uintptr)													写入操作
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)										写入操作

func AddInt32(addr *int32, delta int32) (new int32)												修改操作
func AddInt64(addr *int64, delta int64) (new int64)												修改操作
func AddUint32(addr *uint32, delta uint32) (new uint32)											修改操作
func AddUint64(addr *uint64, delta uint64) (new uint64)											修改操作
func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)										修改操作

func SwapInt32(addr *int32, new int32) (old int32)												交换操作
func SwapInt64(addr *int64, new int64) (old int64)												交换操作
func SwapUint32(addr *uint32, new uint32) (old uint32)											交换操作
func SwapUint64(addr *uint64, new uint64) (old uint64)											交换操作
func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)										交换操作
func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)					交换操作

func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)							比较并交换操作
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)							比较并交换操作
func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)							比较并交换操作
func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)							比较并交换操作
func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)						比较并交换操作
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)		比较并交换操作
*/
//atomic包提供了底层的原子级内存操作，对于同步算法的实现很有用。这些函数必须谨慎地保证正确使用。除了某些特殊的底层应用，使用通道或者 sync 包的函数/类型实现同步更好。

// 十四、练习题 - 并发
// 练习题
// 1、使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序，例如生成随机数61345，计算其每个位数上的数字之和为19。
// 2、开启一个 goroutine 循环生成int64类型的随机数，发送到jobChan
// 3、开启24个 goroutine 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
// 4、主 goroutine 从resultChan取出结果并打印到终端输出

// 十五、处理并发错误 ******
//recover goroutine中的panic

// errgroup
// errgroup.Group 提供了Go和Wait两个方法。
// func (g *Group) Go(f func() error)
// 1、 Go 函数会在新的 goroutine 中调用传入的函数f。
// 2、 第一个返回非零错误的调用将取消该Group；下面的Wait方法会返回该错误
// func (g *Group) Wait() error
//Wait 会阻塞直至由上述 Go 方法调用的所有函数都返回，然后从它们返回第一个非nil的错误（如果有）。

// 十六、网络编程
