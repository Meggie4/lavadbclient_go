package main

import (
	"fmt"
	"./protocols"
	"math/big"
	"net"
	"os"
	"strconv"
	"strings"
	"./lavadb"
	"time"
)

func PrintHelpmsg(args []string) {
	fmt.Printf("%v get keyrange\n", args[0])
	fmt.Printf("%v set keyrange value\n", args[0])
	fmt.Printf("%v del keyrange\n", args[0])
	fmt.Printf("%v delr prefix\n", args[0])
	fmt.Printf("%v list prefix [marker]\n", args[0])
	fmt.Println("-----------------------------------")
	fmt.Println("You can set the shell env, such as 'export LAVADB_DEBUG=1', for following variable:")
	fmt.Printf("\tLAVADB_DEBUG\n")
	fmt.Printf("\tLAVADB_KEYHASH (default: coslist)\n")
	fmt.Printf("\tLAVADB_TID\n")
	fmt.Printf("\tLAVADB_CID (default: 1)\n")
	fmt.Printf("\tLAVADB_IP\n")
	fmt.Printf("\tLAVADB_PORT\n")
	fmt.Printf("\tLAVADB_KEYHASH\n")
	fmt.Printf("\tLAVADB_KEYHASH_DELIMITER (default: no delimiter)\n")
}

func main() {
	args := os.Args 
	if len(args) < 2 {
		PrintHelpmsg(args)
		return 
	}
	var debug_val, tid_val, cid_val, port_val int64
	var ip_val, keyhash_val, keyhash_delimiter_val, keyrange string
	if debug := os.Getenv("LAVADB_DEBUG"); debug != "" {
		if mydebug, err := strconv.Atoi(debug); err == nil {
			fmt.Println("debug error", err)
			return
		} else {
			debug_val = int64(mydebug)
		}
	} else {
		debug_val = 0
	}

	if tid := os.Getenv("LAVADB_TID"); tid != "" {
		if mytid, err := strconv.Atoi(tid); err == nil {
			fmt.Println("tid error", err)
			return
		} else {
			tid_val = int64(mytid)
		}
	} else {
		tid_val = 900001
	}

	if cid := os.Getenv("LAVADB_CID"); cid != "" {
		if mycid, err := strconv.Atoi(cid); err == nil {
			fmt.Println("cid error", err)
			return
		} else {
			cid_val = int64(mycid)
		}
	} else {
		cid_val = 15
	}

	if port := os.Getenv("LAVADB_PORT"); port != "" {
		if myport, err := strconv.Atoi(port); err == nil {
			fmt.Println("port error", err)
			return
		} else {
			port_val = int64(myport)
		}
	} else {
		port_val = 9090
	}

	if ip_val := os.Getenv("LAVADB_IP"); ip_val == "" {
		ip_val = "10.58.90.158"
	} 

	if keyhash_val := os.Getenv("LAVADB_KEYHASH"); keyhash_val == "" {
		keyhash_val = "test_hash"
	} 

	if keyhash_delimiter_val := os.Getenv("LAVADB_KEYHASH_DELIMITER");
			keyhash_delimiter_val == "" {
		keyhash_delimiter_val = "."
	} 

	if len(args) >= 3 {
		if len(keyhash_delimiter_val) != 0 {
			key := args[2]
			idx := strings.Index(keyhash_delimiter_val, key)
			if idx >= 0 {
				keyhash_val = key[0: idx]
				delimiter_len := len(keyhash_delimiter_val)
				keyrange = key[idx + delimiter_len:]
			} else {
				keyrange = args[2]
			}
		} else {
			keyrange = args[2]
		}
	}

	if debug_val != 0 {
		fmt.Printf("LAVADB_TID=%d, LAVADB_CID=%d\n", tid_val, cid_val)
		fmt.Printf("LAVADB_IP=%s, LAVADB_PORT=%d\n", ip_val, port_val)
		fmt.Printf("keyhash=%s, keyrange=%s\n", keyhash_val, keyrange)
		fmt.Println("----------------------------------------")
	}

	uip := big.NewInt(0)
	uip.SetBytes(net.ParseIP(ip_val).To4())
	lava := lavadb_client.Lavadb {
		Ip_:		ip_val,
		Uip_: 		uip.Int64(),
		Port_:  	port_val,
		Conn_:     	nil,
		Tid_ :   	tid_val,
		Cid_ :      cid_val,
		LastVisit_: time.Now(),
		CurrVisit_: nil,
	}

	var ret int64
	if strings.Compare(args[1], "get") == 0 {
		if len(args) < 3 {
			PrintHelpmsg(args)
			return
		}
		var val string
		var rsp = new(lavadb_protocol.RspLavaDBGetRecord)
		var err error
		rsp, err = lava.Get(keyrange, keyhash_val)
		if err != nil {
			fmt.Println("get failed,", err)
			return
		}
		ret = rsp.Retcode
		if ret != 0 && ret != lavadb_client.E_CELL_NO_RECORD {
			fmt.Printf("get %s failed, ret = %d, retmsg = %s\n",
					keyrange, ret, string(rsp.Retmsg))
			if ret > 0 {
				ret = -ret
			}
			return
		}
		val = string(rsp.Value)
		
		if ret == lavadb_client.E_CELL_NO_RECORD {
			fmt.Println("no record")
			return 
		}
		fmt.Println("val:", val)
		return
	}
	if strings.Compare(args[1], "list") == 0 {
		if len(args) < 3 {
			PrintHelpmsg(args)
			return
		}

		iter := lavadb_client.LavadbListIter{

		}
		
		if len(args) > 3 {
			iter.SetMarker(args[3])
		}

		/*
		TODO
		 */
		return
	}
	if strings.Compare(args[1], "set") == 0 {
		if debug_val == 0 {
			fmt.Println("LAVADB_DEBUG not set or false")
			return
		}
		if len(args) < 4 {
			PrintHelpmsg(args)
			return
		}
		var val string
		var rsp = new(lavadb_protocol.RspLavaDBSetRecord)
		val = args[3]
		var err error
		rsp, err = lava.Set(keyrange, val, keyhash_val)
		if err != nil {
			fmt.Println("set failed,", err)
			return
		}
		ret = rsp.Retcode
		if ret != 0 {
			retmsg := string(rsp.Retmsg)
			fmt.Printf("set [%s]: data size[%d] failed, ret = %d, retmsg = %s\n",
						keyrange, len(val), ret, retmsg)
		}
		
		if ret == 0 {
			fmt.Println("", args[3])
		}
		return
	}
	if strings.Compare(args[1], "del") == 0 {
		if debug_val == 0 {
			fmt.Println("LAVADB_DEBUG not set or false")
			return
		}
		if len(args) < 3 {
			PrintHelpmsg(args)
			return
		}
		myret := lava.Del(keyrange, keyhash_val)
		fmt.Printf("deleted: %d\n", myret)
		return
	}

	if strings.Compare(args[1], "delr") == 0 {
		if debug_val == 0 {
			fmt.Println("LAVADB_DEBUG not set or false")
			return
		}
		if len(args) < 3 {
			PrintHelpmsg(args)
			return
		}
		ret := lava.DelRecurse(keyrange, nil, 1000, keyhash_val)
		fmt.Printf("deleted: %d\n", ret)
		return
	}

	fmt.Printf("not support %s operation\n", args[1])
	return
}

