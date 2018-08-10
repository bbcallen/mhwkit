build:
	go build ./pkg/mhw

test:
	go test ./pkg/mhw

fmt:
	go fmt ./pkg/mhw/*.go

checkinline:
	go build -gcflags=-m ./pkg/mhw

lib4c:
	go build -v -o mhwkit.dylib -buildmode=c-shared ./pkg/mhw


testv:
	go test -v ./pkg/mhw

testf:
	go test -v ./pkg/mhw -run FindSolution

testp:
	go test -v ./pkg/mhw -run ExecuteProcedure

tests:
	go test -v ./pkg/mhw -run LoadSkills

testd:
	go test -v ./pkg/mhw -run LoadDecorations

testa:
	go test -v ./pkg/mhw -run LoadArmors

testas:
	go test -v ./pkg/mhw -run LoadArmorSetBonuses

testsc:
	go test -v ./pkg/mhw -run LoadSlotCombinations

testc:
	go test -v ./pkg/mhw -run LoadCharms

testidm:
	go test -v ./pkg/mhw -run EvaluateCandidateArmorList

benchcpu:
	go test -v ./pkg/mhw -bench . -run Dummy -cpuprofile cpuprofile.out
	# go tool pprof bench.test cpuprofile.out

benchmem:
	go test -v ./pkg/mhw -bench . -run Dummy -benchmem -memprofile memprofile.out
	# go tool pprof --alloc_space bench.test memprofile.out