package main

import "github.com/schulterklopfer/cyphermate2000/waveshare"

func main() {

  println("EPD_4IN2_test Demo")
  if(waveshare.DEV_Module_Init()!=0){
    panic( "NAARGHG!!" )
  }

  println("e-Paper Init and Clear...")
  waveshare.EPD_4IN2_Init()
  waveshare.EPD_4IN2_Clear()
  waveshare.DEV_Delay_ms(500)

  //imagesize := waveshare.EPD_4IN2_WIDTH * waveshare.EPD_4IN2_HEIGHT

  //blackImage := make( []byte, imagesize )

  //waveshare.Paint_NewImage( blackImage )
  waveshare.Paint_Clear( waveshare.WHITE )
  waveshare.Paint_Clear( waveshare.BLACK )
  //waveshare.EPD_4IN2_Clear()

  waveshare.DEV_Module_Exit()
}