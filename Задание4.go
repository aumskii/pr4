package main
import (
	"fmt"
	"sync"
	"net/http"
)
func check(id int, jobs<-chan string, wg *sync.WaitGroup){
	defer wg.Done()

	for url:= range jobs {
		resp, err:= http.Get(url)
		if err != nil {
			fmt.Println("Ошибка")
			continue
		}
		fmt.Printf("Статус %s: %s\n", url, resp.Status)
		resp.Body.Close()
	}
}
func main() {
	urls:=[]string{
		"https://github.com",
		"https://vk.ru",
		"https://ggsel.net",
		"https://www.dns-shop.ru",
		"https://yandex.ru",
		"https://www.twitch.tv",
	}
	jobs := make(chan string, len(urls))
	var wg sync.WaitGroup
	
	for w:=1; w<=3; w++ {
		wg.Add(1)
		go check(w,jobs, &wg)
	}
	for _, u := range urls {
		jobs <-u
	}
	close(jobs)
	wg.Wait()
}