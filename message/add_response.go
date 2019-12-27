package message

import "fmt"

//
//        AddResponse ::= [APPLICATION 9] LDAPResult

func (l *AddResponse) SetResultCode(code int) {
	l.resultCode = ENUMERATED(code)
}

func (l *AddResponse) SetMatchedDN(dn string) {
	l.matchedDN = LDAPDN(dn)
}

func (l *AddResponse) SetDiagnosticMessage(msg string) {
	l.diagnosticMessage = LDAPString(msg)
}

func readAddResponse(bytes *Bytes) (ret AddResponse, err error) {
	var res LDAPResult
	res, err = readTaggedLDAPResult(bytes, classApplication, TagAddResponse)
	if err != nil {
		err = LdapError{fmt.Sprintf("readAddResponse:\n%s", err.Error())}
		return
	}
	ret = AddResponse(res)
	return
}

func (a AddResponse) size() int {
	return LDAPResult(a).sizeTagged(TagAddResponse)
}
func (a AddResponse) write(bytes *Bytes) int {
	return LDAPResult(a).writeTagged(bytes, classApplication, TagAddResponse)
}
