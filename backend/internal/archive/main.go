package main

import (
	"fmt"
	"net"
	"strings"

	// "os"
	// utilsErr "sample/utils/error"
	// utilsNet "sample/utils/net"
	"os/exec"
	"time"
)

type Station struct {
	BranchesOrStops map[string]BranchOrStop
}
type BranchOrStop struct {
	name   string
	Worker WorkerDevice
}
type WorkerDevice interface {
	Setter(bool) error
	Getter() bool
}

//ESP32など
type RemoteWorkerDevice struct {
	status  bool
	address string
}

func (w RemoteWorkerDevice) Getter() bool {
	return w.status
}
func (w RemoteWorkerDevice) Setter(s bool) error {
	//TODO
	return nil
}

//本体に繋いだやつ
type LocalWorkerDevice struct {
	status bool
	pin    string
}

func (w LocalWorkerDevice) Getter() bool {
	return w.status
}
func (w LocalWorkerDevice) Setter(s bool) error {
	var arg string
	w.status = s
	if s {
		arg = "1"
	} else {
		arg = "0"
	}
	err := exec.Command("python3", "controlpin.py", w.pin, arg).Start()
	return err
}

var Stations = map[string]Station{"chofu": Chofu, "meidaimae": Meidaimae, "sasazuka": Sasazuka, "kitano": Kitano, "temp": Temp}
var Chofu Station
var Meidaimae Station
var Sasazuka Station
var Kitano Station
var Temp Station

