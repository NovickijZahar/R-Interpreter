# Класс S3
person <- list(name = "Иван", age = 30)
class(person) <- "Person"

# Метод для класса S3
print.Person <- function(obj) {
  cat("Имя:", obj$name, "\nВозраст:", obj$age, "\n")
}
print(person)

# Класс S4
setClass("Car",
         slots = list(make = "character", year = "numeric"))

# Метод для класса S4
setMethod("show", "Car", function(object) {
  cat("Марка:", object@make, "\nГод выпуска:", object@year, "\n")
})

my_car <- new("Car", make = "Toyota", year = 2020)

show(my_car)

# Reference класс
PersonRef <- setRefClass("PersonRef",
                         fields = list(name = "character", age = "numeric"),
                         methods = list(
                           greet = function() {
                             cat("Привет, меня зовут", name, 
                                "и мне", age, "лет.\n")
                           }
                         ))

person_ref <- PersonRef$new(name = "Анна", age = 25)

person_ref$greet()

# R6 класс
library(R6)
PersonR6 <- R6Class("PersonR6",
                    public = list(
                      name = NULL,
                      age = NULL,
                      initialize = function(name, age) {
                        self$name <- name
                        self$age <- age
                      },
                      greet = function() {
                        cat("Привет, меня зовут", self$name, "и мне", self$age, "лет.\n")
                      }
                    ))
person_r6 <- PersonR6$new(name = "Олег", age = 35)
person_r6$greet()