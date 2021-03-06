package filestorage_test

import (
	"testing"

	"github.com/linkai-io/am/am"

	"github.com/linkai-io/am/pkg/convert"
	"github.com/linkai-io/am/pkg/filestorage"
)

func TestShardName(t *testing.T) {
	in := "abcd"
	_, err := filestorage.ShardName(in)
	if err != filestorage.ErrNameTooSmall {
		t.Fatalf("expected ErrNameToSmall")
	}

	in = "abcdefgh"
	expected := "/a/b/c/d/e/abcdefgh"
	out, _ := filestorage.ShardName(in)
	if out != expected {
		t.Fatalf("expected %v got %v", expected, out)
	}
}

func TestPathFromData(t *testing.T) {
	expected := "/2/6/8/a/0/268a0a588b41ac3726ecb5e7d5edf738b037b15b"
	data := []byte("asldkfja;sldkfjasd;lfkjasd;lfkajsdfl;kajdsf;lakjdsf")

	address := &am.ScanGroupAddress{
		OrgID:   1,
		GroupID: 1,
	}
	hashName := convert.HashData(data)
	out := filestorage.PathFromData(address, hashName)
	if out != expected {
		t.Fatalf("expected %v got %v\n", expected, out)
	}

	expected = "null"
	dataStr := ""

	out = filestorage.PathFromData(address, dataStr)
	if out != expected {
		t.Fatalf("expected %v got %v\n", expected, out)
	}
}
