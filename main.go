package main

import (
	"log"

	"github.com/Rahmanwghazi/Monefy/app/middlewares"
	_userController "github.com/Rahmanwghazi/Monefy/app/presenter/users"
	"github.com/Rahmanwghazi/Monefy/app/routes"
	_userUseCase "github.com/Rahmanwghazi/Monefy/business/users"
	_userDB "github.com/Rahmanwghazi/Monefy/repository/databases/users"
	_userRepository "github.com/Rahmanwghazi/Monefy/repository/databases/users"

	_incomeController "github.com/Rahmanwghazi/Monefy/app/presenter/income"
	_incomeUseCase "github.com/Rahmanwghazi/Monefy/business/income"
	_incomeDB "github.com/Rahmanwghazi/Monefy/repository/databases/income"
	_incomeRepository "github.com/Rahmanwghazi/Monefy/repository/databases/income"

	_expenseController "github.com/Rahmanwghazi/Monefy/app/presenter/expenses"
	_expenseUseCase "github.com/Rahmanwghazi/Monefy/business/expenses"
	_expenseDB "github.com/Rahmanwghazi/Monefy/repository/databases/expenses"
	_expenseRepository "github.com/Rahmanwghazi/Monefy/repository/databases/expenses"

	_investplanController "github.com/Rahmanwghazi/Monefy/app/presenter/investplans"
	_investplanUseCase "github.com/Rahmanwghazi/Monefy/business/investplans"
	_investplanDB "github.com/Rahmanwghazi/Monefy/repository/databases/investplans"
	_investplanRepository "github.com/Rahmanwghazi/Monefy/repository/databases/investplans"

	_productRepository "github.com/Rahmanwghazi/Monefy/repository/thirdparties/ojk_invest"

	_mysqlDriver "github.com/Rahmanwghazi/Monefy/repository/mysql"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Server RUN on Debug mode")
	}
}
func DBMigration(db *gorm.DB) {
	db.AutoMigrate(&_userDB.User{})
	db.AutoMigrate(&_incomeDB.Income{})
	db.AutoMigrate(&_expenseDB.Expense{})
	db.AutoMigrate(&_investplanDB.InvestPlan{})
}

func main() {
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.password`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	configJWT := middlewares.ConfigJWT{
		SecretJwt:       viper.GetString("jwt.secret"),
		ExpiredDuration: viper.GetInt("jwt.expired"),
	}

	connection := configDB.InitialDB()
	DBMigration(connection)

	echo := echo.New()
	middlewares.LogMiddleware(echo)

	productRepository := _productRepository.NewOJKAPI()

	userRepository := _userRepository.NewMysqlUserRepository(connection)
	userUseCase := _userUseCase.NewUserUsecase(userRepository, &configJWT)
	userController := _userController.NewUserController(userUseCase)

	incomeRepository := _incomeRepository.NewMysqlIncomeRepository(connection)
	incomeUseCase := _incomeUseCase.NewIncomeUsecase(incomeRepository)
	incomeController := _incomeController.NewIncomeController(incomeUseCase)

	expenseRepository := _expenseRepository.NewMysqlExpenseRepository(connection)
	expenseUseCase := _expenseUseCase.NewExpenseUsecase(expenseRepository)
	expenseController := _expenseController.NewExpenseController(expenseUseCase)

	investplanRepository := _investplanRepository.NewMysqlnvestPlanRepository(connection)
	investplanUseCase := _investplanUseCase.NewInvestPlanUsecase(investplanRepository, productRepository)
	investplanController := _investplanController.NewInvestPlanController(investplanUseCase)

	routesInit := routes.ControllerList{
		JWTMiddleware:        configJWT.Init(),
		UserController:       *userController,
		IncomeController:     *incomeController,
		ExpenseController:    *expenseController,
		InvestPlanController: *investplanController,
	}

	routesInit.Routes(echo)
	log.Fatal(echo.Start(viper.GetString("server.address")))
}
