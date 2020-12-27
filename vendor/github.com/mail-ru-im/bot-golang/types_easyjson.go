// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package botgolang

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6601e8cdDecodeGithubComMailRuImBotGolang(in *jlexer.Lexer, out *eventsResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ok":
			out.OK = bool(in.Bool())
		case "events":
			if in.IsNull() {
				in.Skip()
				out.Events = nil
			} else {
				in.Delim('[')
				if out.Events == nil {
					if !in.IsDelim(']') {
						out.Events = make([]*Event, 0, 8)
					} else {
						out.Events = []*Event{}
					}
				} else {
					out.Events = (out.Events)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Event
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Event)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Events = append(out.Events, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComMailRuImBotGolang(out *jwriter.Writer, in eventsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ok\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.OK))
	}
	{
		const prefix string = ",\"events\":"
		out.RawString(prefix)
		if in.Events == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Events {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v eventsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v eventsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *eventsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *eventsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang(l, v)
}
func easyjson6601e8cdDecodeGithubComMailRuImBotGolang1(in *jlexer.Lexer, out *UsersListResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "users":
			if in.IsNull() {
				in.Skip()
				out.List = nil
			} else {
				in.Delim('[')
				if out.List == nil {
					if !in.IsDelim(']') {
						out.List = make([]User, 0, 4)
					} else {
						out.List = []User{}
					}
				} else {
					out.List = (out.List)[:0]
				}
				for !in.IsDelim(']') {
					var v4 User
					(v4).UnmarshalEasyJSON(in)
					out.List = append(out.List, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComMailRuImBotGolang1(out *jwriter.Writer, in UsersListResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"users\":"
		out.RawString(prefix[1:])
		if in.List == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.List {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UsersListResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UsersListResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UsersListResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UsersListResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang1(l, v)
}
func easyjson6601e8cdDecodeGithubComMailRuImBotGolang2(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "userId":
			out.ID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComMailRuImBotGolang2(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"userId\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang2(l, v)
}
func easyjson6601e8cdDecodeGithubComMailRuImBotGolang3(in *jlexer.Lexer, out *Response) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ok":
			out.OK = bool(in.Bool())
		case "description":
			out.Description = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComMailRuImBotGolang3(out *jwriter.Writer, in Response) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ok\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.OK))
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang3(l, v)
}
func easyjson6601e8cdDecodeGithubComMailRuImBotGolang4(in *jlexer.Lexer, out *Photo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "url":
			out.URL = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComMailRuImBotGolang4(out *jwriter.Writer, in Photo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix[1:])
		out.String(string(in.URL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Photo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Photo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Photo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Photo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang4(l, v)
}
func easyjson6601e8cdDecodeGithubComMailRuImBotGolang5(in *jlexer.Lexer, out *PartPayload) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "firstName":
			out.FirstName = string(in.String())
		case "lastName":
			out.LastName = string(in.String())
		case "userId":
			out.UserID = string(in.String())
		case "fileId":
			out.FileID = string(in.String())
		case "caption":
			out.Caption = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "message":
			(out.PartMessage).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComMailRuImBotGolang5(out *jwriter.Writer, in PartPayload) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"firstName\":"
		out.RawString(prefix[1:])
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"lastName\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"userId\":"
		out.RawString(prefix)
		out.String(string(in.UserID))
	}
	{
		const prefix string = ",\"fileId\":"
		out.RawString(prefix)
		out.String(string(in.FileID))
	}
	{
		const prefix string = ",\"caption\":"
		out.RawString(prefix)
		out.String(string(in.Caption))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		(in.PartMessage).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PartPayload) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PartPayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PartPayload) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PartPayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang5(l, v)
}
func easyjson6601e8cdDecodeGithubComMailRuImBotGolang6(in *jlexer.Lexer, out *PartMessage) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "from":
			(out.From).UnmarshalEasyJSON(in)
		case "msgId":
			out.MsgID = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "timestamp":
			out.Timestamp = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComMailRuImBotGolang6(out *jwriter.Writer, in PartMessage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"from\":"
		out.RawString(prefix[1:])
		(in.From).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"msgId\":"
		out.RawString(prefix)
		out.String(string(in.MsgID))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"timestamp\":"
		out.RawString(prefix)
		out.Int(int(in.Timestamp))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PartMessage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PartMessage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PartMessage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PartMessage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang6(l, v)
}
func easyjson6601e8cdDecodeGithubComMailRuImBotGolang7(in *jlexer.Lexer, out *Part) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			out.Type = PartType(in.String())
		case "payload":
			(out.Payload).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComMailRuImBotGolang7(out *jwriter.Writer, in Part) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix[1:])
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"payload\":"
		out.RawString(prefix)
		(in.Payload).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Part) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Part) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Part) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Part) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang7(l, v)
}
func easyjson6601e8cdDecodeGithubComMailRuImBotGolang8(in *jlexer.Lexer, out *MembersListResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "members":
			if in.IsNull() {
				in.Skip()
				out.List = nil
			} else {
				in.Delim('[')
				if out.List == nil {
					if !in.IsDelim(']') {
						out.List = make([]ChatMember, 0, 2)
					} else {
						out.List = []ChatMember{}
					}
				} else {
					out.List = (out.List)[:0]
				}
				for !in.IsDelim(']') {
					var v7 ChatMember
					(v7).UnmarshalEasyJSON(in)
					out.List = append(out.List, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComMailRuImBotGolang8(out *jwriter.Writer, in MembersListResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"members\":"
		out.RawString(prefix[1:])
		if in.List == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.List {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MembersListResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MembersListResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MembersListResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MembersListResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang8(l, v)
}
func easyjson6601e8cdDecodeGithubComMailRuImBotGolang9(in *jlexer.Lexer, out *EventPayload) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "msgId":
			out.MsgID = string(in.String())
		case "chat":
			(out.Chat).UnmarshalEasyJSON(in)
		case "from":
			(out.From).UnmarshalEasyJSON(in)
		case "text":
			out.Text = string(in.String())
		case "parts":
			if in.IsNull() {
				in.Skip()
				out.Parts = nil
			} else {
				in.Delim('[')
				if out.Parts == nil {
					if !in.IsDelim(']') {
						out.Parts = make([]Part, 0, 1)
					} else {
						out.Parts = []Part{}
					}
				} else {
					out.Parts = (out.Parts)[:0]
				}
				for !in.IsDelim(']') {
					var v10 Part
					(v10).UnmarshalEasyJSON(in)
					out.Parts = append(out.Parts, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "timestamp":
			out.Timestamp = int(in.Int())
		case "queryId":
			out.QueryID = string(in.String())
		case "callbackData":
			out.CallbackData = string(in.String())
		case "leftMembers":
			if in.IsNull() {
				in.Skip()
				out.LeftMembers = nil
			} else {
				in.Delim('[')
				if out.LeftMembers == nil {
					if !in.IsDelim(']') {
						out.LeftMembers = make([]Contact, 0, 1)
					} else {
						out.LeftMembers = []Contact{}
					}
				} else {
					out.LeftMembers = (out.LeftMembers)[:0]
				}
				for !in.IsDelim(']') {
					var v11 Contact
					(v11).UnmarshalEasyJSON(in)
					out.LeftMembers = append(out.LeftMembers, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "newMembers":
			if in.IsNull() {
				in.Skip()
				out.NewMembers = nil
			} else {
				in.Delim('[')
				if out.NewMembers == nil {
					if !in.IsDelim(']') {
						out.NewMembers = make([]Contact, 0, 1)
					} else {
						out.NewMembers = []Contact{}
					}
				} else {
					out.NewMembers = (out.NewMembers)[:0]
				}
				for !in.IsDelim(']') {
					var v12 Contact
					(v12).UnmarshalEasyJSON(in)
					out.NewMembers = append(out.NewMembers, v12)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "addedBy":
			(out.AddedBy).UnmarshalEasyJSON(in)
		case "removedBy":
			(out.RemovedBy).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComMailRuImBotGolang9(out *jwriter.Writer, in EventPayload) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"msgId\":"
		out.RawString(prefix[1:])
		out.String(string(in.MsgID))
	}
	{
		const prefix string = ",\"chat\":"
		out.RawString(prefix)
		(in.Chat).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"from\":"
		out.RawString(prefix)
		(in.From).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"parts\":"
		out.RawString(prefix)
		if in.Parts == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v13, v14 := range in.Parts {
				if v13 > 0 {
					out.RawByte(',')
				}
				(v14).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"timestamp\":"
		out.RawString(prefix)
		out.Int(int(in.Timestamp))
	}
	{
		const prefix string = ",\"queryId\":"
		out.RawString(prefix)
		out.String(string(in.QueryID))
	}
	{
		const prefix string = ",\"callbackData\":"
		out.RawString(prefix)
		out.String(string(in.CallbackData))
	}
	{
		const prefix string = ",\"leftMembers\":"
		out.RawString(prefix)
		if in.LeftMembers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v15, v16 := range in.LeftMembers {
				if v15 > 0 {
					out.RawByte(',')
				}
				(v16).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"newMembers\":"
		out.RawString(prefix)
		if in.NewMembers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v17, v18 := range in.NewMembers {
				if v17 > 0 {
					out.RawByte(',')
				}
				(v18).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"addedBy\":"
		out.RawString(prefix)
		(in.AddedBy).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"removedBy\":"
		out.RawString(prefix)
		(in.RemovedBy).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventPayload) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventPayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventPayload) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventPayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang9(l, v)
}
func easyjson6601e8cdDecodeGithubComMailRuImBotGolang10(in *jlexer.Lexer, out *Event) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "eventId":
			out.EventID = int(in.Int())
		case "type":
			out.Type = EventType(in.String())
		case "payload":
			(out.Payload).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComMailRuImBotGolang10(out *jwriter.Writer, in Event) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"eventId\":"
		out.RawString(prefix[1:])
		out.Int(int(in.EventID))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"payload\":"
		out.RawString(prefix)
		(in.Payload).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Event) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Event) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Event) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Event) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang10(l, v)
}
func easyjson6601e8cdDecodeGithubComMailRuImBotGolang11(in *jlexer.Lexer, out *Contact) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "firstName":
			out.FirstName = string(in.String())
		case "lastName":
			out.LastName = string(in.String())
		case "userId":
			out.ID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComMailRuImBotGolang11(out *jwriter.Writer, in Contact) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"firstName\":"
		out.RawString(prefix[1:])
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"lastName\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"userId\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Contact) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Contact) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Contact) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Contact) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang11(l, v)
}
func easyjson6601e8cdDecodeGithubComMailRuImBotGolang12(in *jlexer.Lexer, out *ChatMember) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "creator":
			out.Creator = bool(in.Bool())
		case "admin":
			out.Admin = bool(in.Bool())
		case "userId":
			out.ID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComMailRuImBotGolang12(out *jwriter.Writer, in ChatMember) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"creator\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Creator))
	}
	{
		const prefix string = ",\"admin\":"
		out.RawString(prefix)
		out.Bool(bool(in.Admin))
	}
	{
		const prefix string = ",\"userId\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ChatMember) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ChatMember) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ChatMember) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ChatMember) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang12(l, v)
}
func easyjson6601e8cdDecodeGithubComMailRuImBotGolang13(in *jlexer.Lexer, out *BotInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "nick":
			out.Nick = string(in.String())
		case "firstName":
			out.FirstName = string(in.String())
		case "about":
			out.About = string(in.String())
		case "photo":
			if in.IsNull() {
				in.Skip()
				out.Photo = nil
			} else {
				in.Delim('[')
				if out.Photo == nil {
					if !in.IsDelim(']') {
						out.Photo = make([]Photo, 0, 4)
					} else {
						out.Photo = []Photo{}
					}
				} else {
					out.Photo = (out.Photo)[:0]
				}
				for !in.IsDelim(']') {
					var v19 Photo
					(v19).UnmarshalEasyJSON(in)
					out.Photo = append(out.Photo, v19)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "userId":
			out.ID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComMailRuImBotGolang13(out *jwriter.Writer, in BotInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nick\":"
		out.RawString(prefix[1:])
		out.String(string(in.Nick))
	}
	{
		const prefix string = ",\"firstName\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"about\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	{
		const prefix string = ",\"photo\":"
		out.RawString(prefix)
		if in.Photo == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.Photo {
				if v20 > 0 {
					out.RawByte(',')
				}
				(v21).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"userId\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BotInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BotInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BotInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BotInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang13(l, v)
}
func easyjson6601e8cdDecodeGithubComMailRuImBotGolang14(in *jlexer.Lexer, out *AdminsListResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "admins":
			if in.IsNull() {
				in.Skip()
				out.List = nil
			} else {
				in.Delim('[')
				if out.List == nil {
					if !in.IsDelim(']') {
						out.List = make([]ChatMember, 0, 2)
					} else {
						out.List = []ChatMember{}
					}
				} else {
					out.List = (out.List)[:0]
				}
				for !in.IsDelim(']') {
					var v22 ChatMember
					(v22).UnmarshalEasyJSON(in)
					out.List = append(out.List, v22)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComMailRuImBotGolang14(out *jwriter.Writer, in AdminsListResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"admins\":"
		out.RawString(prefix[1:])
		if in.List == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v23, v24 := range in.List {
				if v23 > 0 {
					out.RawByte(',')
				}
				(v24).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AdminsListResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang14(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AdminsListResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComMailRuImBotGolang14(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AdminsListResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang14(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AdminsListResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComMailRuImBotGolang14(l, v)
}
