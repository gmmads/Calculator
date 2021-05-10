package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gmmads/Calculator/calculator"
	"github.com/gmmads/Calculator/controller"
	"github.com/gmmads/Calculator/history"
	router "github.com/gmmads/Calculator/http"
)

var (
	calcHistory     history.History       = history.NewCalcHistory()
	basicCalculator calculator.Calculator = calculator.NewBasicCalculator(calcHistory)

	calculationController controller.CalculationController = controller.NewCalculationController(basicCalculator)
	httpRouter            router.Router                    = router.NewChiRouter()
)

func main() {
	var port string
	if len(os.Args) > 1 && os.Args[1] == "test" {
		port = ":8000"
	} else {
		port = ":" + os.Getenv("PORT")
	}

	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	})

	httpRouter.GET("/calculations", calculationController.GetHistory)
	httpRouter.POST("/calculations", calculationController.AddCalculation)

	httpRouter.SERVE(port)
}
