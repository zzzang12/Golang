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

func main() {
	items := make([]item, 5)
	buyer := NewBuyer()

	items[0] = item{"텀블러", 10000, 30}
	items[1] = item{"내셔널지오그래픽 롱패딩", 500000, 20}
	items[2] = item{"디스커버리 백팩", 400000, 20}
	items[3] = item{"나이키 운동화", 150000, 50}
	items[4] = item{"빼빼로", 1200, 500}

	for {
		menu := 0

		fmt.Println("1. 구매")
		fmt.Println("2. 잔여 수량 확인")
		fmt.Println("3. 잔여 마일리지 확인")
		fmt.Println("4. 배송 상태 확인")
		fmt.Println("5. 장바구니 확인")
		fmt.Println("6. 프로그램 종료")
		fmt.Print("실행할 기능을 입력하시오: ")

		fmt.Scanln(&menu)
		fmt.Println()

		if menu == 1 { // 물건 구매
			fmt.Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 2 { // 남은 수량 확인
			fmt.Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 3 { // 잔여 마일리지 확인
			fmt.Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 4 { // 배송 상태 확인
			fmt.Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 5 { // 장바구니 확인
			fmt.Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 6 { // 프로그램 종료
			fmt.Print("프로그램을 종료합니다.")
			return
		} else {
			fmt.Println("잘못된 입력입니다. 다시 입력해주세요.")
		}
	}
}
