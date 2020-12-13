package test

import (
	"12306test/fastiml/core/controller"
	"12306test/fastiml/model"
	"math/rand"
	"sync"
	"testing"
)

// 限制goroutine
type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

func NewPool(size int) *Pool {
	if size <= 0 {
		size = 1
	}
	return &Pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

// 添加线程,也可以是减少
func (p *Pool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	for i := 0; i > delta; i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

// 线程完成
func (p *Pool) Done() {
	<-p.queue
	p.wg.Done()
}

// 等待完成
func (p *Pool) Wait() {
	p.wg.Wait()
}

type testMachine struct {
}

func (m testMachine) SearchTicketNum(info *controller.BuyTicketInfo) int {
	return 0
}

func (m testMachine) BuyTicket(infos []*controller.BuyTicketInfo) []*controller.Ticket {
	res := make([]*controller.Ticket, 0, len(infos))
	for _, v := range infos {
		res = append(res, controller.NewTicket(
			v.StartStation(),
			v.EndStation(),
			"00A",
		))
	}
	return res
}

// 测试machine的正确性
func BenchmarkMachine(b *testing.B) {
	//b.Log(b.N)
	var machine controller.TicketMachine
	// TODO 创建machine
	machine = testMachine{}

	btInfos := genTestData(b)
	tickets := make([]*controller.Ticket, 0)
	tt := make([][]*controller.Ticket, 0) // 零时存票
	// 100个goroutine
	pool := NewPool(50)
	b.StartTimer()
	for i := 0; i < 2000; i++ {
		pool.Add(1)
		tt = append(tt, make([]*controller.Ticket, 0))
		i := i
		go func() {
			tt[i] = append(tt[i], buy(btInfos, machine)...)
			pool.Done()
		}()
	}
	pool.Wait()
	b.StopTimer()
	for _, v := range tt {
		tickets = append(tickets, v...)
	}
	b.Log(len(tickets))
	checkTickets(tickets, b, machine)
}

func buy(btInfos [][]*controller.BuyTicketInfo, machine controller.TicketMachine) []*controller.Ticket {
	res := make([]*controller.Ticket, 0)
	for _, v := range btInfos {
		res = append(res, machine.BuyTicket(v)...)
	}
	return res
}

func genTestData(b *testing.B) [][]*controller.BuyTicketInfo {
	model.InitDB()
	// 每次生成两万个数据
	tsInfos := make([][]*controller.BuyTicketInfo, 20000)
	train := model.GetAllTrains()[0]
	b.Log(train)
	ts := model.GetStopInfos(map[string]interface{}{"train_id": train.ID})
	stations := make([]*model.StopInfo, len(ts))
	// 通过seq进行排序
	for _, v := range ts {
		stations[v.Seq] = v
	}

	// 本次购票数,起始站,下车站
	var length, s, e int
	for i := 0; i < len(tsInfos); i++ {
		length = rand.Intn(5) // 最高一下购买5张票
		tsInfos[i] = make([]*controller.BuyTicketInfo, length)
		for j := 0; j < length; j++ {
			e = rand.Intn(len(stations))
			if e <= 0 {
				e = 1
			}
			s = rand.Intn(e)
			tsInfos[i][j] = controller.NewBuyTicketInfo(
				stations[s].Station.Name,
				stations[e].Station.Name,
				getRandSeatType(),
			)
		}
	}

	defer model.Close()
	return tsInfos
}

func getRandSeatType() string {
	i := rand.Intn(3)
	switch i {
	case 0:
		return "A"
	case 1:
		return "B"
	case 2:
		return "C"
	}
	return ""
}

func checkTickets(tickets []*controller.Ticket, b *testing.B, machine controller.TicketMachine) {
	for _, v := range tickets {
		if v.SeatInfo() == "" &&
			machine.SearchTicketNum(
				controller.NewBuyTicketInfo(v.StartStation(), v.EndStation(), "A")) > 0 &&
			machine.SearchTicketNum(
				controller.NewBuyTicketInfo(v.StartStation(), v.EndStation(), "B")) > 0 &&
			machine.SearchTicketNum(
				controller.NewBuyTicketInfo(v.StartStation(), v.EndStation(), "C")) > 0 {
			b.Fatal("有票但是没卖")
		}
		for _, vj := range tickets {
			if v.Equals(vj) {
				b.Fatal("出现相同票")
			}
		}
	}
}
