package ui

import (
	"fmt"
	"strconv"

	"github.com/dann-merlin/binprehend/src/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type StructureTreeView struct {
	*fyne.Container
	tree *widget.Tree
}

func toID(id string) uint64 {
	parsed, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		fmt.Printf("Failed to convert string id to uint64 id: %s\n", err)
		return 0
	}
	return parsed
}

func getChildIDs(id widget.TreeNodeID) []widget.TreeNodeID {
	if id == "" {
		return []widget.TreeNodeID{"0"}
	}
	node := model.GetNodeWithID(toID(id))
	if cont, ok := node.(model.STContainer); ok {
		var result []widget.TreeNodeID
		for _, c := range cont.GetChildren() {
			result = append(result, string(c.GetID()))
		}
		return result
	}
	fmt.Println("This should never happen anyways, I think.")
	return []widget.TreeNodeID{}
}

func isBranch(id widget.TreeNodeID) bool {
	if id == "" {
		return true
	}
	node := model.GetNodeWithID(toID(id))
	_, ok := node.(model.STContainer)
	return ok
}

func create(branch bool) fyne.CanvasObject {
	return NewSTVEntry(branch)
}

func update(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
	o.(*STVEntry).Update(model.GetNodeWithID(toID(id)))
}

func (stv *StructureTreeView) onDataChanged(idNmbr uint64) {
	// id := string(idNmbr)
	stv.tree.Refresh()
}

func NewStructureTreeView() (StructureTreeView, error) {
	addnewView := NewAddnewView()
	tree := widget.NewTree(getChildIDs, isBranch, create, update)
	cont := container.NewBorder(addnewView, nil, nil, nil, tree)
	stv := StructureTreeView{Container: cont, tree: tree}
	model.AddChangedCallback(stv.onDataChanged)
	return stv, nil
}

