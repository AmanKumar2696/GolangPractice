package main
import(
	"fmt"
	"sync"
)

var wg=sync.WaitGroup{}

func disp(i int, str string){
	for j:=0;j<i;j++{
		fmt.Println(str, i)
	}
	wg.Done()
}

func main(){
	wg.Add(2)
	go disp(5, "Fucntion 1 ")
	disp(5, "Function 2 ")
	wg.Wait()
}