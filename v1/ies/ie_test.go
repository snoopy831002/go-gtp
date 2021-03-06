// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	v1 "github.com/wmnsk/go-gtp/v1"
	"github.com/wmnsk/go-gtp/v1/ies"
)

func TestIEs(t *testing.T) {
	cases := []struct {
		description string
		structured  *ies.IE
		serialized  []byte
	}{
		{
			"IMSI",
			ies.NewIMSI("123451234567890"),
			[]byte{0x02, 0x21, 0x43, 0x15, 0x32, 0x54, 0x76, 0x98, 0xf0},
		}, {
			"PacketTMSI",
			ies.NewPacketTMSI(0xbeebee),
			[]byte{0x05, 0x00, 0xbe, 0xeb, 0xee},
		}, {
			"AuthenticationTriplet",
			ies.NewAuthenticationTriplet(
				[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
				[]byte{0xde, 0xad, 0xbe, 0xef},
				[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77},
			),
			[]byte{
				0x09,
				0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
				0xde, 0xad, 0xbe, 0xef,
				0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77,
			},
		}, {
			"MAPCause",
			ies.NewMAPCause(v1.MAPCauseSystemFailure),
			[]byte{0x0b, 0x22},
		}, {
			"PTMSISignature",
			ies.NewPTMSISignature(0xbeebee),
			[]byte{0x0c, 0xbe, 0xeb, 0xee},
		}, {
			"MSValidated",
			ies.NewMSValidated(true),
			[]byte{0x0d, 0xff},
		}, {
			"Recovery",
			ies.NewRecovery(1),
			[]byte{0x0e, 0x01},
		}, {
			"SelectionMode",
			ies.NewSelectionMode(v1.SelectionModeMSorNetworkProvidedAPNSubscribedVerified),
			[]byte{0x0f, 0xf0},
		}, {
			"TEIDDataI",
			ies.NewTEIDDataI(0xdeadbeef),
			[]byte{0x10, 0xde, 0xad, 0xbe, 0xef},
		}, {
			"TEIDCPlane",
			ies.NewTEIDCPlane(0xdeadbeef),
			[]byte{0x11, 0xde, 0xad, 0xbe, 0xef},
		}, {
			"TEIDDataII",
			ies.NewTEIDDataII(0xdeadbeef),
			[]byte{0x12, 0xde, 0xad, 0xbe, 0xef},
		}, {
			"TeardownInd",
			ies.NewTeardownInd(true),
			[]byte{0x13, 0xff},
		}, {
			"NSAPI",
			ies.NewNSAPI(0x05),
			[]byte{0x14, 0x05},
		}, {
			"RANAPCause",
			ies.NewRANAPCause(v1.MAPCauseUnknownSubscriber),
			[]byte{0x15, 0x01},
		}, {
			"EndUserAddress/v4",
			ies.NewEndUserAddress("1.1.1.1"),
			[]byte{0x80, 0x00, 0x06, 0xf1, 0x21, 0x01, 0x01, 0x01, 0x01},
		}, {
			"EndUserAddress/v6",
			ies.NewEndUserAddress("2001::1"),
			[]byte{
				0x80, 0x00, 0x12, 0x00,
				0x57, 0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			},
		}, {
			"AccessPointName",
			ies.NewAccessPointName("some.apn.example"),
			[]byte{
				0x83, 0x00, 0x11,
				0x04, 0x73, 0x6f, 0x6d, 0x65, 0x03, 0x61, 0x70, 0x6e, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
			},
		}, {
			"GSNAddressV4",
			ies.NewGSNAddress("1.1.1.1"),
			[]byte{0x85, 0x00, 0x04, 0x01, 0x01, 0x01, 0x01},
		}, {
			"GSNAddressV6",
			ies.NewGSNAddress("2001::1"),
			[]byte{
				0x85, 0x00, 0x10,
				0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			},
		}, {
			"MSISDN",
			ies.NewMSISDN("818012345678"),
			[]byte{0x86, 0x00, 0x07, 0x91, 0x18, 0x08, 0x21, 0x43, 0x65, 0x87},
		}, {
			"AuthenticationQuintuplet",
			ies.NewAuthenticationQuintuplet(
				[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
				[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
				[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
				[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
				[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			),
			[]byte{
				0x88, 0x00, 0x52,
				0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
				0x10,
				0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
				0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
				0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
				0x10,
				0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
			},
		}, {
			"CommonFlags",
			ies.NewCommonFlags(0, 1, 0, 0, 0, 0, 0, 0),
			[]byte{0x94, 0x00, 0x01, 0x40},
		}, {
			"APNRestriction",
			ies.NewAPNRestriction(v1.APNRestrictionPrivate1),
			[]byte{0x95, 0x00, 0x01, 0x03},
		}, {
			"RATType",
			ies.NewRATType(v1.RatTypeEUTRAN),
			[]byte{0x97, 0x00, 0x01, 0x06},
		}, {
			"UserLocationInformationWithCGI",
			ies.NewUserLocationInformationWithCGI("123", "45", 0xff, 0),
			[]byte{0x98, 0x00, 0x08, 0x00, 0x21, 0xf3, 0x54, 0x00, 0xff, 0x00, 0x00},
		}, {
			"UserLocationInformationWithSAI",
			ies.NewUserLocationInformationWithSAI("123", "45", 0xff, 0),
			[]byte{0x98, 0x00, 0x08, 0x01, 0x21, 0xf3, 0x54, 0x00, 0xff, 0x00, 0x00},
		}, {
			"UserLocationInformationWithRAI",
			ies.NewUserLocationInformationWithRAI("123", "45", 0xff, 0),
			[]byte{0x98, 0x00, 0x07, 0x02, 0x21, 0xf3, 0x54, 0x00, 0xff, 0x00},
		}, {
			"MSTimeZone",
			ies.NewMSTimeZone(9*time.Hour, 0), // XXX - should be updated with more realistic value
			[]byte{0x99, 0x00, 0x02, 0x63, 0x00},
		}, {
			"IMEISV",
			ies.NewIMEISV("123450123456789"),
			[]byte{0x9a, 0x00, 0x08, 0x21, 0x43, 0x05, 0x21, 0x43, 0x65, 0x87, 0xf9},
		}, {
			"ULITimestamp",
			ies.NewULITimestamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			[]byte{0xd6, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00},
		},
	}

	for _, c := range cases {
		t.Run("Marshal/"+c.description, func(t *testing.T) {
			got, err := c.structured.Marshal()
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, c.serialized); diff != "" {
				t.Error(diff)
			}
		})

		t.Run("Parse/"+c.description, func(t *testing.T) {
			got, err := ies.Parse(c.serialized)
			if err != nil {
				t.Fatal(err)
			}

			opt := cmp.AllowUnexported(*got, *c.structured)
			if diff := cmp.Diff(got, c.structured, opt); diff != "" {
				t.Error(diff)
			}
		})
	}
}