func main() {
    chofu_b1 := BranchOrStop{name: "chofu_b1", Worker: LocalWorkerDevice{status: false, pin: "1"}}
    chofu_b2 := BranchOrStop{name: "chofu_b2", Worker: LocalWorkerDevice{status: false, pin: "1"}}
    chofu_b3 := BranchOrStop{name: "chofu_b3", Worker: LocalWorkerDevice{status: false, pin: "1"}}
    chofu_s1 := BranchOrStop{name: "chofu_s1", Worker: LocalWorkerDevice{status: false, pin: "1"}}
    chofu_s2 := BranchOrStop{name: "chofu_s2", Worker: LocalWorkerDevice{status: false, pin: "1"}}
    chofu_s3 := BranchOrStop{name: "chofu_s3", Worker: LocalWorkerDevice{status: false, pin: "1"}}
    chofu_s4 := BranchOrStop{name: "chofu_s4", Worker: LocalWorkerDevice{status: false, pin: "1"}}
	Chofu.BranchesOrStops = map[string]BranchOrStop{chofu_b1.name: chofu_b1, chofu_b2.name: chofu_b2, chofu_b3.name: chofu_b3, chofu_s1.name: chofu_s1, chofu_s2.name: chofu_s2, chofu_s3.name: chofu_s3, chofu_s4.name: chofu_s4}

	meidaimae_s1 := BranchOrStop{name: "meidaimae_s1", Worker: RemoteWorkerDevice{status: false, address: "localhost"}}
	meidaimae_s2 := BranchOrStop{name: "meidaimae_s2", Worker: RemoteWorkerDevice{status: false, address: "localhost"}}
	Meidaimae.BranchesOrStops = map[string]BranchOrStop{meidaimae_s1.name: meidaimae_s1, meidaimae_s2.name: meidaimae_s2}

	sasazuka_b1 := BranchOrStop{name: "sasazuka_b1", Worker: RemoteWorkerDevice{status: false, address: "localhost"}}
	sasazuka_s1 := BranchOrStop{name: "sasazuka_s1", Worker: RemoteWorkerDevice{status: false, address: "localhost"}}
	sasazuka_s2 := BranchOrStop{name: "sasazuka_s2", Worker: RemoteWorkerDevice{status: false, address: "localhost"}}
	sasazuka_s4 := BranchOrStop{name: "sasazuka_s4", Worker: RemoteWorkerDevice{status: false, address: "localhost"}}
	Sasazuka.BranchesOrStops = map[string]BranchOrStop{sasazuka_b1.name: sasazuka_b1, sasazuka_s1.name: sasazuka_s1, sasazuka_s2.name: sasazuka_s2, sasazuka_s4.name: sasazuka_s4}

	kitano_b1 := BranchOrStop{name: "kitano_b1", Worker: RemoteWorkerDevice{status: false, address: "localhost"}}
	kitano_s2 := BranchOrStop{name: "kitano_s2", Worker: RemoteWorkerDevice{status: false, address: "localhost"}}
	kitano_s3 := BranchOrStop{name: "kitano_s3", Worker: RemoteWorkerDevice{status: false, address: "localhost"}}
	Kitano.BranchesOrStops = map[string]BranchOrStop{kitano_b1.name: kitano_b1, kitano_s2.name: kitano_s2, kitano_s3.name: kitano_s3}

	// 今回はサーバー/クライアント共にローカルマシン
	// ポートは適当に
	serverIP := "150.95.182.89"
	serverPort := "8081"
	clientIP := "localhost"
	clientPort := 5005

	// tcpで接続
	serverAdd, _ := net.ResolveTCPAddr("tcp", serverIP+":"+serverPort)
	// utilsErr.CheckWithExit(err)
	clientAddr := new(net.TCPAddr)
	clientAddr.IP = net.ParseIP(clientIP)
	clientAddr.Port = clientPort
	conn, err := net.DialTCP("tcp", clientAddr, serverAdd)

	if err != nil {
		fmt.Println(err)
		return
	}

	// 読み取り用
	go handleClient(conn)
	// utilsErr.CheckWithExit(err)
	// defer conn.Close()

	// メッセージをバイトに変換して送信
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.Write([]byte("Device1\nADD chofu\nREG chofu chofu_b1 OnOff\nREG chofu chofu_b2 OnOff\nREG chofu chofu_b3 OnOff\nREG chofu chofu_s1 OnOff\nREG chofu chofu_s2 OnOff\nREG chofu chofu_s3 OnOff\nREG chofu chofu_s4 OnOff\nADD meidaimae\nREG meidaimae meidaimae_s1 OnOff\nREG meidaimae meidaimae_s2 OnOff\nAdd sasazuka\nREG sasazuka sasazuka_b1 OnOff\nREG sasazuka sasazuka_s1 OnOff\nREG sasazuka sasazuka_s2 OnOff\nREG sasazuka sasazuka_s4 OnOff\nADD kitano\nREG kitano kitano_b1 OnOff\nREG kitano kitano_s2 OnOff\nREG kitano kitano_s3 OnOff\nFIN\n"))
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("sleeped")
		// conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		// conn.Write([]byte("\n"))
	}
	// fmt.Println("end")
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))

		messageBuf := make([]byte, 1024)
		messageLen, _ := conn.Read(messageBuf)
		// utilsErr.CheckWithExit(err)
		message := string(messageBuf[:messageLen])
		fmt.Println("message", messageLen)
		fmt.Println(message)
		// この書き込み処理がないと、連続してclientからメッセージを送信すると以下のエラーが発生する
		// fatal: error: dial tcp 192.168.3.22:50031->192.168.3.22:50030: bind: address already in useexit status 1
		// 理由は理解していない・・
		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		conn.Write([]byte("\n"))
		if messageLen == 0 {
			continue
		}
		for _, m := range strings.Split(message, "\n") {
			op := strings.Split(m, " ")
			var onOrOff bool
			if len(op) != 3 {
				fmt.Println("warning:", op)
			}
			if op[2] == "On" {
				onOrOff = true
			} else if op[2] == "Off" {
				onOrOff = false
			} else {
				fmt.Println("warning:", op)
			}
			s, isCorrect := Stations[op[0]]
			if !isCorrect {
				fmt.Println("warning:", op)
				continue
			}
			b, isCorrect := s.BranchesOrStops[op[1]]
			if !isCorrect {
				fmt.Println("warning:", op)
				continue
			}
			b.Worker.Setter(onOrOff)
		}
	}
}
