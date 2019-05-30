package main

import (
	"github.com/containernetworking/cni/pkg/skel"
	"testing"
)

func TestCmdAdd(t *testing.T) {
	cmdAdd(&skel.CmdArgs{
		Args: "hello world",
	})
}
