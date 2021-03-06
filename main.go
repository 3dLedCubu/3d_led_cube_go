package main

import (
	"bufio"
	"flag"
	"3d_led_cube_go/ledlib"
	"3d_led_cube_go/ledlib/servicegateway"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"strings"
	"time"
)

func getUnixNano() int64 {
	return time.Now().UnixNano()
}

func main() {
	var (
		optDestination   = flag.String("d", "localhost:9001", "Specify IP and port of Led Cube. if opt is not set, launch simulator.")
		optIdentifier    = flag.String("i", "", "Identifier for this process. Audio module use this identifier to manage audio.")
		optAudigo        = flag.String("a", "audigo.local", "Specify IP and port of device which Audigo is installed.")
		optRealsense     = flag.String("r", "127.0.0.1:5501", "Specify IP and port of server main_realsense_serivce.py running.")
		optStarupOrder   = flag.String("s", "", "Specify show order which will run when application launch.")
		optTestInputMode = flag.Bool("t", false, "Specify flag provide stdin which you can input order.")
	)
	flag.Parse()
	if *optDestination == "" {
		runtime.LockOSThread()
		ledlib.GetLed().EnableSimulator(true)
	} else {
		ledlib.GetLed().EnableSimulator(false)
		ledlib.GetLed().SetUrl(*optDestination)
	}
	log.Println("udp target: " + *optDestination)
	go func() {
		//		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	/*
	 setup audigo
	*/
	servicegateway.InitAudigoSeriveGateway("http://"+*optAudigo, *optIdentifier)
	ledlib.InitSeriveGatewayRealsense("tcp://" + *optRealsense)

	/*
		setup renderer
	*/
	renderer := ledlib.NewLedBlockRenderer()
	renderer.Start()

	ledlib.SetUpWebAPIforCommon(renderer)
	ledlib.SetUpWebAPIforPainting(renderer)

	log.Println("led framework is running ...  on port 5001")
	if *optStarupOrder != "" {
		log.Println("[INFO]default order: " + *optStarupOrder)
		renderer.Show(*optStarupOrder)
	}

	runServerCommand := func() {
		log.Fatal(http.ListenAndServe(":5001", nil))
	}
	if *optTestInputMode {
		go runServerCommand()

		for {
			sc := bufio.NewScanner(os.Stdin)
			log.Print(">>")
			if sc.Scan() {
				input := sc.Text()
				log.Println("input:" + input)
				switch {
				case strings.HasPrefix(input, "show"):
					renderer.Show(strings.Replace(input, "show:", "", 1))
				case strings.HasPrefix(input, "abort"):
					renderer.Abort()
				}
			}
		}
	} else {
		runServerCommand()
	}
	renderer.Terminate()

}
