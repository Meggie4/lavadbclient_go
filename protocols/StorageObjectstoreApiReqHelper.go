package lavadb_protocol

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"log"
)

const VERSION = 1 
const ROUTE_DEST_ID = 0x0000
const ROUTE_DEST_KEY = 0x0001
const ROUTE_REMOTE_ID = 0x0000
const ROUTE_REMOTE_IPPORT = 0x0001

func CheckAsnComplete(data []byte) int64 {
	len := (uint64)(len(data))
	MAX_LEN := (uint64)(32<<10<<10)
	if len < 4 {
		return 0
	}
	if data[0] != '0' {
		return -1
	}

	var head_len, cont_len uint64
	
	if (uint64)(data[1]) > 128 {
		len_len := (uint64)(data[1] & 127)
		head_len = 2 + len_len
		cont_len = binary.LittleEndian.Uint64(data[2: len_len + 2])
	} else {
		head_len = 2
		cont_len = (uint64) (data[1])
	}

	pkt_len := head_len + cont_len
	if MAX_LEN < pkt_len {
		log.Println("pkt_len > MAX_LEN")
		return -1
	}

	if len > pkt_len {
		if data[pkt_len] != '0' {
			return -1
		}
		return int64(pkt_len)
	} else if len == pkt_len {
		return int64(pkt_len)
	} else {
		return 0
	}
}

//get dest_key
func GetSerial(key_hash []byte, key_range []byte) int32 {
	var buffer bytes.Buffer
	buffer.Write(key_hash)
	buffer.Write(key_range)
	key := buffer.Bytes()
	md5Ctx := md5.New()
	md5Ctx.Write(key)
	cipherStr := md5Ctx.Sum(nil)
	
	id := int32(binary.BigEndian.Uint32(cipherStr[4:13]))
	return id
}

func GetRouteinfo(
	srcid      *int64,
	dest_id    *int64,
	dest_key   *int64,
	remoteip   *int64,
	remoteport *int64,
	remoteid   *int64) ROUTEINFO {
	var route ROUTEINFO
	route.Srcid = *srcid
	route.Version = VERSION
	if dest_id != nil {
		route.Destid = *dest_id
		route.Desttype = ROUTE_DEST_ID
	} else {
		route.Destkey = *dest_key
		route.Desttype = ROUTE_DEST_KEY
	}

	if remoteid != nil {
		route.Remoteid = *remoteid
		route.Remotetype = ROUTE_REMOTE_ID
	} else {
		route.Remoteip = *remoteip
		route.Remoteport = *remoteport
		route.Remotetype = ROUTE_REMOTE_IPPORT
	}
	
	return route
}

//set
func GetReqLavaDBSetRecord(
	out_msg *StorageMessage,
	msg_seq int64,
	smcd_id int64,
	dest_ip int64,
	dest_port int64,
	dbid int64,
	tableid int64,
	key_hash []byte,
	key_range []byte,
	value []byte) {
	dest_key := (int64)(GetSerial(key_hash, key_range))
	route_info := GetRouteinfo(&smcd_id, nil, &dest_key,
					&dest_ip, &dest_port, nil)
	*out_msg = StorageMessage{
		Version:   VERSION,
		Seq:       msg_seq,
		Routeinfo: route_info,
		Msgtype:   0,
		Body: OBJECTSTOREPkt{
			Version:  VERSION,
			Echodata: []byte(""),
			Body: ReqLavaDBSetRecord{
				Dbid:      dbid,
				Tableid:   tableid,
				Key_hash:  key_hash,
				Key_range: key_range,
				Value: 	   value,
				Timetolive: 0,
			},
		},
	}
}

//set version
func GetReqLavaDBSetVerRecord(
	out_msg *StorageMessage,
	msg_seq int64,
	smcd_id int64,
	dest_ip int64,
	dest_port int64,
	dbid int64,
	tableid int64,
	key_hash []byte,
	key_range []byte,
	value []byte,
	ver   []byte,
	add_flag  int64) {
	dest_key := (int64)(GetSerial(key_hash, key_range))
	route_info := GetRouteinfo(&smcd_id, nil, &dest_key,
					&dest_ip, &dest_port, nil)
	*out_msg = StorageMessage{
		Version:   VERSION,
		Seq:       msg_seq,
		Routeinfo: route_info,
		Msgtype:   0,
		Body: OBJECTSTOREPkt{
			Version:  VERSION,
			Echodata: []byte(""),
			Body: ReqLavaDBSetVerRecord{
				Dbid:      dbid,
				Tableid:   tableid,
				Key_hash:  key_hash,
				Key_range: key_range,
				Value: 	   value,
				Ver:	   ver,
				Add_flag:  add_flag,
				Timetolive: 0,
			},
		},
	}
}

