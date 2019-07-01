package lavadb_protocol

type OBJECTSTOREPkt struct {
	Version		int64
	Echodata	[]byte
	Body	interface{} `asn1:"choice:body"`
}
type ReqLavaDBSetRecord struct {
	Dbid		int64
	Tableid		int64
	Key_hash	[]byte
	Key_range	[]byte
	Value		[]byte
	Timetolive	int64
}
type RspLavaDBSetRecord struct {
	Retcode		int64
	Retmsg		[]byte
	Ver		[]byte
	Key_hash	[]byte
	Key_range	[]byte
}
type ReqLavaDBSetVerRecord struct {
	Dbid		int64
	Tableid		int64
	Key_hash	[]byte
	Key_range	[]byte
	Value		[]byte
	Ver		[]byte
	Add_flag	int64
	Timetolive	int64
}
type RspLavaDBSetVerRecord struct {
	Retcode		int64
	Retmsg		[]byte
	Ver		[]byte
	Key_hash	[]byte
	Key_range	[]byte
}
type ReqLavaDBGetRecord struct {
	Dbid		int64
	Tableid		int64
	Key_hash	[]byte
	Key_range	[]byte
}
type RspLavaDBGetRecord struct {
	Retcode		int64
	Retmsg		[]byte
	Value		[]byte
	Ver		[]byte
	Key_hash	[]byte
	Key_range	[]byte
}
type ReqLavaDBListRecord struct {
	Dbid		int64
	Tableid		int64
	Key_hash	[]byte
	Prefix		[]byte
	Delimiter	[]byte
	Marker		[]byte
	Value_offset	int64
	Value_size	int64
	Maxnum		int64
	Type		int64
}
type ReqLavaDBListRecordWithEndMarker struct {
	Dbid		int64
	Tableid		int64
	Key_hash	[]byte
	Prefix		[]byte
	Delimiter	[]byte
	Marker		[]byte
	End_marker	[]byte
	Value_offset	int64
	Value_size	int64
	Maxnum		int64
	Type		int64
}
type Record struct {
	Key_hash	[]byte
	Partial_key	[]byte
	Value		[]byte
	Ver		[]byte
}
type RecordList []Record
type RspLavaDBListRecord struct {
	Retcode		int64
	Retmsg		[]byte
	Record_list	RecordList
	Prefix		[]byte
	Marker		[]byte
	Next_marker	[]byte
	Is_truncate	int64
}
type RspLavaDBListRecordWithEndMarker struct {
	Retcode		int64
	Retmsg		[]byte
	Record_list	RecordList
	Prefix		[]byte
	Marker		[]byte
	End_marker	[]byte
	Next_marker	[]byte
	Is_truncate	int64
}
type ReqLavaDBDelRecord struct {
	Dbid		int64
	Tableid		int64
	Key_hash	[]byte
	Key_range	[]byte
}
type RspLavaDBDelRecord struct {
	Retcode		int64
	Retmsg		[]byte
	Key_hash	[]byte
	Key_range	[]byte
}
type ReqLavaDBDelVerRecord struct {
	Dbid		int64
	Tableid		int64
	Key_hash	[]byte
	Key_range	[]byte
	Ver		[]byte
}
type RspLavaDBDelVerRecord struct {
	Retcode		int64
	Retmsg		[]byte
	Key_hash	[]byte
	Key_range	[]byte
}
type ReqLavaDBHeartBeat struct {
	Seq int64
}
type RspLavaDBHeartBeat struct {
	Retcode		int64
	Seq		int64
	Ver1		int64
	Ver2		int64
	Ver3		int64
	Starttime	int64
	Resv		[]byte
}
type ReqTFSHeartBeat struct {
	Seq int64
}
type RspTFSHeartBeat struct {
	Retcode		int64
	Seq		int64
	Ver1		int64
	Ver2		int64
	Ver3		int64
	Starttime	int64
}
type ReqTFSUploadData struct {
	Size		int64
	Data		[]byte
	Expdeltime	int64	`asn1:"optional"`
	Ieee_crc32	int64	`asn1:"optional,explicit,tag:1"`
	Is_return_crc32	int64	`asn1:"optional,explicit,tag:2"`
	File_key	[]byte	`asn1:"optional,explicit,tag:3"`
}
type RspTFSUploadData struct {
	Retcode		int64
	Size		int64
	Chid		int64
	Fid		int64
	Ieee_crc32	int64	`asn1:"optional"`
}
type ReqTFSDownloadData struct {
	Chid		int64
	Fid		int64
	Offset		int64
	Size		int64
	File_key	[]byte	`asn1:"optional"`
}
type RspTFSDownloadData struct {
	Retcode	int64
	Chid	int64
	Fid	int64
	Offset	int64
	Size	int64
	Data	[]byte
}
type ReqTFSDeleteData struct {
	Chid		int64
	Fid		int64
	File_key	[]byte	`asn1:"optional"`
}
type RspTFSDeleteData struct {
	Retcode	int64
	Chid	int64
	Fid	int64
}
type ReqTFSRefreshCfg struct {
	Cfgfile		[]byte
	Cfgcontent	[]byte
}
type RspTFSRefreshCfg struct {
	Retcode	int64
	Retmsg	[]byte
}
