// Code generated by protoc-gen-go. DO NOT EDIT.
// source: build.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BuildStep struct {
	Argv                 []string `protobuf:"bytes,1,rep,name=argv" json:"argv,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildStep) Reset()         { *m = BuildStep{} }
func (m *BuildStep) String() string { return proto.CompactTextString(m) }
func (*BuildStep) ProtoMessage()    {}
func (*BuildStep) Descriptor() ([]byte, []int) {
	return fileDescriptor_14ce178a580e4ede, []int{0}
}

func (m *BuildStep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildStep.Unmarshal(m, b)
}
func (m *BuildStep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildStep.Marshal(b, m, deterministic)
}
func (m *BuildStep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildStep.Merge(m, src)
}
func (m *BuildStep) XXX_Size() int {
	return xxx_messageInfo_BuildStep.Size(m)
}
func (m *BuildStep) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildStep.DiscardUnknown(m)
}

var xxx_messageInfo_BuildStep proto.InternalMessageInfo

func (m *BuildStep) GetArgv() []string {
	if m != nil {
		return m.Argv
	}
	return nil
}

type CBuilder struct {
	ExtraConfigureFlag []string `protobuf:"bytes,1,rep,name=extra_configure_flag,json=extraConfigureFlag" json:"extra_configure_flag,omitempty"`
	ExtraLdflag        []string `protobuf:"bytes,3,rep,name=extra_ldflag,json=extraLdflag" json:"extra_ldflag,omitempty"`
	// Packages which don’t support separate builddirs need to use this
	// kludge. Bugs should be reported with the respective upstream authors.
	CopyToBuilddir       *bool    `protobuf:"varint,2,opt,name=copy_to_builddir,json=copyToBuilddir" json:"copy_to_builddir,omitempty"`
	ExtraMakeFlag        []string `protobuf:"bytes,4,rep,name=extra_make_flag,json=extraMakeFlag" json:"extra_make_flag,omitempty"`
	Autoreconf           *bool    `protobuf:"varint,5,opt,name=autoreconf" json:"autoreconf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CBuilder) Reset()         { *m = CBuilder{} }
func (m *CBuilder) String() string { return proto.CompactTextString(m) }
func (*CBuilder) ProtoMessage()    {}
func (*CBuilder) Descriptor() ([]byte, []int) {
	return fileDescriptor_14ce178a580e4ede, []int{1}
}

func (m *CBuilder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CBuilder.Unmarshal(m, b)
}
func (m *CBuilder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CBuilder.Marshal(b, m, deterministic)
}
func (m *CBuilder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CBuilder.Merge(m, src)
}
func (m *CBuilder) XXX_Size() int {
	return xxx_messageInfo_CBuilder.Size(m)
}
func (m *CBuilder) XXX_DiscardUnknown() {
	xxx_messageInfo_CBuilder.DiscardUnknown(m)
}

var xxx_messageInfo_CBuilder proto.InternalMessageInfo

func (m *CBuilder) GetExtraConfigureFlag() []string {
	if m != nil {
		return m.ExtraConfigureFlag
	}
	return nil
}

func (m *CBuilder) GetExtraLdflag() []string {
	if m != nil {
		return m.ExtraLdflag
	}
	return nil
}

func (m *CBuilder) GetCopyToBuilddir() bool {
	if m != nil && m.CopyToBuilddir != nil {
		return *m.CopyToBuilddir
	}
	return false
}

func (m *CBuilder) GetExtraMakeFlag() []string {
	if m != nil {
		return m.ExtraMakeFlag
	}
	return nil
}

func (m *CBuilder) GetAutoreconf() bool {
	if m != nil && m.Autoreconf != nil {
		return *m.Autoreconf
	}
	return false
}

type CMakeBuilder struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CMakeBuilder) Reset()         { *m = CMakeBuilder{} }
func (m *CMakeBuilder) String() string { return proto.CompactTextString(m) }
func (*CMakeBuilder) ProtoMessage()    {}
func (*CMakeBuilder) Descriptor() ([]byte, []int) {
	return fileDescriptor_14ce178a580e4ede, []int{2}
}

