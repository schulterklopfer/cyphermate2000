package waveshare
// -Wall

// #cgo CFLAGS: -g -I../c/lib/Config -I../c/lib/GUI -I../c/lib/e-Paper
// #cgo LDFLAGS: -lwiringPi -lm ${SRCDIR}/lib/waveshare_4in2.o
// #include <stdlib.h>
// #include "DEV_Config.h"
// #include "GUI_BMPfile.h"
// #include "GUI_Paint.h"
// #include "EPD_4in2.h"
import "C"

import (
  //"fmt"
  "errors"
  "fmt"
  "unsafe"
)

const EPD_4IN2_WIDTH  = C.EPD_4IN2_WIDTH
const EPD_4IN2_HEIGHT = C.EPD_4IN2_HEIGHT
const WHITE           = C.WHITE
const BLACK           = C.BLACK
const ROTATE_0        = C.ROTATE_0
const ROTATE_90       = C.ROTATE_90
const ROTATE_180      = C.ROTATE_180
const ROTATE_270      = C.ROTATE_270

type RedrawArea struct {
  X_start uint16
  Y_start uint16
  X_end uint16
  Y_end uint16
}


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

func EPD_4IN2_Display( imagePtr unsafe.Pointer ) {
  C.EPD_4IN2_Display( (*C.uchar)(imagePtr) )
}

func EPD_4IN2_PartialDisplay( X_start uint16, Y_start uint16, X_end uint16, Y_end uint16, imagePtr unsafe.Pointer ) {
  C.EPD_4IN2_PartialDisplay( C.ushort(X_start), C.ushort(Y_start), C.ushort(X_end), C.ushort(Y_end), (*C.uchar)(imagePtr))
}

func EPD_4IN2_PartialDisplayMulti( redrawAreas *[]RedrawArea, areaCount uint16, imagePtr unsafe.Pointer ) {
  println( (*C.EDP4IN2REDRAWAREA)(unsafe.Pointer( &(*redrawAreas)[0] )) )
  C.EPD_4IN2_PartialDisplayMulti( (*C.EDP4IN2REDRAWAREA)(unsafe.Pointer( &(*redrawAreas)[0] )), C.ushort(areaCount), (*C.uchar)(imagePtr))
}

func DEV_Delay_ms( delay int ) {
  C.DEV_Delay_ms( C.uint(delay) )
}

func Paint_NewImage( imagePtr unsafe.Pointer, rotate C.ushort, color C.ushort ) {
  C.Paint_NewImage( (*C.uchar)(imagePtr), EPD_4IN2_WIDTH, EPD_4IN2_HEIGHT, rotate, color)
}

func Paint_Clear( color C.ushort ) {
  C.Paint_Clear( color )
}

func Paint_ClearWindows( x1 uint16, y1 uint16, x2 uint16, y2 uint16, color C.ushort ) {
  C.Paint_ClearWindows(C.ushort(x1), C.ushort(y1), C.ushort(x2), C.ushort(y2), color)
}

func Paint_SelectImage( imagePtr unsafe.Pointer ) {
  C.Paint_SelectImage( (*C.uchar)(imagePtr) )
}

func DEV_Module_Exit() {
  C.DEV_Module_Exit()
}

func CreateImage( width uint16, height uint16 ) (unsafe.Pointer, error) {
  /* you have to edit the startup_stm32fxxx.s file and set a big enough heap size */

  imagesize := height
  if width % 8 == 0 {
    imagesize*=width / 8
  } else {
    imagesize*=width / 8 + 1
  }
  imagePtr := unsafe.Pointer(C.malloc(C.uint(imagesize)))

  fmt.Printf("Created imagePtr %d\n", imagePtr )

  if imagePtr == nil {
    return nil, errors.New("Failed to apply for black memory");
  }
  return imagePtr,nil
}

func FreeImage( imagePtr unsafe.Pointer ) {
  fmt.Printf("Freeing imagePtr %d\n", imagePtr )
  C.free( imagePtr )
}