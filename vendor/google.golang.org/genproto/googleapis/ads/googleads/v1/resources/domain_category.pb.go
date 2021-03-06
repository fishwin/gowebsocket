// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/domain_category.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A category generated automatically by crawling a domain. If a campaign uses
// the DynamicSearchAdsSetting, then domain categories will be generated for
// the domain. The categories can be targeted using WebpageConditionInfo.
// See: https://support.google.com/google-ads/answer/2471185
type DomainCategory struct {
	// The resource name of the domain category.
	// Domain category resource names have the form:
	//
	//
	// `customers/{customer_id}/domainCategories/{campaign_id}~{category_base64}~{language_code}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The campaign this category is recommended for.
	Campaign *wrappers.StringValue `protobuf:"bytes,2,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Recommended category for the website domain. e.g. if you have a website
	// about electronics, the categories could be "cameras", "televisions", etc.
	Category *wrappers.StringValue `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	// The language code specifying the language of the website. e.g. "en" for
	// English. The language can be specified in the DynamicSearchAdsSetting
	// required for dynamic search ads. This is the language of the pages from
	// your website that you want Google Ads to find, create ads for,
	// and match searches with.
	LanguageCode *wrappers.StringValue `protobuf:"bytes,4,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// The domain for the website. The domain can be specified in the
	// DynamicSearchAdsSetting required for dynamic search ads.
	Domain *wrappers.StringValue `protobuf:"bytes,5,opt,name=domain,proto3" json:"domain,omitempty"`
	// Fraction of pages on your site that this category matches.
	CoverageFraction *wrappers.DoubleValue `protobuf:"bytes,6,opt,name=coverage_fraction,json=coverageFraction,proto3" json:"coverage_fraction,omitempty"`
	// The position of this category in the set of categories. Lower numbers
	// indicate a better match for the domain. null indicates not recommended.
	CategoryRank *wrappers.Int64Value `protobuf:"bytes,7,opt,name=category_rank,json=categoryRank,proto3" json:"category_rank,omitempty"`
	// Indicates whether this category has sub-categories.
	HasChildren *wrappers.BoolValue `protobuf:"bytes,8,opt,name=has_children,json=hasChildren,proto3" json:"has_children,omitempty"`
	// The recommended cost per click for the category.
	RecommendedCpcBidMicros *wrappers.Int64Value `protobuf:"bytes,9,opt,name=recommended_cpc_bid_micros,json=recommendedCpcBidMicros,proto3" json:"recommended_cpc_bid_micros,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}             `json:"-"`
	XXX_unrecognized        []byte               `json:"-"`
	XXX_sizecache           int32                `json:"-"`
}

func (m *DomainCategory) Reset()         { *m = DomainCategory{} }
func (m *DomainCategory) String() string { return proto.CompactTextString(m) }
func (*DomainCategory) ProtoMessage()    {}
func (*DomainCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2d1ae0f79ccb888, []int{0}
}

func (m *DomainCategory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DomainCategory.Unmarshal(m, b)
}
func (m *DomainCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DomainCategory.Marshal(b, m, deterministic)
}
func (m *DomainCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DomainCategory.Merge(m, src)
}
func (m *DomainCategory) XXX_Size() int {
	return xxx_messageInfo_DomainCategory.Size(m)
}
func (m *DomainCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_DomainCategory.DiscardUnknown(m)
}

var xxx_messageInfo_DomainCategory proto.InternalMessageInfo

func (m *DomainCategory) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *DomainCategory) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *DomainCategory) GetCategory() *wrappers.StringValue {
	if m != nil {
		return m.Category
	}
	return nil
}

func (m *DomainCategory) GetLanguageCode() *wrappers.StringValue {
	if m != nil {
		return m.LanguageCode
	}
	return nil
}

func (m *DomainCategory) GetDomain() *wrappers.StringValue {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (m *DomainCategory) GetCoverageFraction() *wrappers.DoubleValue {
	if m != nil {
		return m.CoverageFraction
	}
	return nil
}

func (m *DomainCategory) GetCategoryRank() *wrappers.Int64Value {
	if m != nil {
		return m.CategoryRank
	}
	return nil
}

func (m *DomainCategory) GetHasChildren() *wrappers.BoolValue {
	if m != nil {
		return m.HasChildren
	}
	return nil
}

func (m *DomainCategory) GetRecommendedCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.RecommendedCpcBidMicros
	}
	return nil
}

func init() {
	proto.RegisterType((*DomainCategory)(nil), "google.ads.googleads.v1.resources.DomainCategory")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/domain_category.proto", fileDescriptor_c2d1ae0f79ccb888)
}