func (m *CMakeBuilder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMakeBuilder.Unmarshal(m, b)
}
func (m *CMakeBuilder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMakeBuilder.Marshal(b, m, deterministic)
}
func (m *CMakeBuilder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMakeBuilder.Merge(m, src)
}
func (m *CMakeBuilder) XXX_Size() int {
	return xxx_messageInfo_CMakeBuilder.Size(m)
}
func (m *CMakeBuilder) XXX_DiscardUnknown() {
	xxx_messageInfo_CMakeBuilder.DiscardUnknown(m)
}

var xxx_messageInfo_CMakeBuilder proto.InternalMessageInfo

type PerlBuilder struct {
	// Extra flags to be specified when running Makefile.PL, e.g.
	// EXPATLIBPATH=/ro/expat-2.2.6/buildoutput/lib
	ExtraMakefileFlag    []string `protobuf:"bytes,1,rep,name=extra_makefile_flag,json=extraMakefileFlag" json:"extra_makefile_flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PerlBuilder) Reset()         { *m = PerlBuilder{} }
func (m *PerlBuilder) String() string { return proto.CompactTextString(m) }
func (*PerlBuilder) ProtoMessage()    {}
func (*PerlBuilder) Descriptor() ([]byte, []int) {
	return fileDescriptor_14ce178a580e4ede, []int{3}
}

func (m *PerlBuilder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PerlBuilder.Unmarshal(m, b)
}
func (m *PerlBuilder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PerlBuilder.Marshal(b, m, deterministic)
}
func (m *PerlBuilder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerlBuilder.Merge(m, src)
}
func (m *PerlBuilder) XXX_Size() int {
	return xxx_messageInfo_PerlBuilder.Size(m)
}
func (m *PerlBuilder) XXX_DiscardUnknown() {
	xxx_messageInfo_PerlBuilder.DiscardUnknown(m)
}

var xxx_messageInfo_PerlBuilder proto.InternalMessageInfo

func (m *PerlBuilder) GetExtraMakefileFlag() []string {
	if m != nil {
		return m.ExtraMakefileFlag
	}
	return nil
}

type PythonBuilder struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PythonBuilder) Reset()         { *m = PythonBuilder{} }
func (m *PythonBuilder) String() string { return proto.CompactTextString(m) }
func (*PythonBuilder) ProtoMessage()    {}
func (*PythonBuilder) Descriptor() ([]byte, []int) {
	return fileDescriptor_14ce178a580e4ede, []int{4}
}

func (m *PythonBuilder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PythonBuilder.Unmarshal(m, b)
}
func (m *PythonBuilder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PythonBuilder.Marshal(b, m, deterministic)
}
func (m *PythonBuilder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PythonBuilder.Merge(m, src)
}
func (m *PythonBuilder) XXX_Size() int {
	return xxx_messageInfo_PythonBuilder.Size(m)
}
func (m *PythonBuilder) XXX_DiscardUnknown() {
	xxx_messageInfo_PythonBuilder.DiscardUnknown(m)
}

var xxx_messageInfo_PythonBuilder proto.InternalMessageInfo

type GomodBuilder struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GomodBuilder) Reset()         { *m = GomodBuilder{} }
func (m *GomodBuilder) String() string { return proto.CompactTextString(m) }
func (*GomodBuilder) ProtoMessage()    {}
func (*GomodBuilder) Descriptor() ([]byte, []int) {
	return fileDescriptor_14ce178a580e4ede, []int{5}
}

func (m *GomodBuilder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GomodBuilder.Unmarshal(m, b)
}
func (m *GomodBuilder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GomodBuilder.Marshal(b, m, deterministic)
}
func (m *GomodBuilder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GomodBuilder.Merge(m, src)
}
func (m *GomodBuilder) XXX_Size() int {
	return xxx_messageInfo_GomodBuilder.Size(m)
}
func (m *GomodBuilder) XXX_DiscardUnknown() {
	xxx_messageInfo_GomodBuilder.DiscardUnknown(m)
}

var xxx_messageInfo_GomodBuilder proto.InternalMessageInfo

type Install struct {
	SystemdUnit          []string           `protobuf:"bytes,1,rep,name=systemd_unit,json=systemdUnit" json:"systemd_unit,omitempty"`
	Symlink              []*Install_Symlink `protobuf:"bytes,2,rep,name=symlink" json:"symlink,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Install) Reset()         { *m = Install{} }
func (m *Install) String() string { return proto.CompactTextString(m) }
func (*Install) ProtoMessage()    {}
func (*Install) Descriptor() ([]byte, []int) {
	return fileDescriptor_14ce178a580e4ede, []int{6}
}

func (m *Install) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Install.Unmarshal(m, b)
}
func (m *Install) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Install.Marshal(b, m, deterministic)
}
func (m *Install) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Install.Merge(m, src)
}
func (m *Install) XXX_Size() int {
	return xxx_messageInfo_Install.Size(m)
}
func (m *Install) XXX_DiscardUnknown() {
	xxx_messageInfo_Install.DiscardUnknown(m)
}

