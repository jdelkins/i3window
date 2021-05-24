package main

import (
	"fmt"
	"os"

	"github.com/mdirkse/i3ipc"
)

var output string

func nodeToWorkspace(ipc *i3ipc.IPCSocket, n *i3ipc.I3Node) *i3ipc.Workspace {
	wss, _ := ipc.GetWorkspaces()
	for _, w := range wss {
		if w.Name == n.Name {
			return &w
		}
	}
	return nil
}

func showFocusedName() {
	ipc, _ := i3ipc.GetIPCSocket()
	tree, _ := ipc.GetTree()
	focused := tree.FindFocused()
	if focused == nil {
		fmt.Println("")
		return
	}
	wsn := focused.Workspace()
	if wsn == nil {
		fmt.Println("")
		return
	}
	ws := nodeToWorkspace(ipc, wsn)
	if ws.Output == output {
		fmt.Println(focused.Name)
	}
}

func main() {
	output = os.Args[1]
	i3ipc.StartEventListener()
	w_events, _ := i3ipc.Subscribe(i3ipc.I3WindowEvent)
	ws_events, _ := i3ipc.Subscribe(i3ipc.I3WorkspaceEvent)
	for {
		select {
		case <-w_events:
			showFocusedName()
		case <-ws_events:
			showFocusedName()
		}
	}
}