var fileDescriptor_c2d1ae0f79ccb888 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x6a, 0xd4, 0x4c,
	0x18, 0xc6, 0xd9, 0x6d, 0xbf, 0xfd, 0xda, 0xe9, 0xae, 0x68, 0x3c, 0x30, 0xac, 0x45, 0x5a, 0xa5,
	0xd0, 0xa3, 0x09, 0xd1, 0xa2, 0x12, 0x11, 0xcc, 0x6e, 0xb1, 0x54, 0x50, 0xca, 0x0a, 0x8b, 0xc8,
	0x42, 0x98, 0x9d, 0x99, 0xce, 0x0e, 0x4d, 0xe6, 0x0d, 0x33, 0xc9, 0x8a, 0x67, 0x5e, 0x8b, 0x87,
	0x5e, 0x89, 0x78, 0x29, 0x5e, 0x85, 0x24, 0x33, 0x13, 0x94, 0xa2, 0xdd, 0xb3, 0x37, 0x79, 0x9f,
	0xdf, 0xf3, 0xbc, 0xf3, 0x0f, 0x3d, 0x13, 0x00, 0x22, 0xe7, 0x11, 0x61, 0x26, 0xb2, 0x65, 0x53,
	0xad, 0xe3, 0x48, 0x73, 0x03, 0xb5, 0xa6, 0xdc, 0x44, 0x0c, 0x0a, 0x22, 0x55, 0x46, 0x49, 0xc5,
	0x05, 0xe8, 0xcf, 0xb8, 0xd4, 0x50, 0x41, 0x70, 0x68, 0xd5, 0x98, 0x30, 0x83, 0x3b, 0x10, 0xaf,
	0x63, 0xdc, 0x81, 0xe3, 0x7d, 0xef, 0x5d, 0xca, 0x88, 0x28, 0x05, 0x15, 0xa9, 0x24, 0x28, 0x63,
	0x0d, 0xc6, 0x0f, 0x5c, 0xb7, 0xfd, 0x5a, 0xd6, 0x97, 0xd1, 0x27, 0x4d, 0xca, 0x92, 0x6b, 0xd7,
	0x7f, 0xf8, 0x7d, 0x1b, 0xdd, 0x3a, 0x6d, 0xa3, 0xa7, 0x2e, 0x39, 0x78, 0x84, 0x46, 0xde, 0x3d,
	0x53, 0xa4, 0xe0, 0x61, 0xef, 0xa0, 0x77, 0xbc, 0x3b, 0x1b, 0xfa, 0x9f, 0xef, 0x48, 0xc1, 0x83,
	0xe7, 0x68, 0x87, 0x92, 0xa2, 0x24, 0x52, 0xa8, 0xb0, 0x7f, 0xd0, 0x3b, 0xde, 0x7b, 0xbc, 0xef,
	0x06, 0xc4, 0x3e, 0x0a, 0xbf, 0xaf, 0xb4, 0x54, 0x62, 0x4e, 0xf2, 0x9a, 0xcf, 0x3a, 0xb5, 0x25,
	0x6d, 0x54, 0xb8, 0xb5, 0x19, 0xe9, 0x06, 0x4b, 0xd1, 0x28, 0x27, 0x4a, 0xd4, 0x44, 0xf0, 0x8c,
	0x02, 0xe3, 0xe1, 0xf6, 0x06, 0xf8, 0xd0, 0x23, 0x53, 0x60, 0x3c, 0x38, 0x41, 0x03, 0xbb, 0xd1,
	0xe1, 0x7f, 0x1b, 0xb0, 0x4e, 0x1b, 0x9c, 0xa3, 0x3b, 0x14, 0xd6, 0x5c, 0x37, 0xc1, 0x97, 0x9a,
	0xd0, 0x66, 0x83, 0xc3, 0xc1, 0x5f, 0x0c, 0x4e, 0xa1, 0x5e, 0xe6, 0xdc, 0x1a, 0xdc, 0xf6, 0xd8,
	0x6b, 0x47, 0x05, 0xaf, 0xd0, 0xc8, 0xaf, 0x27, 0xd3, 0x44, 0x5d, 0x85, 0xff, 0xb7, 0x36, 0xf7,
	0xaf, 0xd9, 0x9c, 0xab, 0xea, 0xe9, 0x89, 0x5b, 0x82, 0x27, 0x66, 0x44, 0x5d, 0x05, 0x2f, 0xd1,
	0x70, 0x45, 0x4c, 0x46, 0x57, 0x32, 0x67, 0x9a, 0xab, 0x70, 0xa7, 0x35, 0x18, 0x5f, 0x33, 0x98,
	0x00, 0xe4, 0x96, 0xdf, 0x5b, 0x11, 0x33, 0x75, 0xf2, 0xe0, 0x03, 0x1a, 0x6b, 0x4e, 0xa1, 0x28,
	0xb8, 0x62, 0x9c, 0x65, 0xb4, 0xa4, 0xd9, 0x52, 0xb2, 0xac, 0x90, 0x54, 0x83, 0x09, 0x77, 0x6f,
	0x9e, 0xe6, 0xde, 0x6f, 0xf8, 0xb4, 0xa4, 0x13, 0xc9, 0xde, 0xb6, 0xec, 0xe4, 0x4b, 0x1f, 0x1d,
	0x51, 0x28, 0xf0, 0x8d, 0x57, 0x76, 0x72, 0xf7, 0xcf, 0x1b, 0x77, 0xd1, 0xa4, 0x5c, 0xf4, 0x3e,
	0xbe, 0x71, 0xa4, 0x80, 0xe6, 0xcc, 0x30, 0x68, 0x11, 0x09, 0xae, 0xda, 0x19, 0xfc, 0xab, 0x29,
	0xa5, 0xf9, 0xc7, 0x23, 0x7a, 0xd1, 0x55, 0x5f, 0xfb, 0x5b, 0x67, 0x69, 0xfa, 0xad, 0x7f, 0x78,
	0x66, 0x2d, 0x53, 0x66, 0xb0, 0x2d, 0x9b, 0x6a, 0x1e, 0xe3, 0x99, 0x57, 0xfe, 0xf0, 0x9a, 0x45,
	0xca, 0xcc, 0xa2, 0xd3, 0x2c, 0xe6, 0xf1, 0xa2, 0xd3, 0xfc, 0xec, 0x1f, 0xd9, 0x46, 0x92, 0xa4,
	0xcc, 0x24, 0x49, 0xa7, 0x4a, 0x92, 0x79, 0x9c, 0x24, 0x9d, 0x6e, 0x39, 0x68, 0x87, 0x7d, 0xf2,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xcd, 0x54, 0x36, 0xf0, 0x03, 0x00, 0x00,
}