var xxx_messageInfo_Install proto.InternalMessageInfo

func (m *Install) GetSystemdUnit() []string {
	if m != nil {
		return m.SystemdUnit
	}
	return nil
}

func (m *Install) GetSymlink() []*Install_Symlink {
	if m != nil {
		return m.Symlink
	}
	return nil
}

type Install_Symlink struct {
	Oldname              *string  `protobuf:"bytes,1,opt,name=oldname" json:"oldname,omitempty"`
	Newname              *string  `protobuf:"bytes,2,opt,name=newname" json:"newname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Install_Symlink) Reset()         { *m = Install_Symlink{} }
func (m *Install_Symlink) String() string { return proto.CompactTextString(m) }
func (*Install_Symlink) ProtoMessage()    {}
func (*Install_Symlink) Descriptor() ([]byte, []int) {
	return fileDescriptor_14ce178a580e4ede, []int{6, 0}
}

func (m *Install_Symlink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Install_Symlink.Unmarshal(m, b)
}
func (m *Install_Symlink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Install_Symlink.Marshal(b, m, deterministic)
}
func (m *Install_Symlink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Install_Symlink.Merge(m, src)
}
func (m *Install_Symlink) XXX_Size() int {
	return xxx_messageInfo_Install_Symlink.Size(m)
}
func (m *Install_Symlink) XXX_DiscardUnknown() {
	xxx_messageInfo_Install_Symlink.DiscardUnknown(m)
}

var xxx_messageInfo_Install_Symlink proto.InternalMessageInfo

func (m *Install_Symlink) GetOldname() string {
	if m != nil && m.Oldname != nil {
		return *m.Oldname
	}
	return ""
}

func (m *Install_Symlink) GetNewname() string {
	if m != nil && m.Newname != nil {
		return *m.Newname
	}
	return ""
}

type Claim struct {
	Glob                 *string  `protobuf:"bytes,1,opt,name=glob" json:"glob,omitempty"`
	Dir                  *string  `protobuf:"bytes,2,opt,name=dir" json:"dir,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Claim) Reset()         { *m = Claim{} }
func (m *Claim) String() string { return proto.CompactTextString(m) }
func (*Claim) ProtoMessage()    {}
func (*Claim) Descriptor() ([]byte, []int) {
	return fileDescriptor_14ce178a580e4ede, []int{7}
}

func (m *Claim) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Claim.Unmarshal(m, b)
}
func (m *Claim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Claim.Marshal(b, m, deterministic)
}
func (m *Claim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Claim.Merge(m, src)
}
func (m *Claim) XXX_Size() int {
	return xxx_messageInfo_Claim.Size(m)
}
func (m *Claim) XXX_DiscardUnknown() {
	xxx_messageInfo_Claim.DiscardUnknown(m)
}

