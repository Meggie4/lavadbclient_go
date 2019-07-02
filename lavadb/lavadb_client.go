package lavadb_client

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
	"lavadbclient_go/protocols"
)

const BUFSIZE = 2<<10<<10
const RECONN_SECOND = 45

type LavadbConfig struct {
	Ip     	string
	Port   	int
	Tid    	int
	Cid    	int
}

type Lavadb struct {
	Ip_     	string
	Uip_    	int64
	Port_   	int64
	Conn_     	net.Conn
	Tid_    	int64
	Cid_        int64
	LastVisit_ time.Time
	CurrVisit_ time.Time
}

type LavadbListIter struct {
	Lava_		Lavadb
	Prefix_		string
	NeedValue_	string
	Count_		int
	Hash_		string
	IsListOver_	bool
	NextMarker	string
	Rsp_		lavadb_protocol.RspLavaDBListRecord
}

/*
TODO
 */
func (listIter LavadbListIter) SetMarker(marker string) {
	listIter.NextMarker = marker
}

/*
TODO
 */
func (listIter LavadbListIter) Next () int {
	return 0
}

func (ldb Lavadb) DoRequest(
	reqMsg		lavadb_protocol.StorageMessage,
	rspMsg		*lavadb_protocol.StorageMessage) int {
	ldb.CurrVisit_ = time.Now()
	subS := ldb.CurrVisit_.Sub(ldb.LastVisit_).Seconds()
	if subS > RECONN_SECOND && ldb.Conn_ != nil{
		ldb.Conn_.Close()
		ldb.Conn_ = nil
	}
	ldb.LastVisit_ = ldb.CurrVisit_
	if ldb.Conn_ == nil {
		var err error
		address := ldb.Ip_ + string(ldb.Port_)
		ldb.Conn_, err = net.Dial("tcp", address)
		if err != nil {
			log.Println("dial err:", err)
			return -1
		}
	}

	writebuf, err := lavadb_protocol.Encode(reqMsg)
	if err != nil {
		log.Println("encode failed,", err)
		return -1
	}
	writecount, writeerr := ldb.Conn_.Write(writebuf)
	if writeerr != nil {
		log.Println("conn write err,", writeerr)
		return -1
	}
	fmt.Printf("conn write %d bytes\n", writecount)
	readbuf := make([]byte, BUFSIZE)
	readcount, readerr := ldb.Conn_.Read(readbuf)
	if readerr != nil {
		log.Println("conn read failed,", err)
		return -1
	}
	log.Printf("conn read %d bytes\n", readcount)
	remainbuf, err := lavadb_protocol.Decode(readbuf, rspMsg)
	if err != nil {
		log.Printf("decode failed,", err)
		return -1
	} else {
		log.Printf("after decode, remaining buf len is %d\n",
						len(remainbuf))
		return 0
	}
}


func (ldb Lavadb) Set (
	keyRange		string,
	val				string,
	hash 			string) (rsp *lavadb_protocol.RspLavaDBSetRecord, err error){
	var reqMsg, rspMsg lavadb_protocol.StorageMessage
	rand.Seed(time.Now().Unix())
	seq := rand.Int63()
	lavadb_protocol.GetReqLavaDBSetRecord(&reqMsg, seq,0, ldb.Uip_,
						ldb.Port_, ldb.Tid_, ldb.Cid_, []byte(hash),
						[]byte(keyRange), []byte(val))

	ret := ldb.DoRequest(reqMsg, &rspMsg)

	if ret < 0 {
		*rsp = lavadb_protocol.RspLavaDBSetRecord{}
		err = errors.New("ret < 0")
		return
	}

	pkt, ok := rspMsg.Body.(lavadb_protocol.OBJECTSTOREPkt)
	if ok == false {
		*rsp = lavadb_protocol.RspLavaDBSetRecord{}
		err = errors.New("get pkt failed")
		return
	}

	setRecord, ok := pkt.Body.(lavadb_protocol.RspLavaDBSetRecord)
	if ok == false {
		*rsp = lavadb_protocol.RspLavaDBSetRecord{}
		err = errors.New("get set record failed")
		return
	} else {
		*rsp = setRecord
		err = nil
		return
	}
}
/*
TODO
 */
func (ldb Lavadb) SetVer (
	keyRange		string,
	val				string,
	ver 			string,
	hash 			string) lavadb_protocol.RspLavaDBSetVerRecord {
	return lavadb_protocol.RspLavaDBSetVerRecord{}
}

