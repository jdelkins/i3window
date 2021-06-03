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

func workspaceToNode(ipc *i3ipc.IPCSocket, ws *i3ipc.Workspace) *i3ipc.I3Node {
	root, _ := ipc.GetTree()
	for _, n := range root.Descendents() {
		if n.Type == "workspace" && n.Name == ws.Name {
			return n
		}
	}
	return nil
}

func myWorkspace(ipc *i3ipc.IPCSocket) *i3ipc.Workspace {
	wss, _ := ipc.GetWorkspaces()
	for _, w := range wss {
		if w.Output == output && (w.Visible || w.Focused) {
			return &w
		}
	}
	return nil
}

func wsCons(ipc *i3ipc.IPCSocket, ws *i3ipc.Workspace) []*i3ipc.I3Node {
	mynode := workspaceToNode(ipc, ws)
	return mynode.Descendents()
}

func showFocusedName(ipc *i3ipc.IPCSocket) {
	tree, _ := ipc.GetTree()
	focused := tree.FindFocused()
	if focused == nil {
		if len(wsCons(ipc, myWorkspace(ipc))) == 0 {
			fmt.Println("")
		}
		return
	}
	wsn := focused.Workspace()
	if wsn == nil {
		if len(wsCons(ipc, myWorkspace(ipc))) == 0 {
			fmt.Println("")
		}
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
	ipc, err := i3ipc.GetIPCSocket()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not get i3 IPC socket")
		os.Exit(1)
	}
	w_events, _ := i3ipc.Subscribe(i3ipc.I3WindowEvent)
	ws_events, _ := i3ipc.Subscribe(i3ipc.I3WorkspaceEvent)
	for {
		select {
		case <-w_events:
			showFocusedName(ipc)
		case <-ws_events:
			showFocusedName(ipc)
		}
	}
}