var xxx_messageInfo_Claim proto.InternalMessageInfo

func (m *Claim) GetGlob() string {
	if m != nil && m.Glob != nil {
		return *m.Glob
	}
	return ""
}

func (m *Claim) GetDir() string {
	if m != nil && m.Dir != nil {
		return *m.Dir
	}
	return ""
}

type SplitPackage struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	RuntimeDep           []string `protobuf:"bytes,2,rep,name=runtime_dep,json=runtimeDep" json:"runtime_dep,omitempty"`
	Claim                []*Claim `protobuf:"bytes,3,rep,name=claim" json:"claim,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SplitPackage) Reset()         { *m = SplitPackage{} }
func (m *SplitPackage) String() string { return proto.CompactTextString(m) }
func (*SplitPackage) ProtoMessage()    {}
func (*SplitPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_14ce178a580e4ede, []int{8}
}

func (m *SplitPackage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SplitPackage.Unmarshal(m, b)
}
func (m *SplitPackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SplitPackage.Marshal(b, m, deterministic)
}
func (m *SplitPackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SplitPackage.Merge(m, src)
}
func (m *SplitPackage) XXX_Size() int {
	return xxx_messageInfo_SplitPackage.Size(m)
}
func (m *SplitPackage) XXX_DiscardUnknown() {
	xxx_messageInfo_SplitPackage.DiscardUnknown(m)
}

var xxx_messageInfo_SplitPackage proto.InternalMessageInfo

func (m *SplitPackage) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SplitPackage) GetRuntimeDep() []string {
	if m != nil {
		return m.RuntimeDep
	}
	return nil
}

func (m *SplitPackage) GetClaim() []*Claim {
	if m != nil {
		return m.Claim
	}
	return nil
}

type Build struct {
	// URL to upstream source archive. Currently, only tar.gz archives are supported.
	Source *string `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	// Hash (TODO: describe length, introduce prefixes?) of the upstream source archive.
	Hash *string `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	// Version number, to be concatenated to the package name.
	Version *string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	// Types that are valid to be assigned to Builder:
	//	*Build_Cbuilder
	//	*Build_Cmakebuilder
	//	*Build_Perlbuilder
	//	*Build_Pythonbuilder
	//	*Build_Gomodbuilder
	Builder isBuild_Builder `protobuf_oneof:"builder"`
	// TODO: move this field into a custom builder
	BuildStep []*BuildStep `protobuf:"bytes,4,rep,name=build_step,json=buildStep" json:"build_step,omitempty"`
	// TODO: rename to build_dep
	Dep        []string `protobuf:"bytes,5,rep,name=dep" json:"dep,omitempty"`
	RuntimeDep []string `protobuf:"bytes,9,rep,name=runtime_dep,json=runtimeDep" json:"runtime_dep,omitempty"`
	// URLs to patches to cherry-pick after extracting the source, e.g.:
	// https://git.savannah.gnu.org/cgit/make.git/patch/?id=48c8a116a914a325a0497721f5d8b58d5bba34d4
	CherryPick           []string        `protobuf:"bytes,6,rep,name=cherry_pick,json=cherryPick" json:"cherry_pick,omitempty"`
	Install              *Install        `protobuf:"bytes,8,opt,name=install" json:"install,omitempty"`
	SplitPackage         []*SplitPackage `protobuf:"bytes,11,rep,name=split_package,json=splitPackage" json:"split_package,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Build) Reset()         { *m = Build{} }
func (m *Build) String() string { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()    {}
func (*Build) Descriptor() ([]byte, []int) {
	return fileDescriptor_14ce178a580e4ede, []int{9}
}

func (m *Build) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Build.Unmarshal(m, b)
}
func (m *Build) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Build.Marshal(b, m, deterministic)
}
func (m *Build) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Build.Merge(m, src)
}
func (m *Build) XXX_Size() int {
	return xxx_messageInfo_Build.Size(m)
}
func (m *Build) XXX_DiscardUnknown() {
	xxx_messageInfo_Build.DiscardUnknown(m)
}

