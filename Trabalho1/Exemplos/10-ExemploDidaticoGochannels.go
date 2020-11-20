
package main
import (
    "fmt"
    "time"
)
func say(s string, done chan string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
    done <- "Terminei"
}
func main() {
    done := make(chan string)
    go say("world", done)
    fmt.Println(<-done)
}

//Aqui a thread principal aguarda uma mensagem da função say() para continuar
//função essa que está sendo executada na própria goroutine.
