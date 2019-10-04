package message

func NewLDAPMessageWithProtocolOp(po ProtocolOp) *LDAPMessage {
	m := NewLDAPMessage()
	m.protocolOp = po
	return m
}

func NewFilterEqualityMatch(desc, val string) FilterEqualityMatch {
	return FilterEqualityMatch{
		attributeDesc:  AttributeDescription(desc),
		assertionValue: AssertionValue(val),
	}
}
