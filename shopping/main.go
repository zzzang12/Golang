package main

import (
	"fmt"
)

type item struct {
	name   string
	price  int
	amount int
}

type buyer struct {
	point          int
	shoppingBucket map[string]int
}

func NewBuyer() *buyer {
	b := buyer{}
	b.point = 1000000
	b.shoppingBucket = map[string]int{}
	return &b
}

func ReturnToMenu() {
	fmt.Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
	fmt.Scanln()
}

func main() {
	items := make([]item, 5)
	buyer := NewBuyer()

	items[0] = item{"텀블러", 10000, 30}
	items[1] = item{"내셔널지오그래픽 롱패딩", 500000, 20}
	items[2] = item{"디스커버리 백팩", 400000, 20}
	items[3] = item{"나이키 운동화", 150000, 50}
	items[4] = item{"빼빼로", 1200, 500}

	for {
		menuChoice := 0

		fmt.Println("1. 상품 구매")
		fmt.Println("2. 잔여 수량 확인")
		fmt.Println("3. 잔여 마일리지 확인")
		fmt.Println("4. 배송 상태 확인")
		fmt.Println("5. 장바구니 확인")
		fmt.Println("6. 프로그램 종료")
		fmt.Print("실행할 기능을 입력하세요: ")

		fmt.Scanln(&menuChoice)
		fmt.Println()

		if menuChoice == 1 { // 상품 구매
			ReturnToMenu()
		} else if menuChoice == 2 { // 잔여 수량 확인
			for _, v := range items {
				fmt.Printf("%s의 잔여 수량은 %d개입니다.\n", v.name, v.amount)
			}
			ReturnToMenu()
		} else if menuChoice == 3 { // 잔여 마일리지 확인
			fmt.Printf("현재 잔여 마일리지는 %d점입니다.\n", buyer.point)
			ReturnToMenu()
		} else if menuChoice == 4 { // 배송 상태 확인
			ReturnToMenu()
		} else if menuChoice == 5 { // 장바구니 확인
			ReturnToMenu()
		} else if menuChoice == 6 { // 프로그램 종료
			fmt.Print("프로그램을 종료합니다.")
			return
		} else {
			fmt.Println("잘못된 입력입니다. 다시 입력해주세요.")
		}
	}
}
