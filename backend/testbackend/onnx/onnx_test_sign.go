package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Sign", "TestSign", NewTestSign)
}

// NewTestSign version: 3.
func NewTestSign() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Sign",
		Title:  "TestSign",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x3b, 0xa, 0xc, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x4, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x9, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x5a, 0xf, 0xa, 0x1, 0x78, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0xb, 0x62, 0xf, 0xa, 0x1, 0x79, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0xb, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Sign",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(11),
				tensor.WithBacking([]float32{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(11),
				tensor.WithBacking([]float32{-1, -1, -1, -1, -1, 0, 1, 1, 1, 1, 1}),
			),
		},
	}
}
