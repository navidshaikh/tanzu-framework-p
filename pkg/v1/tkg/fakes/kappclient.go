// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	v1alpha1a "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/packaging/v1alpha1"
	v1alpha1b "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apiserver/apis/datapackaging/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/vmware-tanzu-private/core/pkg/v1/tkg/kappclient"
)

type KappClient struct {
	CreatePackageInstallStub        func(*v1alpha1.PackageInstall, bool, bool) error
	createPackageInstallMutex       sync.RWMutex
	createPackageInstallArgsForCall []struct {
		arg1 *v1alpha1.PackageInstall
		arg2 bool
		arg3 bool
	}
	createPackageInstallReturns struct {
		result1 error
	}
	createPackageInstallReturnsOnCall map[int]struct {
		result1 error
	}
	CreatePackageRepositoryStub        func(*v1alpha1.PackageRepository) error
	createPackageRepositoryMutex       sync.RWMutex
	createPackageRepositoryArgsForCall []struct {
		arg1 *v1alpha1.PackageRepository
	}
	createPackageRepositoryReturns struct {
		result1 error
	}
	createPackageRepositoryReturnsOnCall map[int]struct {
		result1 error
	}
	DeletePackageRepositoryStub        func(*v1alpha1.PackageRepository) error
	deletePackageRepositoryMutex       sync.RWMutex
	deletePackageRepositoryArgsForCall []struct {
		arg1 *v1alpha1.PackageRepository
	}
	deletePackageRepositoryReturns struct {
		result1 error
	}
	deletePackageRepositoryReturnsOnCall map[int]struct {
		result1 error
	}
	GetAppCRStub        func(string, string) (*v1alpha1a.App, error)
	getAppCRMutex       sync.RWMutex
	getAppCRArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getAppCRReturns struct {
		result1 *v1alpha1a.App
		result2 error
	}
	getAppCRReturnsOnCall map[int]struct {
		result1 *v1alpha1a.App
		result2 error
	}
	GetClientStub        func() client.Client
	getClientMutex       sync.RWMutex
	getClientArgsForCall []struct {
	}
	getClientReturns struct {
		result1 client.Client
	}
	getClientReturnsOnCall map[int]struct {
		result1 client.Client
	}
	GetPackageInstallStub        func(string, string) (*v1alpha1.PackageInstall, error)
	getPackageInstallMutex       sync.RWMutex
	getPackageInstallArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getPackageInstallReturns struct {
		result1 *v1alpha1.PackageInstall
		result2 error
	}
	getPackageInstallReturnsOnCall map[int]struct {
		result1 *v1alpha1.PackageInstall
		result2 error
	}
	GetPackageMetadataByNameStub        func(string, string) (*v1alpha1b.PackageMetadata, error)
	getPackageMetadataByNameMutex       sync.RWMutex
	getPackageMetadataByNameArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getPackageMetadataByNameReturns struct {
		result1 *v1alpha1b.PackageMetadata
		result2 error
	}
	getPackageMetadataByNameReturnsOnCall map[int]struct {
		result1 *v1alpha1b.PackageMetadata
		result2 error
	}
	GetPackageRepositoryStub        func(string, string) (*v1alpha1.PackageRepository, error)
	getPackageRepositoryMutex       sync.RWMutex
	getPackageRepositoryArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getPackageRepositoryReturns struct {
		result1 *v1alpha1.PackageRepository
		result2 error
	}
	getPackageRepositoryReturnsOnCall map[int]struct {
		result1 *v1alpha1.PackageRepository
		result2 error
	}
	ListPackageInstallsStub        func(string) (*v1alpha1.PackageInstallList, error)
	listPackageInstallsMutex       sync.RWMutex
	listPackageInstallsArgsForCall []struct {
		arg1 string
	}
	listPackageInstallsReturns struct {
		result1 *v1alpha1.PackageInstallList
		result2 error
	}
	listPackageInstallsReturnsOnCall map[int]struct {
		result1 *v1alpha1.PackageInstallList
		result2 error
	}
	ListPackageMetadataStub        func(string) (*v1alpha1b.PackageMetadataList, error)
	listPackageMetadataMutex       sync.RWMutex
	listPackageMetadataArgsForCall []struct {
		arg1 string
	}
	listPackageMetadataReturns struct {
		result1 *v1alpha1b.PackageMetadataList
		result2 error
	}
	listPackageMetadataReturnsOnCall map[int]struct {
		result1 *v1alpha1b.PackageMetadataList
		result2 error
	}
	ListPackageRepositoriesStub        func(string) (*v1alpha1.PackageRepositoryList, error)
	listPackageRepositoriesMutex       sync.RWMutex
	listPackageRepositoriesArgsForCall []struct {
		arg1 string
	}
	listPackageRepositoriesReturns struct {
		result1 *v1alpha1.PackageRepositoryList
		result2 error
	}
	listPackageRepositoriesReturnsOnCall map[int]struct {
		result1 *v1alpha1.PackageRepositoryList
		result2 error
	}
	ListPackagesStub        func(string, string) (*v1alpha1b.PackageList, error)
	listPackagesMutex       sync.RWMutex
	listPackagesArgsForCall []struct {
		arg1 string
		arg2 string
	}
	listPackagesReturns struct {
		result1 *v1alpha1b.PackageList
		result2 error
	}
	listPackagesReturnsOnCall map[int]struct {
		result1 *v1alpha1b.PackageList
		result2 error
	}
	UpdatePackageRepositoryStub        func(*v1alpha1.PackageRepository) error
	updatePackageRepositoryMutex       sync.RWMutex
	updatePackageRepositoryArgsForCall []struct {
		arg1 *v1alpha1.PackageRepository
	}
	updatePackageRepositoryReturns struct {
		result1 error
	}
	updatePackageRepositoryReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *KappClient) CreatePackageInstall(arg1 *v1alpha1.PackageInstall, arg2 bool, arg3 bool) error {
	fake.createPackageInstallMutex.Lock()
	ret, specificReturn := fake.createPackageInstallReturnsOnCall[len(fake.createPackageInstallArgsForCall)]
	fake.createPackageInstallArgsForCall = append(fake.createPackageInstallArgsForCall, struct {
		arg1 *v1alpha1.PackageInstall
		arg2 bool
		arg3 bool
	}{arg1, arg2, arg3})
	stub := fake.CreatePackageInstallStub
	fakeReturns := fake.createPackageInstallReturns
	fake.recordInvocation("CreatePackageInstall", []interface{}{arg1, arg2, arg3})
	fake.createPackageInstallMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *KappClient) CreatePackageInstallCallCount() int {
	fake.createPackageInstallMutex.RLock()
	defer fake.createPackageInstallMutex.RUnlock()
	return len(fake.createPackageInstallArgsForCall)
}

