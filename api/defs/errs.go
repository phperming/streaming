package defs

type Err struct {
	 Error string `json:"error"`
	 ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error Err
}

//定义错误返回信息
var (
	ErrorRequestBodyParseFailed = ErrorResponse{
		HttpSC: 400,
		Error: Err{
			Error:"Request body is not correct",
			ErrorCode: "001",
		},
	}
	ErrorNoAuthUser = ErrorResponse{
		HttpSC: 401,
		Error: Err{
			Error: "User authentication failed",
			ErrorCode: "002",
		},
	}
	ErrorDBError = ErrorResponse{
		HttpSC: 500,
		Error: Err{
			Error: "DB ops failed",
			ErrorCode: "003",
		},
	}
	ErrorInternalFaults = ErrorResponse{
		HttpSC: 500,
		Error: Err{
			Error: "Internal Service error",
			ErrorCode: "004",
		},
	}
)
