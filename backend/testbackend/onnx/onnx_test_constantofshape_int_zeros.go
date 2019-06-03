package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("ConstantOfShape", "TestConstantofshapeIntZeros", NewTestConstantofshapeIntZeros)
}

// NewTestConstantofshapeIntZeros version: 4.
func NewTestConstantofshapeIntZeros() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "ConstantOfShape",
		Title:  "TestConstantofshapeIntZeros",
		ModelB: []byte{0x8, 0x4, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x7b, 0xa, 0x33, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0xf, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4f, 0x66, 0x53, 0x68, 0x61, 0x70, 0x65, 0x2a, 0x1a, 0xa, 0x5, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0xe, 0x8, 0x1, 0x10, 0x6, 0x2a, 0x1, 0x0, 0x42, 0x5, 0x76, 0x61, 0x6c, 0x75, 0x65, 0xa0, 0x1, 0x4, 0x12, 0x1e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x6f, 0x66, 0x73, 0x68, 0x61, 0x70, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x5f, 0x7a, 0x65, 0x72, 0x6f, 0x73, 0x5a, 0xf, 0xa, 0x1, 0x78, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x62, 0x13, 0xa, 0x1, 0x79, 0x12, 0xe, 0xa, 0xc, 0x8, 0x6, 0x12, 0x8, 0xa, 0x2, 0x8, 0xa, 0xa, 0x2, 0x8, 0x6, 0x42, 0x2, 0x10, 0x9},

		/*

		   &pb.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "ConstantOfShape",
		     Attributes: ([]*pb.AttributeProto) (len=1 cap=1) {
		    (*pb.AttributeProto)(0xc000132900)(name:"value" type:TENSOR t:<dims:1 data_type:6 int32_data:0 name:"value" > )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]int64{10, 6}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(10, 6),
				tensor.WithBacking([]int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
			),
		},
	}
}
