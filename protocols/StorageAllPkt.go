package lavadb_protocol

type StorageMessage struct {
	Version		int64
	Seq		int64
	Routeinfo	ROUTEINFO
	Msgtype		int64
	Body interface{} `asn1:"choice:smbody"`
}
type ROUTEINFO struct {
	Version		int64
	Srcid		int64
	Destkey		int64
	Destid		int64
	Desttype	int64
	Remoteip	int64
	Remoteport	int64
	Remoteid	int64
	Remotetype	int64
}