package errordict

import "github.com/Eitol/yapo/pkg/cerrors"

var (
	ErrUnableToParseAdSearchResponse = cerrors.Error{
		Code: "ErrUnableToParseAdSearchResponse",
	}
	ErrUnableToBuildTLSConfig = cerrors.Error{
		Code: "ErrUnableToBuildTlsConfig",
	}
	ErrReadingTheResponseBody = cerrors.Error{
		Code: "ErrReadingTheResponseBody",
	}
	ErrExecutingTheRequest = cerrors.Error{
		Code: "ErrExecutingTheRequest",
	}
	ErrSendingTheRequest = cerrors.Error{
		Code: "ErrSendingTheRequest",
	}
	ErrBuildingTheRequest = cerrors.Error{
		Code: "ErrBuildingTheRequest",
	}
)
