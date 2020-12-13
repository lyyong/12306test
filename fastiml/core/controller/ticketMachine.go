package controller

// 购票的接口

type TicketMachine interface {
	BuyTicket(infos []*BuyTicketInfo) []*Ticket
	SearchTicketNum(info *BuyTicketInfo) int // 得到某个区间的座位数,与座位类型无关
}

type BuyTicketInfo struct {
	startStation string
	endStation   string
	seatType     string // 座位类型A,B或者C
}

func (b *BuyTicketInfo) SeatType() string {
	return b.seatType
}

func (b *BuyTicketInfo) SetSeatType(seatType string) {
	b.seatType = seatType
}

func (b *BuyTicketInfo) EndStation() string {
	return b.endStation
}

func (b *BuyTicketInfo) SetEndStation(endStation string) {
	b.endStation = endStation
}

func (b *BuyTicketInfo) StartStation() string {
	return b.startStation
}

func (b *BuyTicketInfo) SetStartStation(startStation string) {
	b.startStation = startStation
}

func NewBuyTicketInfo(startStation string, endStation string, seatType string) *BuyTicketInfo {
	return &BuyTicketInfo{startStation: startStation, endStation: endStation, seatType: seatType}
}

type Ticket struct {
	startStation string
	endStation   string
	// 买不到票也要返回一个seatInfo为"",但是有站内容的结构
	seatInfo string
}

func NewTicket(startStation string, endStation string, seatInfo string) *Ticket {
	return &Ticket{startStation: startStation, endStation: endStation, seatInfo: seatInfo}
}

func (t *Ticket) SeatInfo() string {
	return t.seatInfo
}

func (t *Ticket) SetSeatInfo(seatInfo string) {
	t.seatInfo = seatInfo
}

func (t *Ticket) EndStation() string {
	return t.endStation
}

func (t *Ticket) SetEndStation(endStation string) {
	t.endStation = endStation
}

func (t *Ticket) StartStation() string {
	return t.startStation
}

func (t *Ticket) SetStartStation(startStation string) {
	t.startStation = startStation
}

func (t *Ticket) Equals(ta *Ticket) bool {
	if ta == nil {
		return false
	}
	if ta.seatInfo == t.seatInfo && ta.startStation == t.startStation && ta.endStation == t.endStation {
		return true
	}
	return false
}
