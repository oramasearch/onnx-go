package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Conv", "TestConvWithStridesPadding", NewTestConvWithStridesPadding)
}

// NewTestConvWithStridesPadding version: 3.
func NewTestConvWithStridesPadding() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Conv",
		Title:  "TestConvWithStridesPadding",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xc4, 0x1, 0xa, 0x4b, 0xa, 0x1, 0x78, 0xa, 0x1, 0x57, 0x12, 0x1, 0x79, 0x22, 0x4, 0x43, 0x6f, 0x6e, 0x76, 0x2a, 0x15, 0xa, 0xc, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x40, 0x3, 0x40, 0x3, 0xa0, 0x1, 0x7, 0x2a, 0x11, 0xa, 0x4, 0x70, 0x61, 0x64, 0x73, 0x40, 0x1, 0x40, 0x1, 0x40, 0x1, 0x40, 0x1, 0xa0, 0x1, 0x7, 0x2a, 0x10, 0xa, 0x7, 0x73, 0x74, 0x72, 0x69, 0x64, 0x65, 0x73, 0x40, 0x2, 0x40, 0x2, 0xa0, 0x1, 0x7, 0x12, 0x1e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x64, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5a, 0x1b, 0xa, 0x1, 0x78, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x7, 0xa, 0x2, 0x8, 0x5, 0x5a, 0x1b, 0xa, 0x1, 0x57, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x3, 0x62, 0x1b, 0xa, 0x1, 0x79, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x3, 0x42, 0x2, 0x10, 0x9},

		/*

		   &pb.NodeProto{
		     Input:     []string{"x", "W"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Conv",
		     Attributes: ([]*pb.AttributeProto) (len=3 cap=4) {
		    (*pb.AttributeProto)(0xc000126500)(name:"kernel_shape" type:INTS ints:3 ints:3 ),
		    (*pb.AttributeProto)(0xc000126600)(name:"pads" type:INTS ints:1 ints:1 ints:1 ints:1 ),
		    (*pb.AttributeProto)(0xc000126700)(name:"strides" type:INTS ints:2 ints:2 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 1, 7, 5),
				tensor.WithBacking([]float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34}),
			),

			tensor.New(
				tensor.WithShape(1, 1, 3, 3),
				tensor.WithBacking([]float32{1, 1, 1, 1, 1, 1, 1, 1, 1}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 1, 4, 3),
				tensor.WithBacking([]float32{12, 27, 24, 63, 108, 81, 123, 198, 141, 112, 177, 124}),
			),
		},
	}
}
