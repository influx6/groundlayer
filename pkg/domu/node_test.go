package domu

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTree(t *testing.T) {
	base := Element("red", Id("red-div"))
	require.NotNil(t, base)

	textNode := Text("free-bar")
	span := HTMLSpan(Id("content-span"), textNode)
	span2 := HTMLSpan(Id("end-span"), span)

	base.AppendChild(span2)
	require.Equal(t, "/red-div/end-span/content-span", span.RefTree())

	var targetNode, err = base.FindRefNode("red-div/end-span/content-span")
	require.NoError(t, err)
	require.Equal(t, span, targetNode)
	require.NotEqual(t, span2, targetNode)

	var targetNode1, err1 = base.FindRefNode("red-div/end-span")
	require.NoError(t, err1)
	require.Equal(t, span2, targetNode1)
	require.NotEqual(t, span, targetNode1)
}

func TestJSONNodeStream(t *testing.T) {
	base := Element("red", Id("red-div"))
	textNode := Text("free-bar")
	span := HTMLSpan(Id("end-span"), textNode)
	base.AppendChild(span)

	base2 := Element("red", Id("red-div"))
	textNode2 := Text("free-bar -2")
	span2 := HTMLSpan(Id("end-span"), textNode2)
	base2.AppendChild(span2)

	var changes = base2.ReconcileStream(base)
	require.Len(t, changes, 1)
	require.Equal(t, ("end-span"), changes[0].Id)
}

func TestCarrierNode(t *testing.T) {
	base := Carrier()
	HTMLSpan(Id("end-span"), Text("free-bar")).Mount(base)
	HTMLSpan(Id("end-span1"), Text("free-bar2")).Mount(base)
	HTMLSpan(Id("end-span3"), Text("free-bar3")).Mount(base)

	require.Equal(t, 3, base.ChildCount())

	base2 := Element("red", Id("red-div"))
	base2.AppendChild(base)
	require.Equal(t, 3, base2.ChildCount())

	base3 := Element("red", Id("red-div"))
	base.Mount(base3)

	require.Equal(t, 0, base3.ChildCount())
}

func TestReplicatorNode(t *testing.T) {
	base := Replicate()
	HTMLSpan(Id("end-span"), Text("free-bar")).Mount(base)
	HTMLSpan(Id("end-span1"), Text("free-bar2")).Mount(base)
	HTMLSpan(Id("end-span3"), Text("free-bar3")).Mount(base)

	require.Equal(t, 3, base.ChildCount())

	base2 := Element("red", Id("red-div"))
	base2.AppendChild(base)
	require.Equal(t, 3, base2.ChildCount())

	base3 := Element("red", Id("red-div"))
	base.Mount(base3)

	require.Equal(t, 3, base3.ChildCount())
	require.False(t, base2.Reconcile(base3))
}

func TestJSONNodeStream2(t *testing.T) {
	base := Element("red", Id("red-div"))
	HTMLSpan(Id("end-span"), Text("free-bar")).Mount(base)
	HTMLSpan(Id("end-span1"), Text("free-bar2")).Mount(base)
	HTMLSpan(Id("end-span3"), Text("free-bar3")).Mount(base)

	base2 := Element("red", Id("red-div"))
	HTMLSpan(Id("end-span"), Text("free-bar-luba")).Mount(base2)
	HTMLSpan(Id("end-span1"), Text("free-bar2")).Mount(base2)
	HTMLSpan(Id("end-span3"), Text("free-bar3-rec-a")).Mount(base2)

	var changes = base2.ReconcileStream(base)
	require.Len(t, changes, 2)
	require.Equal(t, ("end-span"), changes[0].Id)
	require.Equal(t, ("end-span3"), changes[1].Id)
}

func TestReconcile(t *testing.T) {
	base := Element("red", Id("red-div"))
	HTMLSpan(Id("end-span"), Text("free-bar")).Mount(base)
	base2 := Element("red", Id("red-div2"))
	HTMLSpan(Id("end-span"), Text("free-bar")).Mount(base2)
	HTMLSpan(Id("end-span1"), Text("free-bar2")).Mount(base2)
	HTMLSpan(Id("end-span3"), Text("free-bar3")).Mount(base2)
	base2.Mount(base)

	baser := Element("red", Id("red-div"))
	HTMLSpan(Id("end-span"), Text("free-bar-leno")).Mount(baser)
	baser2 := Element("red", Id("red-div2"))
	HTMLSpan(Id("end-span"), Text("free-bar")).Mount(baser2)
	HTMLSpan(Id("end-span1"), Text("free-bar2")).Mount(baser2)
	HTMLSpan(Id("end-span3"), Text("free-bar3-a")).Mount(baser2)
	baser2.Mount(baser)

	var changes = baser.Reconcile(base)
	require.True(t, changes)

	var bm bytes.Buffer
	require.NoError(t, baser.RenderHTMLDiff(&bm, true))
}

