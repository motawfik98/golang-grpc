// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: calculator.proto

package proto

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

var File_calculator_proto protoreflect.FileDescriptor

var file_calculator_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x73, 0x75, 0x6d, 0x1a, 0x09, 0x73, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x09, 0x61, 0x76, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x6d, 0x61, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x73, 0x71, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xf6, 0x01, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x0f,
	0x2e, 0x73, 0x75, 0x6d, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x73, 0x75, 0x6d, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x30, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x75, 0x6d,
	0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x73, 0x75, 0x6d, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x12, 0x2a, 0x0a, 0x03, 0x41, 0x76, 0x67, 0x12, 0x0f, 0x2e, 0x73, 0x75, 0x6d,
	0x2e, 0x41, 0x76, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x75,
	0x6d, 0x2e, 0x41, 0x76, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12,
	0x2c, 0x0a, 0x03, 0x4d, 0x61, 0x78, 0x12, 0x0f, 0x2e, 0x73, 0x75, 0x6d, 0x2e, 0x4d, 0x61, 0x78,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x75, 0x6d, 0x2e, 0x4d, 0x61,
	0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x2b, 0x0a,
	0x04, 0x53, 0x71, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x75, 0x6d, 0x2e, 0x53, 0x71, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x75, 0x6d, 0x2e, 0x53, 0x71,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x6f,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_calculator_proto_goTypes = []interface{}{
	(*SumRequest)(nil),    // 0: sum.SumRequest
	(*PrimeRequest)(nil),  // 1: sum.PrimeRequest
	(*AvgRequest)(nil),    // 2: sum.AvgRequest
	(*MaxRequest)(nil),    // 3: sum.MaxRequest
	(*SqrtRequest)(nil),   // 4: sum.SqrtRequest
	(*SumResponse)(nil),   // 5: sum.SumResponse
	(*PrimeResponse)(nil), // 6: sum.PrimeResponse
	(*AvgResponse)(nil),   // 7: sum.AvgResponse
	(*MaxResponse)(nil),   // 8: sum.MaxResponse
	(*SqrtResponse)(nil),  // 9: sum.SqrtResponse
}
var file_calculator_proto_depIdxs = []int32{
	0, // 0: sum.CalculatorService.Sum:input_type -> sum.SumRequest
	1, // 1: sum.CalculatorService.Prime:input_type -> sum.PrimeRequest
	2, // 2: sum.CalculatorService.Avg:input_type -> sum.AvgRequest
	3, // 3: sum.CalculatorService.Max:input_type -> sum.MaxRequest
	4, // 4: sum.CalculatorService.Sqrt:input_type -> sum.SqrtRequest
	5, // 5: sum.CalculatorService.Sum:output_type -> sum.SumResponse
	6, // 6: sum.CalculatorService.Prime:output_type -> sum.PrimeResponse
	7, // 7: sum.CalculatorService.Avg:output_type -> sum.AvgResponse
	8, // 8: sum.CalculatorService.Max:output_type -> sum.MaxResponse
	9, // 9: sum.CalculatorService.Sqrt:output_type -> sum.SqrtResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_proto_init() }
func file_calculator_proto_init() {
	if File_calculator_proto != nil {
		return
	}
	file_sum_proto_init()
	file_prime_proto_init()
	file_avg_proto_init()
	file_max_proto_init()
	file_sqrt_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_proto_depIdxs,
	}.Build()
	File_calculator_proto = out.File
	file_calculator_proto_rawDesc = nil
	file_calculator_proto_goTypes = nil
	file_calculator_proto_depIdxs = nil
}
