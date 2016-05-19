package iscsilt

import (
	"testing"
//	"io"
	"bytes"
	"fmt"
)

func TestSession(t *testing.T) {
	var p ISCSIConnection

	for _, el := range pks {
		fmt.Println("Case is: ", el.Name)
		r := bytes.NewReader(el.Request)
		n, err := p.ReadFrom(r)
		fmt.Println(n, err)
		if n, r := compareSlice(p.DataW, el.Responce); r {
			fmt.Println("Case ", el.Name, " is OK!.")
		} else {
			fmt.Println(len(p.DataW), len(el.Responce), n)
			fmt.Printf("%x\n%x\n", p.DataW, el.Responce)
			t.Fatalf("Case "+el.Name+" is BAD!!.")
		}
	}
}

func compareSlice(s1, s2 []byte) (int, bool) {
	for i, el := range s1 {
		if i == 5 || i == 6 || i == 7 || i >= 48 {
			continue
		}
		if el != s2[i] {
			return i, false
		}
	}
	return 0, true
}

type packExample struct {
	Name		string
	Request		[]byte
	Responce	[]byte
}
var (
pks []packExample = []packExample{{
Name: "ISCSILoginRequest", Request: []byte{	/* len 296 bytes */
0x43, 0x87, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf7, /* C....... */
0x00, 0x02, 0x3d, 0x00, 0x00, 0x00, 0x00, 0x00, /* ..=..... */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, /* Initiato */
0x72, 0x4e, 0x61, 0x6d, 0x65, 0x3d, 0x69, 0x71, /* rName=iq */
0x6e, 0x2e, 0x31, 0x39, 0x39, 0x33, 0x2d, 0x30, /* n.1993-0 */
0x38, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x64, 0x65, /* 8.org.de */
0x62, 0x69, 0x61, 0x6e, 0x3a, 0x30, 0x31, 0x3a, /* bian:01: */
0x31, 0x38, 0x37, 0x61, 0x39, 0x31, 0x61, 0x66, /* 187a91af */
0x34, 0x30, 0x00, 0x49, 0x6e, 0x69, 0x74, 0x69, /* 40.Initi */
0x61, 0x74, 0x6f, 0x72, 0x41, 0x6c, 0x69, 0x61, /* atorAlia */
0x73, 0x3d, 0x73, 0x69, 0x74, 0x2d, 0x31, 0x39, /* s=sit-19 */
0x32, 0x30, 0x00, 0x53, 0x65, 0x73, 0x73, 0x69, /* 20.Sessi */
0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x3d, 0x44, /* onType=D */
0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, /* iscovery */
0x00, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x44, /* .HeaderD */
0x69, 0x67, 0x65, 0x73, 0x74, 0x3d, 0x4e, 0x6f, /* igest=No */
0x6e, 0x65, 0x00, 0x44, 0x61, 0x74, 0x61, 0x44, /* ne.DataD */
0x69, 0x67, 0x65, 0x73, 0x74, 0x3d, 0x4e, 0x6f, /* igest=No */
0x6e, 0x65, 0x00, 0x44, 0x65, 0x66, 0x61, 0x75, /* ne.Defau */
0x6c, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x57, /* ltTime2W */
0x61, 0x69, 0x74, 0x3d, 0x32, 0x00, 0x44, 0x65, /* ait=2.De */
0x66, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x69, 0x6d, /* faultTim */
0x65, 0x32, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6e, /* e2Retain */
0x3d, 0x30, 0x00, 0x49, 0x46, 0x4d, 0x61, 0x72, /* =0.IFMar */
0x6b, 0x65, 0x72, 0x3d, 0x4e, 0x6f, 0x00, 0x4f, /* ker=No.O */
0x46, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x3d, /* FMarker= */
0x4e, 0x6f, 0x00, 0x45, 0x72, 0x72, 0x6f, 0x72, /* No.Error */
0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, /* Recovery */
0x4c, 0x65, 0x76, 0x65, 0x6c, 0x3d, 0x30, 0x00, /* Level=0. */
0x4d, 0x61, 0x78, 0x52, 0x65, 0x63, 0x76, 0x44, /* MaxRecvD */
0x61, 0x74, 0x61, 0x53, 0x65, 0x67, 0x6d, 0x65, /* ataSegme */
0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, /* ntLength */
0x3d, 0x33, 0x32, 0x37, 0x36, 0x38, 0x00, 0x00},  /* =32768.. */

Responce: []byte{0x23, 0x87, 0x00, 0x00, 0x00, 0x00, /*   #..... */
0x00, 0x8e, 0x00, 0x02, 0x3d, 0x00, 0x00, 0x00, /* ....=... */
0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, /* ..Target */
0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x47, 0x72, /* PortalGr */
0x6f, 0x75, 0x70, 0x54, 0x61, 0x67, 0x3d, 0x31, /* oupTag=1 */
0x00, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x44, /* .HeaderD */
0x69, 0x67, 0x65, 0x73, 0x74, 0x3d, 0x4e, 0x6f, /* igest=No */
0x6e, 0x65, 0x00, 0x44, 0x61, 0x74, 0x61, 0x44, /* ne.DataD */
0x69, 0x67, 0x65, 0x73, 0x74, 0x3d, 0x4e, 0x6f, /* igest=No */
0x6e, 0x65, 0x00, 0x44, 0x65, 0x66, 0x61, 0x75, /* ne.Defau */
0x6c, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x57, /* ltTime2W */
0x61, 0x69, 0x74, 0x3d, 0x32, 0x00, 0x44, 0x65, /* ait=2.De */
0x66, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x69, 0x6d, /* faultTim */
0x65, 0x32, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6e, /* e2Retain */
0x3d, 0x30, 0x00, 0x49, 0x46, 0x4d, 0x61, 0x72, /* =0.IFMar */
0x6b, 0x65, 0x72, 0x3d, 0x4e, 0x6f, 0x00, 0x4f, /* ker=No.O */
0x46, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x3d, /* FMarker= */
0x4e, 0x6f, 0x00, 0x45, 0x72, 0x72, 0x6f, 0x72, /* No.Error */
0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, /* Recovery */
0x4c, 0x65, 0x76, 0x65, 0x6c, 0x3d, 0x30, 0x00, /* Level=0. */
0x00, 0x00}},                                      /* .. */

{Name: "TextCommandRequest", Request: []byte{
0x04, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, /* ........ */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x01, 0xff, 0xff, 0xff, 0xff, /* ........ */
0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, /* ........ */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x53, 0x65, 0x6e, 0x64, 0x54, 0x61, 0x72, 0x67, /* SendTarg */
0x65, 0x74, 0x73, 0x3d, 0x41, 0x6c, 0x6c, 0x00},  /* ets=All. */

Responce: []byte{0x24, 0x80, 0x00, 0x00, 0x00, 0x00, /*   $..... */
0x00, 0x72, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* .r...... */
0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xff, 0xff, /* ........ */
0xff, 0xff, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, /* ........ */
0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, /* ..Target */
0x4e, 0x61, 0x6d, 0x65, 0x3d, 0x69, 0x71, 0x6e, /* Name=iqn */
0x2e, 0x32, 0x30, 0x31, 0x35, 0x2d, 0x30, 0x39, /* .2015-09 */
0x2e, 0x6e, 0x70, 0x70, 0x2e, 0x61, 0x6e, 0x64, /* .npp.and */
0x79, 0x75, 0x3a, 0x73, 0x74, 0x6f, 0x72, 0x61, /* yu:stora */
0x67, 0x65, 0x2e, 0x6c, 0x75, 0x6e, 0x31, 0x00, /* ge.lun1. */
0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x64, /* TargetAd */
0x64, 0x72, 0x65, 0x73, 0x73, 0x3d, 0x31, 0x37, /* dress=17 */
0x32, 0x2e, 0x32, 0x34, 0x2e, 0x31, 0x39, 0x30, /* 2.24.190 */
0x2e, 0x32, 0x35, 0x33, 0x3a, 0x33, 0x32, 0x36, /* .253:326 */
0x30, 0x2c, 0x31, 0x00, 0x54, 0x61, 0x72, 0x67, /* 0,1.Targ */
0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, /* etAddres */
0x73, 0x3d, 0x31, 0x37, 0x32, 0x2e, 0x31, 0x37, /* s=172.17 */
0x2e, 0x30, 0x2e, 0x31, 0x3a, 0x33, 0x32, 0x36, /* .0.1:326 */
0x30, 0x2c, 0x31, 0x00, 0x00, 0x00}},              /* 0,1... */

{Name: "ISCSILoginRequest", Request: []byte{
0x43, 0x87, 0x00, 0x00, 0x00, 0x00, 0x01, 0xc3, /* C....... */
0x00, 0x02, 0x3d, 0x01, 0x00, 0x00, 0x00, 0x00, /* ..=..... */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, /* Initiato */
0x72, 0x4e, 0x61, 0x6d, 0x65, 0x3d, 0x69, 0x71, /* rName=iq */
0x6e, 0x2e, 0x31, 0x39, 0x39, 0x33, 0x2d, 0x30, /* n.1993-0 */
0x38, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x64, 0x65, /* 8.org.de */
0x62, 0x69, 0x61, 0x6e, 0x3a, 0x30, 0x31, 0x3a, /* bian:01: */
0x31, 0x38, 0x37, 0x61, 0x39, 0x31, 0x61, 0x66, /* 187a91af */
0x34, 0x30, 0x00, 0x49, 0x6e, 0x69, 0x74, 0x69, /* 40.Initi */
0x61, 0x74, 0x6f, 0x72, 0x41, 0x6c, 0x69, 0x61, /* atorAlia */
0x73, 0x3d, 0x73, 0x69, 0x74, 0x2d, 0x31, 0x39, /* s=sit-19 */
0x32, 0x30, 0x00, 0x54, 0x61, 0x72, 0x67, 0x65, /* 20.Targe */
0x74, 0x4e, 0x61, 0x6d, 0x65, 0x3d, 0x69, 0x71, /* tName=iq */
0x6e, 0x2e, 0x32, 0x30, 0x31, 0x35, 0x2d, 0x30, /* n.2015-0 */
0x39, 0x2e, 0x6e, 0x70, 0x70, 0x2e, 0x61, 0x6e, /* 9.npp.an */
0x64, 0x79, 0x75, 0x3a, 0x73, 0x74, 0x6f, 0x72, /* dyu:stor */
0x61, 0x67, 0x65, 0x2e, 0x6c, 0x75, 0x6e, 0x31, /* age.lun1 */
0x00, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, /* .Session */
0x54, 0x79, 0x70, 0x65, 0x3d, 0x4e, 0x6f, 0x72, /* Type=Nor */
0x6d, 0x61, 0x6c, 0x00, 0x48, 0x65, 0x61, 0x64, /* mal.Head */
0x65, 0x72, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, /* erDigest */
0x3d, 0x4e, 0x6f, 0x6e, 0x65, 0x00, 0x44, 0x61, /* =None.Da */
0x74, 0x61, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, /* taDigest */
0x3d, 0x4e, 0x6f, 0x6e, 0x65, 0x00, 0x44, 0x65, /* =None.De */
0x66, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x69, 0x6d, /* faultTim */
0x65, 0x32, 0x57, 0x61, 0x69, 0x74, 0x3d, 0x32, /* e2Wait=2 */
0x00, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, /* .Default */
0x54, 0x69, 0x6d, 0x65, 0x32, 0x52, 0x65, 0x74, /* Time2Ret */
0x61, 0x69, 0x6e, 0x3d, 0x30, 0x00, 0x49, 0x46, /* ain=0.IF */
0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x3d, 0x4e, /* Marker=N */
0x6f, 0x00, 0x4f, 0x46, 0x4d, 0x61, 0x72, 0x6b, /* o.OFMark */
0x65, 0x72, 0x3d, 0x4e, 0x6f, 0x00, 0x45, 0x72, /* er=No.Er */
0x72, 0x6f, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x76, /* rorRecov */
0x65, 0x72, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, /* eryLevel */
0x3d, 0x30, 0x00, 0x49, 0x6e, 0x69, 0x74, 0x69, /* =0.Initi */
0x61, 0x6c, 0x52, 0x32, 0x54, 0x3d, 0x4e, 0x6f, /* alR2T=No */
0x00, 0x49, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, /* .Immedia */
0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x3d, 0x59, /* teData=Y */
0x65, 0x73, 0x00, 0x4d, 0x61, 0x78, 0x42, 0x75, /* es.MaxBu */
0x72, 0x73, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, /* rstLengt */
0x68, 0x3d, 0x31, 0x36, 0x37, 0x37, 0x36, 0x31, /* h=167761 */
0x39, 0x32, 0x00, 0x46, 0x69, 0x72, 0x73, 0x74, /* 92.First */
0x42, 0x75, 0x72, 0x73, 0x74, 0x4c, 0x65, 0x6e, /* BurstLen */
0x67, 0x74, 0x68, 0x3d, 0x32, 0x36, 0x32, 0x31, /* gth=2621 */
0x34, 0x34, 0x00, 0x4d, 0x61, 0x78, 0x4f, 0x75, /* 44.MaxOu */
0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, /* tstandin */
0x67, 0x52, 0x32, 0x54, 0x3d, 0x31, 0x00, 0x4d, /* gR2T=1.M */
0x61, 0x78, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, /* axConnec */
0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3d, 0x31, 0x00, /* tions=1. */
0x44, 0x61, 0x74, 0x61, 0x50, 0x44, 0x55, 0x49, /* DataPDUI */
0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x3d, 0x59, /* nOrder=Y */
0x65, 0x73, 0x00, 0x44, 0x61, 0x74, 0x61, 0x53, /* es.DataS */
0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, /* equenceI */
0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x3d, 0x59, /* nOrder=Y */
0x65, 0x73, 0x00, 0x4d, 0x61, 0x78, 0x52, 0x65, /* es.MaxRe */
0x63, 0x76, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, /* cvDataSe */
0x67, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, /* gmentLen */
0x67, 0x74, 0x68, 0x3d, 0x32, 0x36, 0x32, 0x31, /* gth=2621 */
0x34, 0x34, 0x00, 0x00},                          /* 44.. */

Responce: []byte{0x23, 0x87, 0x00, 0x00, 0x00, 0x00, /*   #..... */
0x01, 0x2c, 0x00, 0x02, 0x3d, 0x01, 0x00, 0x00, /* .,..=... */
0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, /* ..Target */
0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x47, 0x72, /* PortalGr */
0x6f, 0x75, 0x70, 0x54, 0x61, 0x67, 0x3d, 0x31, /* oupTag=1 */
0x00, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x44, /* .HeaderD */
0x69, 0x67, 0x65, 0x73, 0x74, 0x3d, 0x4e, 0x6f, /* igest=No */
0x6e, 0x65, 0x00, 0x44, 0x61, 0x74, 0x61, 0x44, /* ne.DataD */
0x69, 0x67, 0x65, 0x73, 0x74, 0x3d, 0x4e, 0x6f, /* igest=No */
0x6e, 0x65, 0x00, 0x44, 0x65, 0x66, 0x61, 0x75, /* ne.Defau */
0x6c, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x57, /* ltTime2W */
0x61, 0x69, 0x74, 0x3d, 0x32, 0x00, 0x44, 0x65, /* ait=2.De */
0x66, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x69, 0x6d, /* faultTim */
0x65, 0x32, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6e, /* e2Retain */
0x3d, 0x30, 0x00, 0x49, 0x46, 0x4d, 0x61, 0x72, /* =0.IFMar */
0x6b, 0x65, 0x72, 0x3d, 0x4e, 0x6f, 0x00, 0x4f, /* ker=No.O */
0x46, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x3d, /* FMarker= */
0x4e, 0x6f, 0x00, 0x45, 0x72, 0x72, 0x6f, 0x72, /* No.Error */
0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, /* Recovery */
0x4c, 0x65, 0x76, 0x65, 0x6c, 0x3d, 0x30, 0x00, /* Level=0. */
0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x52, /* InitialR */
0x32, 0x54, 0x3d, 0x59, 0x65, 0x73, 0x00, 0x49, /* 2T=Yes.I */
0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, /* mmediate */
0x44, 0x61, 0x74, 0x61, 0x3d, 0x59, 0x65, 0x73, /* Data=Yes */
0x00, 0x4d, 0x61, 0x78, 0x42, 0x75, 0x72, 0x73, /* .MaxBurs */
0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x3d, /* tLength= */
0x32, 0x36, 0x32, 0x31, 0x34, 0x34, 0x00, 0x46, /* 262144.F */
0x69, 0x72, 0x73, 0x74, 0x42, 0x75, 0x72, 0x73, /* irstBurs */
0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x3d, /* tLength= */
0x36, 0x35, 0x35, 0x33, 0x36, 0x00, 0x4d, 0x61, /* 65536.Ma */
0x78, 0x4f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, /* xOutstan */
0x64, 0x69, 0x6e, 0x67, 0x52, 0x32, 0x54, 0x3d, /* dingR2T= */
0x31, 0x00, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6e, /* 1.MaxCon */
0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, /* nections */
0x3d, 0x31, 0x00, 0x44, 0x61, 0x74, 0x61, 0x50, /* =1.DataP */
0x44, 0x55, 0x49, 0x6e, 0x4f, 0x72, 0x64, 0x65, /* DUInOrde */
0x72, 0x3d, 0x59, 0x65, 0x73, 0x00, 0x44, 0x61, /* r=Yes.Da */
0x74, 0x61, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, /* taSequen */
0x63, 0x65, 0x49, 0x6e, 0x4f, 0x72, 0x64, 0x65, /* ceInOrde */
0x72, 0x3d, 0x59, 0x65, 0x73, 0x00}},              /* r=Yes. */

{Name: "SCSICommandRequest", Request: []byte{
            0x01, 0xc1, 0x00, 0x00, 0x00, 0x00, /*   ...... */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x24, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* .$...... */
0x00, 0x01, 0x12, 0x00, 0x00, 0x00, 0x24, 0x00, /* ......$. */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00},                                      /* .. */

Responce: []byte{0x25, 0x81, 0x00, 0x00, 0x00, 0x00, /*   %..... */
0x00, 0x24, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* .$...... */
0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xff, 0xff, /* ........ */
0xff, 0xff, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, /* ........ */
0x00, 0x01, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00, /* .....!.. */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x00, 0x04, 0x52, 0x3b, 0x00, /* .....R;. */
0x00, 0x02, 0x49, 0x45, 0x54, 0x20, 0x20, 0x20, /* ..IET    */
0x20, 0x20, 0x56, 0x49, 0x52, 0x54, 0x55, 0x41, /*   VIRTUA */
0x4c, 0x2d, 0x44, 0x49, 0x53, 0x4b, 0x20, 0x20, /* L-DISK   */
0x20, 0x20, 0x30, 0x20, 0x20, 0x20}},              /*   0    */

{Name: "SCSICommandRequest", Request: []byte{
	    0x01, 0xc1, 0x00, 0x00, 0x00, 0x00, /*   ...... */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x40, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, /* .@...... */
0x00, 0x02, 0x12, 0x00, 0x00, 0x00, 0x40, 0x00, /* ......@. */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00},                                      /* .. */

Responce: []byte{
	    0x25, 0x81, 0x00, 0x00, 0x00, 0x00, /*   %..... */
0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* .@...... */
0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0xff, 0xff, /* ........ */
0xff, 0xff, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, /* ........ */
0x00, 0x02, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, /* .....".. */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x00, 0x04, 0x52, 0x3b, 0x00, /* .....R;. */
0x00, 0x02, 0x49, 0x45, 0x54, 0x20, 0x20, 0x20, /* ..IET    */
0x20, 0x20, 0x56, 0x49, 0x52, 0x54, 0x55, 0x41, /*   VIRTUA */
0x4c, 0x2d, 0x44, 0x49, 0x53, 0x4b, 0x20, 0x20, /* L-DISK   */
0x20, 0x20, 0x30, 0x20, 0x20, 0x20, 0x00, 0x00, /*   0   .. */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x00, 0x03, 0x20, 0x09, 0x60, /* ..... .` */
0x03, 0x00}},                                      /* .. */

{Name: "SCSICommandRequest", Request: []byte{
	    0x01, 0xc1, 0x00, 0x00, 0x00, 0x00, /*   ...... */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0xff, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, /* ........ */
0x00, 0x03, 0x12, 0x01, 0x00, 0x00, 0xff, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00},                                      /* .. */

Responce: []byte{
	    0x25, 0x83, 0x00, 0x00, 0x00, 0x00, /*   %..... */
0x00, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0xff, 0xff, /* ........ */
0xff, 0xff, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, /* ........ */
0x00, 0x03, 0x00, 0x00, 0x00, 0x23, 0x00, 0x00, /* .....#.. */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0xf8, 0x00, 0x00, 0x00, 0x03, 0x00, 0x80, /* ........ */
0x83, 0x00}},                                      /* .. */

{Name: "SCSICommandTestUnitReady", Request: []byte{
	    0x01, 0x81, 0x00, 0x00, 0x00, 0x00, /*   ...... */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x00, /* ..|..... */
0x00, 0x00, 0x00, 0x00, 0x00, 0x7b, 0x00, 0x00, /* .....{.. */
0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* .|...... */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00},                                      /* .. */

Responce: []byte{
	    0x21, 0x80, 0x00, 0x00, 0x00, 0x00, /*   !..... */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x00, /* ..|..... */
0x00, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, /* .....|.. */
0x00, 0x7c, 0x00, 0x00, 0x00, 0x9c, 0x00, 0x00, /* .|...... */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00}},                                      /* .. */

{Name: "SCSICommandRequest", Request: []byte{
	    0x01, 0x81, 0x00, 0x00, 0x00, 0x00, /*   ...... */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x7d, 0x00, 0x00, 0x00, 0x00, 0x00, /* ..}..... */
0x00, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, /* .....|.. */
0x00, 0x7d, 0x85, 0x06, 0x20, 0x00, 0x05, 0x00, /* .}.. ... */
0xfe, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, /* .......@ */
0xef, 0x00},                                      /* .. */
Responce: []byte{
	    0x21, 0x80, 0x00, 0x02, 0x00, 0x00, /*   !..... */
0x00, 0x14, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x7d, 0x00, 0x00, 0x00, 0x00, 0x00, /* ..}..... */
0x00, 0x00, 0x00, 0x00, 0x00, 0x7d, 0x00, 0x00, /* .....}.. */
0x00, 0x7d, 0x00, 0x00, 0x00, 0x9d, 0x00, 0x00, /* .}...... */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x12, 0xf0, 0x00, 0x05, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x20, 0x00, 0x00, 0x00, 0x00, 0x00}},              /*  ..... */
{Name: "SCSICommandRequest", Request: []byte{
	    0x01, 0xc1, 0x00, 0x00, 0x00, 0x00, /*   ...... */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x7e, 0x00, 0x00, 0x00, 0x00, 0x00, /* ..~..... */
0x02, 0x00, 0x00, 0x00, 0x00, 0x7d, 0x00, 0x00, /* .....}.. */
0x00, 0x7e, 0x85, 0x08, 0x0e, 0x00, 0x00, 0x00, /* .~...... */
0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, /* .......@ */
0xec, 0x00},                                      /* .. */
Responce: []byte{
	    0x21, 0x82, 0x00, 0x02, 0x00, 0x00, /*   !..... */
0x00, 0x14, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x00, 0x00, 0x7e, 0x00, 0x00, 0x00, 0x00, 0x00, /* ..~..... */
0x00, 0x00, 0x00, 0x00, 0x00, 0x7e, 0x00, 0x00, /* .....~.. */
0x00, 0x7e, 0x00, 0x00, 0x00, 0x9e, 0x00, 0x00, /* .~...... */
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x02, 0x00, 0x00, 0x12, 0xf0, 0x00, 0x05, 0x00, /* ........ */
0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x00, /* ........ */
0x20, 0x00, 0x00, 0x00, 0x00, 0x00}}}              /*  ..... */

)