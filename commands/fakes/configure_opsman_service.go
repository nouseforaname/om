// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type ConfigureOpsmanService struct {
	UpdatePivnetTokenStub        func(string) error
	updatePivnetTokenMutex       sync.RWMutex
	updatePivnetTokenArgsForCall []struct {
		arg1 string
	}
	updatePivnetTokenReturns struct {
		result1 error
	}
	updatePivnetTokenReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateSSLCertificateStub        func(api.SSLCertificateInput) error
	updateSSLCertificateMutex       sync.RWMutex
	updateSSLCertificateArgsForCall []struct {
		arg1 api.SSLCertificateInput
	}
	updateSSLCertificateReturns struct {
		result1 error
	}
	updateSSLCertificateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ConfigureOpsmanService) UpdatePivnetToken(arg1 string) error {
	fake.updatePivnetTokenMutex.Lock()
	ret, specificReturn := fake.updatePivnetTokenReturnsOnCall[len(fake.updatePivnetTokenArgsForCall)]
	fake.updatePivnetTokenArgsForCall = append(fake.updatePivnetTokenArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("UpdatePivnetToken", []interface{}{arg1})
	fake.updatePivnetTokenMutex.Unlock()
	if fake.UpdatePivnetTokenStub != nil {
		return fake.UpdatePivnetTokenStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updatePivnetTokenReturns
	return fakeReturns.result1
}

func (fake *ConfigureOpsmanService) UpdatePivnetTokenCallCount() int {
	fake.updatePivnetTokenMutex.RLock()
	defer fake.updatePivnetTokenMutex.RUnlock()
	return len(fake.updatePivnetTokenArgsForCall)
}

func (fake *ConfigureOpsmanService) UpdatePivnetTokenCalls(stub func(string) error) {
	fake.updatePivnetTokenMutex.Lock()
	defer fake.updatePivnetTokenMutex.Unlock()
	fake.UpdatePivnetTokenStub = stub
}

func (fake *ConfigureOpsmanService) UpdatePivnetTokenArgsForCall(i int) string {
	fake.updatePivnetTokenMutex.RLock()
	defer fake.updatePivnetTokenMutex.RUnlock()
	argsForCall := fake.updatePivnetTokenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigureOpsmanService) UpdatePivnetTokenReturns(result1 error) {
	fake.updatePivnetTokenMutex.Lock()
	defer fake.updatePivnetTokenMutex.Unlock()
	fake.UpdatePivnetTokenStub = nil
	fake.updatePivnetTokenReturns = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureOpsmanService) UpdatePivnetTokenReturnsOnCall(i int, result1 error) {
	fake.updatePivnetTokenMutex.Lock()
	defer fake.updatePivnetTokenMutex.Unlock()
	fake.UpdatePivnetTokenStub = nil
	if fake.updatePivnetTokenReturnsOnCall == nil {
		fake.updatePivnetTokenReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updatePivnetTokenReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureOpsmanService) UpdateSSLCertificate(arg1 api.SSLCertificateInput) error {
	fake.updateSSLCertificateMutex.Lock()
	ret, specificReturn := fake.updateSSLCertificateReturnsOnCall[len(fake.updateSSLCertificateArgsForCall)]
	fake.updateSSLCertificateArgsForCall = append(fake.updateSSLCertificateArgsForCall, struct {
		arg1 api.SSLCertificateInput
	}{arg1})
	fake.recordInvocation("UpdateSSLCertificate", []interface{}{arg1})
	fake.updateSSLCertificateMutex.Unlock()
	if fake.UpdateSSLCertificateStub != nil {
		return fake.UpdateSSLCertificateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateSSLCertificateReturns
	return fakeReturns.result1
}

func (fake *ConfigureOpsmanService) UpdateSSLCertificateCallCount() int {
	fake.updateSSLCertificateMutex.RLock()
	defer fake.updateSSLCertificateMutex.RUnlock()
	return len(fake.updateSSLCertificateArgsForCall)
}

func (fake *ConfigureOpsmanService) UpdateSSLCertificateCalls(stub func(api.SSLCertificateInput) error) {
	fake.updateSSLCertificateMutex.Lock()
	defer fake.updateSSLCertificateMutex.Unlock()
	fake.UpdateSSLCertificateStub = stub
}

func (fake *ConfigureOpsmanService) UpdateSSLCertificateArgsForCall(i int) api.SSLCertificateInput {
	fake.updateSSLCertificateMutex.RLock()
	defer fake.updateSSLCertificateMutex.RUnlock()
	argsForCall := fake.updateSSLCertificateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigureOpsmanService) UpdateSSLCertificateReturns(result1 error) {
	fake.updateSSLCertificateMutex.Lock()
	defer fake.updateSSLCertificateMutex.Unlock()
	fake.UpdateSSLCertificateStub = nil
	fake.updateSSLCertificateReturns = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureOpsmanService) UpdateSSLCertificateReturnsOnCall(i int, result1 error) {
	fake.updateSSLCertificateMutex.Lock()
	defer fake.updateSSLCertificateMutex.Unlock()
	fake.UpdateSSLCertificateStub = nil
	if fake.updateSSLCertificateReturnsOnCall == nil {
		fake.updateSSLCertificateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateSSLCertificateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureOpsmanService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.updatePivnetTokenMutex.RLock()
	defer fake.updatePivnetTokenMutex.RUnlock()
	fake.updateSSLCertificateMutex.RLock()
	defer fake.updateSSLCertificateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ConfigureOpsmanService) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
