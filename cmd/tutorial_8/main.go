package main

import "fmt"
import "time"
import "math/rand"

func main() {
    var c = make(chan int, 5)
    go process(c)

    for i := range c {
        fmt.Println(i)
        time.Sleep(time.Second*1)
    }

    chickenExample()
    
}

func process(c chan int) {
    defer close(c)
    for i := 0; i < 5; i++ {
        c <- i
    }
    fmt.Println("Exiting the process")
}


func chickenExample() {
    var chickenChannel = make(chan string)
    var tofuChannel = make(chan string)

    var websites = []string {"walmart", "costo", "wholefoods"}

    for i := range websites {
        go checkChickenPrices(websites[i], chickenChannel)
        go checkTofuPrices(websites[i], tofuChannel)
    }
    sendMessage(chickenChannel, tofuChannel)
}

const MAX_CHICKEN_PRICE float32 = 5
const MAX_TOFU_PRICE float32 = 2

func checkChickenPrices(
    website string, 
    chickenChannel chan string,
    ) {
    for {
        time.Sleep(time.Second*1)
        var chickenPrice = rand.Float32()*20
        if chickenPrice <= MAX_CHICKEN_PRICE {
            chickenChannel <- website
            break
        }
    }
}

func checkTofuPrices(website string, tofuChannel chan string) {
    for {
        time.Sleep(time.Second*1)
        var tofuPrice = rand.Float32()*20
        if  tofuPrice <= MAX_TOFU_PRICE {
            tofuChannel <- website
            break
        }
    }
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
    select {
    case website := <-tofuChannel:
        fmt.Printf("\nTofu found, %v", website)
    case website := <-chickenChannel:
        fmt.Printf("\nChicken found, %v", website)
    }
}