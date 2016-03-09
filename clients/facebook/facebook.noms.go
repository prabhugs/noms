// This file was generated by nomdl/codegen.

package main

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __mainPackageInFile_facebook_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.Type{
		types.MakeStructType("User",
			[]types.Field{
				types.Field{"Id", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"Name", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"Photos", types.MakeCompoundType(types.SetKind, types.MakeCompoundType(types.RefKind, types.MakeType(ref.Parse("sha1-10004087fdbc623873c649d28aa59f4e066d374e"), 0))), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{
		ref.Parse("sha1-10004087fdbc623873c649d28aa59f4e066d374e"),
	})
	__mainPackageInFile_facebook_CachedRef = types.RegisterPackage(&p)
}

// User

type User struct {
	_Id     string
	_Name   string
	_Photos SetOfRefOfRemotePhoto

	ref *ref.Ref
}

func NewUser() User {
	return User{
		_Id:     "",
		_Name:   "",
		_Photos: NewSetOfRefOfRemotePhoto(),

		ref: &ref.Ref{},
	}
}

type UserDef struct {
	Id     string
	Name   string
	Photos SetOfRefOfRemotePhotoDef
}

func (def UserDef) New() User {
	return User{
		_Id:     def.Id,
		_Name:   def.Name,
		_Photos: def.Photos.New(),
		ref:     &ref.Ref{},
	}
}

func (s User) Def() (d UserDef) {
	d.Id = s._Id
	d.Name = s._Name
	d.Photos = s._Photos.Def()
	return
}

var __typeForUser types.Type

func (m User) Type() types.Type {
	return __typeForUser
}

func init() {
	__typeForUser = types.MakeType(__mainPackageInFile_facebook_CachedRef, 0)
	types.RegisterStruct(__typeForUser, builderForUser, readerForUser)
}

func builderForUser(values []types.Value) types.Value {
	i := 0
	s := User{ref: &ref.Ref{}}
	s._Id = values[i].(types.String).String()
	i++
	s._Name = values[i].(types.String).String()
	i++
	s._Photos = values[i].(SetOfRefOfRemotePhoto)
	i++
	return s
}

func readerForUser(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(User)
	values = append(values, types.NewString(s._Id))
	values = append(values, types.NewString(s._Name))
	values = append(values, s._Photos)
	return values
}

func (s User) Equals(other types.Value) bool {
	return other != nil && __typeForUser.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s User) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s User) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeForUser.Chunks()...)
	chunks = append(chunks, s._Photos.Chunks()...)
	return
}

func (s User) ChildValues() (ret []types.Value) {
	ret = append(ret, types.NewString(s._Id))
	ret = append(ret, types.NewString(s._Name))
	ret = append(ret, s._Photos)
	return
}

func (s User) Id() string {
	return s._Id
}

func (s User) SetId(val string) User {
	s._Id = val
	s.ref = &ref.Ref{}
	return s
}

func (s User) Name() string {
	return s._Name
}

func (s User) SetName(val string) User {
	s._Name = val
	s.ref = &ref.Ref{}
	return s
}

func (s User) Photos() SetOfRefOfRemotePhoto {
	return s._Photos
}

func (s User) SetPhotos(val SetOfRefOfRemotePhoto) User {
	s._Photos = val
	s.ref = &ref.Ref{}
	return s
}

// RefOfUser

type RefOfUser struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfUser(target ref.Ref) RefOfUser {
	return RefOfUser{target, &ref.Ref{}}
}

func (r RefOfUser) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfUser) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfUser) Equals(other types.Value) bool {
	return other != nil && __typeForRefOfUser.Equals(other.Type()) && r.Ref() == other.Ref()
}

func (r RefOfUser) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, r.Type().Chunks()...)
	chunks = append(chunks, r.target)
	return
}

func (r RefOfUser) ChildValues() []types.Value {
	return nil
}

// A Noms Value that describes RefOfUser.
var __typeForRefOfUser types.Type

