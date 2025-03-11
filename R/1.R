if (TRUE) {


  # Базовые типы данных
  numeric_var <- -2.0  # Числовой
  integer_var <- 12L   # Целочисленный
  character_var <- "Hello, R!"  # Символьный
  logical_var <- TRUE  # Логический
  complex_var <- 3 + 4i  # Комплексный
  raw_var <- charToRaw("R")  # Raw

  print(numeric_var)

  # Условные операторы
  if (numeric_var > 40) {
    print("Число больше 40")
  } else if (numeric_var == 40) {
    print("Число равно 40")
  } else {
    print("Число меньше 40")
  }

  # Функция ifelse
  result <- ifelse(logical_var=1, "Это правда", "Это ложь")
  print(result)

  # Функция switch
  day <- 3
  day_name <- switch(day,
                    "Понедельник",
                    "Вторник",
                    "Среда",
                    "Четверг",
                    "Пятница",
                    "Суббота",
                    "Воскресенье")
  print(day_name)
}