func (ldb Lavadb) Get (
	keyRange		string,
	hash 			string) (rsp *lavadb_protocol.RspLavaDBGetRecord, err error) {
	var reqMsg, rspMsg lavadb_protocol.StorageMessage
	seq := rand.Int63()
	lavadb_protocol.GetReqLavaDBGetRecord(&reqMsg, seq, 0,
									ldb.Uip_, ldb.Port_, ldb.Tid_,
									ldb.Cid_, []byte(hash), []byte(keyRange))
	ret := ldb.DoRequest(reqMsg, &rspMsg)
	if ret < 0 {
		*rsp = lavadb_protocol.RspLavaDBGetRecord{}
		err = errors.New("do request failed")
		return
	}

	pkt, ok := rspMsg.Body.(lavadb_protocol.OBJECTSTOREPkt)
	if ok == false {
		*rsp = lavadb_protocol.RspLavaDBGetRecord{}
		err = errors.New("get pkt failed")
		return
	}
	getRecord, ok := pkt.Body.(lavadb_protocol.RspLavaDBGetRecord)
	if ok == false {
		*rsp = lavadb_protocol.RspLavaDBGetRecord{}
		err = errors.New("get record failed")
		return
	} else {
		*rsp = getRecord
		err = nil
		return
	}

}

func (ldb Lavadb) Del (
	keyRange		string,
	hash 			string) int {
	var reqMsg, rspMsg lavadb_protocol.StorageMessage
	seq := rand.Int63()
	lavadb_protocol.GetReqLavaDBDelRecord(&reqMsg, seq, 0,
		ldb.Uip_, ldb.Port_, ldb.Tid_,
		ldb.Cid_, []byte(hash), []byte(keyRange))
	ret := ldb.DoRequest(reqMsg, &rspMsg)
	if ret < 0 {
		return -1
	}

	pkt, ok := rspMsg.Body.(lavadb_protocol.OBJECTSTOREPkt)
	if ok == false {
		return -1
	}
	delRecord, ok := pkt.Body.(lavadb_protocol.RspLavaDBDelRecord)
	if ok == false {
		return -1
	} else {
		retcode := delRecord.Retcode
		if retcode == E_CELL_NO_RECORD {
			log.Printf("no record: %s\n", keyRange)
			return -1
		} else if retcode != 0 {
			log.Printf("retcode:%d, retmsg:%s\n", retcode,
							string(delRecord.Retmsg))
			return - 1
		} else {
			return 0
		}
	}

}

/*
TODO
 */
func (ldb Lavadb) DelRecurse (
	prefix			string,
	delCount		*int,
	count 			int,
	hash 			string) int {
	return 0
}

func (ldb Lavadb) List (
	prefix		string,
	nextMarker string,
	needValue	int64,
	count 		int64,
	hash 		string) (rsp *lavadb_protocol.RspLavaDBListRecord, err error) {
	var reqMsg, rspMsg lavadb_protocol.StorageMessage
	seq := rand.Int63()
	lavadb_protocol.GetReqLavaDBListRecord(&reqMsg, seq, 0,
										ldb.Uip_, ldb.Port_, ldb.Tid_,
										ldb.Cid_, []byte(hash), []byte(prefix),
										[]byte{}, []byte(nextMarker),0,
										-1, count, needValue)
	ret := ldb.DoRequest(reqMsg, &rspMsg)
	if ret < 0 {
		*rsp = lavadb_protocol.RspLavaDBListRecord{}
		err = errors.New("ret < 0")
		return
	}

	pkt, ok := rspMsg.Body.(lavadb_protocol.OBJECTSTOREPkt)
	if ok == false {
		*rsp = lavadb_protocol.RspLavaDBListRecord{}
		err = errors.New("get pkt failed")
		return
	}
	listRecord, ok := pkt.Body.(lavadb_protocol.RspLavaDBListRecord)
	if ok == false {
		*rsp = lavadb_protocol.RspLavaDBListRecord{}
		err = errors.New("get record failed")
		return
	} else {
		*rsp = listRecord
		err = nil
		return
	}
}

/*
TODO
 */
func (ldb Lavadb) Count(
	prefix		string,
	count 		int,
	limit 		int,
	hash 		string) int {
	return 0
}

