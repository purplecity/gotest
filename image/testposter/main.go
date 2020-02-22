package main

import (
	"flag"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/golang/freetype"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"syscall"
)

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	size     = flag.Float64("size", 50, "font size in points")
	PosterPath = "/Users/ludongdong/Downloads/"
	BackgroundImage ="/Users/ludongdong/Downloads/xintuiguag.png"
	QrCodeTextPrefix = "https://app-hpoption-download.azfaster.com?code=HPOPTIONcode"

)




func generatePoster(iv,uid string) (code int) {
	//生成二维码图片
	//创建目录
	syscall.Umask(0)
	err := os.MkdirAll(PosterPath+uid,0755)
	if err != nil {
		log.Printf("MkdirAll failed %+v\n",err)
		return 1 //创建目录失败
	}

	//创建二维码文件
	qrurl := PosterPath+uid+"/qr.png"
	qf, err := os.Create(qrurl)
	if err != nil {
		log.Printf("create qr png failed %+v\n",err)
		os.RemoveAll(PosterPath+uid)
		return 1 //创建二维码文件失败
	}
	defer qf.Close()

	//生成二维码
	qImage, _ := qr.Encode(QrCodeTextPrefix+iv, qr.M, qr.Auto)
	qImage, _ = barcode.Scale(qImage, 410, 410)
	err = png.Encode(qf, qImage)
	if err != nil {
		log.Printf("generate qr code failed %+v\n",err)
		os.RemoveAll(PosterPath+uid)
		return 1 //生成二维码失败
	}

	//打开二维码图片
	/*
		qfo,_ := os.Open(qrurl)
		defer qfo.Close()
		qImage, err := png.Decode(qf)
		if err != nil {
			log.Printf("open qrcode failed %+v\n",err)
			os.RemoveAll(PosterPath+uid)
			return 1 //打开二维码图片失败
		}

	*/

	//打开背景图片
	bf,_ := os.Open(BackgroundImage)
	defer bf.Close()
	bImage, err := png.Decode(bf)
	if err != nil {
		log.Printf("open BackgroundImage failed %+v\n",err)
		os.RemoveAll(PosterPath+uid)
		return 1 //打开背景图片失败
	}


	//创建画布
	jpg := image.NewRGBA(image.Rect(0, 0, 750, 1334))  //获取画布大小 应该跟背景图尺寸一致
	//fmt.Printf("bf bounds:%+v,%+v,%+v,%+v   qf bounds :%+v,%+v,%+v,%+v\n",bImage.Bounds().Min.X,bImage.Bounds().Min.Y,bImage.Bounds().Max.X,bImage.Bounds().Max.Y, qImage.Bounds().Min.X,qImage.Bounds().Min.Y,qImage.Bounds().Max.X,qImage.Bounds().Max.Y)

	//画图
	draw.Draw(jpg, jpg.Bounds(), bImage, bImage.Bounds().Min, draw.Over)
	draw.Draw(jpg, jpg.Bounds(), qImage, qImage.Bounds().Min.Sub(image.Pt(88, 821)), draw.Over)  //中间为偏移量  显然应该是画布x1- qf的wight 再除以2 来居中 高偏移的话随意了

	bg := image.White

	ftBytes,err := ioutil.ReadFile(PosterPath +"hp.ttf")
	if err != nil {
		log.Printf("open font failed %+v\n",err)
		os.RemoveAll(PosterPath+uid)
		return 1 //打开字体失败
	}

	fonts, err := freetype.ParseFont(ftBytes)
	if err != nil {
		log.Printf("parse font failed %+v\n",err)
		return 1 //解析字体失败
	}


	c := freetype.NewContext()
	c.SetDPI(*dpi)
	c.SetFont(fonts)
	c.SetFontSize(*size)
	c.SetClip(jpg.Bounds())
	c.SetDst(jpg)
	c.SetSrc(bg) //白色字体
	//在这个偏移量开始画字体
	pt := freetype.Pt(74, 780)
	_,err = c.DrawString(iv,pt)
	if err != nil {
		log.Printf("draw string failed %+v\n",err)
		os.RemoveAll(PosterPath+uid)
		return 1 //绘画字体失败
	}

	mf, err := os.Create(PosterPath+uid+"/dst.png")
	if err != nil {
		log.Printf("generate Poster failed %+v\n",err)
		os.RemoveAll(PosterPath+uid)
		return 1 //生成海报失败
	}

	defer mf.Close()

	jpeg.Encode(mf, jpg, nil)
	os.Chmod(PosterPath+uid+"/dst.png",0755)
	return 0

}








func main() {

	Invitationcode := "OAANL9"
	Uid := "1227108124377231360"
	c := generatePoster(Invitationcode,Uid)
	if c != 0 {
		log.Printf("ERROR----%+v,%+v",Invitationcode,Uid)
	}

}


