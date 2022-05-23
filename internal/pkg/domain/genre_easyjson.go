// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package domain

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

func easyjson52fdb84bDecodeCodexInternalPkgDomain(in *jlexer.Lexer, out *GenreWithMovies) {
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
		case "href":
			out.Href = string(in.String())
		case "imgsrc":
			out.Imgsrc = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "movielist":
			if in.IsNull() {
				in.Skip()
				out.MovieList = nil
			} else {
				in.Delim('[')
				if out.MovieList == nil {
					if !in.IsDelim(']') {
						out.MovieList = make([]MovieBasic, 0, 0)
					} else {
						out.MovieList = []MovieBasic{}
					}
				} else {
					out.MovieList = (out.MovieList)[:0]
				}
				for !in.IsDelim(']') {
					var v1 MovieBasic
					easyjson52fdb84bDecodeCodexInternalPkgDomain1(in, &v1)
					out.MovieList = append(out.MovieList, v1)
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
func easyjson52fdb84bEncodeCodexInternalPkgDomain(out *jwriter.Writer, in GenreWithMovies) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"href\":"
		out.RawString(prefix[1:])
		out.String(string(in.Href))
	}
	{
		const prefix string = ",\"imgsrc\":"
		out.RawString(prefix)
		out.String(string(in.Imgsrc))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"movielist\":"
		out.RawString(prefix)
		if in.MovieList == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.MovieList {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjson52fdb84bEncodeCodexInternalPkgDomain1(out, v3)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GenreWithMovies) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson52fdb84bEncodeCodexInternalPkgDomain(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GenreWithMovies) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson52fdb84bDecodeCodexInternalPkgDomain(l, v)
}
func easyjson52fdb84bDecodeCodexInternalPkgDomain1(in *jlexer.Lexer, out *MovieBasic) {
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
		case "ID":
			out.Id = string(in.String())
		case "poster":
			out.Poster = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "rating":
			out.Rating = string(in.String())
		case "info":
			out.Info = string(in.String())
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
func easyjson52fdb84bEncodeCodexInternalPkgDomain1(out *jwriter.Writer, in MovieBasic) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"poster\":"
		out.RawString(prefix)
		out.String(string(in.Poster))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.String(string(in.Rating))
	}
	{
		const prefix string = ",\"info\":"
		out.RawString(prefix)
		out.String(string(in.Info))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	out.RawByte('}')
}
func easyjson52fdb84bDecodeCodexInternalPkgDomain2(in *jlexer.Lexer, out *GenreInMovie) {
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
		case "href":
			out.Href = string(in.String())
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
func easyjson52fdb84bEncodeCodexInternalPkgDomain2(out *jwriter.Writer, in GenreInMovie) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"href\":"
		out.RawString(prefix[1:])
		out.String(string(in.Href))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GenreInMovie) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson52fdb84bEncodeCodexInternalPkgDomain2(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GenreInMovie) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson52fdb84bDecodeCodexInternalPkgDomain2(l, v)
}
func easyjson52fdb84bDecodeCodexInternalPkgDomain3(in *jlexer.Lexer, out *Genre) {
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
		case "href":
			out.Href = string(in.String())
		case "imgsrc":
			out.Imgsrc = string(in.String())
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
func easyjson52fdb84bEncodeCodexInternalPkgDomain3(out *jwriter.Writer, in Genre) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"href\":"
		out.RawString(prefix[1:])
		out.String(string(in.Href))
	}
	{
		const prefix string = ",\"imgsrc\":"
		out.RawString(prefix)
		out.String(string(in.Imgsrc))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Genre) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson52fdb84bEncodeCodexInternalPkgDomain3(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Genre) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson52fdb84bDecodeCodexInternalPkgDomain3(l, v)
}