func TestJSONNodeStream3(t *testing.T) {
	base := Element("red", Id("red-div"))
	HTMLSpan(Id("end-span"), Text("free-bar")).Mount(base)
	base2 := Element("red", Id("red-div2"))
	HTMLSpan(Id("end-span"), Text("free-bar")).Mount(base2)
	HTMLSpan(Id("end-span1"), Text("free-bar2")).Mount(base2)
	HTMLSpan(Id("end-span3"), Text("free-bar3")).Mount(base2)
	base2.Mount(base)

	baser := Element("red", Id("red-div"))
	HTMLSpan(Id("end-span"), Text("free-bar-leno")).Mount(baser)
	baser2 := Element("red", Id("red-div2"))
	HTMLSpan(Id("end-span"), Text("free-bar")).Mount(baser2)
	HTMLSpan(Id("end-span1"), Text("free-bar2")).Mount(baser2)
	HTMLSpan(Id("end-span3"), Text("free-bar3-a")).Mount(baser2)
	baser2.Mount(baser)

	var changes = baser.ReconcileStream(base)
	require.Len(t, changes, 2)
	require.Equal(t, ("end-span"), changes[0].Id)
	require.Equal(t, ("end-span3"), changes[1].Id)
}

func TestJSONNodeStream4(t *testing.T) {
	base := Element("red", Id("red-div"))
	HTMLSpan(Id("end-span"), Text("free-bar")).Mount(base)
	base2 := Element("red", Id("red-div"))
	HTMLSpan(Id("end-span"), Text("free-bar")).Mount(base2)
	HTMLSpan(Id("end-span1"), Text("free-bar2")).Mount(base2)
	HTMLSpan(Id("end-span3"), Text("free-bar3")).Mount(base2)
	base2.Mount(base)

	baser := Element("red", Id("red-div"))
	HTMLSpan(Id("end-span"), Text("free-bar")).Mount(baser)
	baser2 := Element("red", Id("red-div"))
	HTMLSpan(Id("end-span"), Text("free-bar")).Mount(baser2)
	HTMLSpan(Id("end-span1"), Text("free-bar2")).Mount(baser2)
	HTMLSpan(Id("end-span3"), Text("free-bar3")).Mount(baser2)
	baser2.Mount(baser)

	var changes = baser.ReconcileStream(base)
	require.Len(t, changes, 0)
}

func TestJSONNodeStream5(t *testing.T) {
	base := Element("red", Id("red-div"))
	HTMLSpan(Id("end-span"), Text("free-bar")).Mount(base)
	base2 := Element("red", Id("red-div2"))
	HTMLSpan(Id("end-span"), Text("free-bar")).Mount(base2)
	HTMLSpan(Id("end-span1"), Text("free-bar2")).Mount(base2)
	HTMLSpan(Id("end-span3"), Text("free-bar3")).Mount(base2)
	base2.Mount(base)

	baser := Element("red", Id("red-div"))
	HTMLSpan(Id("end-span"), Text("free-bar")).Mount(baser)
	baser2 := Element("red", Id("red-div2"))
	HTMLSpan(Id("end-span"), Text("free-bar")).Mount(baser2)
	HTMLSpan(Id("end-span1"), Text("free-bar2")).Mount(baser2)
	HTMLSpan(Id("end-span3"), Text("free-bar3-a")).Mount(baser2)
	baser2.Mount(baser)

	var changes = baser.ReconcileStream(base)
	require.Len(t, changes, 1)
	require.Equal(t, ("end-span3"), changes[0].Id)
}

func TestTreeClone(t *testing.T) {
	base := Element("red", Id("red-div"))
	require.NotNil(t, base)

	textNode := Text("free-bar")
	span := HTMLSpan(Id("content-span"), textNode)
	span2 := HTMLSpan(Id("end-span"), span)

	base.AppendChild(span2)
	require.Equal(t, "/red-div/end-span/content-span", span.RefTree())

	var baseClone = base.Clone(true)
	require.Equal(t, base.atid, baseClone.atid)
	require.Equal(t, base.tid, baseClone.tid)

	var baseContent, baseCloneContent strings.Builder
	require.NoError(t, base.RenderShallowHTML(&baseContent, false))
	require.NoError(t, baseClone.RenderShallowHTML(&baseCloneContent, false))
	require.Equal(t, baseContent.String(), baseCloneContent.String())

	baseContent.Reset()
	baseCloneContent.Reset()

	require.NoError(t, base.RenderHTMLTo(&baseContent, false, true))
	require.NoError(t, baseClone.RenderHTMLTo(&baseCloneContent, false, true))
	require.Equal(t, baseContent.String(), baseCloneContent.String())

	baseContent.Reset()
	baseCloneContent.Reset()
}

func TestDeeperChangedJSON(t *testing.T) {
	base := NewNode(ElementNode, "red")
	Id("red-div").Mount(base)

	bag := HTMLDiv(Id("content-span"), Text("free-bar"))

	require.NotNil(t, base)
	base.AppendChild(HTMLSection(Id("end-section"), bag))

	baseClone := base.Clone(true)
	var baseContent, baseCloneContent strings.Builder
	require.NoError(t, base.RenderShallowHTML(&baseContent, true))
	require.NoError(t, baseClone.RenderShallowHTML(&baseCloneContent, true))
	require.Equal(t, baseContent.String(), baseCloneContent.String())

	section, err := baseClone.Get(0)
	require.NoError(t, err)
	require.NotNil(t, section)

	sectionBag, err := section.Get(0)
	require.NoError(t, err)
	require.NotNil(t, sectionBag)

	sectionBag.AppendChild(HTMLDiv(Id("red-bag")))

	baseContent.Reset()
}