func (fake *KappClient) CreatePackageInstallCalls(stub func(*v1alpha1.PackageInstall, bool, bool) error) {
	fake.createPackageInstallMutex.Lock()
	defer fake.createPackageInstallMutex.Unlock()
	fake.CreatePackageInstallStub = stub
}

func (fake *KappClient) CreatePackageInstallArgsForCall(i int) (*v1alpha1.PackageInstall, bool, bool) {
	fake.createPackageInstallMutex.RLock()
	defer fake.createPackageInstallMutex.RUnlock()
	argsForCall := fake.createPackageInstallArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *KappClient) CreatePackageInstallReturns(result1 error) {
	fake.createPackageInstallMutex.Lock()
	defer fake.createPackageInstallMutex.Unlock()
	fake.CreatePackageInstallStub = nil
	fake.createPackageInstallReturns = struct {
		result1 error
	}{result1}
}

func (fake *KappClient) CreatePackageInstallReturnsOnCall(i int, result1 error) {
	fake.createPackageInstallMutex.Lock()
	defer fake.createPackageInstallMutex.Unlock()
	fake.CreatePackageInstallStub = nil
	if fake.createPackageInstallReturnsOnCall == nil {
		fake.createPackageInstallReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createPackageInstallReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *KappClient) CreatePackageRepository(arg1 *v1alpha1.PackageRepository) error {
	fake.createPackageRepositoryMutex.Lock()
	ret, specificReturn := fake.createPackageRepositoryReturnsOnCall[len(fake.createPackageRepositoryArgsForCall)]
	fake.createPackageRepositoryArgsForCall = append(fake.createPackageRepositoryArgsForCall, struct {
		arg1 *v1alpha1.PackageRepository
	}{arg1})
	stub := fake.CreatePackageRepositoryStub
	fakeReturns := fake.createPackageRepositoryReturns
	fake.recordInvocation("CreatePackageRepository", []interface{}{arg1})
	fake.createPackageRepositoryMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *KappClient) CreatePackageRepositoryCallCount() int {
	fake.createPackageRepositoryMutex.RLock()
	defer fake.createPackageRepositoryMutex.RUnlock()
	return len(fake.createPackageRepositoryArgsForCall)
}

func (fake *KappClient) CreatePackageRepositoryCalls(stub func(*v1alpha1.PackageRepository) error) {
	fake.createPackageRepositoryMutex.Lock()
	defer fake.createPackageRepositoryMutex.Unlock()
	fake.CreatePackageRepositoryStub = stub
}

func (fake *KappClient) CreatePackageRepositoryArgsForCall(i int) *v1alpha1.PackageRepository {
	fake.createPackageRepositoryMutex.RLock()
	defer fake.createPackageRepositoryMutex.RUnlock()
	argsForCall := fake.createPackageRepositoryArgsForCall[i]
	return argsForCall.arg1
}

func (fake *KappClient) CreatePackageRepositoryReturns(result1 error) {
	fake.createPackageRepositoryMutex.Lock()
	defer fake.createPackageRepositoryMutex.Unlock()
	fake.CreatePackageRepositoryStub = nil
	fake.createPackageRepositoryReturns = struct {
		result1 error
	}{result1}
}

func (fake *KappClient) CreatePackageRepositoryReturnsOnCall(i int, result1 error) {
	fake.createPackageRepositoryMutex.Lock()
	defer fake.createPackageRepositoryMutex.Unlock()
	fake.CreatePackageRepositoryStub = nil
	if fake.createPackageRepositoryReturnsOnCall == nil {
		fake.createPackageRepositoryReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createPackageRepositoryReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *KappClient) DeletePackageRepository(arg1 *v1alpha1.PackageRepository) error {
	fake.deletePackageRepositoryMutex.Lock()
	ret, specificReturn := fake.deletePackageRepositoryReturnsOnCall[len(fake.deletePackageRepositoryArgsForCall)]
	fake.deletePackageRepositoryArgsForCall = append(fake.deletePackageRepositoryArgsForCall, struct {
		arg1 *v1alpha1.PackageRepository
	}{arg1})
	stub := fake.DeletePackageRepositoryStub
	fakeReturns := fake.deletePackageRepositoryReturns
	fake.recordInvocation("DeletePackageRepository", []interface{}{arg1})
	fake.deletePackageRepositoryMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *KappClient) DeletePackageRepositoryCallCount() int {
	fake.deletePackageRepositoryMutex.RLock()
	defer fake.deletePackageRepositoryMutex.RUnlock()
	return len(fake.deletePackageRepositoryArgsForCall)
}

func (fake *KappClient) DeletePackageRepositoryCalls(stub func(*v1alpha1.PackageRepository) error) {
	fake.deletePackageRepositoryMutex.Lock()
	defer fake.deletePackageRepositoryMutex.Unlock()
	fake.DeletePackageRepositoryStub = stub
}

func (fake *KappClient) DeletePackageRepositoryArgsForCall(i int) *v1alpha1.PackageRepository {
	fake.deletePackageRepositoryMutex.RLock()
	defer fake.deletePackageRepositoryMutex.RUnlock()
	argsForCall := fake.deletePackageRepositoryArgsForCall[i]
	return argsForCall.arg1
}

func (fake *KappClient) DeletePackageRepositoryReturns(result1 error) {
	fake.deletePackageRepositoryMutex.Lock()
	defer fake.deletePackageRepositoryMutex.Unlock()
	fake.DeletePackageRepositoryStub = nil
	fake.deletePackageRepositoryReturns = struct {
		result1 error
	}{result1}
}

func (fake *KappClient) DeletePackageRepositoryReturnsOnCall(i int, result1 error) {
	fake.deletePackageRepositoryMutex.Lock()
	defer fake.deletePackageRepositoryMutex.Unlock()
	fake.DeletePackageRepositoryStub = nil
	if fake.deletePackageRepositoryReturnsOnCall == nil {
		fake.deletePackageRepositoryReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deletePackageRepositoryReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *KappClient) GetAppCR(arg1 string, arg2 string) (*v1alpha1a.App, error) {
	fake.getAppCRMutex.Lock()
	ret, specificReturn := fake.getAppCRReturnsOnCall[len(fake.getAppCRArgsForCall)]
	fake.getAppCRArgsForCall = append(fake.getAppCRArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.GetAppCRStub
	fakeReturns := fake.getAppCRReturns
	fake.recordInvocation("GetAppCR", []interface{}{arg1, arg2})
	fake.getAppCRMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *KappClient) GetAppCRCallCount() int {
	fake.getAppCRMutex.RLock()
	defer fake.getAppCRMutex.RUnlock()
	return len(fake.getAppCRArgsForCall)
}

func (fake *KappClient) GetAppCRCalls(stub func(string, string) (*v1alpha1a.App, error)) {
	fake.getAppCRMutex.Lock()
	defer fake.getAppCRMutex.Unlock()
	fake.GetAppCRStub = stub
}

func (fake *KappClient) GetAppCRArgsForCall(i int) (string, string) {
	fake.getAppCRMutex.RLock()
	defer fake.getAppCRMutex.RUnlock()
	argsForCall := fake.getAppCRArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *KappClient) GetAppCRReturns(result1 *v1alpha1a.App, result2 error) {
	fake.getAppCRMutex.Lock()
	defer fake.getAppCRMutex.Unlock()
	fake.GetAppCRStub = nil
	fake.getAppCRReturns = struct {
		result1 *v1alpha1a.App
		result2 error
	}{result1, result2}
}

func (fake *KappClient) GetAppCRReturnsOnCall(i int, result1 *v1alpha1a.App, result2 error) {
	fake.getAppCRMutex.Lock()
	defer fake.getAppCRMutex.Unlock()
	fake.GetAppCRStub = nil
	if fake.getAppCRReturnsOnCall == nil {
		fake.getAppCRReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1a.App
			result2 error
		})
	}
	fake.getAppCRReturnsOnCall[i] = struct {
		result1 *v1alpha1a.App
		result2 error
	}{result1, result2}
}

func (fake *KappClient) GetClient() client.Client {
	fake.getClientMutex.Lock()
	ret, specificReturn := fake.getClientReturnsOnCall[len(fake.getClientArgsForCall)]
	fake.getClientArgsForCall = append(fake.getClientArgsForCall, struct {
	}{})
	stub := fake.GetClientStub
	fakeReturns := fake.getClientReturns
	fake.recordInvocation("GetClient", []interface{}{})
	fake.getClientMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *KappClient) GetClientCallCount() int {
	fake.getClientMutex.RLock()
	defer fake.getClientMutex.RUnlock()
	return len(fake.getClientArgsForCall)
}

func (fake *KappClient) GetClientCalls(stub func() client.Client) {
	fake.getClientMutex.Lock()
	defer fake.getClientMutex.Unlock()
	fake.GetClientStub = stub
}

func (fake *KappClient) GetClientReturns(result1 client.Client) {
	fake.getClientMutex.Lock()
	defer fake.getClientMutex.Unlock()
	fake.GetClientStub = nil
	fake.getClientReturns = struct {
		result1 client.Client
	}{result1}
}

func (fake *KappClient) GetClientReturnsOnCall(i int, result1 client.Client) {
	fake.getClientMutex.Lock()
	defer fake.getClientMutex.Unlock()
	fake.GetClientStub = nil
	if fake.getClientReturnsOnCall == nil {
		fake.getClientReturnsOnCall = make(map[int]struct {
			result1 client.Client
		})
	}
	fake.getClientReturnsOnCall[i] = struct {
		result1 client.Client
	}{result1}
}

func (fake *KappClient) GetPackageInstall(arg1 string, arg2 string) (*v1alpha1.PackageInstall, error) {
	fake.getPackageInstallMutex.Lock()
	ret, specificReturn := fake.getPackageInstallReturnsOnCall[len(fake.getPackageInstallArgsForCall)]
	fake.getPackageInstallArgsForCall = append(fake.getPackageInstallArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.GetPackageInstallStub
	fakeReturns := fake.getPackageInstallReturns
	fake.recordInvocation("GetPackageInstall", []interface{}{arg1, arg2})
	fake.getPackageInstallMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *KappClient) GetPackageInstallCallCount() int {
	fake.getPackageInstallMutex.RLock()
	defer fake.getPackageInstallMutex.RUnlock()
	return len(fake.getPackageInstallArgsForCall)
}

func (fake *KappClient) GetPackageInstallCalls(stub func(string, string) (*v1alpha1.PackageInstall, error)) {
	fake.getPackageInstallMutex.Lock()
	defer fake.getPackageInstallMutex.Unlock()
	fake.GetPackageInstallStub = stub
}

func (fake *KappClient) GetPackageInstallArgsForCall(i int) (string, string) {
	fake.getPackageInstallMutex.RLock()
	defer fake.getPackageInstallMutex.RUnlock()
	argsForCall := fake.getPackageInstallArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *KappClient) GetPackageInstallReturns(result1 *v1alpha1.PackageInstall, result2 error) {
	fake.getPackageInstallMutex.Lock()
	defer fake.getPackageInstallMutex.Unlock()
	fake.GetPackageInstallStub = nil
	fake.getPackageInstallReturns = struct {
		result1 *v1alpha1.PackageInstall
		result2 error
	}{result1, result2}
}

func (fake *KappClient) GetPackageInstallReturnsOnCall(i int, result1 *v1alpha1.PackageInstall, result2 error) {
	fake.getPackageInstallMutex.Lock()
	defer fake.getPackageInstallMutex.Unlock()
	fake.GetPackageInstallStub = nil
	if fake.getPackageInstallReturnsOnCall == nil {
		fake.getPackageInstallReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.PackageInstall
			result2 error
		})
	}
	fake.getPackageInstallReturnsOnCall[i] = struct {
		result1 *v1alpha1.PackageInstall
		result2 error
	}{result1, result2}
}

func (fake *KappClient) GetPackageMetadataByName(arg1 string, arg2 string) (*v1alpha1b.PackageMetadata, error) {
	fake.getPackageMetadataByNameMutex.Lock()
	ret, specificReturn := fake.getPackageMetadataByNameReturnsOnCall[len(fake.getPackageMetadataByNameArgsForCall)]
	fake.getPackageMetadataByNameArgsForCall = append(fake.getPackageMetadataByNameArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.GetPackageMetadataByNameStub
	fakeReturns := fake.getPackageMetadataByNameReturns
	fake.recordInvocation("GetPackageMetadataByName", []interface{}{arg1, arg2})
	fake.getPackageMetadataByNameMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *KappClient) GetPackageMetadataByNameCallCount() int {
	fake.getPackageMetadataByNameMutex.RLock()
	defer fake.getPackageMetadataByNameMutex.RUnlock()
	return len(fake.getPackageMetadataByNameArgsForCall)
}

func (fake *KappClient) GetPackageMetadataByNameCalls(stub func(string, string) (*v1alpha1b.PackageMetadata, error)) {
	fake.getPackageMetadataByNameMutex.Lock()
	defer fake.getPackageMetadataByNameMutex.Unlock()
	fake.GetPackageMetadataByNameStub = stub
}

func (fake *KappClient) GetPackageMetadataByNameArgsForCall(i int) (string, string) {
	fake.getPackageMetadataByNameMutex.RLock()
	defer fake.getPackageMetadataByNameMutex.RUnlock()
	argsForCall := fake.getPackageMetadataByNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *KappClient) GetPackageMetadataByNameReturns(result1 *v1alpha1b.PackageMetadata, result2 error) {
	fake.getPackageMetadataByNameMutex.Lock()
	defer fake.getPackageMetadataByNameMutex.Unlock()
	fake.GetPackageMetadataByNameStub = nil
	fake.getPackageMetadataByNameReturns = struct {
		result1 *v1alpha1b.PackageMetadata
		result2 error
	}{result1, result2}
}

func (fake *KappClient) GetPackageMetadataByNameReturnsOnCall(i int, result1 *v1alpha1b.PackageMetadata, result2 error) {
	fake.getPackageMetadataByNameMutex.Lock()
	defer fake.getPackageMetadataByNameMutex.Unlock()
	fake.GetPackageMetadataByNameStub = nil
	if fake.getPackageMetadataByNameReturnsOnCall == nil {
		fake.getPackageMetadataByNameReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1b.PackageMetadata
			result2 error
		})
	}
	fake.getPackageMetadataByNameReturnsOnCall[i] = struct {
		result1 *v1alpha1b.PackageMetadata
		result2 error
	}{result1, result2}
}

