package main

import (
	"fmt"
	"time"
)

type Item struct {
	name   string
	price  int
	amount int
}

type Buyer struct {
	point  int
	bucket []*Item
}

type Truck struct {
	status   string
	packages []*Item
}

type Delivery struct {
	numOrder     int
	deliveryList []*Item
	trucks       []*Truck
}

func NewBuyer() *Buyer {
	buyer := Buyer{}
	buyer.point = 1000000
	buyer.bucket = make([]*Item, 0, 5)
	return &buyer
}

func NewDelivery() *Delivery {
	delivery := Delivery{}
	delivery.numOrder = 0
	delivery.deliveryList = make([]*Item, 0, 5)
	delivery.trucks = make([]*Truck, 5)
	for i := range delivery.trucks {
		delivery.trucks[i] = &Truck{"", make([]*Item, 0, 5)}
	}
	return &delivery
}

func ReturnToMenu() {
	fmt.Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
	fmt.Scanln()
}

func PrintItems(items []*Item) {
	for i, v := range items {
		fmt.Printf("상품%d: %s, 가격: %d원, 잔여 수량: %d개\n", i+1, v.name, v.price, v.amount)
	}
}

func ChoiceItem(items []*Item, buyer *Buyer) (item *Item, buyAmount int, err error) {
	for {
		fmt.Print("구매할 상품을 선택하세요: ")
		itemChoice := 0
		fmt.Scanln(&itemChoice)

		if itemChoice >= 1 && itemChoice <= 5 {
			item = items[itemChoice-1]
			break
		} else {
			fmt.Println("잘못된 입력입니다. 다시 입력해주세요.")
		}
	}
	fmt.Print("구매할 수량을 입력하세요: ")
	fmt.Scanln(&buyAmount)
	fmt.Println()

	if buyAmount <= 0 {
		err = fmt.Errorf("올바른 수량을 입력하세요")
	} else if buyAmount > item.amount {
		err = fmt.Errorf("남은 수량이 부족합니다")
	} else if item.price*buyAmount > buyer.point {
		err = fmt.Errorf("마일리지가 부족합니다")
	}
	return
}

func Contains(items []*Item, item *Item) (bool, int) {
	for i, v := range items {
		if v.name == item.name {
			return true, i
		}
	}
	return false, -1
}

func BuyItem(buyer *Buyer, item *Item, buyAmount int, delivery *Delivery, deliveryChan chan bool) (err error) {
	for {
		fmt.Println("1. 바로 구매")
		fmt.Println("2. 장바구니에 담기")
		fmt.Print("구매할 방법을 선택하세요: ")
		buyChoice := 0
		fmt.Scanln(&buyChoice)
		fmt.Println()

		switch buyChoice {
		case 1:
			if delivery.numOrder <= 5 {
				buyer.point -= buyAmount * item.price
				item.amount -= buyAmount
				fmt.Println("상품의 주문이 접수되었습니다.")

				delivery.deliveryList = append(delivery.deliveryList, &Item{item.name, item.price, buyAmount})
				deliveryChan <- true
				delivery.numOrder++
			} else {
				err = fmt.Errorf("배송 한도를 초과했습니다. 배송이 완료되면 주문하세요")
			}
			return
		case 2:
			if isContain, index := Contains(buyer.bucket, item); isContain {
				if buyer.bucket[index].amount+buyAmount > item.amount {
					err = fmt.Errorf("잔여 수량을 초과했습니다")
				} else {
					buyer.bucket[index].amount += buyAmount
				}
			} else {
				buyer.bucket = append(buyer.bucket, &Item{item.name, item.price, buyAmount})
			}
			fmt.Println("상품이 장바구니에 추가되었습니다.")
			return
		default:
			fmt.Println("잘못된 입력입니다. 다시 입력해주세요.")
		}
	}
}

func CheckRemainingMileage(buyer *Buyer) {
	fmt.Printf("현재 잔여 마일리지는 %d점입니다.\n", buyer.point)
}

func CheckDeliveryStatus(delivery *Delivery) {
	total := 0
	for _, v := range delivery.trucks {
		total += len(v.packages)
	}
	if total == 0 {
		fmt.Println("배송 중인 상품이 없습니다.")
	} else {
		for _, v := range delivery.trucks {
			if len(v.packages) != 0 {
				for _, val := range v.packages {
					fmt.Printf("%s %d개\n", val.name, val.amount)
				}
				fmt.Printf("배송 상황: %s\n", v.status)
				fmt.Println()
			}
		}
	}
}

func PrintBucket(bucket []*Item) {
	for _, v := range bucket {
		fmt.Printf("상품: %s, 수량: %d\n", v.name, v.amount)
	}
}

func CheckBucketEmpty(bucket []*Item) (isEmpty bool) {
	if len(bucket) == 0 {
		fmt.Println("장바구니가 비었습니다.")
		isEmpty = true
	}
	return
}

func BucketTotal(bucket []*Item) (bucketTotal int) {
	for _, v := range bucket {
		bucketTotal += v.price * v.amount
	}
	return
}

