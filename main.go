package main

import (
	"fmt"
	"math"
)

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	dummy := &listNode{}
// 	current := dummy
// 	for list1 != nill && list2 != nill {
// 		if list1.Val < list2.Val {
// 			current.Next = list1
// 			list1 = list1.Next
// 		}
// 		else{
// 			current.Next = list2
// 			list2 = list2.Next
// 		}
// 		current = current.Next
// 	}
// 	if list1 != nil {
//         current.Next = list1
//     } else {
//         current.Next = list2
//     }

//     // Return the merged list starting from the next node after dummy
//     return dummy.Next
// }

// func topKFrequent(nums []int, k int) []int {
//     frequencyMap := Make(map(int)int)
// 	for _, num := range nums {
// 		frequencyMap[num]++
// 	}
// 	bucket := make([][]int, len(nums)+1)
//     for num, freq := range frequencyMap {
//         bucket[freq] = append(bucket[freq], num)
//     }

//    result := make([]int, 0, k)
//     for i := len(bucket) - 1; i >= 0 && len(result) < k; i-- {
//         if len(bucket[i]) > 0 {
//             result = append(result, bucket[i]...)
//         }
//     }

//	    return result[:k]
//	}
type Circle struct {
	radius float64
}
type Employee struct {
	name     string
	age      int
	position string
}

func tinhcvtron(circle Circle) float64 {
	return 2 * circle.radius * math.Pi
}
func tinhdttron(circle Circle) float64 {
	return circle.radius * circle.radius * math.Pi
}
func getHighestAge(employee []Employee) string {
	result := ""
	highestAge := 0
	for _, e := range employee {
		if e.age > highestAge {
			highestAge = e.age
			result = e.name
		}
	}
	return result
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Vehicle interface {
	Start()
	Stop()
}
type Car struct {
	name string
}
type Bike struct {
	name string
}
type Rectangle struct {
	width   float64
	heaight float64
}

func (circle Circle) Area() float64 {
	return 3.14 * circle.radius * circle.radius
}
func (rectangle Rectangle) Area() float64 {
	return rectangle.width * rectangle.heaight
}
func (rectangle Rectangle) Perimeter() float64 {
	return (rectangle.width + rectangle.heaight) * 2
}
func (circle Circle) Perimeter() float64 {
	return 2 * circle.radius * math.Pi
}
func (car Car) Start() {
	fmt.Println(car.name + " start...")
}

func (car Car) Stop() {
	fmt.Println(car.name + " stop...")
}

func (bike Bike) Start() {
	fmt.Println(bike.name + " start...")
}

func (bike Bike) Stop() {
	fmt.Println(bike.name + " stop...")
}
func Change(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
func GiaiThua(n int) int {
	if n == 0 {
		return 1
	}
	return n * GiaiThua(n-1)
}
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
func PrintFibonacci(n int) {
	if n <= 0 {
		return
	}
	PrintFibonacci(n - 1)          // Gọi đệ quy trước
	fmt.Print(Fibonacci(n-1), " ") // In ra số Fibonacci thứ n-1
}
func divide(a, b int) int {
	if b == 0 {
		panic("Lỗi: chia cho 0 không hợp lệ!") // kích hoạt panic
	}
	return a / b
}
func sayGoodbye() {
	fmt.Println("Goodbye!")
}

func greetAndPanic() {
	// Đặt defer để đảm bảo "Goodbye!" được in ra khi hàm kết thúc
	defer sayGoodbye()

	fmt.Println("Hello!")

	// Gọi panic để mô phỏng lỗi xảy ra
	panic("Panic xảy ra")
}
func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func Max(nums ...int) int {
	if len(nums) == 0 {
		// Trả về giá trị mặc định nếu không có số nào
		return 0
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
func main() {
	employees := []Employee{
		{"Alice", 30, "Engineer"},
		{"Bob", 45, "Manager"},
		{"Charlie", 40, "Designer"},
	}

	car := Car{"G63"}
	bike := Bike{"Martin"}
	rectangle := Rectangle{float64(4), float64(6)}
	circle := Circle{float64(5)}
	fmt.Println("câu 1: " + getHighestAge(employees))
	fmt.Println("câu 2 chu vi, diện tích: "+fmt.Sprintf("%.2f", tinhcvtron(circle)), tinhdttron(circle))
	fmt.Println("Interface 1 Area: "+fmt.Sprintf("%.2f", circle.Area()), circle.Perimeter())
	fmt.Println("Interface 1 Area retangle: "+fmt.Sprintf("%.2f", rectangle.Area()), rectangle.Perimeter())
	car.Start()
	car.Stop()
	bike.Start()
	bike.Stop()
	z := 10
	k := 5

	Change(&z, &k)
	fmt.Println("Sau hoán đổi: z =", z, "k =", k)
	fmt.Println("5! =", GiaiThua(5))
	n := 10
	fmt.Printf("%d số đầu tiên trong dãy Fibonacci: ", n)
	PrintFibonacci(n)
	// fmt.Println(divide(10, 0))
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Đã phục hồi từ panic:", r)
	// 	}
	// }()

	// greetAndPanic()
	fmt.Println("Tổng của các số 1, 2, 3, 4:", Sum(1, 2, 3, 4))
	fmt.Println("Giá trị lớn nhất trong 3, 5, 7, 2:", Max(3, 5, 7, 2)) // Kết quả: 7
}