func (fake *KappClient) GetPackageRepository(arg1 string, arg2 string) (*v1alpha1.PackageRepository, error) {
	fake.getPackageRepositoryMutex.Lock()
	ret, specificReturn := fake.getPackageRepositoryReturnsOnCall[len(fake.getPackageRepositoryArgsForCall)]
	fake.getPackageRepositoryArgsForCall = append(fake.getPackageRepositoryArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.GetPackageRepositoryStub
	fakeReturns := fake.getPackageRepositoryReturns
	fake.recordInvocation("GetPackageRepository", []interface{}{arg1, arg2})
	fake.getPackageRepositoryMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *KappClient) GetPackageRepositoryCallCount() int {
	fake.getPackageRepositoryMutex.RLock()
	defer fake.getPackageRepositoryMutex.RUnlock()
	return len(fake.getPackageRepositoryArgsForCall)
}

func (fake *KappClient) GetPackageRepositoryCalls(stub func(string, string) (*v1alpha1.PackageRepository, error)) {
	fake.getPackageRepositoryMutex.Lock()
	defer fake.getPackageRepositoryMutex.Unlock()
	fake.GetPackageRepositoryStub = stub
}

func (fake *KappClient) GetPackageRepositoryArgsForCall(i int) (string, string) {
	fake.getPackageRepositoryMutex.RLock()
	defer fake.getPackageRepositoryMutex.RUnlock()
	argsForCall := fake.getPackageRepositoryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *KappClient) GetPackageRepositoryReturns(result1 *v1alpha1.PackageRepository, result2 error) {
	fake.getPackageRepositoryMutex.Lock()
	defer fake.getPackageRepositoryMutex.Unlock()
	fake.GetPackageRepositoryStub = nil
	fake.getPackageRepositoryReturns = struct {
		result1 *v1alpha1.PackageRepository
		result2 error
	}{result1, result2}
}

func (fake *KappClient) GetPackageRepositoryReturnsOnCall(i int, result1 *v1alpha1.PackageRepository, result2 error) {
	fake.getPackageRepositoryMutex.Lock()
	defer fake.getPackageRepositoryMutex.Unlock()
	fake.GetPackageRepositoryStub = nil
	if fake.getPackageRepositoryReturnsOnCall == nil {
		fake.getPackageRepositoryReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.PackageRepository
			result2 error
		})
	}
	fake.getPackageRepositoryReturnsOnCall[i] = struct {
		result1 *v1alpha1.PackageRepository
		result2 error
	}{result1, result2}
}

