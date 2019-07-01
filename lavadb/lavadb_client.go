package lavadb_client

import (
	"../protocols"
	"crypto"
	"crypto/x509"
	"errors"
	"expvar"
	"fmt"
	"github.com/golang/tools/go/analysis/passes/atomic/testdata/src/a"
	"github.com/golang/tools/go/ssa"
	"log"
	"math/rand"
	"net"
	"time"
)

const BUFSIZE = 2<<10<<10
const RECONN_SECOND = 45

type LavadbConfig struct {
	ip     	string
	port   	int
	tid    	int
	cid    	int
}

type Lavadb struct {
	ip_     	string
	uip_    	int64
	port_   	int64
	conn_     	net.Conn
	tid_    	int64
	cid_		int64
	lastVisit_ time.Time
	currVisit_ time.Time
}

type LavadbListIter struct {
	lava_		Lavadb
	prefix_		string
	needValue_	string
	count_		int
	hash_		string
	isListOver_	bool
	nextMarker	string
	rsp_		lavadb_protocol.RspLavaDBListRecord
}

func (listIter LavadbListIter) setMarker(marker string) {
	listIter.nextMarker = marker
}

func (listIter LavadbListIter) next () int {

}

func (ldb Lavadb) doRequest(
	reqMsg		lavadb_protocol.StorageMessage,
	rspMsg		*lavadb_protocol.StorageMessage) int {
	ldb.currVisit_ = time.Now()
	subS := ldb.currVisit_.Sub(ldb.lastVisit_).Seconds()
	if subS > RECONN_SECOND && ldb.conn_ != nil{
		ldb.conn_.Close()
		ldb.conn_ = nil
	}
	ldb.lastVisit_ = ldb.currVisit_
	if ldb.conn_ == nil {
		var err error
		address := ldb.ip_ + string(ldb.port_)
		ldb.conn_, err = net.Dial("tcp", address)
		if err != nil {
			log.Println("dial err:", err)
			return -1
		}
	}
	if len(ldb.buf_) <= 0 {
		log.Println("buf_ size is zero")
		return -1
	} else {
		writebuf, err := lavadb_protocol.Encode(reqMsg)
		if err != nil {
			log.Println("encode failed,", err)
			return -1
		}
		writecount, writeerr := ldb.conn_.Write(writebuf)
		if writeerr != nil {
			log.Println("conn write err,", writeerr)
			return -1
		} else {
			fmt.Printf("conn write %d bytes\n", writecount)
			readbuf := make([]byte, BUFSIZE)
			readcount, readerr := ldb.conn_.Read(readbuf)
			if readerr != nil {
				log.Println("conn read failed,", err)
				return -1
			} else {
				log.Printf("conn read %d bytes\n", readcount)
				remainbuf, err := lavadb_protocol.Decode(readbuf, rspMsg)
				if err != nil {
					log.Printf("decode failed,", err)
					return -1
				} else {
					log.Printf("after decode, remaining buf len is %d\n", len(remainbuf))
					return 0
				}
			}
		}
	}
}


func (ldb Lavadb) set (
	keyRange		string,
	val				string,
	hash 			string) (lavadb_protocol.RspLavaDBSetRecord, error){
	var reqMsg, rspMsg lavadb_protocol.StorageMessage
	rand.Seed(time.Now().Unix())
	seq := rand.Int63()
	lavadb_protocol.GetReqLavaDBSetRecord(&reqMsg, seq,0, ldb.uip_,
						ldb.port_, ldb.tid_, ldb.cid_, []byte(hash),
						[]byte(keyRange), []byte(val))

	ret := ldb.doRequest(reqMsg, &rspMsg)

	if ret < 0 {
		return lavadb_protocol.RspLavaDBSetRecord{}, errors.New("ret < 0")
	} else {
		pkt, ok := rspMsg.Body.(lavadb_protocol.OBJECTSTOREPkt)
		if ok == false {
			return lavadb_protocol.RspLavaDBSetRecord{}, errors.New("get pkt failed")
		} else {
			setRecord, ok := pkt.Body.(lavadb_protocol.RspLavaDBSetRecord)
			if ok == false {
				return lavadb_protocol.RspLavaDBSetRecord{}, errors.New("get set record failed")
			} else {
				return setRecord, nil
			}
		}
	}
}

