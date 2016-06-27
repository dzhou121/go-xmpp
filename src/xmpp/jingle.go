package xmpp

import "encoding/xml"

type Jingle struct {
	XMLName   xml.Name `xml:"urn:xmpp:jingle:1 jingle"`
	Action    string   `xml:"action,attr"`
	Initiator string   `xml:"initiator,attr,omitempty"`
	Responder string   `xml:"responder,attr,omitempty"`
	Sid       string   `xml:"sid,attr"`

	Contents []*JingleContent `xml:"content"`
	Ringing  *JingleRinging
}

func (_ Jingle) Name() string {
	return "jingle"
}

func (_ Jingle) FullName() string {
	return "urn:xmpp:jingle:1 jingle"
}

func (j Jingle) String() string {
	return ""
}

type JingleContent struct {
	Name         string              `xml:"name,attr"`
	Creator      string              `xml:"creator,attr"`
	Description  *JingleDesc         `xml:"urn:xmpp:jingle:apps:rtp:1 description"`
	Transport    *JingleTransport    `xml:"urn:xmpp:jingle:transports:ice-udp:1 transport"`
	RawTransport *JingleRawTransport `xml:"urn:xmpp:jingle:transports:raw-udp:1 transport"`
}

type RtcpMux struct {
}

type JingleSourceParam struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type JingleSource struct {
	Ssrc   string               `xml:"ssrc,attr"`
	Params []*JingleSourceParam `xml:"parameter"`
}

type RtpHdrext struct {
	Id  string `xml:"id,attr"`
	Uri string `xml:"uri,attr"`
}

type JingleDesc struct {
	Media      string          `xml:"media,attr"`
	Payloads   []*PayloadType  `xml:"payload-type"`
	Sources    []*JingleSource `xml:"urn:xmpp:jingle:apps:rtp:ssma:0 source"`
	RtpHdrexts []*RtpHdrext    `xml:"urn:xmpp:jingle:apps:rtp:rtp-hdrext:0"`
}

type JingleParameter struct {
	Value string `xml:"value,attr"`
	Name  string `xml:"name,attr"`
}

type PayloadType struct {
	Id        string             `xml:"id,attr"`
	Name      string             `xml:"name,attr"`
	ClockRate string             `xml:"clockrate,attr"`
	Channels  string             `xml:"channels,attr,omitempty"`
	Parameter []*JingleParameter `xml:"parameter"`
}

type JingleFingerprint struct {
	Hash  string `xml:"hash,attr"`
	Setup string `xml:"setup,attr"`
	Text  string `xml:",chardata"`
}

type JingleRawTransport struct {
	Pwd         string             `xml:"pwd,attr,omitempty"`
	Ufrag       string             `xml:"ufrag,attr,omitempty"`
	Fingerprint *JingleFingerprint `xml:"urn:xmpp:jingle:apps:dtls:0 fingerprint,omitempty"`
	Candidates  []*Candidate       `xml:"candidate"`
}

type JingleTransport struct {
	Pwd         string             `xml:"pwd,attr,omitempty"`
	Ufrag       string             `xml:"ufrag,attr,omitempty"`
	Fingerprint *JingleFingerprint `xml:"urn:xmpp:jingle:apps:dtls:0 fingerprint,omitempty"`
	Candidates  []*Candidate       `xml:"candidate"`
}

type Candidate struct {
	Id         string `xml:"id,attr"`
	Component  string `xml:"component,attr"`
	Foundation string `xml:"foundation,attr,omitempty"`
	Generation string `xml:"generation,attr"`
	Ip         string `xml:"ip,attr"`
	Network    string `xml:"network,attr,omitempty"`
	Port       string `xml:"port,attr"`
	Priority   string `xml:"priority,attr,omitempty"`
	Protocol   string `xml:"protocol,attr,omitempty"`
	RelAddr    string `xml:"rel-addr,attr,omitempty"`
	RelPort    string `xml:"rel-port,attr,omitempty"`
	Type       string `xml:"type,attr"`
}

type JingleRinging struct {
	XMLName xml.Name `xml:"urn:xmpp:jingle:apps:rtp:info:1 ringing"`
}
