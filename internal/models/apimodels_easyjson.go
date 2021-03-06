// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels(in *jlexer.Lexer, out *Vote) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "nickname":
			out.Nickname = string(in.String())
		case "voice":
			out.Voice = int32(in.Int32())
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
func easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels(out *jwriter.Writer, in Vote) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nickname\":"
		out.RawString(prefix[1:])
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"voice\":"
		out.RawString(prefix)
		out.Int32(int32(in.Voice))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Vote) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Vote) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Vote) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Vote) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels(l, v)
}
func easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels1(in *jlexer.Lexer, out *Users) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Users, 0, 1)
			} else {
				*out = Users{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 User
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels1(out *jwriter.Writer, in Users) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Users) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Users) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Users) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Users) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels1(l, v)
}
func easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels2(in *jlexer.Lexer, out *UserUpdate) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "about":
			out.About = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "fullname":
			out.Fullname = string(in.String())
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
func easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels2(out *jwriter.Writer, in UserUpdate) {
	out.RawByte('{')
	first := true
	_ = first
	if in.About != "" {
		const prefix string = ",\"about\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.About))
	}
	if in.Email != "" {
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	if in.Fullname != "" {
		const prefix string = ",\"fullname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Fullname))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserUpdate) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserUpdate) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserUpdate) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserUpdate) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels2(l, v)
}
func easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels3(in *jlexer.Lexer, out *User) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "nickname":
			out.Nickname = string(in.String())
		case "about":
			out.About = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "fullname":
			out.Fullname = string(in.String())
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
func easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels3(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nickname\":"
		out.RawString(prefix[1:])
		out.String(string(in.Nickname))
	}
	if in.About != "" {
		const prefix string = ",\"about\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	if in.Email != "" {
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	if in.Fullname != "" {
		const prefix string = ",\"fullname\":"
		out.RawString(prefix)
		out.String(string(in.Fullname))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels3(l, v)
}
func easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels4(in *jlexer.Lexer, out *Threads) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Threads, 0, 0)
			} else {
				*out = Threads{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v4 Thread
			(v4).UnmarshalEasyJSON(in)
			*out = append(*out, v4)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels4(out *jwriter.Writer, in Threads) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in {
			if v5 > 0 {
				out.RawByte(',')
			}
			(v6).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Threads) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Threads) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Threads) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Threads) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels4(l, v)
}
func easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels5(in *jlexer.Lexer, out *ThreadUpdate) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "message":
			out.Message = string(in.String())
		case "title":
			out.Title = string(in.String())
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
func easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels5(out *jwriter.Writer, in ThreadUpdate) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Message != "" {
		const prefix string = ",\"message\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Message))
	}
	if in.Title != "" {
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ThreadUpdate) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ThreadUpdate) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ThreadUpdate) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ThreadUpdate) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels5(l, v)
}
func easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels6(in *jlexer.Lexer, out *Thread) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int64(in.Int64())
		case "author":
			out.Author = string(in.String())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
		case "forum":
			out.Forum = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "slug":
			out.Slug = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "votes":
			out.Votes = int(in.Int())
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
func easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels6(out *jwriter.Writer, in Thread) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Id))
	}
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		out.String(string(in.Author))
	}
	if true {
		const prefix string = ",\"created\":"
		out.RawString(prefix)
		out.Raw((in.Created).MarshalJSON())
	}
	if in.Forum != "" {
		const prefix string = ",\"forum\":"
		out.RawString(prefix)
		out.String(string(in.Forum))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	if in.Slug != "" {
		const prefix string = ",\"slug\":"
		out.RawString(prefix)
		out.String(string(in.Slug))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	if in.Votes != 0 {
		const prefix string = ",\"votes\":"
		out.RawString(prefix)
		out.Int(int(in.Votes))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Thread) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Thread) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Thread) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Thread) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels6(l, v)
}
func easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels7(in *jlexer.Lexer, out *Status) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "forum":
			out.Forum = int(in.Int())
		case "post":
			out.Post = int(in.Int())
		case "thread":
			out.Thread = int(in.Int())
		case "user":
			out.User = int(in.Int())
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
func easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels7(out *jwriter.Writer, in Status) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"forum\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Forum))
	}
	{
		const prefix string = ",\"post\":"
		out.RawString(prefix)
		out.Int(int(in.Post))
	}
	{
		const prefix string = ",\"thread\":"
		out.RawString(prefix)
		out.Int(int(in.Thread))
	}
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		out.Int(int(in.User))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Status) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Status) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Status) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Status) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels7(l, v)
}
func easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels8(in *jlexer.Lexer, out *Posts) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Posts, 0, 0)
			} else {
				*out = Posts{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v7 Post
			(v7).UnmarshalEasyJSON(in)
			*out = append(*out, v7)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels8(out *jwriter.Writer, in Posts) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v8, v9 := range in {
			if v8 > 0 {
				out.RawByte(',')
			}
			(v9).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Posts) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Posts) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Posts) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Posts) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels8(l, v)
}
func easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels9(in *jlexer.Lexer, out *PostUpdate) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "message":
			out.Message = string(in.String())
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
func easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels9(out *jwriter.Writer, in PostUpdate) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Message != "" {
		const prefix string = ",\"message\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostUpdate) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostUpdate) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostUpdate) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostUpdate) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels9(l, v)
}
func easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels10(in *jlexer.Lexer, out *PostDetails) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "author":
			if in.IsNull() {
				in.Skip()
				out.Author = nil
			} else {
				if out.Author == nil {
					out.Author = new(User)
				}
				(*out.Author).UnmarshalEasyJSON(in)
			}
		case "forum":
			if in.IsNull() {
				in.Skip()
				out.Forum = nil
			} else {
				if out.Forum == nil {
					out.Forum = new(Forum)
				}
				(*out.Forum).UnmarshalEasyJSON(in)
			}
		case "post":
			if in.IsNull() {
				in.Skip()
				out.Post = nil
			} else {
				if out.Post == nil {
					out.Post = new(Post)
				}
				(*out.Post).UnmarshalEasyJSON(in)
			}
		case "thread":
			if in.IsNull() {
				in.Skip()
				out.Thread = nil
			} else {
				if out.Thread == nil {
					out.Thread = new(Thread)
				}
				(*out.Thread).UnmarshalEasyJSON(in)
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
func easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels10(out *jwriter.Writer, in PostDetails) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Author != nil {
		const prefix string = ",\"author\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Author).MarshalEasyJSON(out)
	}
	if in.Forum != nil {
		const prefix string = ",\"forum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Forum).MarshalEasyJSON(out)
	}
	if in.Post != nil {
		const prefix string = ",\"post\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Post).MarshalEasyJSON(out)
	}
	if in.Thread != nil {
		const prefix string = ",\"thread\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Thread).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostDetails) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostDetails) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostDetails) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostDetails) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels10(l, v)
}
func easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels11(in *jlexer.Lexer, out *Post) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int(in.Int())
		case "author":
			out.Author = string(in.String())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
		case "forum":
			out.Forum = string(in.String())
		case "isEdited":
			out.IsEdited = bool(in.Bool())
		case "message":
			out.Message = string(in.String())
		case "parent":
			out.Parent = int(in.Int())
		case "thread":
			out.Thread = int(in.Int())
		case "tree":
			if in.IsNull() {
				in.Skip()
				out.Tree = nil
			} else {
				in.Delim('[')
				if out.Tree == nil {
					if !in.IsDelim(']') {
						out.Tree = make([]int, 0, 8)
					} else {
						out.Tree = []int{}
					}
				} else {
					out.Tree = (out.Tree)[:0]
				}
				for !in.IsDelim(']') {
					var v10 int
					v10 = int(in.Int())
					out.Tree = append(out.Tree, v10)
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
func easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels11(out *jwriter.Writer, in Post) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		out.String(string(in.Author))
	}
	if true {
		const prefix string = ",\"created\":"
		out.RawString(prefix)
		out.Raw((in.Created).MarshalJSON())
	}
	if in.Forum != "" {
		const prefix string = ",\"forum\":"
		out.RawString(prefix)
		out.String(string(in.Forum))
	}
	if in.IsEdited {
		const prefix string = ",\"isEdited\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsEdited))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	if in.Parent != 0 {
		const prefix string = ",\"parent\":"
		out.RawString(prefix)
		out.Int(int(in.Parent))
	}
	if in.Thread != 0 {
		const prefix string = ",\"thread\":"
		out.RawString(prefix)
		out.Int(int(in.Thread))
	}
	if len(in.Tree) != 0 {
		const prefix string = ",\"tree\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v11, v12 := range in.Tree {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v12))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Post) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Post) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Post) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Post) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels11(l, v)
}
func easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels12(in *jlexer.Lexer, out *Parents) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ParentId":
			out.ParentId = int(in.Int())
		case "Tree":
			if in.IsNull() {
				in.Skip()
				out.Tree = nil
			} else {
				in.Delim('[')
				if out.Tree == nil {
					if !in.IsDelim(']') {
						out.Tree = make([]int, 0, 8)
					} else {
						out.Tree = []int{}
					}
				} else {
					out.Tree = (out.Tree)[:0]
				}
				for !in.IsDelim(']') {
					var v13 int
					v13 = int(in.Int())
					out.Tree = append(out.Tree, v13)
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
func easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels12(out *jwriter.Writer, in Parents) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ParentId\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ParentId))
	}
	{
		const prefix string = ",\"Tree\":"
		out.RawString(prefix)
		if in.Tree == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Tree {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v15))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Parents) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Parents) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Parents) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Parents) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels12(l, v)
}
func easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels13(in *jlexer.Lexer, out *Params) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Limit":
			out.Limit = int(in.Int())
		case "Since":
			out.Since = string(in.String())
		case "Desc":
			out.Desc = bool(in.Bool())
		case "Sort":
			out.Sort = string(in.String())
		case "Related":
			out.Related = string(in.String())
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
func easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels13(out *jwriter.Writer, in Params) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Limit\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Limit))
	}
	{
		const prefix string = ",\"Since\":"
		out.RawString(prefix)
		out.String(string(in.Since))
	}
	{
		const prefix string = ",\"Desc\":"
		out.RawString(prefix)
		out.Bool(bool(in.Desc))
	}
	{
		const prefix string = ",\"Sort\":"
		out.RawString(prefix)
		out.String(string(in.Sort))
	}
	{
		const prefix string = ",\"Related\":"
		out.RawString(prefix)
		out.String(string(in.Related))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Params) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Params) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Params) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Params) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels13(l, v)
}
func easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels14(in *jlexer.Lexer, out *Forum) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "posts":
			out.Posts = int(in.Int())
		case "slug":
			out.Slug = string(in.String())
		case "threads":
			out.Threads = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "user":
			out.User = string(in.String())
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
func easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels14(out *jwriter.Writer, in Forum) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Posts != 0 {
		const prefix string = ",\"posts\":"
		first = false
		out.RawString(prefix[1:])
		out.Int(int(in.Posts))
	}
	{
		const prefix string = ",\"slug\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Slug))
	}
	if in.Threads != 0 {
		const prefix string = ",\"threads\":"
		out.RawString(prefix)
		out.Int(int(in.Threads))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		out.String(string(in.User))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Forum) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels14(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Forum) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d796ea4EncodeGithubComShysaTPDBMSInternalModels14(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Forum) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels14(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Forum) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d796ea4DecodeGithubComShysaTPDBMSInternalModels14(l, v)
}