func (ldb Lavadb) setVer (
	keyRange		string,
	val				string,
	ver 			string,
	hash 			string) lavadb_protocol.RspLavaDBSetVerRecord {
}

func (ldb Lavadb) get (
	keyRange		string,
	hash 			string) (lavadb_protocol.RspLavaDBGetRecord, error) {
	var reqMsg, rspMsg lavadb_protocol.StorageMessage
	seq := rand.Int63()
	lavadb_protocol.GetReqLavaDBGetRecord(&reqMsg, seq, 0,
									ldb.uip_, ldb.port_, ldb.tid_,
									ldb.cid_, []byte(hash), []byte(keyRange))
	ret := ldb.doRequest(reqMsg, &rspMsg)
	if ret < 0 {
		return lavadb_protocol.RspLavaDBGetRecord{}, errors.New("do request failed")
	} else {
		pkt, ok := rspMsg.Body.(lavadb_protocol.OBJECTSTOREPkt)
		if ok == false {
			return lavadb_protocol.RspLavaDBGetRecord{}, errors.New("get pkt failed")
		} else {
			getRecord, ok := pkt.Body.(lavadb_protocol.RspLavaDBGetRecord)
			if ok == false {
				return lavadb_protocol.RspLavaDBGetRecord{}, errors.New("get record failed")
			} else {
				return getRecord, nil
			}
		}
	}
}

func (ldb Lavadb) del (
	keyRange		string,
	hash 			string) int {
	var reqMsg, rspMsg lavadb_protocol.StorageMessage
	seq := rand.Int63()
	lavadb_protocol.GetReqLavaDBDelRecord(&reqMsg, seq, 0,
		ldb.uip_, ldb.port_, ldb.tid_,
		ldb.cid_, []byte(hash), []byte(keyRange))
	ret := ldb.doRequest(reqMsg, &rspMsg)
	if ret < 0 {
		return -1
	} else {
		pkt, ok := rspMsg.Body.(lavadb_protocol.OBJECTSTOREPkt)
		if ok == false {
			return -1
		} else {
			delRecord, ok := pkt.Body.(lavadb_protocol.RspLavaDBDelRecord)
			if ok == false {
				return -1
			} else {
				retcode := delRecord.Retcode
				if retcode == E_CELL_NO_RECORD {
					log.Printf("no record: %s\n", keyRange)
					return -1
				} else if retcode != 0 {
					log.Printf("retcode:%d, retmsg:%s\n", retcode, string(delRecord.Retmsg))
					return - 1
				} else {
					return 0
				}
			}
		}
	}
}

func (ldb Lavadb) delRecurse (
	prefix			string,
	delCount		int,
	count 			int,
	hash 			string) int {

}

func (ldb Lavadb) list (
	prefix		string,
	nextMarker string,
	needValue	int64,
	count 		int64,
	hash 		string) (lavadb_protocol.RspLavaDBListRecord, error) {
	var reqMsg, rspMsg lavadb_protocol.StorageMessage
	seq := rand.Int63()
	lavadb_protocol.GetReqLavaDBListRecord(&reqMsg, seq, 0,
										ldb.uip_, ldb.port_, ldb.tid_,
										ldb.cid_, []byte(hash), []byte(prefix),
										[]byte{}, []byte(nextMarker),0,
										-1, count, needValue)
	ret := ldb.doRequest(reqMsg, &rspMsg)
	if ret < 0 {
		return -1
	} else {

	}
}

func (ldb Lavadb) count(
	prefix		string,
	count 		int,
	limit 		int,
	hash 		string) int {

}

