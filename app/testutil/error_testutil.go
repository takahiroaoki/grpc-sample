package testutil

import "github.com/takahiroaoki/grpc-sample/app/domain/domerr"

func SameDomainErrors(err1, err2 domerr.DomErr) bool {
	return (err1.Error() == err2.Error()) && (err1.Cause() == err2.Cause()) && (err1.LogLevel() == err2.LogLevel())
}
