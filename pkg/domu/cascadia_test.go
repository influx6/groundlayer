package domu

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCascadiaQuery(t *testing.T) {
	firstChild := Element("div", Id("div1"), Element("span", Id("span-2"), Element("span", Id("span-3"))))
	require.NotNil(t, firstChild)

	query, err := Query("div")
	require.NoError(t, err)
	require.NotNil(t, query)

	require.True(t, query.Match(firstChild))

	query2, err := Query("div span#span-2")
	require.NoError(t, err)
	require.NotNil(t, query2)

	dn := query2.MatchAll(firstChild)
	require.Len(t, dn, 1)
}

var elems = []string{"section", "div", "span", "em", "p", "a", "img"}

func generateRandomNodeTree(n int, base *Node) {
	node := NewNode(ElementNode, "section")
	Id("section-43").Mount(node)
	for i := 0; i < n; i++ {
		elem := elems[rand.Intn(len(elems))]
		var nd = NewNode(ElementNode, elem)
		Id(strconv.Itoa(i)).Mount(nd)
		node.AppendChild(nd)
	}
	base.AppendChild(node)
}
