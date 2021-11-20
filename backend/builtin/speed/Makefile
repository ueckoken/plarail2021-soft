start: startPy startGo

startGo: 
	./main

startPy:
	python3 control_pin.py

build_arm:
	GOOS=linux GOARCH=arm GOARM=6 go build -o main cmd/main.go
