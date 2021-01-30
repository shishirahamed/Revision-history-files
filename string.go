package main
import "fmt"

func main(){

line1 :="how are you"
line2 :="fine you"
fmt.Println(line1)
fmt.Println(line2)

//slicing string
sohel := line1[2:]
fmt.Println(sohel)

masud := line2[1:4]
fmt.Println(masud)

kajol := line1{3:}
fmt.Println(kajol)

character := line1{0}
fmt.Println(character)




}