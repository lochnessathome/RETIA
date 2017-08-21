# RETIA
RETIA - акроним для RElaTIonal Algebra

## Примеры использования

### Кортежи

`TUPLE { age integer 18, name char "Ivan" }` - создать неименнованный кортеж.

`boy := TUPLE { age integer 16, name char "Alexander" }` - создать именованный кортеж.

`boy` - прочитать переменную.

### Отношения

`RELATION { TUPLE { age integer 18, name char "Ivan" } }` - создать неименнованное отношение.

`RELATION { TUPLE { age integer 18, name char "Ivan" }, boy }` - создать неименнованное отношение, включив ранее заданный кортеж, указав его имя.

`boys := RELATION { TUPLE { age integer 16, name char "Alexander" }, TUPLE { age integer 13, name char "John" } }` - создать именованное отношение.

`boys` - прочитать переменную.

### Сокращение (WHERE)

`RELATION { TUPLE { age integer 16, name char "Alexander" }, TUPLE { age integer 13, name char "John" } } WHERE ( age >= 16 )` - произвести сокращение, создать новое неименованное отношение.

`boys WHERE ( age > 16 )` - произвести сокращение, создать новое неименованное отношение.

`boys := boys WHERE ( age > 16 )` - произвести сокращение, перезаписать отношение.

### Проекция (PROJECT)

`RELATION { TUPLE { age integer 16, name char "Alexander" }, TUPLE { age integer 13, name char "John" } } PROJECT (name)`

### Переименование (RENAME)

`RELATION { TUPLE { age integer 16, name char "Alexander" } RENAME (name AS fullname)`

### Пересечение (INTERSECT)

`boys INTERSECT RELATION { TUPLE { age integer 16, name char "Alexander" } }` - атрибуты отношений должны быть одинаковыми.

### Объединение (UNION)

`boys UNION RELATION { TUPLE { age integer 18, name char "Ivan" } }` - атрибуты отношений должны быть одинаковыми.

### Разность (MINUS)

`boys MINUS RELATION { TUPLE { age integer 16, name char "Alexander" } }` - найти все элементы первого отношения, которых нет во втором; атрибуты отношений должны быть одинаковыми.

### Декартово произведение (TIMES)

`boys TIMES { TUPLE { gender char "Man"} }` - добавляет к каждому элементу первого отношения каждый элемент второго; атрибуты отношений должны быть разыми.

### Натуральное соединение (JOIN)

`boys JOIN RELATION { TUPLE { age integer 16, gender char "Teenager" }, TUPLE { age integer 13, gender char "Boy" } }` - объдиняет те кортежи обоих отношений, которые имеют одинаковые значения для общих атрибутов.

## Обработка ошибок

При ошибке обычно выводится сообщение, в управляющую функцию возвращается `nil`, затем программа ожидает новый ввод.

* Проверяется соответствие значения заявленному типу.

* Проверяется что типы всех кортежей, входящих в отношение, одинаковы.

* Проверяется что множество кортежей, входящих в отношение, ненулевое.

* Проверяется, что аргументы в сокращении принадлежат одному типу.

* Проверяется, что оператор в сокращении применим к типу аргументов.