func (r RefOfUser) Type() types.Type {
	return __typeForRefOfUser
}

func (r RefOfUser) Less(other types.OrderedValue) bool {
	return r.TargetRef().Less(other.(types.RefBase).TargetRef())
}

func init() {
	__typeForRefOfUser = types.MakeCompoundType(types.RefKind, types.MakeType(__mainPackageInFile_facebook_CachedRef, 0))
	types.RegisterRef(__typeForRefOfUser, builderForRefOfUser)
}

func builderForRefOfUser(r ref.Ref) types.Value {
	return NewRefOfUser(r)
}

func (r RefOfUser) TargetValue(vr types.ValueReader) User {
	return vr.ReadValue(r.target).(User)
}

func (r RefOfUser) SetTargetValue(val User, vw types.ValueWriter) RefOfUser {
	return NewRefOfUser(vw.WriteValue(val))
}

// SetOfRefOfRemotePhoto

type SetOfRefOfRemotePhoto struct {
	s   types.Set
	ref *ref.Ref
}

func NewSetOfRefOfRemotePhoto() SetOfRefOfRemotePhoto {
	return SetOfRefOfRemotePhoto{types.NewTypedSet(__typeForSetOfRefOfRemotePhoto), &ref.Ref{}}
}

type SetOfRefOfRemotePhotoDef map[ref.Ref]bool

func (def SetOfRefOfRemotePhotoDef) New() SetOfRefOfRemotePhoto {
	l := make([]types.Value, len(def))
	i := 0
	for d, _ := range def {
		l[i] = NewRefOfRemotePhoto(d)
		i++
	}
	return SetOfRefOfRemotePhoto{types.NewTypedSet(__typeForSetOfRefOfRemotePhoto, l...), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Def() SetOfRefOfRemotePhotoDef {
	def := make(map[ref.Ref]bool, s.Len())
	s.s.Iter(func(v types.Value) bool {
		def[v.(RefOfRemotePhoto).TargetRef()] = true
		return false
	})
	return def
}

func (s SetOfRefOfRemotePhoto) Equals(other types.Value) bool {
	return other != nil && __typeForSetOfRefOfRemotePhoto.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s SetOfRefOfRemotePhoto) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s SetOfRefOfRemotePhoto) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, s.Type().Chunks()...)
	chunks = append(chunks, s.s.Chunks()...)
	return
}

func (s SetOfRefOfRemotePhoto) ChildValues() []types.Value {
	return append([]types.Value{}, s.s.ChildValues()...)
}

// A Noms Value that describes SetOfRefOfRemotePhoto.
var __typeForSetOfRefOfRemotePhoto types.Type

func (m SetOfRefOfRemotePhoto) Type() types.Type {
	return __typeForSetOfRefOfRemotePhoto
}

func init() {
	__typeForSetOfRefOfRemotePhoto = types.MakeCompoundType(types.SetKind, types.MakeCompoundType(types.RefKind, types.MakeType(ref.Parse("sha1-10004087fdbc623873c649d28aa59f4e066d374e"), 0)))
	types.RegisterValue(__typeForSetOfRefOfRemotePhoto, builderForSetOfRefOfRemotePhoto, readerForSetOfRefOfRemotePhoto)
}

func builderForSetOfRefOfRemotePhoto(v types.Value) types.Value {
	return SetOfRefOfRemotePhoto{v.(types.Set), &ref.Ref{}}
}

func readerForSetOfRefOfRemotePhoto(v types.Value) types.Value {
	return v.(SetOfRefOfRemotePhoto).s
}

func (s SetOfRefOfRemotePhoto) Empty() bool {
	return s.s.Empty()
}

func (s SetOfRefOfRemotePhoto) Len() uint64 {
	return s.s.Len()
}

func (s SetOfRefOfRemotePhoto) Has(p RefOfRemotePhoto) bool {
	return s.s.Has(p)
}

type SetOfRefOfRemotePhotoIterCallback func(p RefOfRemotePhoto) (stop bool)

