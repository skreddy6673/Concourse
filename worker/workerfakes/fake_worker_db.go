// This file was generated by counterfeiter
package workerfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/worker"
)

type FakeWorkerDB struct {
	WorkersStub        func() ([]db.SavedWorker, error)
	workersMutex       sync.RWMutex
	workersArgsForCall []struct{}
	workersReturns     struct {
		result1 []db.SavedWorker
		result2 error
	}
	GetWorkerStub        func(string) (db.SavedWorker, bool, error)
	getWorkerMutex       sync.RWMutex
	getWorkerArgsForCall []struct {
		arg1 string
	}
	getWorkerReturns struct {
		result1 db.SavedWorker
		result2 bool
		result3 error
	}
	CreateContainerStub        func(container db.Container, ttl time.Duration, maxLifetime time.Duration, volumeHandles []string) (db.SavedContainer, error)
	createContainerMutex       sync.RWMutex
	createContainerArgsForCall []struct {
		container     db.Container
		ttl           time.Duration
		maxLifetime   time.Duration
		volumeHandles []string
	}
	createContainerReturns struct {
		result1 db.SavedContainer
		result2 error
	}
	UpdateContainerTTLToBeRemovedStub        func(container db.Container, ttl time.Duration, maxLifetime time.Duration) (db.SavedContainer, error)
	updateContainerTTLToBeRemovedMutex       sync.RWMutex
	updateContainerTTLToBeRemovedArgsForCall []struct {
		container   db.Container
		ttl         time.Duration
		maxLifetime time.Duration
	}
	updateContainerTTLToBeRemovedReturns struct {
		result1 db.SavedContainer
		result2 error
	}
	GetContainerStub        func(string) (db.SavedContainer, bool, error)
	getContainerMutex       sync.RWMutex
	getContainerArgsForCall []struct {
		arg1 string
	}
	getContainerReturns struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}
	FindContainerByIdentifierStub        func(db.ContainerIdentifier) (db.SavedContainer, bool, error)
	findContainerByIdentifierMutex       sync.RWMutex
	findContainerByIdentifierArgsForCall []struct {
		arg1 db.ContainerIdentifier
	}
	findContainerByIdentifierReturns struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}
	UpdateExpiresAtOnContainerStub        func(handle string, ttl time.Duration) error
	updateExpiresAtOnContainerMutex       sync.RWMutex
	updateExpiresAtOnContainerArgsForCall []struct {
		handle string
		ttl    time.Duration
	}
	updateExpiresAtOnContainerReturns struct {
		result1 error
	}
	ReapContainerStub        func(handle string) error
	reapContainerMutex       sync.RWMutex
	reapContainerArgsForCall []struct {
		handle string
	}
	reapContainerReturns struct {
		result1 error
	}
	GetPipelineByIDStub        func(pipelineID int) (db.SavedPipeline, error)
	getPipelineByIDMutex       sync.RWMutex
	getPipelineByIDArgsForCall []struct {
		pipelineID int
	}
	getPipelineByIDReturns struct {
		result1 db.SavedPipeline
		result2 error
	}
	GetVolumeTTLStub        func(volumeHandle string) (time.Duration, bool, error)
	getVolumeTTLMutex       sync.RWMutex
	getVolumeTTLArgsForCall []struct {
		volumeHandle string
	}
	getVolumeTTLReturns struct {
		result1 time.Duration
		result2 bool
		result3 error
	}
	ReapVolumeStub        func(handle string) error
	reapVolumeMutex       sync.RWMutex
	reapVolumeArgsForCall []struct {
		handle string
	}
	reapVolumeReturns struct {
		result1 error
	}
	SetVolumeTTLStub        func(string, time.Duration) error
	setVolumeTTLMutex       sync.RWMutex
	setVolumeTTLArgsForCall []struct {
		arg1 string
		arg2 time.Duration
	}
	setVolumeTTLReturns struct {
		result1 error
	}
	AcquireVolumeCreatingLockStub        func(lager.Logger, int) (db.Lock, bool, error)
	acquireVolumeCreatingLockMutex       sync.RWMutex
	acquireVolumeCreatingLockArgsForCall []struct {
		arg1 lager.Logger
		arg2 int
	}
	acquireVolumeCreatingLockReturns struct {
		result1 db.Lock
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWorkerDB) Workers() ([]db.SavedWorker, error) {
	fake.workersMutex.Lock()
	fake.workersArgsForCall = append(fake.workersArgsForCall, struct{}{})
	fake.recordInvocation("Workers", []interface{}{})
	fake.workersMutex.Unlock()
	if fake.WorkersStub != nil {
		return fake.WorkersStub()
	} else {
		return fake.workersReturns.result1, fake.workersReturns.result2
	}
}

