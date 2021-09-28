package main

import (
	"log"

	"github.com/Rahmanwghazi/Monefy/app/middlewares"
	"github.com/Rahmanwghazi/Monefy/app/routes"
	_userUseCase "github.com/Rahmanwghazi/Monefy/business/users"
	_userController "github.com/Rahmanwghazi/Monefy/controllers/users"
	_userDB "github.com/Rahmanwghazi/Monefy/drivers/databases/users"
	_userRepository "github.com/Rahmanwghazi/Monefy/drivers/databases/users"

	_incomeUseCase "github.com/Rahmanwghazi/Monefy/business/income"
	_incomeController "github.com/Rahmanwghazi/Monefy/controllers/income"
	_incomeDB "github.com/Rahmanwghazi/Monefy/drivers/databases/income"
	_incomeRepository "github.com/Rahmanwghazi/Monefy/drivers/databases/income"

	_mysqlDriver "github.com/Rahmanwghazi/Monefy/drivers/mysql"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
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

	userRepository := _userRepository.NewMysqlUserRepository(connection)
	userUseCase := _userUseCase.NewUserUsecase(userRepository, &configJWT)
	userController := _userController.NewUserController(userUseCase)

	incomeRepository := _incomeRepository.NewMysqlIncomeRepository(connection)
	incomeUseCase := _incomeUseCase.NewIncomeUsecase(incomeRepository)
	incomeController := _incomeController.NewIncomeController(incomeUseCase)

	routesInit := routes.ControllerList{
		JWTMiddleware:    configJWT.Init(),
		UserController:   *userController,
		IncomeController: *incomeController,
	}

	routesInit.Routes(echo)
	log.Fatal(echo.Start(viper.GetString("server.address")))
}
