# Цикл for
for (i in 1:10) {
  print(i)
}


# Цикл while
counter <- 1
while (counter <= 3) {
  print(paste("Счётчик:", counter))
  counter <- counter + 1
}

# Цикл repeat с break и next
repeat {
  counter <- counter - 1
  if (counter == 0) {
    next  # Пропустить итерацию
  }
  print(paste("Повтор:", counter))
  if (counter < -3) {
    break  # Выйти из цикла
  }
}

# Пользовательская функция с аргументами
my_function <- function(a, b = 10, ...) {
  args <- list(...)
  print(paste("a =", a, "b =", b))
  print("Дополнительные аргументы:")
  print(args)
}

my_function(5, c = "Дополнительный аргумент")


`%add%` <- function(a, b) {
  return(a + b)
}

print(1 %add% 2 + a)
