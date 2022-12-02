package xmlcompare

import "encoding/xml"

type Node struct {
	XMLName xml.Name
	Attr    []xml.Attr `xml:",any,attr"`
	Content []byte     `xml:",innerxml"`
	Nodes   []Node     `xml:",any"`
}

func (n *Node) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type inNode Node

	err := d.DecodeElement((*inNode)(n), &start)
	if err != nil {
		return err
	}

	if len(n.Nodes) > 0 {
		n.Content = nil
	}
	return nil
}