func (fake *KappClient) ListPackageInstalls(arg1 string) (*v1alpha1.PackageInstallList, error) {
	fake.listPackageInstallsMutex.Lock()
	ret, specificReturn := fake.listPackageInstallsReturnsOnCall[len(fake.listPackageInstallsArgsForCall)]
	fake.listPackageInstallsArgsForCall = append(fake.listPackageInstallsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ListPackageInstallsStub
	fakeReturns := fake.listPackageInstallsReturns
	fake.recordInvocation("ListPackageInstalls", []interface{}{arg1})
	fake.listPackageInstallsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *KappClient) ListPackageInstallsCallCount() int {
	fake.listPackageInstallsMutex.RLock()
	defer fake.listPackageInstallsMutex.RUnlock()
	return len(fake.listPackageInstallsArgsForCall)
}

func (fake *KappClient) ListPackageInstallsCalls(stub func(string) (*v1alpha1.PackageInstallList, error)) {
	fake.listPackageInstallsMutex.Lock()
	defer fake.listPackageInstallsMutex.Unlock()
	fake.ListPackageInstallsStub = stub
}

func (fake *KappClient) ListPackageInstallsArgsForCall(i int) string {
	fake.listPackageInstallsMutex.RLock()
	defer fake.listPackageInstallsMutex.RUnlock()
	argsForCall := fake.listPackageInstallsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *KappClient) ListPackageInstallsReturns(result1 *v1alpha1.PackageInstallList, result2 error) {
	fake.listPackageInstallsMutex.Lock()
	defer fake.listPackageInstallsMutex.Unlock()
	fake.ListPackageInstallsStub = nil
	fake.listPackageInstallsReturns = struct {
		result1 *v1alpha1.PackageInstallList
		result2 error
	}{result1, result2}
}

func (fake *KappClient) ListPackageInstallsReturnsOnCall(i int, result1 *v1alpha1.PackageInstallList, result2 error) {
	fake.listPackageInstallsMutex.Lock()
	defer fake.listPackageInstallsMutex.Unlock()
	fake.ListPackageInstallsStub = nil
	if fake.listPackageInstallsReturnsOnCall == nil {
		fake.listPackageInstallsReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.PackageInstallList
			result2 error
		})
	}
	fake.listPackageInstallsReturnsOnCall[i] = struct {
		result1 *v1alpha1.PackageInstallList
		result2 error
	}{result1, result2}
}