func (fake *FakeWorkerDB) WorkersCallCount() int {
	fake.workersMutex.RLock()
	defer fake.workersMutex.RUnlock()
	return len(fake.workersArgsForCall)
}

func (fake *FakeWorkerDB) WorkersReturns(result1 []db.SavedWorker, result2 error) {
	fake.WorkersStub = nil
	fake.workersReturns = struct {
		result1 []db.SavedWorker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerDB) GetWorker(arg1 string) (db.SavedWorker, bool, error) {
	fake.getWorkerMutex.Lock()
	fake.getWorkerArgsForCall = append(fake.getWorkerArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetWorker", []interface{}{arg1})
	fake.getWorkerMutex.Unlock()
	if fake.GetWorkerStub != nil {
		return fake.GetWorkerStub(arg1)
	} else {
		return fake.getWorkerReturns.result1, fake.getWorkerReturns.result2, fake.getWorkerReturns.result3
	}
}

func (fake *FakeWorkerDB) GetWorkerCallCount() int {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return len(fake.getWorkerArgsForCall)
}

func (fake *FakeWorkerDB) GetWorkerArgsForCall(i int) string {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return fake.getWorkerArgsForCall[i].arg1
}

func (fake *FakeWorkerDB) GetWorkerReturns(result1 db.SavedWorker, result2 bool, result3 error) {
	fake.GetWorkerStub = nil
	fake.getWorkerReturns = struct {
		result1 db.SavedWorker
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerDB) CreateContainer(container db.Container, ttl time.Duration, maxLifetime time.Duration, volumeHandles []string) (db.SavedContainer, error) {
	var volumeHandlesCopy []string
	if volumeHandles != nil {
		volumeHandlesCopy = make([]string, len(volumeHandles))
		copy(volumeHandlesCopy, volumeHandles)
	}
	fake.createContainerMutex.Lock()
	fake.createContainerArgsForCall = append(fake.createContainerArgsForCall, struct {
		container     db.Container
		ttl           time.Duration
		maxLifetime   time.Duration
		volumeHandles []string
	}{container, ttl, maxLifetime, volumeHandlesCopy})
	fake.recordInvocation("CreateContainer", []interface{}{container, ttl, maxLifetime, volumeHandlesCopy})
	fake.createContainerMutex.Unlock()
	if fake.CreateContainerStub != nil {
		return fake.CreateContainerStub(container, ttl, maxLifetime, volumeHandles)
	} else {
		return fake.createContainerReturns.result1, fake.createContainerReturns.result2
	}
}

func (fake *FakeWorkerDB) CreateContainerCallCount() int {
	fake.createContainerMutex.RLock()
	defer fake.createContainerMutex.RUnlock()
	return len(fake.createContainerArgsForCall)
}

func (fake *FakeWorkerDB) CreateContainerArgsForCall(i int) (db.Container, time.Duration, time.Duration, []string) {
	fake.createContainerMutex.RLock()
	defer fake.createContainerMutex.RUnlock()
	return fake.createContainerArgsForCall[i].container, fake.createContainerArgsForCall[i].ttl, fake.createContainerArgsForCall[i].maxLifetime, fake.createContainerArgsForCall[i].volumeHandles
}

func (fake *FakeWorkerDB) CreateContainerReturns(result1 db.SavedContainer, result2 error) {
	fake.CreateContainerStub = nil
	fake.createContainerReturns = struct {
		result1 db.SavedContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerDB) UpdateContainerTTLToBeRemoved(container db.Container, ttl time.Duration, maxLifetime time.Duration) (db.SavedContainer, error) {
	fake.updateContainerTTLToBeRemovedMutex.Lock()
	fake.updateContainerTTLToBeRemovedArgsForCall = append(fake.updateContainerTTLToBeRemovedArgsForCall, struct {
		container   db.Container
		ttl         time.Duration
		maxLifetime time.Duration
	}{container, ttl, maxLifetime})
	fake.recordInvocation("UpdateContainerTTLToBeRemoved", []interface{}{container, ttl, maxLifetime})
	fake.updateContainerTTLToBeRemovedMutex.Unlock()
	if fake.UpdateContainerTTLToBeRemovedStub != nil {
		return fake.UpdateContainerTTLToBeRemovedStub(container, ttl, maxLifetime)
	} else {
		return fake.updateContainerTTLToBeRemovedReturns.result1, fake.updateContainerTTLToBeRemovedReturns.result2
	}
}

func (fake *FakeWorkerDB) UpdateContainerTTLToBeRemovedCallCount() int {
	fake.updateContainerTTLToBeRemovedMutex.RLock()
	defer fake.updateContainerTTLToBeRemovedMutex.RUnlock()
	return len(fake.updateContainerTTLToBeRemovedArgsForCall)
}

func (fake *FakeWorkerDB) UpdateContainerTTLToBeRemovedArgsForCall(i int) (db.Container, time.Duration, time.Duration) {
	fake.updateContainerTTLToBeRemovedMutex.RLock()
	defer fake.updateContainerTTLToBeRemovedMutex.RUnlock()
	return fake.updateContainerTTLToBeRemovedArgsForCall[i].container, fake.updateContainerTTLToBeRemovedArgsForCall[i].ttl, fake.updateContainerTTLToBeRemovedArgsForCall[i].maxLifetime
}

func (fake *FakeWorkerDB) UpdateContainerTTLToBeRemovedReturns(result1 db.SavedContainer, result2 error) {
	fake.UpdateContainerTTLToBeRemovedStub = nil
	fake.updateContainerTTLToBeRemovedReturns = struct {
		result1 db.SavedContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerDB) GetContainer(arg1 string) (db.SavedContainer, bool, error) {
	fake.getContainerMutex.Lock()
	fake.getContainerArgsForCall = append(fake.getContainerArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetContainer", []interface{}{arg1})
	fake.getContainerMutex.Unlock()
	if fake.GetContainerStub != nil {
		return fake.GetContainerStub(arg1)
	} else {
		return fake.getContainerReturns.result1, fake.getContainerReturns.result2, fake.getContainerReturns.result3
	}
}

func (fake *FakeWorkerDB) GetContainerCallCount() int {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return len(fake.getContainerArgsForCall)
}

func (fake *FakeWorkerDB) GetContainerArgsForCall(i int) string {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return fake.getContainerArgsForCall[i].arg1
}

func (fake *FakeWorkerDB) GetContainerReturns(result1 db.SavedContainer, result2 bool, result3 error) {
	fake.GetContainerStub = nil
	fake.getContainerReturns = struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerDB) FindContainerByIdentifier(arg1 db.ContainerIdentifier) (db.SavedContainer, bool, error) {
	fake.findContainerByIdentifierMutex.Lock()
	fake.findContainerByIdentifierArgsForCall = append(fake.findContainerByIdentifierArgsForCall, struct {
		arg1 db.ContainerIdentifier
	}{arg1})
	fake.recordInvocation("FindContainerByIdentifier", []interface{}{arg1})
	fake.findContainerByIdentifierMutex.Unlock()
	if fake.FindContainerByIdentifierStub != nil {
		return fake.FindContainerByIdentifierStub(arg1)
	} else {
		return fake.findContainerByIdentifierReturns.result1, fake.findContainerByIdentifierReturns.result2, fake.findContainerByIdentifierReturns.result3
	}
}

func (fake *FakeWorkerDB) FindContainerByIdentifierCallCount() int {
	fake.findContainerByIdentifierMutex.RLock()
	defer fake.findContainerByIdentifierMutex.RUnlock()
	return len(fake.findContainerByIdentifierArgsForCall)
}

func (fake *FakeWorkerDB) FindContainerByIdentifierArgsForCall(i int) db.ContainerIdentifier {
	fake.findContainerByIdentifierMutex.RLock()
	defer fake.findContainerByIdentifierMutex.RUnlock()
	return fake.findContainerByIdentifierArgsForCall[i].arg1
}

func (fake *FakeWorkerDB) FindContainerByIdentifierReturns(result1 db.SavedContainer, result2 bool, result3 error) {
	fake.FindContainerByIdentifierStub = nil
	fake.findContainerByIdentifierReturns = struct {
		result1 db.SavedContainer
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerDB) UpdateExpiresAtOnContainer(handle string, ttl time.Duration) error {
	fake.updateExpiresAtOnContainerMutex.Lock()
	fake.updateExpiresAtOnContainerArgsForCall = append(fake.updateExpiresAtOnContainerArgsForCall, struct {
		handle string
		ttl    time.Duration
	}{handle, ttl})
	fake.recordInvocation("UpdateExpiresAtOnContainer", []interface{}{handle, ttl})
	fake.updateExpiresAtOnContainerMutex.Unlock()
	if fake.UpdateExpiresAtOnContainerStub != nil {
		return fake.UpdateExpiresAtOnContainerStub(handle, ttl)
	} else {
		return fake.updateExpiresAtOnContainerReturns.result1
	}
}

func (fake *FakeWorkerDB) UpdateExpiresAtOnContainerCallCount() int {
	fake.updateExpiresAtOnContainerMutex.RLock()
	defer fake.updateExpiresAtOnContainerMutex.RUnlock()
	return len(fake.updateExpiresAtOnContainerArgsForCall)
}

func (fake *FakeWorkerDB) UpdateExpiresAtOnContainerArgsForCall(i int) (string, time.Duration) {
	fake.updateExpiresAtOnContainerMutex.RLock()
	defer fake.updateExpiresAtOnContainerMutex.RUnlock()
	return fake.updateExpiresAtOnContainerArgsForCall[i].handle, fake.updateExpiresAtOnContainerArgsForCall[i].ttl
}

func (fake *FakeWorkerDB) UpdateExpiresAtOnContainerReturns(result1 error) {
	fake.UpdateExpiresAtOnContainerStub = nil
	fake.updateExpiresAtOnContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerDB) ReapContainer(handle string) error {
	fake.reapContainerMutex.Lock()
	fake.reapContainerArgsForCall = append(fake.reapContainerArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("ReapContainer", []interface{}{handle})
	fake.reapContainerMutex.Unlock()
	if fake.ReapContainerStub != nil {
		return fake.ReapContainerStub(handle)
	} else {
		return fake.reapContainerReturns.result1
	}
}

func (fake *FakeWorkerDB) ReapContainerCallCount() int {
	fake.reapContainerMutex.RLock()
	defer fake.reapContainerMutex.RUnlock()
	return len(fake.reapContainerArgsForCall)
}

func (fake *FakeWorkerDB) ReapContainerArgsForCall(i int) string {
	fake.reapContainerMutex.RLock()
	defer fake.reapContainerMutex.RUnlock()
	return fake.reapContainerArgsForCall[i].handle
}

func (fake *FakeWorkerDB) ReapContainerReturns(result1 error) {
	fake.ReapContainerStub = nil
	fake.reapContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerDB) GetPipelineByID(pipelineID int) (db.SavedPipeline, error) {
	fake.getPipelineByIDMutex.Lock()
	fake.getPipelineByIDArgsForCall = append(fake.getPipelineByIDArgsForCall, struct {
		pipelineID int
	}{pipelineID})
	fake.recordInvocation("GetPipelineByID", []interface{}{pipelineID})
	fake.getPipelineByIDMutex.Unlock()
	if fake.GetPipelineByIDStub != nil {
		return fake.GetPipelineByIDStub(pipelineID)
	} else {
		return fake.getPipelineByIDReturns.result1, fake.getPipelineByIDReturns.result2
	}
}

func (fake *FakeWorkerDB) GetPipelineByIDCallCount() int {
	fake.getPipelineByIDMutex.RLock()
	defer fake.getPipelineByIDMutex.RUnlock()
	return len(fake.getPipelineByIDArgsForCall)
}

func (fake *FakeWorkerDB) GetPipelineByIDArgsForCall(i int) int {
	fake.getPipelineByIDMutex.RLock()
	defer fake.getPipelineByIDMutex.RUnlock()
	return fake.getPipelineByIDArgsForCall[i].pipelineID
}

func (fake *FakeWorkerDB) GetPipelineByIDReturns(result1 db.SavedPipeline, result2 error) {
	fake.GetPipelineByIDStub = nil
	fake.getPipelineByIDReturns = struct {
		result1 db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerDB) GetVolumeTTL(volumeHandle string) (time.Duration, bool, error) {
	fake.getVolumeTTLMutex.Lock()
	fake.getVolumeTTLArgsForCall = append(fake.getVolumeTTLArgsForCall, struct {
		volumeHandle string
	}{volumeHandle})
	fake.recordInvocation("GetVolumeTTL", []interface{}{volumeHandle})
	fake.getVolumeTTLMutex.Unlock()
	if fake.GetVolumeTTLStub != nil {
		return fake.GetVolumeTTLStub(volumeHandle)
	} else {
		return fake.getVolumeTTLReturns.result1, fake.getVolumeTTLReturns.result2, fake.getVolumeTTLReturns.result3
	}
}

func (fake *FakeWorkerDB) GetVolumeTTLCallCount() int {
	fake.getVolumeTTLMutex.RLock()
	defer fake.getVolumeTTLMutex.RUnlock()
	return len(fake.getVolumeTTLArgsForCall)
}

func (fake *FakeWorkerDB) GetVolumeTTLArgsForCall(i int) string {
	fake.getVolumeTTLMutex.RLock()
	defer fake.getVolumeTTLMutex.RUnlock()
	return fake.getVolumeTTLArgsForCall[i].volumeHandle
}

func (fake *FakeWorkerDB) GetVolumeTTLReturns(result1 time.Duration, result2 bool, result3 error) {
	fake.GetVolumeTTLStub = nil
	fake.getVolumeTTLReturns = struct {
		result1 time.Duration
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerDB) ReapVolume(handle string) error {
	fake.reapVolumeMutex.Lock()
	fake.reapVolumeArgsForCall = append(fake.reapVolumeArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("ReapVolume", []interface{}{handle})
	fake.reapVolumeMutex.Unlock()
	if fake.ReapVolumeStub != nil {
		return fake.ReapVolumeStub(handle)
	} else {
		return fake.reapVolumeReturns.result1
	}
}

func (fake *FakeWorkerDB) ReapVolumeCallCount() int {
	fake.reapVolumeMutex.RLock()
	defer fake.reapVolumeMutex.RUnlock()
	return len(fake.reapVolumeArgsForCall)
}

func (fake *FakeWorkerDB) ReapVolumeArgsForCall(i int) string {
	fake.reapVolumeMutex.RLock()
	defer fake.reapVolumeMutex.RUnlock()
	return fake.reapVolumeArgsForCall[i].handle
}

func (fake *FakeWorkerDB) ReapVolumeReturns(result1 error) {
	fake.ReapVolumeStub = nil
	fake.reapVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerDB) SetVolumeTTL(arg1 string, arg2 time.Duration) error {
	fake.setVolumeTTLMutex.Lock()
	fake.setVolumeTTLArgsForCall = append(fake.setVolumeTTLArgsForCall, struct {
		arg1 string
		arg2 time.Duration
	}{arg1, arg2})
	fake.recordInvocation("SetVolumeTTL", []interface{}{arg1, arg2})
	fake.setVolumeTTLMutex.Unlock()
	if fake.SetVolumeTTLStub != nil {
		return fake.SetVolumeTTLStub(arg1, arg2)
	} else {
		return fake.setVolumeTTLReturns.result1
	}
}

func (fake *FakeWorkerDB) SetVolumeTTLCallCount() int {
	fake.setVolumeTTLMutex.RLock()
	defer fake.setVolumeTTLMutex.RUnlock()
	return len(fake.setVolumeTTLArgsForCall)
}

func (fake *FakeWorkerDB) SetVolumeTTLArgsForCall(i int) (string, time.Duration) {
	fake.setVolumeTTLMutex.RLock()
	defer fake.setVolumeTTLMutex.RUnlock()
	return fake.setVolumeTTLArgsForCall[i].arg1, fake.setVolumeTTLArgsForCall[i].arg2
}

func (fake *FakeWorkerDB) SetVolumeTTLReturns(result1 error) {
	fake.SetVolumeTTLStub = nil
	fake.setVolumeTTLReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerDB) AcquireVolumeCreatingLock(arg1 lager.Logger, arg2 int) (db.Lock, bool, error) {
	fake.acquireVolumeCreatingLockMutex.Lock()
	fake.acquireVolumeCreatingLockArgsForCall = append(fake.acquireVolumeCreatingLockArgsForCall, struct {
		arg1 lager.Logger
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("AcquireVolumeCreatingLock", []interface{}{arg1, arg2})
	fake.acquireVolumeCreatingLockMutex.Unlock()
	if fake.AcquireVolumeCreatingLockStub != nil {
		return fake.AcquireVolumeCreatingLockStub(arg1, arg2)
	} else {
		return fake.acquireVolumeCreatingLockReturns.result1, fake.acquireVolumeCreatingLockReturns.result2, fake.acquireVolumeCreatingLockReturns.result3
	}
}

func (fake *FakeWorkerDB) AcquireVolumeCreatingLockCallCount() int {
	fake.acquireVolumeCreatingLockMutex.RLock()
	defer fake.acquireVolumeCreatingLockMutex.RUnlock()
	return len(fake.acquireVolumeCreatingLockArgsForCall)
}

func (fake *FakeWorkerDB) AcquireVolumeCreatingLockArgsForCall(i int) (lager.Logger, int) {
	fake.acquireVolumeCreatingLockMutex.RLock()
	defer fake.acquireVolumeCreatingLockMutex.RUnlock()
	return fake.acquireVolumeCreatingLockArgsForCall[i].arg1, fake.acquireVolumeCreatingLockArgsForCall[i].arg2
}

func (fake *FakeWorkerDB) AcquireVolumeCreatingLockReturns(result1 db.Lock, result2 bool, result3 error) {
	fake.AcquireVolumeCreatingLockStub = nil
	fake.acquireVolumeCreatingLockReturns = struct {
		result1 db.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.workersMutex.RLock()
	defer fake.workersMutex.RUnlock()
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	fake.createContainerMutex.RLock()
	defer fake.createContainerMutex.RUnlock()
	fake.updateContainerTTLToBeRemovedMutex.RLock()
	defer fake.updateContainerTTLToBeRemovedMutex.RUnlock()
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	fake.findContainerByIdentifierMutex.RLock()
	defer fake.findContainerByIdentifierMutex.RUnlock()
	fake.updateExpiresAtOnContainerMutex.RLock()
	defer fake.updateExpiresAtOnContainerMutex.RUnlock()
	fake.reapContainerMutex.RLock()
	defer fake.reapContainerMutex.RUnlock()
	fake.getPipelineByIDMutex.RLock()
	defer fake.getPipelineByIDMutex.RUnlock()
	fake.getVolumeTTLMutex.RLock()
	defer fake.getVolumeTTLMutex.RUnlock()
	fake.reapVolumeMutex.RLock()
	defer fake.reapVolumeMutex.RUnlock()
	fake.setVolumeTTLMutex.RLock()
	defer fake.setVolumeTTLMutex.RUnlock()
	fake.acquireVolumeCreatingLockMutex.RLock()
	defer fake.acquireVolumeCreatingLockMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeWorkerDB) recordInvocation(key string, args []interface{}) {
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

var _ worker.WorkerDB = new(FakeWorkerDB)
