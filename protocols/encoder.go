package lavadb_protocol

import (
	"github.com/PromonLogicalis/asn1"
	"reflect"
)

var (
	ctx *asn1.Context
)

func init() {
	ctx = asn1.NewContext()
	if err := ctx.AddChoice("smbody", []asn1.Choice{
		{Options: "tag:30000,explicit", Type: reflect.TypeOf(OBJECTSTOREPkt{})},
	}); err != nil {
		panic(err.Error())
	}
	if err := ctx.AddChoice("body", []asn1.Choice{
		{Options: "tag:2000,explicit", Type: reflect.TypeOf(ReqLavaDBSetRecord{})},
		{Options: "tag:2001,explicit", Type: reflect.TypeOf(RspLavaDBSetRecord{})},

		{Options: "tag:2010,explicit", Type: reflect.TypeOf(ReqLavaDBGetRecord{})},
		{Options: "tag:2011,explicit", Type: reflect.TypeOf(RspLavaDBGetRecord{})},

		{Options: "tag:2020,explicit", Type: reflect.TypeOf(ReqLavaDBDelRecord{})},
		{Options: "tag:2021,explicit", Type: reflect.TypeOf(RspLavaDBDelRecord{})},

		{Options: "tag:4016,explicit", Type: reflect.TypeOf(ReqLavaDBHeartBeat{})},
		{Options: "tag:4017,explicit", Type: reflect.TypeOf(RspLavaDBHeartBeat{})},

		{Options: "tag:8100,explicit", Type: reflect.TypeOf(ReqLavaDBListRecord{})},
		{Options: "tag:8101,explicit", Type: reflect.TypeOf(RspLavaDBListRecord{})},

		{Options: "tag:8102,explicit", Type: reflect.TypeOf(ReqLavaDBSetVerRecord{})},
		{Options: "tag:8103,explicit", Type: reflect.TypeOf(RspLavaDBSetVerRecord{})},

		{Options: "tag:8104,explicit", Type: reflect.TypeOf(ReqLavaDBDelVerRecord{})},
		{Options: "tag:8105,explicit", Type: reflect.TypeOf(RspLavaDBDelVerRecord{})},

		{Options: "tag:8106,explicit", Type: reflect.TypeOf(ReqLavaDBListRecordWithEndMarker{})},
		{Options: "tag:8107,explicit", Type: reflect.TypeOf(RspLavaDBListRecordWithEndMarker{})},

		// -- TFSProxy interface invoked by preupload and http --
		{Options: "tag:8200,explicit", Type: reflect.TypeOf(ReqTFSUploadData{})},
		{Options: "tag:8201,explicit", Type: reflect.TypeOf(RspTFSUploadData{})},

		{Options: "tag:8202,explicit", Type: reflect.TypeOf(ReqTFSDownloadData{})},
		{Options: "tag:8203,explicit", Type: reflect.TypeOf(RspTFSDownloadData{})},

		{Options: "tag:8204,explicit", Type: reflect.TypeOf(ReqTFSDeleteData{})},
		{Options: "tag:8205,explicit", Type: reflect.TypeOf(RspTFSDeleteData{})},

		{Options: "tag:8206,explicit", Type: reflect.TypeOf(ReqTFSHeartBeat{})},
		{Options: "tag:8207,explicit", Type: reflect.TypeOf(RspTFSHeartBeat{})},

		{Options: "tag:8208,explicit", Type: reflect.TypeOf(ReqTFSRefreshCfg{})},
		{Options: "tag:8209,explicit", Type: reflect.TypeOf(RspTFSRefreshCfg{})},
	}); err != nil {
		panic(err.Error())
	}

}

func Encode(obj interface{}) (data []byte, err error) {
	return ctx.Encode(obj)
}

func Decode(data []byte, obj interface{}) (rest []byte, err error) {
	return ctx.Decode(data, obj)
}