func (fake *KappClient) ListPackageMetadata(arg1 string) (*v1alpha1b.PackageMetadataList, error) {
	fake.listPackageMetadataMutex.Lock()
	ret, specificReturn := fake.listPackageMetadataReturnsOnCall[len(fake.listPackageMetadataArgsForCall)]
	fake.listPackageMetadataArgsForCall = append(fake.listPackageMetadataArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ListPackageMetadataStub
	fakeReturns := fake.listPackageMetadataReturns
	fake.recordInvocation("ListPackageMetadata", []interface{}{arg1})
	fake.listPackageMetadataMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *KappClient) ListPackageMetadataCallCount() int {
	fake.listPackageMetadataMutex.RLock()
	defer fake.listPackageMetadataMutex.RUnlock()
	return len(fake.listPackageMetadataArgsForCall)
}

func (fake *KappClient) ListPackageMetadataCalls(stub func(string) (*v1alpha1b.PackageMetadataList, error)) {
	fake.listPackageMetadataMutex.Lock()
	defer fake.listPackageMetadataMutex.Unlock()
	fake.ListPackageMetadataStub = stub
}

func (fake *KappClient) ListPackageMetadataArgsForCall(i int) string {
	fake.listPackageMetadataMutex.RLock()
	defer fake.listPackageMetadataMutex.RUnlock()
	argsForCall := fake.listPackageMetadataArgsForCall[i]
	return argsForCall.arg1
}

func (fake *KappClient) ListPackageMetadataReturns(result1 *v1alpha1b.PackageMetadataList, result2 error) {
	fake.listPackageMetadataMutex.Lock()
	defer fake.listPackageMetadataMutex.Unlock()
	fake.ListPackageMetadataStub = nil
	fake.listPackageMetadataReturns = struct {
		result1 *v1alpha1b.PackageMetadataList
		result2 error
	}{result1, result2}
}

func (fake *KappClient) ListPackageMetadataReturnsOnCall(i int, result1 *v1alpha1b.PackageMetadataList, result2 error) {
	fake.listPackageMetadataMutex.Lock()
	defer fake.listPackageMetadataMutex.Unlock()
	fake.ListPackageMetadataStub = nil
	if fake.listPackageMetadataReturnsOnCall == nil {
		fake.listPackageMetadataReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1b.PackageMetadataList
			result2 error
		})
	}
	fake.listPackageMetadataReturnsOnCall[i] = struct {
		result1 *v1alpha1b.PackageMetadataList
		result2 error
	}{result1, result2}
}

