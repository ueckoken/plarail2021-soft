start: startGo startPy

startGo: 
	./speedControl &

startPy:
	python3 control_pin.py

build_arm:
	GOOS=linux GOARCH=arm GOARM=6 go build -o speedControl cmd/main.go
