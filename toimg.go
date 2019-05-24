package helpers

 import (
         "bytes"
         "fmt"
         "image"
         "image/png"
         "os"
 )

 func MakeImg(source []byte) {
	fmt.Println("THE SOURCE IS : ", source)
	// convert []byte to image for saving to file
	img, _, _ := image.Decode(bytes.NewReader(source))
	fmt.Println(">>>>>decoded")
	//save the imgByte to file
	out, err := os.Create("./tmp/QRImg.png")

	fmt.Println(">>>>>tmpimage")
	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}

	err = png.Encode(out, img)
	fmt.Println(">>>>>encoded")
	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}

	// everything ok
	fmt.Println("QR code generated and saved to QRimg.png")

 }