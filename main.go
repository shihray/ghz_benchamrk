package main

import (
	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
	"github.com/golang/protobuf/proto"
	pb "github.com/lctech-tw/gin-proto/dist/go/star/direct"
	"log"
	"os"
)

func main() {
	// 组装BinaryData
	item := pb.TagListReq{Token: "b5g5PgZXUw6drUcoLBNBMgXXY7DxfabE"}
	buf := proto.Buffer{}
	err := buf.EncodeMessage(&item)
	if err != nil {
		log.Fatal(err)
		return
	}
	report, err := runner.Run(
		"star.direct.direct.DirectService.StarList",
		"localhost:8080",
		runner.WithBinaryData(buf.Bytes()),
		runner.WithInsecure(true),
		runner.WithAsync(true),
		runner.WithTotalRequests(10000),

		//runner.WithConcurrency(100),
		runner.WithConcurrencySchedule(runner.ScheduleLine),
		runner.WithConcurrencyStep(10),
		runner.WithConcurrencyStart(5),
		runner.WithConcurrencyEnd(100),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	// output path
	file, err := os.Create("report.html")
	if err != nil {
		log.Fatal(err)
		return
	}
	rp := printer.ReportPrinter{
		Out:    file,
		Report: report,
	}
	// output format
	_ = rp.Print("html")
}
