package sysinfo

import (
	. "base"
	client "client"
	tool "tool"
)

const (
	KindCliList string = "/sysinfo/clientlist"
)

func init() {
	Factory := GetSingleFactory()
	Factory.Add(KindCliList, func() EntryIntf { return &ClientListStu{EntryStu: NewEmptyEntry(KindCliList)} })

	ProRso := GetSingleProblemResolveMap()
	ProRso.Add(KindCliList, KindCliList)
}

type ClientListStu struct {
	*EntryStu
	EntryRst *EntryStu
	*tool.TwLogStu
}

func (r *ClientListStu) InitPrn() (ret int) {
	r.TwLogStu = tool.NewTwLogStu("%v\t%v\t\n")
	return
}

func (r *ClientListStu) ScanClients() (ret int) {
	Clients := client.GetClientsStu()
	for k, v := range Clients.Clients {
		r.Print(k, v.GetId())
	}
	return
}
func (r *ClientListStu) CreateRst() (ret int) {
	r.EntryRst = NewEmptyEntry("")
	r.EntryRst.SetBody(r.Flush())
	return
}
func (r *ClientListStu) Execute() (EntryOut *EntryStu, ret int) {
	defer func() {
		EntryOut = r.EntryRst
	}()

	ret = r.InitPrn()
	r.Print("clilist", "--list--")
	if ret != 0 {
		r.Print("error", "r.InitPrn")
		return
	}
	ret = r.ScanClients()
	if ret != 0 {
		r.Print("error", "r.ScanClients")
		return
	}
	ret = r.CreateRst()
	if ret != 0 {
		r.Print("error", "r.CreateRst")
		return
	}
	return
}
