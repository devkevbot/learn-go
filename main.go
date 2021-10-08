package main

import (
	"flag"
)

func main() {
	modulePtr := flag.String("m", "", "The module to run")
	flag.Parse()
	run(*modulePtr)
}

func run(moduleName string) {
	switch moduleName {
	case "1":
		_001TypeInference()
		_001TypeConversion()
	case "2":
		_002ForLoop()
		_002Conditional()
		_002Switch()
		_002Defer()
	case "3":
		_003Pointers()
	case "4":
		_004Structs()
	case "5":
		_005Array()
		_005Slice()
	case "6":
		_006Range()
	case "7":
		_007Maps()
		_007MutatingMaps()
	case "8":
		_008FunctionValues()
		_008FunctionClosures()
	case "9":
		_009Interfaces()
	case "10":
		_010TypeAssertions()
		_010TypeSwitches()
	case "11":
		_011Errors()
	case "12":
		_012Readers()
	case "13":
		_013GoRoutines()
		_013UnbufferedChannels()
		_013BufferedChannels()
		_013RangeAndClose()
		_013Select()
		_013SelectDefault()
		_013Mutex()
	default:
		panic("Unknown module provided!")
	}
}
