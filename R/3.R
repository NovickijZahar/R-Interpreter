# Reference класс
PersonRef <- setRefClass("PersonRef",
                         fields = list(name = "character", age = "numeric"))

# Создание экземпляра класса Reference
person_ref <- PersonRef$new(name = "Анна", age = 25)

# R6 класс
library(R6)
PersonR6 <- R6Class("PersonR6",
                    public = list(
                      name = "character",
                      age = "numeric"
                    ))

# Создание экземпляра класса R6
person_r6 <- PersonR6$new(name = "Олег", age = 35)

# Другие классы без методов

# Reference класс для животных
AnimalRef <- setRefClass("AnimalRef",
                         fields = list(species = "character", age = "numeric"))

# Создание экземпляра класса AnimalRef
animal_ref <- AnimalRef$new(species = "Собака", age = 5)

# R6 класс для автомобилей
CarR6 <- R6Class("CarR6",
                 public = list(
                   make = "character",
                   model = "character",
                   year = "numeric"
                 ))

# Создание экземпляра класса CarR6
car_r6 <- CarR6$new(make = "Тойота", model = "Камри", year = 2020)

# Reference класс для книг
BookRef <- setRefClass("BookRef",
                       fields = list(title = "character", author = "character"))

# Создание экземпляра класса BookRef
book_ref <- BookRef$new(title = "1984", author = "Джордж Оруэлл")

# R6 класс для сотрудников
EmployeeR6 <- R6Class("EmployeeR6",
                      public = list(
                        name = "character",
                        position = "character"
                      ))

# Создание экземпляра класса EmployeeR6
employee_r6 <- EmployeeR6$new(name = "Иван", position = "Менеджер")