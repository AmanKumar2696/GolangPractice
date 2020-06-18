package main
import "fmt"
func bytesandchar(str string){
	for i:=0; i<len(str);i++{
		fmt.Println(str[i], string(str[i]))
	}
}
func main(){
	str:="Golang"
	fmt.Printf("Type: %T, Value: %v\nBytes and Characters:", str, str)

	bytesandchar(str)
	str2:="Python"
	str3:="Golang"
	fmt.Println("str2:",str2,"str3:",str3,"\nstr2+str3: ",str2+str3)
	fmt.Println("str==str2: ", str==str2, "\nstr==str3: ", str==str3)


}