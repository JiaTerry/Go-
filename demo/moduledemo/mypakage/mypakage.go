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