func IsAmountOver(items, bucket []*Item) (isAmountOver bool, overList []Item) {
	for _, v := range bucket {
		if _, index := Contains(items, v); v.amount > items[index].amount {
			isAmountOver = true
			overList = append(overList, Item{v.name, v.price, v.amount - items[index].amount})
		}
	}
	return isAmountOver, overList
}

func ClearBucket(buyer *Buyer) {
	buyer.bucket = make([]*Item, 0, 5)
}

func BuyBucket(buyer *Buyer, items []*Item, delivery *Delivery, deliveryChan chan bool) (err error) {
	if bucketTotal := BucketTotal(buyer.bucket); bucketTotal > buyer.point {
		err = fmt.Errorf("마일리지가 %d점 부족합니다", bucketTotal-buyer.point)
	} else if isAmountOver, overList := IsAmountOver(items, buyer.bucket); isAmountOver {
		var errStr string
		for _, v := range overList {
			errStr += fmt.Sprintf("%s가 %d개 초과했습니다\n", v.name, v.amount)
		}
		err = fmt.Errorf(errStr)
	} else if len(buyer.bucket) == 0 {
		err = fmt.Errorf("주문 가능한 목록이 없습니다")
	} else if delivery.numOrder > 5 {
		err = fmt.Errorf("배송 한도를 초과했습니다. 배송이 완료되면 주문하세요")
	} else {
		for _, v := range buyer.bucket {
			delivery.deliveryList = append(delivery.deliveryList, &Item{v.name, v.price, v.amount})
			buyer.point -= v.amount * v.price
			_, index := Contains(items, v)
			items[index].amount -= v.amount
		}
		fmt.Println("상품의 주문이 접수되었습니다.")

		deliveryChan <- true
		delivery.numOrder++

		ClearBucket(buyer)
	}
	return
}

func DeliveryStatus(deliveryChan chan bool, delivery *Delivery, i int) {
	for {
		if <-deliveryChan {
			for _, v := range delivery.deliveryList {
				delivery.trucks[i].packages = append(delivery.trucks[i].packages, v)
			}
			delivery.deliveryList = make([]*Item, 0, 5)

			delivery.trucks[i].status = "주문 접수"
			time.Sleep(10 * time.Second)

			delivery.trucks[i].status = "배송 중"
			time.Sleep(30 * time.Second)

			delivery.trucks[i].status = "배송 완료"
			time.Sleep(10 * time.Second)

			delivery.trucks[i].status = ""
			delivery.numOrder--
			delivery.trucks[i].packages = make([]*Item, 0, 5)
		}
	}
}

func main() {
	items := make([]*Item, 5)
	buyer := NewBuyer()
	delivery := NewDelivery()
	deliveryChan := make(chan bool)

	items[0] = &Item{"텀블러", 10000, 30}
	items[1] = &Item{"내셔널지오그래픽 롱패딩", 500000, 20}
	items[2] = &Item{"디스커버리 백팩", 400000, 20}
	items[3] = &Item{"나이키 운동화", 150000, 50}
	items[4] = &Item{"빼빼로", 1200, 500}

	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond)
		go DeliveryStatus(deliveryChan, delivery, i)
	}

	for {
		fmt.Println("1. 상품 확인")
		fmt.Println("2. 상품 구매")
		fmt.Println("3. 잔여 마일리지 확인")
		fmt.Println("4. 배송 상태 확인")
		fmt.Println("5. 장바구니 확인")
		fmt.Println("6. 장바구니 초기화")
		fmt.Println("7. 장바구니 상품 주문")
		fmt.Println("8. 프로그램 종료")
		fmt.Print("실행할 기능을 입력하세요: ")
		menuChoice := 0
		fmt.Scanln(&menuChoice)
		fmt.Println()

		switch menuChoice {
		case 1: // 상품 확인
			PrintItems(items)
			ReturnToMenu()
		case 2: // 상품 구매
			if item, buyAmount, err := ChoiceItem(items, buyer); err != nil {
				fmt.Println(err)
				ReturnToMenu()
				break
			} else if err := BuyItem(buyer, item, buyAmount, delivery, deliveryChan); err != nil {
				fmt.Println(err)
			}
			ReturnToMenu()
		case 3: // 잔여 마일리지 확인
			CheckRemainingMileage(buyer)
			ReturnToMenu()
		case 4: // 배송 상태 확인
			CheckDeliveryStatus(delivery)
			ReturnToMenu()
		case 5: // 장바구니 확인
			if isEmpty := CheckBucketEmpty(buyer.bucket); isEmpty {
				ReturnToMenu()
				break
			} else {
				PrintBucket(buyer.bucket)
			}
			ReturnToMenu()
		case 6: // 장바구니 초기화
			ClearBucket(buyer)
			ReturnToMenu()
		case 7: // 장바구니 상품 주문
			if err := BuyBucket(buyer, items, delivery, deliveryChan); err != nil {
				fmt.Println(err)
			}
			ReturnToMenu()
		case 8: // 프로그램 종료
			fmt.Print("프로그램을 종료합니다.")
			return
		default:
			fmt.Println("잘못된 입력입니다. 다시 입력해주세요.")
		}
	}
}