var xxx_messageInfo_Build proto.InternalMessageInfo

func (m *Build) GetSource() string {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return ""
}

func (m *Build) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func (m *Build) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

type isBuild_Builder interface {
	isBuild_Builder()
}

type Build_Cbuilder struct {
	Cbuilder *CBuilder `protobuf:"bytes,7,opt,name=cbuilder,oneof"`
}

type Build_Cmakebuilder struct {
	Cmakebuilder *CMakeBuilder `protobuf:"bytes,14,opt,name=cmakebuilder,oneof"`
}

type Build_Perlbuilder struct {
	Perlbuilder *PerlBuilder `protobuf:"bytes,10,opt,name=perlbuilder,oneof"`
}

type Build_Pythonbuilder struct {
	Pythonbuilder *PythonBuilder `protobuf:"bytes,12,opt,name=pythonbuilder,oneof"`
}

type Build_Gomodbuilder struct {
	Gomodbuilder *GomodBuilder `protobuf:"bytes,13,opt,name=gomodbuilder,oneof"`
}

func (*Build_Cbuilder) isBuild_Builder() {}

func (*Build_Cmakebuilder) isBuild_Builder() {}

func (*Build_Perlbuilder) isBuild_Builder() {}

func (*Build_Pythonbuilder) isBuild_Builder() {}

func (*Build_Gomodbuilder) isBuild_Builder() {}

func (m *Build) GetBuilder() isBuild_Builder {
	if m != nil {
		return m.Builder
	}
	return nil
}

func (m *Build) GetCbuilder() *CBuilder {
	if x, ok := m.GetBuilder().(*Build_Cbuilder); ok {
		return x.Cbuilder
	}
	return nil
}

func (m *Build) GetCmakebuilder() *CMakeBuilder {
	if x, ok := m.GetBuilder().(*Build_Cmakebuilder); ok {
		return x.Cmakebuilder
	}
	return nil
}

func (m *Build) GetPerlbuilder() *PerlBuilder {
	if x, ok := m.GetBuilder().(*Build_Perlbuilder); ok {
		return x.Perlbuilder
	}
	return nil
}

func (m *Build) GetPythonbuilder() *PythonBuilder {
	if x, ok := m.GetBuilder().(*Build_Pythonbuilder); ok {
		return x.Pythonbuilder
	}
	return nil
}

func (m *Build) GetGomodbuilder() *GomodBuilder {
	if x, ok := m.GetBuilder().(*Build_Gomodbuilder); ok {
		return x.Gomodbuilder
	}
	return nil
}

func (m *Build) GetBuildStep() []*BuildStep {
	if m != nil {
		return m.BuildStep
	}
	return nil
}

func (m *Build) GetDep() []string {
	if m != nil {
		return m.Dep
	}
	return nil
}

func (m *Build) GetRuntimeDep() []string {
	if m != nil {
		return m.RuntimeDep
	}
	return nil
}

func (m *Build) GetCherryPick() []string {
	if m != nil {
		return m.CherryPick
	}
	return nil
}

func (m *Build) GetInstall() *Install {
	if m != nil {
		return m.Install
	}
	return nil
}

func (m *Build) GetSplitPackage() []*SplitPackage {
	if m != nil {
		return m.SplitPackage
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Build) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Build_OneofMarshaler, _Build_OneofUnmarshaler, _Build_OneofSizer, []interface{}{
		(*Build_Cbuilder)(nil),
		(*Build_Cmakebuilder)(nil),
		(*Build_Perlbuilder)(nil),
		(*Build_Pythonbuilder)(nil),
		(*Build_Gomodbuilder)(nil),
	}
}

