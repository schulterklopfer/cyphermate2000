package waveshare

// #cgo CFLAGS: -g -Wall -I../c/lib/Config -I../c/lib/GUI -I../c/lib/e-Paper
// #cgo LDFLAGS: -lwiringPi -lm ${SRCDIR}/lib/waveshare_4in2.o
// #include <stdlib.h>
// #include "DEV_Config.h"
// #include "GUI_BMPfile.h"
// #include "GUI_Paint.h"
// #include "EPD_4in2.h"
import "C"

import (
//"fmt"
  "unsafe"
)

const EPD_4IN2_WIDTH  = C.EPD_4IN2_WIDTH
const EPD_4IN2_HEIGHT = C.EPD_4IN2_HEIGHT
const WHITE           = C.WHITE
const BLACK           = C.BLACK


func DEV_Module_Init() int {
  result := C.DEV_Module_Init()
  return int(result)
}

func EPD_4IN2_Init() {
  C.EPD_4IN2_Init()
}

func EPD_4IN2_Clear() {
  C.EPD_4IN2_Clear()
}

func DEV_Delay_ms( delay int ) {
  C.DEV_Delay_ms( C.uint(delay) )
}

func Paint_NewImage( data []byte ) {
  C.Paint_NewImage((*C.uchar)(unsafe.Pointer(&data[0])), EPD_4IN2_WIDTH, EPD_4IN2_HEIGHT, 0, WHITE)
}

func Paint_Clear( color C.ushort ) {
  C.Paint_Clear( color )
}

func DEV_Module_Exit() {
  C.DEV_Module_Exit()
}
