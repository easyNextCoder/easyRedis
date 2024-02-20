package server

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

//并发过高会导致redis产生read tcp 127.0.0.1:xxx->127.0.0.1.6379: i/o timeout

type concurrentController struct {
	ch chan int
}

func (c *concurrentController) get() {
	ccr.ch <- 0
}

var ccr concurrentController = concurrentController{
	ch: make(chan int, 200),
}

func init() {
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for {
			select {
			case <-ticker.C:
				for i := 0; i < len(ccr.ch); i++ {
					select {
					case <-ccr.ch:
					default:

					}
				}
			}
		}
	}()
}

func grabRedBag() {
	wg := sync.WaitGroup{}

	n := 1000

	gs := make([]*grabber, n)

	//模拟1000个用户取抢红包,每秒200的并发量
	for i := range gs {
		gs[i] = &grabber{
			idx:      i,
			gotMoney: 0,
		}
		wg.Add(1)
		ccr.get()
		go gs[i].grab(&wg)
	}

	wg.Wait()

	totalGrab := 0

	for _, v := range gs {
		totalGrab += v.gotMoney
	}

	log.Println("grabRedBag all done totalGrab ", totalGrab)
}

type grabber struct {
	idx      int
	gotMoney int
}

func (g *grabber) grab(pwg *sync.WaitGroup) {
	defer pwg.Done()

	var body io.Reader = strings.NewReader("get")

	sleepTime := rand.Intn(500)
	maxUseTime := 2000 - sleepTime

	time.Sleep(time.Millisecond * time.Duration(sleepTime))

	req, err := http.NewRequest("POST", "http://127.0.0.1:6378/grab", body)

	if err != nil {
		log.Printf("newRequest err %s\n", err)
		return
	}

	client := http.Client{Timeout: time.Millisecond * time.Duration(maxUseTime)}
	rsp, err := client.Do(req)
	if err != nil {
		log.Printf("client do err %s\n", err)
		return
	}

	all, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Printf("io.ReadAll err %s\n", err)
		return
	}

	s := string(all)
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("atoi err %s\n", err)
		return
	}

	log.Printf("第%d个用户获得%d元红包\n", g.idx, n)

	g.gotMoney += n

}
