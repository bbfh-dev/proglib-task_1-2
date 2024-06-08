# Домашнее задание

В далекой звездной системе встретились две флотилии космических кораблей. Корабли могут передвигаться по всему пространству звездной системы по прямой, поворачиваться против и по часовой стрелке, стрелять фотонными торпедами. Попадание фотонной торпеды в корабль выводит его из строя.

От каждой флотилии в сражении принимают участие по три космических корабля.

Победу в битве одерживает та флотилия, которая первой выведет из строя все корабли соперника.

Управление флотилиями осуществляется игрокам компьютерными программами (то есть не с клавиатуры).

Концептуально игра состоит из трех подсистем:

1. Игровой сервер, где реализуется вся игровая логика.

2. Player - консольное приложение, на котором отображается конкретная битва.

3. Агент - приложение, которое запускает программу управления танками от имени игрока и отправляет управляющие команды на игровой сервер.

## ДЗ№1 Реализовать движение объектов на игровом поле в рамках подсистемы Игровой сервер.

# Definition of Done критерии

## Complexity Level #1:

1. Прямолинейное равномерное движение без деформации.
    1. Само движение реализовано в виде отдельного класса
    1. Для движущихся объектов определен интерфейс, устойчивый к появлению новых видов движущихся объектов
1. Поворот объекта вокруг оси.
    1. Сам поворот реализован в виде отдельного класса
    1. Для поворачивающегося объекта определен интерфейс, устойчивый к появлению новых видов движущихся объектов
1. Код решения опубликован на github/gitlab в отдельной ветке.
1. Код компилируется без ошибок

## Complexity Level #2:

1. Реализован тест: Для объекта, находящегося в точке (12, 5) и движущегося со скоростью (-7, 3) движение меняет положение объекта на (5, 8)
1. Реализован тест: Попытка сдвинуть объект, у которого невозможно прочитать значение мгновенной скорости, приводит к ошибке
1. Реализован тест: Попытка сдвинуть объект, у которого невозможно прочитать значение мгновенной скорости, приводит к ошибке
1. Реализован тест: Попытка сдвинуть объект, у которого невозможно изменить положение в пространстве, приводит к ошибке
1. Все тесты успешно выполняются

## Complexity Level #3:

1. Реализованы тесты для поворота корабля вокруг собственной оси.
1. Настроен расчет покрытия кода тестами.
1. Настроен CI, который умеет собирать проект и прогонять тесты, вычислять покрытие кода тестами.
1. Покрытие кода тестами 100%.
1. Пайплайн “зеленый”