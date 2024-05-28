package main

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"log"
	"math/rand"
	"os"

	"github.com/gin-gonic/gin"
	repoprovider "github.com/pandhu/hehemock/app/providers/repository"
	usecaseprovider "github.com/pandhu/hehemock/app/providers/usecase"
	"github.com/pandhu/hehemock/config"
	"github.com/pandhu/hehemock/database"
	"github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
	"gopkg.in/ukautz/clif.v1"
)

func init() {
	gotenv.Load()

	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetReportCaller(true)

	log.SetOutput(os.Stdout)

	var b [8]byte
	cryptorand.Read(b[:])
	seed := binary.LittleEndian.Uint64(b[:])
	rand.Seed(int64(seed))
}

func main() {
	conf := config.All()

	if conf.App.ENV != "local" {

		defer func() {
			tracer.Stop()
			profiler.Stop()
		}()

		gin.SetMode(gin.ReleaseMode)
	}

	connDB := database.Init(conf)

	repoProvider := repoprovider.InitRepo(connDB)
	usecases := usecaseprovider.InitUsecase(repoProvider)

	if runCli(conf, usecases, repoProvider) {
		return
	}

	server := httpKTA.InitServer(usecases)
	server.Serve()
}

func runCli(conf *config.Configuration, u *usecaseprovider.Usecase, repo *repoprovider.Repo) bool {
	// No need to run CLI if there is no argument
	if len(os.Args) == 1 {
		return false
	}

	cli := clif.New("Go - Flip KTA", "1.0.0", "KTA service - Flip.id")
	cli.Run()

	return true
}
