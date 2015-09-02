package util

import (
	//	"bytes"
	"encoding/hex"
	"testing"
)

func TestUnDecodeLength(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 3},
		{6, 4},
		{7, 5},
		{8, 6},
		{9, 6},
		{10, 7},
		{11, 8},
		{38, 28},
	}

	for _, c := range cases {
		got := UnDecodeLength(c.in)
		if got != c.want {
			t.Errorf("UnDecodeLength(),given:%d, want %d, got %d", c.in, c.want, got)
		}
	}
}

func TestUnEncodeLength(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 6},
		{5, 7},
		{6, 8},
		{7, 10},
		{8, 11},
		{9, 12},
		{10, 14},
		{11, 15},
	}

	for _, c := range cases {
		got := UnEncodeLength(c.in)
		if got != c.want {
			t.Errorf("UnEncodeLength(),given:%d, want %d, got %d", c.in, c.want, got)
		}
	}
}

func TestUuDecode(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{

		{"202020202030202029523c202020202020202020202020202020202020202020202020202020", "00000001000027270000000000000000000000000000000000000000"},
		{"20202020202020202952402020202020202020235a202020202f4820202020232020202020202020202020", "000000000000272800000000000003e8000000fa000000030000000000000000"},
		{"20202020202020202952402020202021202020235a202020202f4820202020232020202020202020202020", "000000000000272800000001000003e8000000fa000000030000000000000000"},
		{"20202020203020203d3320202020202020202020202020202020202020202020202020202020", "00000001000075300000000000000000000000000000000000000000"},
		{"20202020202020202952402020202022202020235a202020202f4820202020232020202020202020202020", "000000000000272800000002000003e8000000fa000000030000000000000000"},
		{"202020203024394f3c462d4535272949383630402a25312d2a3020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020", "00000040466f72636554726961642028544d290000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		{"20202020203020202026302020202021", "000000010000006400000001"},
	}

	for _, c := range cases {
		input, _ := hex.DecodeString(c.in)
		got := UuDecode(input)
		output := hex.EncodeToString(got)
		if c.want != output {
			t.Errorf("UuDecode()\ngiven\t%s\nwant\t%s\ngot\t%s", c.in, c.want, output)
		}
	}

}

func TestUuEnocde(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"00000001000027270000000000000000000000000000000000000000", "202020202030202029523c202020202020202020202020202020202020202020202020202020"},
	}

	for _, c := range cases {
		input, _ := hex.DecodeString(c.in)
		got := UuEnocde(input)
		output := hex.EncodeToString(got)
		if c.want != output {
			//		if !bytes.Equal(got, c.want) {
			t.Errorf("UuEnocde()\ngiven\t%x\nwant\t%x\ngot\t%x", c.in, c.want, got)
		}
	}

}
