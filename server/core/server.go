package core

import (
	"fmt"
	"time"


	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/initialize"
	"github.com/ebedevelopment/next-gen-tms/server/service/system"
)

// server
type server interface {
	ListenAndServe() error
	ListenAndServeTLS(certFile string, keyFile string)error
}

// RunWindowsServer run Windows server
func RunWindowsServer() error{


	// from db load jwt data
	if global.GvaDB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	// Router.Static("/form-generator", "./resource/page")

	// address := fmt.Sprintf(":%d", global.GvaConfig.System.Addr)

	address := global.GvaConfig.System.Addr 
	s := initServer(address, Router)
	// ensure language format
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GvaLog.Info("server run success on "+ address)

	fmt.Println(`
	welcome use github.com/ebedevelopment/next-gen-tms/server
	version:V2
	1.0.0
`, address)

	global.GvaLog.Error(s.ListenAndServe().Error())

	// global.GvaLog.Error( s.ListenAndServeTLS(global.GvaConfig.CertificateConfig.CertificateFile, global.GvaConfig.CertificateConfig.KeyFile).Error())


   return nil

}





