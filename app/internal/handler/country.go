package handler

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	_ "github.com/spashev/superapp/internal/models"
	"github.com/spashev/superapp/internal/repository"
	"github.com/spashev/superapp/internal/service"
)

// GetCountry retrieves a list of all countries
// @Summary Get all countries
// @Description Fetches all available countries
// @Tags Countries/Cities
// @Produce json
// @Success 200 {array} models.Country "List of countries"
// @Failure 500 {object} fiber.Map "Failed to fetch countries"
// @Router /countries [get]
func GetCountry(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		repo := repository.NewCountryRepository(db)
		categoryService := service.NewCountryService(repo)

		categories, err := categoryService.GetAllCountries()
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch countries",
			})
		}

		return c.Status(fiber.StatusOK).JSON(categories)
	}
}

// GetCountry retrieves a list of all cities
// @Summary Get all cities
// @Description Fetches all available cities
// @Tags Countries/Cities
// @Produce json
// @Success 200 {array} models.City "List of cities"
// @Failure 500 {object} fiber.Map "Failed to fetch cities"
// @Router /cities [get]
func GetCity(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		repo := repository.NewCountryRepository(db)
		categoryService := service.NewCountryService(repo)

		categories, err := categoryService.GetAllCities()
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch cities",
			})
		}

		return c.Status(fiber.StatusOK).JSON(categories)
	}
}

// GetCityByCountryId retrieves a list of all cities in the specified country
// @Summary Get all cities in the specified country
// @Description Get all cities in the specified country by their countryID
// @Tags Countries/Cities
// @Accept json
// @Produce json
// @Param id path int true "Country ID"
// @Success 200 {array} models.City "List of cities"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Server error"
// @Router /country/{id}/cities [get]
func GetCityByCountryId(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		params := c.Params("id")
		if params == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Id is required",
			})
		}

		id, err := strconv.Atoi(params)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Id must be an integer",
			})
		}

		repo := repository.NewCountryRepository(db)
		countryService := service.NewCountryService(repo)

		cities, err := countryService.GetAllCitiesByCountryId(id)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch all cities by country ID",
			})
		}

		return c.Status(fiber.StatusOK).JSON(cities)
	}
}
