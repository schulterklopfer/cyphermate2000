package main

import "github.com/schulterklopfer/cyphermate2000/waveshare"

func main() {

  println("EPD_4IN2_test Demo")
  if(waveshare.DEV_Module_Init()!=0){
    panic( "NAARGHG!!" )
  }

  println("e-Paper Init and Clear...")
  waveshare.EPD_4IN2_Init()
  //waveshare.EPD_4IN2_Clear()
  //waveshare.DEV_Delay_ms(500)

  blackImage, err := waveshare.CreateImage( waveshare.EPD_4IN2_WIDTH, waveshare.EPD_4IN2_HEIGHT )
  defer waveshare.FreeImage( blackImage )

  if err != nil {
    println( "Error: "+err.Error() )
    return
  }

  waveshare.Paint_NewImage( blackImage, waveshare.ROTATE_0, waveshare.BLACK )
  waveshare.Paint_SelectImage( blackImage )


  waveshare.Paint_Clear( waveshare.BLACK )
  waveshare.EPD_4IN2_Display( blackImage )

  waveshare.DEV_Delay_ms(500)

  waveshare.Paint_Clear( waveshare.WHITE )
  waveshare.EPD_4IN2_Display( blackImage )

  waveshare.DEV_Delay_ms(500)

  waveshare.Paint_ClearWindows( 30,30,100,100, waveshare.BLACK )
  waveshare.EPD_4IN2_PartialDisplay( 30,30,100,100, blackImage )

  waveshare.DEV_Delay_ms(500)

  waveshare.Paint_ClearWindows( 30,30,100,100, waveshare.WHITE )
  waveshare.EPD_4IN2_PartialDisplay( 30,30,100,100, blackImage )

  //waveshare.EPD_4IN2_Clear()

  waveshare.DEV_Module_Exit()
}