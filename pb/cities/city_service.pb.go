// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: city_service.proto

package cities

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_city_service_proto protoreflect.FileDescriptor

var file_city_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x12, 0x63, 0x69,
	0x74, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x64, 0x0a, 0x0d, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2c, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x12, 0x0c, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x1a,
	0x0c, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x25, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x12, 0x0a, 0x2e, 0x63, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x43, 0x69, 0x74, 0x79, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x62, 0x2f, 0x63, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x3b, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_city_service_proto_goTypes = []any{
	(*City)(nil), // 0: cities.City
	(*Id)(nil),   // 1: cities.Id
}
var file_city_service_proto_depIdxs = []int32{
	0, // 0: cities.CitiesService.GetCitylocal:input_type -> cities.City
	1, // 1: cities.CitiesService.GetCity:input_type -> cities.Id
	0, // 2: cities.CitiesService.GetCitylocal:output_type -> cities.City
	0, // 3: cities.CitiesService.GetCity:output_type -> cities.City
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_city_service_proto_init() }
func file_city_service_proto_init() {
	if File_city_service_proto != nil {
		return
	}
	file_city_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_service_proto_goTypes,
		DependencyIndexes: file_city_service_proto_depIdxs,
	}.Build()
	File_city_service_proto = out.File
	file_city_service_proto_rawDesc = nil
	file_city_service_proto_goTypes = nil
	file_city_service_proto_depIdxs = nil
}