func _Build_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Build)
	// builder
	switch x := m.Builder.(type) {
	case *Build_Cbuilder:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Cbuilder); err != nil {
			return err
		}
	case *Build_Cmakebuilder:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Cmakebuilder); err != nil {
			return err
		}
	case *Build_Perlbuilder:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Perlbuilder); err != nil {
			return err
		}
	case *Build_Pythonbuilder:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Pythonbuilder); err != nil {
			return err
		}
	case *Build_Gomodbuilder:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Gomodbuilder); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Build.Builder has unexpected type %T", x)
	}
	return nil
}

func _Build_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Build)
	switch tag {
	case 7: // builder.cbuilder
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CBuilder)
		err := b.DecodeMessage(msg)
		m.Builder = &Build_Cbuilder{msg}
		return true, err
	case 14: // builder.cmakebuilder
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CMakeBuilder)
		err := b.DecodeMessage(msg)
		m.Builder = &Build_Cmakebuilder{msg}
		return true, err
	case 10: // builder.perlbuilder
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PerlBuilder)
		err := b.DecodeMessage(msg)
		m.Builder = &Build_Perlbuilder{msg}
		return true, err
	case 12: // builder.pythonbuilder
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PythonBuilder)
		err := b.DecodeMessage(msg)
		m.Builder = &Build_Pythonbuilder{msg}
		return true, err
	case 13: // builder.gomodbuilder
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GomodBuilder)
		err := b.DecodeMessage(msg)
		m.Builder = &Build_Gomodbuilder{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Build_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Build)
	// builder
	switch x := m.Builder.(type) {
	case *Build_Cbuilder:
		s := proto.Size(x.Cbuilder)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Build_Cmakebuilder:
		s := proto.Size(x.Cmakebuilder)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Build_Perlbuilder:
		s := proto.Size(x.Perlbuilder)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Build_Pythonbuilder:
		s := proto.Size(x.Pythonbuilder)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Build_Gomodbuilder:
		s := proto.Size(x.Gomodbuilder)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*BuildStep)(nil), "pb.BuildStep")
	proto.RegisterType((*CBuilder)(nil), "pb.CBuilder")
	proto.RegisterType((*CMakeBuilder)(nil), "pb.CMakeBuilder")
	proto.RegisterType((*PerlBuilder)(nil), "pb.PerlBuilder")
	proto.RegisterType((*PythonBuilder)(nil), "pb.PythonBuilder")
	proto.RegisterType((*GomodBuilder)(nil), "pb.GomodBuilder")
	proto.RegisterType((*Install)(nil), "pb.Install")
	proto.RegisterType((*Install_Symlink)(nil), "pb.Install.Symlink")
	proto.RegisterType((*Claim)(nil), "pb.Claim")
	proto.RegisterType((*SplitPackage)(nil), "pb.SplitPackage")
	proto.RegisterType((*Build)(nil), "pb.Build")
}

func init() { proto.RegisterFile("build.proto", fileDescriptor_14ce178a580e4ede) }

