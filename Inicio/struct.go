package main

import "fmt"

type Person struct{
    Name string
    Age int
}

func (p *Person) SayHello(){
    fmt.Println("Hello,", p.Name)
}

func main(){
    var guy = new(Person)
    guy.Name = "Henrique"
    guy.SayHello()
}
