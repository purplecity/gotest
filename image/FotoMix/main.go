package main

import (
	"flag"
	"fmt"
	"github.com/golang/freetype"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

// NOTE: cd go/src/github.com/golang  然后 git clone https://github.com/golang/image

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	size     = flag.Float64("size", 50, "font size in points")
)


func main() {

	mf, err := os.Create("/Users/ludongdong/go/src/gotest/image/dst.png")
	if err != nil {
		fmt.Println(err)
	}
	defer mf.Close()


	bf,_ := os.Open("/Users/ludongdong/go/src/gotest/image/beijing.png")
	defer bf.Close()
	bImage, err := png.Decode(bf)
	if err != nil {
		fmt.Println("bimage",err)
		return
	}

	qf,_ := os.Open("/Users/ludongdong/go/src/gotest/image/qr.png")
	defer qf.Close()
	qImage, err := png.Decode(qf)
	if err != nil {
		fmt.Println("qimage",err)
		return
	}


	//画布
	jpg := image.NewRGBA(image.Rect(0, 0, 864, 1536))  //获取画布大小 应该跟背景图尺寸一致
	fmt.Printf("bf bounds:%+v,%+v,%+v,%+v   qf bounds :%+v,%+v,%+v,%+v\n",bImage.Bounds().Min.X,bImage.Bounds().Min.Y,bImage.Bounds().Max.X,bImage.Bounds().Max.Y,
		qImage.Bounds().Min.X,qImage.Bounds().Min.Y,qImage.Bounds().Max.X,qImage.Bounds().Max.Y)

	draw.Draw(jpg, jpg.Bounds(), bImage, bImage.Bounds().Min, draw.Over)
	draw.Draw(jpg, jpg.Bounds(), qImage, qImage.Bounds().Min.Sub(image.Pt(220, 740)), draw.Over)  //中间为偏移量  显然应该是画布x1- qf的wight 再除以2 来居中 高偏移的话随意了

	bg := image.White

	ftBytes,err := ioutil.ReadFile("/Users/ludongdong/go/src/gotest/image/hp.ttf")
	if err != nil {
		log.Println(err)
		return
	}
	fonts, err := freetype.ParseFont(ftBytes)
	if err != nil {
		log.Println(err)
		return
	}
	c := freetype.NewContext()
	c.SetDPI(*dpi)
	c.SetFont(fonts)
	c.SetFontSize(*size)
	c.SetClip(jpg.Bounds())
	c.SetDst(jpg)
	c.SetSrc(bg)
	pt := freetype.Pt(430, 738)
	text := "HIJKLM"
	_,err = c.DrawString(text,pt)
	if err != nil {
		fmt.Println("text",err)
	}

	jpeg.Encode(mf, jpg, nil)

}
