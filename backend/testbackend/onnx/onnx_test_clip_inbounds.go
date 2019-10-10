package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Clip", "TestClipInbounds", NewTestClipInbounds)
}

// NewTestClipInbounds version: 3.
func NewTestClipInbounds() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Clip",
		Title:  "TestClipInbounds",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x62, 0xa, 0x2a, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x4, 0x43, 0x6c, 0x69, 0x70, 0x2a, 0xd, 0xa, 0x3, 0x6d, 0x61, 0x78, 0x15, 0x0, 0x0, 0xa0, 0x40, 0xa0, 0x1, 0x1, 0x2a, 0xd, 0xa, 0x3, 0x6d, 0x69, 0x6e, 0x15, 0x0, 0x0, 0xa0, 0xc0, 0xa0, 0x1, 0x1, 0x12, 0x12, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x70, 0x5f, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x5a, 0xf, 0xa, 0x1, 0x78, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x62, 0xf, 0xa, 0x1, 0x79, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Clip",
		     Attributes: ([]*ir.AttributeProto) (len=2 cap=2) {
		    (*ir.AttributeProto)(0xc000133a00)(name:"max" type:FLOAT f:5 ),
		    (*ir.AttributeProto)(0xc000133b00)(name:"min" type:FLOAT f:-5 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]float32{-1, 0, 1}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]float32{-1, 0, 1}),
			),
		},
	}
}
