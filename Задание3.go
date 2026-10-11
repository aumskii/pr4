package main
import(
	"fmt"
	"time"
)
func main(){
	requests:=15
	tick :=time.Tick(200*time.Millisecond)
	for i:=1; i<=requests;i++ {
		<-tick
		fmt.Printf("Обработан запрос %d\n",i)
	}
}