package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("LogSoftmax", "TestLogsoftmaxNegativeAxis", NewTestLogsoftmaxNegativeAxis)
}

/*
&ir.ModelProto{
    IrVersion:   6,
    OpsetImport: {
        &ir.OperatorSetIdProto{Domain:"", Version:11},
    },
    ProducerName:    "backend-test",
    ProducerVersion: "",
    Domain:          "",
    ModelVersion:    0,
    DocString:       "",
    Graph:           &ir.GraphProto{
        Node: {
            &ir.NodeProto{
                Input:     {"x"},
                Output:    {"y"},
                Name:      "",
                OpType:    "LogSoftmax",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:        "axis",
                        RefAttrName: "",
                        DocString:   "",
                        Type:        2,
                        F:           0,
                        I:           -1,
                        S:           nil,
                        T:           (*ir.TensorProto)(nil),
                        G:           (*ir.GraphProto)(nil),
                        Floats:      nil,
                        Ints:        nil,
                        Strings:     nil,
                        Tensors:     nil,
                        Graphs:      nil,
                    },
                },
                DocString: "",
            },
        },
        Name:        "test_logsoftmax_negative_axis",
        Initializer: nil,
        DocString:   "",
        Input:       {
            &ir.ValueInfoProto{
                Name: "x",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:4},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:5},
                                        Denotation: "",
                                    },
                                },
                            },
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
        },
        Output: {
            &ir.ValueInfoProto{
                Name: "y",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:4},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:5},
                                        Denotation: "",
                                    },
                                },
                            },
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
        },
        ValueInfo:              nil,
        QuantizationAnnotation: nil,
    },
    MetadataProps: nil,
}
*/

// NewTestLogsoftmaxNegativeAxis version: 6.
func NewTestLogsoftmaxNegativeAxis() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "LogSoftmax",
		Title:  "TestLogsoftmaxNegativeAxis",
		ModelB: []byte{0x8, 0x6, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x7b, 0xa, 0x28, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0xa, 0x4c, 0x6f, 0x67, 0x53, 0x6f, 0x66, 0x74, 0x6d, 0x61, 0x78, 0x2a, 0x14, 0xa, 0x4, 0x61, 0x78, 0x69, 0x73, 0x18, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1, 0xa0, 0x1, 0x2, 0x12, 0x1d, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x6f, 0x66, 0x74, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x5a, 0x17, 0xa, 0x1, 0x78, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x62, 0x17, 0xa, 0x1, 0x79, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0xb},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "LogSoftmax",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc0001281c0)(name:"axis" type:INT i:-1 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{1.7640524, 0.4001572, 0.978738, 2.2408931, 1.867558, 0.9772779, 0.95008844, 0.1513572, 0.10321885, 0.41059852, 0.14404356, 1.4542735, 0.7610377, 0.121675014, 0.44386324, 0.33367434, 1.4940791, 0.20515826, 0.3130677, 0.85409576, 2.5529897, 0.6536186, 0.8644362, 0.742165, 2.2697546, 1.4543657, 0.045758516, 0.18718386, 1.5327792, 1.4693588, 0.15494743, 0.37816253, 0.88778573, 1.9807965, 0.34791216, 0.15634897, 1.2302907, 1.2023798, 0.3873268, 0.30230275, 1.048553, 1.420018, 1.7062702, 1.9507754, 0.5096522, 0.4380743, 1.2527953, 0.7774904, 1.6138978, 0.21274029, 0.89546657, 0.3869025, 0.51080513, 1.1806322, 0.028182229, 0.42833188, 0.06651722, 0.3024719, 0.6343221, 0.36274117}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{-1.4887761, -2.8526711, -2.2740903, -1.0119354, -1.3852705, -1.2225009, -1.2496903, -2.0484216, -2.09656, -1.7891803, -2.1846874, -0.8744575, -1.5676932, -2.207056, -1.8848677, -2.0454924, -0.88508755, -2.1740084, -2.066099, -1.5250709, -0.8115128, -2.710884, -2.5000663, -2.6223373, -1.0947479, -1.2872427, -2.69585, -2.5544245, -1.2088292, -1.2722496, -2.4640212, -2.240806, -1.7311828, -0.6381721, -2.2710564, -2.2181375, -1.1441958, -1.1721066, -1.9871596, -2.0721836, -2.005351, -1.6338862, -1.347634, -1.1031288, -2.544252, -2.1635222, -1.3488011, -1.8241062, -0.98769873, -2.3888562, -1.3942904, -1.9028544, -1.7789519, -1.1091248, -2.2615747, -1.5567553, -1.91857, -1.6826153, -1.3507651, -1.622346}),
			),
		},
	}
}
