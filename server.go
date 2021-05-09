package main

import (
	"fmt"
	"net/http"

	"github.com/gmmads/Calculator/calculator"
	"github.com/gmmads/Calculator/controller"
	router "github.com/gmmads/Calculator/http"
	"github.com/gmmads/Calculator/repository"
)

var (
	calculationRepository repository.CalculationRepository = repository.NewFirestoreRepository()
	basicCalculator       calculator.Calculator            = calculator.NewBasicCalculator(calculationRepository)

	calculationController controller.CalculationController = controller.NewCalculationController(basicCalculator)
	httpRouter            router.Router                    = router.NewChiRouter()
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	})

	httpRouter.GET("/calculate", calculationController.GetHistory)
	httpRouter.POST("/calculate", calculationController.AddCalculation)

	httpRouter.SERVE(port)
}