func (fake *KappClient) ListPackageRepositories(arg1 string) (*v1alpha1.PackageRepositoryList, error) {
	fake.listPackageRepositoriesMutex.Lock()
	ret, specificReturn := fake.listPackageRepositoriesReturnsOnCall[len(fake.listPackageRepositoriesArgsForCall)]
	fake.listPackageRepositoriesArgsForCall = append(fake.listPackageRepositoriesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ListPackageRepositoriesStub
	fakeReturns := fake.listPackageRepositoriesReturns
	fake.recordInvocation("ListPackageRepositories", []interface{}{arg1})
	fake.listPackageRepositoriesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *KappClient) ListPackageRepositoriesCallCount() int {
	fake.listPackageRepositoriesMutex.RLock()
	defer fake.listPackageRepositoriesMutex.RUnlock()
	return len(fake.listPackageRepositoriesArgsForCall)
}

func (fake *KappClient) ListPackageRepositoriesCalls(stub func(string) (*v1alpha1.PackageRepositoryList, error)) {
	fake.listPackageRepositoriesMutex.Lock()
	defer fake.listPackageRepositoriesMutex.Unlock()
	fake.ListPackageRepositoriesStub = stub
}

func (fake *KappClient) ListPackageRepositoriesArgsForCall(i int) string {
	fake.listPackageRepositoriesMutex.RLock()
	defer fake.listPackageRepositoriesMutex.RUnlock()
	argsForCall := fake.listPackageRepositoriesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *KappClient) ListPackageRepositoriesReturns(result1 *v1alpha1.PackageRepositoryList, result2 error) {
	fake.listPackageRepositoriesMutex.Lock()
	defer fake.listPackageRepositoriesMutex.Unlock()
	fake.ListPackageRepositoriesStub = nil
	fake.listPackageRepositoriesReturns = struct {
		result1 *v1alpha1.PackageRepositoryList
		result2 error
	}{result1, result2}
}

func (fake *KappClient) ListPackageRepositoriesReturnsOnCall(i int, result1 *v1alpha1.PackageRepositoryList, result2 error) {
	fake.listPackageRepositoriesMutex.Lock()
	defer fake.listPackageRepositoriesMutex.Unlock()
	fake.ListPackageRepositoriesStub = nil
	if fake.listPackageRepositoriesReturnsOnCall == nil {
		fake.listPackageRepositoriesReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.PackageRepositoryList
			result2 error
		})
	}
	fake.listPackageRepositoriesReturnsOnCall[i] = struct {
		result1 *v1alpha1.PackageRepositoryList
		result2 error
	}{result1, result2}
}

