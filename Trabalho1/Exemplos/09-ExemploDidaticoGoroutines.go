package main
import (
    "fmt"
    "time"
)
func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}
func main() {
    go say("world")
    say("hello")
}

//Aqui se fizéssemos a  chamada da função sem a goroutine: say(“hello”), nada seria impresso
//já que essa rotina é independente da thread principal
//sendo assim, ela terminaria antes que houvesse o print dos valores.
