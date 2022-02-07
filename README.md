#### Запуск: `make run`

### Задание 1 
#### Генератор случайных чисел

Сгенерировать n-количество случайных чисел с помощью пакета *math/rand*, и поместить в канал.

### Задание 2 
#### Конвейер чисел

Даны два канала. В первый пишутся числа. Нужно, чтобы числа читались из первого по мере поступления, что-то с ними происходило (допустим, возводились в квадрат) и результат записывался во второй канал.
После этого нужно вывести в консоль содержимое второго канала.

### Задание 3 
#### Слить N каналов в один
Даны n каналов типа chan int. Надо написать функцию, которая смерджит все данные из этих каналов в один и вернет его.