var fileDescriptor_14ce178a580e4ede = []byte{
	// 653 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xcd, 0xa5, 0xa9, 0xe3, 0xb1, 0x93, 0xb6, 0x5b, 0x84, 0x2c, 0x1e, 0x68, 0xb0, 0x04, 0x8a,
	0x10, 0x8d, 0x50, 0x10, 0x48, 0x3c, 0xf4, 0xa5, 0x41, 0x10, 0x24, 0x90, 0xa2, 0x2d, 0x3c, 0x5b,
	0x8e, 0xbd, 0x75, 0x56, 0x59, 0x7b, 0x57, 0x6b, 0xa7, 0x90, 0x5f, 0xe1, 0x7b, 0xf8, 0x01, 0xfe,
	0x08, 0xed, 0x78, 0x9d, 0xba, 0x88, 0xb7, 0xd9, 0x33, 0xe7, 0xcc, 0x25, 0x73, 0x62, 0xf0, 0xd6,
	0x3b, 0x2e, 0xd2, 0x99, 0xd2, 0xb2, 0x92, 0xa4, 0xa7, 0xd6, 0xe1, 0x05, 0xb8, 0xd7, 0x06, 0xba,
	0xa9, 0x98, 0x22, 0x04, 0x8e, 0x62, 0x9d, 0xdd, 0x05, 0xdd, 0x49, 0x7f, 0xea, 0x52, 0x8c, 0xc3,
	0x3f, 0x5d, 0x18, 0x2e, 0x90, 0xc2, 0x34, 0x79, 0x0d, 0x8f, 0xd8, 0xcf, 0x4a, 0xc7, 0x51, 0x22,
	0x8b, 0x5b, 0x9e, 0xed, 0x34, 0x8b, 0x6e, 0x45, 0x9c, 0x59, 0x01, 0xc1, 0xdc, 0xa2, 0x49, 0x7d,
	0x14, 0x71, 0x46, 0x9e, 0x81, 0x5f, 0x2b, 0x44, 0x8a, 0xcc, 0x3e, 0x32, 0x3d, 0xc4, 0xbe, 0x20,
	0x44, 0xa6, 0x70, 0x9a, 0x48, 0xb5, 0x8f, 0x2a, 0x19, 0xe1, 0x74, 0x29, 0xd7, 0x41, 0x6f, 0xd2,
	0x9d, 0x0e, 0xe9, 0xd8, 0xe0, 0xdf, 0xe4, 0xb5, 0x45, 0xc9, 0x0b, 0x38, 0xa9, 0x8b, 0xe5, 0xf1,
	0xd6, 0x76, 0x3e, 0xc2, 0x7a, 0x23, 0x84, 0xbf, 0xc6, 0xdb, 0xba, 0xe9, 0x53, 0x80, 0x78, 0x57,
	0x49, 0xcd, 0xcc, 0x98, 0xc1, 0x00, 0x6b, 0xb5, 0x90, 0x70, 0x0c, 0xfe, 0xc2, 0x90, 0xed, 0x5a,
	0xe1, 0x15, 0x78, 0x2b, 0xa6, 0x45, 0xb3, 0xe5, 0x0c, 0xce, 0xef, 0xdb, 0xdc, 0x72, 0xf1, 0x60,
	0xc9, 0xb3, 0x43, 0x2b, 0x93, 0x31, 0xed, 0xc2, 0x13, 0x18, 0xad, 0xf6, 0xd5, 0x46, 0x16, 0x4d,
	0xbd, 0x31, 0xf8, 0x9f, 0x64, 0x2e, 0xd3, 0xe6, 0xfd, 0xab, 0x0b, 0xce, 0xe7, 0xa2, 0xac, 0x62,
	0x21, 0xcc, 0x0f, 0x52, 0xee, 0xcb, 0x8a, 0xe5, 0x69, 0xb4, 0x2b, 0x78, 0x65, 0xab, 0x7a, 0x16,
	0xfb, 0x5e, 0xf0, 0x8a, 0x5c, 0x82, 0x53, 0xee, 0x73, 0xc1, 0x8b, 0x6d, 0xd0, 0x9b, 0xf4, 0xa7,
	0xde, 0xfc, 0x7c, 0xa6, 0xd6, 0x33, 0x5b, 0x60, 0x76, 0x53, 0xa7, 0x68, 0xc3, 0x79, 0x72, 0x05,
	0x8e, 0xc5, 0x48, 0x00, 0x8e, 0x14, 0x69, 0x11, 0xe7, 0x2c, 0xe8, 0x4e, 0xba, 0x53, 0x97, 0x36,
	0x4f, 0x93, 0x29, 0xd8, 0x0f, 0xcc, 0xf4, 0xea, 0x8c, 0x7d, 0x86, 0x97, 0x30, 0x58, 0x88, 0x98,
	0xe7, 0xe6, 0xfa, 0x99, 0x90, 0x6b, 0xab, 0xc4, 0x98, 0x9c, 0x42, 0xbf, 0x39, 0x87, 0x4b, 0x4d,
	0x18, 0xa6, 0xe0, 0xdf, 0x28, 0xc1, 0xab, 0x55, 0x9c, 0x6c, 0xe3, 0x8c, 0x19, 0x55, 0xab, 0x1f,
	0xc6, 0xe4, 0x02, 0x3c, 0xbd, 0x2b, 0x2a, 0x9e, 0xb3, 0x28, 0x65, 0x0a, 0x97, 0x70, 0x29, 0x58,
	0xe8, 0x03, 0x53, 0xe4, 0x02, 0x06, 0x89, 0xe9, 0x89, 0x76, 0xf0, 0xe6, 0xae, 0xd9, 0x0f, 0x87,
	0xa0, 0x35, 0x1e, 0xfe, 0x3e, 0x82, 0x01, 0xfe, 0x7a, 0xe4, 0x31, 0x1c, 0x97, 0x72, 0xa7, 0x93,
	0xa6, 0x83, 0x7d, 0x99, 0xbe, 0x9b, 0xb8, 0xdc, 0xd8, 0xd1, 0x30, 0x36, 0x4b, 0xde, 0x31, 0x5d,
	0x72, 0x59, 0x04, 0xfd, 0x7a, 0x49, 0xfb, 0x24, 0x2f, 0x61, 0x98, 0xac, 0xeb, 0x6b, 0x04, 0xce,
	0xa4, 0x3b, 0xf5, 0xe6, 0x3e, 0xf6, 0xb4, 0x17, 0x5a, 0x76, 0xe8, 0x21, 0x4f, 0xde, 0x81, 0x9f,
	0x98, 0xcb, 0x37, 0xfc, 0x31, 0xf2, 0x4f, 0x91, 0xdf, 0x72, 0xcd, 0xb2, 0x43, 0x1f, 0xf0, 0xc8,
	0x1b, 0xf0, 0x14, 0xd3, 0xa2, 0x91, 0x01, 0xca, 0x4e, 0x8c, 0xac, 0x65, 0xae, 0x65, 0x87, 0xb6,
	0x59, 0xe4, 0x3d, 0x8c, 0x14, 0x7a, 0xa7, 0x91, 0xf9, 0x28, 0x3b, 0x43, 0x59, 0xdb, 0x54, 0xcb,
	0x0e, 0x7d, 0xc8, 0x34, 0x73, 0x66, 0xc6, 0x65, 0x8d, 0x72, 0x74, 0x3f, 0x67, 0xdb, 0x7d, 0x66,
	0xce, 0x36, 0x8f, 0xbc, 0x02, 0xc0, 0x30, 0x2a, 0x2b, 0xa6, 0xf0, 0x0f, 0xe4, 0xcd, 0x47, 0x46,
	0x75, 0xf8, 0x10, 0x50, 0x77, 0x7d, 0xf8, 0x26, 0x18, 0x07, 0x30, 0x15, 0x0c, 0xf0, 0x86, 0x26,
	0xfc, 0xf7, 0xba, 0xee, 0x7f, 0xae, 0xeb, 0x25, 0x1b, 0xa6, 0xf5, 0x3e, 0x52, 0x3c, 0xd9, 0x06,
	0xc7, 0x35, 0xa1, 0x86, 0x56, 0x3c, 0xd9, 0x92, 0xe7, 0xe0, 0xf0, 0xda, 0xcd, 0xc1, 0x10, 0x87,
	0xf6, 0x5a, 0x06, 0xa7, 0x4d, 0x8e, 0xbc, 0x85, 0x51, 0x69, 0xac, 0x16, 0xa9, 0xda, 0x6b, 0x81,
	0x87, 0xb3, 0xe2, 0x86, 0x6d, 0x0f, 0x52, 0xbf, 0x6c, 0xbd, 0xae, 0x5d, 0x70, 0xec, 0xaa, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xbf, 0x17, 0xdd, 0xef, 0x04, 0x00, 0x00,
}