func (fake *KappClient) ListPackages(arg1 string, arg2 string) (*v1alpha1b.PackageList, error) {
	fake.listPackagesMutex.Lock()
	ret, specificReturn := fake.listPackagesReturnsOnCall[len(fake.listPackagesArgsForCall)]
	fake.listPackagesArgsForCall = append(fake.listPackagesArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.ListPackagesStub
	fakeReturns := fake.listPackagesReturns
	fake.recordInvocation("ListPackages", []interface{}{arg1, arg2})
	fake.listPackagesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *KappClient) ListPackagesCallCount() int {
	fake.listPackagesMutex.RLock()
	defer fake.listPackagesMutex.RUnlock()
	return len(fake.listPackagesArgsForCall)
}

func (fake *KappClient) ListPackagesCalls(stub func(string, string) (*v1alpha1b.PackageList, error)) {
	fake.listPackagesMutex.Lock()
	defer fake.listPackagesMutex.Unlock()
	fake.ListPackagesStub = stub
}

func (fake *KappClient) ListPackagesArgsForCall(i int) (string, string) {
	fake.listPackagesMutex.RLock()
	defer fake.listPackagesMutex.RUnlock()
	argsForCall := fake.listPackagesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *KappClient) ListPackagesReturns(result1 *v1alpha1b.PackageList, result2 error) {
	fake.listPackagesMutex.Lock()
	defer fake.listPackagesMutex.Unlock()
	fake.ListPackagesStub = nil
	fake.listPackagesReturns = struct {
		result1 *v1alpha1b.PackageList
		result2 error
	}{result1, result2}
}

func (fake *KappClient) ListPackagesReturnsOnCall(i int, result1 *v1alpha1b.PackageList, result2 error) {
	fake.listPackagesMutex.Lock()
	defer fake.listPackagesMutex.Unlock()
	fake.ListPackagesStub = nil
	if fake.listPackagesReturnsOnCall == nil {
		fake.listPackagesReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1b.PackageList
			result2 error
		})
	}
	fake.listPackagesReturnsOnCall[i] = struct {
		result1 *v1alpha1b.PackageList
		result2 error
	}{result1, result2}
}