func TestShallowRender(t *testing.T) {
	base := NewNode(ElementNode, "red")
	Id("red-div").Mount(base)

	require.NotNil(t, base)
	base.AppendChild(HTMLSpan(Id("end-span"), HTMLSpan(Id("content-span"), Text("free-bar"))))

	var baseContent strings.Builder
	require.NoError(t, base.RenderShallowHTML(&baseContent, true))

	var expected = `<red _tid="` + base.tid + `" _ref="/red-div" id="red-div">` + "\n\n" + `</red>`
	require.Equal(t, expected, baseContent.String())
}

func TestChangedJSON(t *testing.T) {
	base := NewNode(ElementNode, "red")
	Id("red-div").Mount(base)

	require.NotNil(t, base)
	base.AppendChild(HTMLSpan(Id("end-span"), HTMLSpan(Id("content-span"), Text("free-bar"))))

	baseClone := base.Clone(true)

	var baseContent, baseCloneContent strings.Builder
	require.NoError(t, base.RenderShallowHTML(&baseContent, true))
	require.NoError(t, baseClone.RenderShallowHTML(&baseCloneContent, true))
	require.Equal(t, baseContent.String(), baseCloneContent.String())

	baseContent.Reset()
}

func TestNodeEvent(t *testing.T) {
	var parent = NewJSONEvent("click", "yo", "yodo", "yobo")

	require.Equal(t, `event-click="yo|yodo|yobo"`, parent.Text())
}

func TestNode(t *testing.T) {
	base := Element("red", Id("767h"))
	require.NotNil(t, base)

	for i := 0; i < 1000; i++ {
		base.AppendChild(Element(fmt.Sprintf("red%d", i), Id("65jnj")))
	}

	require.Equal(t, 1000, base.ChildCount())
}

func TestNode_Remove(t *testing.T) {
	base := Element("red", Id("767h"))
	require.NotNil(t, base)

	for i := 0; i < 1000; i++ {
		base.AppendChild(Element(fmt.Sprintf("red%d", i), Id("65jnj")))
	}

	require.Equal(t, 1000, base.ChildCount())

	deletes := 500
	for i := 0; i < deletes; i++ {
		var target = rand.Intn(deletes)
		if node, err := base.Get(target); err == nil {
			node.Remove()
		}
	}

	require.Equal(t, 500, base.ChildCount())
}

func TestNode_Remove_Balance(t *testing.T) {
	base := Element("red", Id("767h"))
	require.NotNil(t, base)

	for i := 0; i < 1000; i++ {
		base.AppendChild(Element(fmt.Sprintf("red%d", i), Id("65jnj")))
	}

	require.Equal(t, 1000, base.ChildCount())

	deletes := 500
	for i := 0; i < deletes; i++ {
		var target = rand.Intn(deletes)
		if node, err := base.Get(target); err == nil {
			node.Remove()
			node.Balance()
		}
	}

	require.Equal(t, 500, base.ChildCount())
}

func BenchmarkNode_Append(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	base := Element("red", Id("232"))
	for i := 0; i < 1001; i++ {
		base.AppendChild(Element(strconv.Itoa(i), Id("45g")))
	}
}

func BenchmarkNode_Append_Remove_Balance(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	base := Element("red", Id("232"))

	var count = 4001
	for i := 0; i < count; i++ {
		base.AppendChild(Element("sf", Id(strconv.Itoa(i))))
	}

	var deletes = count / 2
	for i := 0; i < deletes; i++ {
		var target = rand.Intn(deletes)
		if node, err := base.Get(target); err == nil {
			node.Remove()
		}
	}
	base.Balance()
}

func BenchmarkNode_Append_Remove_With_Balance(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	base := Element("red", Id("232"))

	var count = 4001
	for i := 0; i < count; i++ {
		base.AppendChild(Element(strconv.Itoa(i), Id("343f")))
	}

	var deletes = count / 2
	for i := 0; i < deletes; i++ {
		var target = rand.Intn(deletes)
		if node, err := base.Get(target); err == nil {
			node.Remove()
			base.Balance()
		}
	}
}

func BenchmarkNode_Append_SwapAll(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	base := Element("red", Id("322"))

	var child *Node
	for i := 0; i < 1001; i++ {
		newChild := Element(strconv.Itoa(i), Id("331"))
		if child == nil {
			base.AppendChild(newChild)
			child = newChild
			continue
		}

		child.SwapAll(newChild)
		child = newChild
	}
}

func BenchmarkNode_Append_SwapNode(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	base := Element("red", Id("112"))

	var child *Node
	for i := 0; i < 1001; i++ {
		newChild := Element(strconv.Itoa(i), Id("132"))
		if child == nil {
			base.AppendChild(newChild)
			child = newChild
			continue
		}

		child.SwapNode(newChild)
		child = newChild
	}
}