//get
func GetReqLavaDBGetRecord(
	out_msg *StorageMessage, 
	msg_seq int64,
	smcd_id int64,
	dest_ip int64,
	dest_port int64,
	dbid int64,
	tableid int64,
	key_hash []byte,
	key_range []byte) {
	dest_key := (int64)(GetSerial(key_hash, key_range))
	route_info := GetRouteinfo(&smcd_id, nil, &dest_key,
						&dest_ip, &dest_port, nil)
	*out_msg = StorageMessage{
		Version:   VERSION,
		Seq:       msg_seq,
		Routeinfo: route_info,
		Msgtype:   0,
		Body: OBJECTSTOREPkt{
			Version:  VERSION,
			Echodata: []byte(""),
			Body: ReqLavaDBGetRecord{
				Dbid:      dbid,
				Tableid:   tableid,
				Key_hash:  key_hash,
				Key_range: key_range,
			},
		},
	}
}

//delete
func GetReqLavaDBDelRecord(
	out_msg *StorageMessage,
	msg_seq int64,
	smcd_id int64,
	dest_ip int64,
	dest_port int64,
	dbid int64,
	tableid int64,
	key_hash []byte,
	key_range []byte) {
	dest_key := (int64)(GetSerial(key_hash, key_range))
	route_info := GetRouteinfo(&smcd_id, nil, &dest_key,
					&dest_ip, &dest_port, nil)
	*out_msg = StorageMessage{
		Version:   VERSION,
		Seq:       msg_seq,
		Routeinfo: route_info,
		Msgtype:   0,
		Body: OBJECTSTOREPkt{
			Version:  VERSION,
			Echodata: []byte(""),
			Body: ReqLavaDBDelRecord{
				Dbid:      dbid,
				Tableid:   tableid,
				Key_hash:  key_hash,
				Key_range: key_range,
			},
		},
	}
}

//delete version
func GetReqLavaDBDelVerRecord(
	out_msg *StorageMessage,
	msg_seq int64,
	smcd_id int64,
	dest_ip int64,
	dest_port int64,
	dbid int64,
	tableid int64,
	key_hash []byte,
	key_range []byte,
	ver 	 []byte) {
	dest_key := (int64)(GetSerial(key_hash, key_range))
	route_info := GetRouteinfo(&smcd_id, nil, &dest_key,
				&dest_ip, &dest_port, nil)
	*out_msg = StorageMessage{
		Version:   VERSION,
		Seq:       msg_seq,
		Routeinfo: route_info,
		Msgtype:   0,
		Body: OBJECTSTOREPkt{
			Version:  VERSION,
			Echodata: []byte(""),
			Body: ReqLavaDBDelVerRecord{
				Dbid:      dbid,
				Tableid:   tableid,
				Key_hash:  key_hash,
				Key_range: key_range,
				Ver:	   ver,
			},
		},
	}
}

//list
func GetReqLavaDBListRecord(
	out_msg 	*StorageMessage,
	msg_seq		int64,
	smcd_id 	int64,
	dest_ip 	int64,
	dest_port 	int64,
	dbid 		int64,
	tableid 	int64,
	key_hash 	[]byte,
	prefix		[]byte,
	delimiter	[]byte,
	marker 		[]byte,
	value_offset int64,
	value_size   int64,
	maxnum       int64,
	thetype      int64) {
	dest_key := (int64)(GetSerial(key_hash, prefix))
	route_info := GetRouteinfo(&smcd_id, nil, &dest_key,
					&dest_ip, &dest_port, nil)
	*out_msg = StorageMessage{
		Version:   VERSION,
		Seq:       msg_seq,
		Routeinfo: route_info,
		Msgtype:   0,
		Body: OBJECTSTOREPkt{
			Version:  VERSION,
			Echodata: []byte(""),
			Body: ReqLavaDBListRecord{
				Dbid:      dbid,
				Tableid:   tableid,
				Key_hash:  key_hash,
				Prefix:    prefix,
				Delimiter: delimiter,
				Marker:    marker,
				Value_offset: value_offset,
				Value_size:   value_size,
				Maxnum:       maxnum,
				Type:         thetype,
			},
		},
	}
}