func (fake *KappClient) UpdatePackageRepository(arg1 *v1alpha1.PackageRepository) error {
	fake.updatePackageRepositoryMutex.Lock()
	ret, specificReturn := fake.updatePackageRepositoryReturnsOnCall[len(fake.updatePackageRepositoryArgsForCall)]
	fake.updatePackageRepositoryArgsForCall = append(fake.updatePackageRepositoryArgsForCall, struct {
		arg1 *v1alpha1.PackageRepository
	}{arg1})
	stub := fake.UpdatePackageRepositoryStub
	fakeReturns := fake.updatePackageRepositoryReturns
	fake.recordInvocation("UpdatePackageRepository", []interface{}{arg1})
	fake.updatePackageRepositoryMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *KappClient) UpdatePackageRepositoryCallCount() int {
	fake.updatePackageRepositoryMutex.RLock()
	defer fake.updatePackageRepositoryMutex.RUnlock()
	return len(fake.updatePackageRepositoryArgsForCall)
}

func (fake *KappClient) UpdatePackageRepositoryCalls(stub func(*v1alpha1.PackageRepository) error) {
	fake.updatePackageRepositoryMutex.Lock()
	defer fake.updatePackageRepositoryMutex.Unlock()
	fake.UpdatePackageRepositoryStub = stub
}

func (fake *KappClient) UpdatePackageRepositoryArgsForCall(i int) *v1alpha1.PackageRepository {
	fake.updatePackageRepositoryMutex.RLock()
	defer fake.updatePackageRepositoryMutex.RUnlock()
	argsForCall := fake.updatePackageRepositoryArgsForCall[i]
	return argsForCall.arg1
}

func (fake *KappClient) UpdatePackageRepositoryReturns(result1 error) {
	fake.updatePackageRepositoryMutex.Lock()
	defer fake.updatePackageRepositoryMutex.Unlock()
	fake.UpdatePackageRepositoryStub = nil
	fake.updatePackageRepositoryReturns = struct {
		result1 error
	}{result1}
}

func (fake *KappClient) UpdatePackageRepositoryReturnsOnCall(i int, result1 error) {
	fake.updatePackageRepositoryMutex.Lock()
	defer fake.updatePackageRepositoryMutex.Unlock()
	fake.UpdatePackageRepositoryStub = nil
	if fake.updatePackageRepositoryReturnsOnCall == nil {
		fake.updatePackageRepositoryReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updatePackageRepositoryReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *KappClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createPackageInstallMutex.RLock()
	defer fake.createPackageInstallMutex.RUnlock()
	fake.createPackageRepositoryMutex.RLock()
	defer fake.createPackageRepositoryMutex.RUnlock()
	fake.deletePackageRepositoryMutex.RLock()
	defer fake.deletePackageRepositoryMutex.RUnlock()
	fake.getAppCRMutex.RLock()
	defer fake.getAppCRMutex.RUnlock()
	fake.getClientMutex.RLock()
	defer fake.getClientMutex.RUnlock()
	fake.getPackageInstallMutex.RLock()
	defer fake.getPackageInstallMutex.RUnlock()
	fake.getPackageMetadataByNameMutex.RLock()
	defer fake.getPackageMetadataByNameMutex.RUnlock()
	fake.getPackageRepositoryMutex.RLock()
	defer fake.getPackageRepositoryMutex.RUnlock()
	fake.listPackageInstallsMutex.RLock()
	defer fake.listPackageInstallsMutex.RUnlock()
	fake.listPackageMetadataMutex.RLock()
	defer fake.listPackageMetadataMutex.RUnlock()
	fake.listPackageRepositoriesMutex.RLock()
	defer fake.listPackageRepositoriesMutex.RUnlock()
	fake.listPackagesMutex.RLock()
	defer fake.listPackagesMutex.RUnlock()
	fake.updatePackageRepositoryMutex.RLock()
	defer fake.updatePackageRepositoryMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *KappClient) recordInvocation(key string, args []interface{}) {
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

var _ kappclient.Client = new(KappClient)