func (s SetOfRefOfRemotePhoto) Iter(cb SetOfRefOfRemotePhotoIterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(v.(RefOfRemotePhoto))
	})
}

type SetOfRefOfRemotePhotoIterAllCallback func(p RefOfRemotePhoto)

func (s SetOfRefOfRemotePhoto) IterAll(cb SetOfRefOfRemotePhotoIterAllCallback) {
	s.s.IterAll(func(v types.Value) {
		cb(v.(RefOfRemotePhoto))
	})
}

func (s SetOfRefOfRemotePhoto) IterAllP(concurrency int, cb SetOfRefOfRemotePhotoIterAllCallback) {
	s.s.IterAllP(concurrency, func(v types.Value) {
		cb(v.(RefOfRemotePhoto))
	})
}

type SetOfRefOfRemotePhotoFilterCallback func(p RefOfRemotePhoto) (keep bool)

func (s SetOfRefOfRemotePhoto) Filter(cb SetOfRefOfRemotePhotoFilterCallback) SetOfRefOfRemotePhoto {
	out := s.s.Filter(func(v types.Value) bool {
		return cb(v.(RefOfRemotePhoto))
	})
	return SetOfRefOfRemotePhoto{out, &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Insert(p ...RefOfRemotePhoto) SetOfRefOfRemotePhoto {
	return SetOfRefOfRemotePhoto{s.s.Insert(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Remove(p ...RefOfRemotePhoto) SetOfRefOfRemotePhoto {
	return SetOfRefOfRemotePhoto{s.s.Remove(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Union(others ...SetOfRefOfRemotePhoto) SetOfRefOfRemotePhoto {
	return SetOfRefOfRemotePhoto{s.s.Union(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) First() RefOfRemotePhoto {
	return s.s.First().(RefOfRemotePhoto)
}

func (s SetOfRefOfRemotePhoto) fromStructSlice(p []SetOfRefOfRemotePhoto) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfRefOfRemotePhoto) fromElemSlice(p []RefOfRemotePhoto) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

// RefOfRemotePhoto

type RefOfRemotePhoto struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfRemotePhoto(target ref.Ref) RefOfRemotePhoto {
	return RefOfRemotePhoto{target, &ref.Ref{}}
}

func (r RefOfRemotePhoto) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfRemotePhoto) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfRemotePhoto) Equals(other types.Value) bool {
	return other != nil && __typeForRefOfRemotePhoto.Equals(other.Type()) && r.Ref() == other.Ref()
}

func (r RefOfRemotePhoto) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, r.Type().Chunks()...)
	chunks = append(chunks, r.target)
	return
}

func (r RefOfRemotePhoto) ChildValues() []types.Value {
	return nil
}

// A Noms Value that describes RefOfRemotePhoto.
var __typeForRefOfRemotePhoto types.Type

func (r RefOfRemotePhoto) Type() types.Type {
	return __typeForRefOfRemotePhoto
}

func (r RefOfRemotePhoto) Less(other types.OrderedValue) bool {
	return r.TargetRef().Less(other.(types.RefBase).TargetRef())
}

func init() {
	__typeForRefOfRemotePhoto = types.MakeCompoundType(types.RefKind, types.MakeType(ref.Parse("sha1-10004087fdbc623873c649d28aa59f4e066d374e"), 0))
	types.RegisterRef(__typeForRefOfRemotePhoto, builderForRefOfRemotePhoto)
}

func builderForRefOfRemotePhoto(r ref.Ref) types.Value {
	return NewRefOfRemotePhoto(r)
}

func (r RefOfRemotePhoto) TargetValue(vr types.ValueReader) RemotePhoto {
	return vr.ReadValue(r.target).(RemotePhoto)
}

func (r RefOfRemotePhoto) SetTargetValue(val RemotePhoto, vw types.ValueWriter) RefOfRemotePhoto {
	return NewRefOfRemotePhoto(vw.WriteValue(val))
}
