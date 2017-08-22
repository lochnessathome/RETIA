# RETIA
RETIA - акроним для RElaTIonal Algebra


## Введение

Большая часть операций производится с отношениями (Relation; SQL аналог - таблица). Отношение - это множество кортежей (Tuple; SQL аналог - строка).

Реляционная алгебра предполагает 8 базовых операций над отношениями: объединение (union), пересечение (intersection), соединение (join), произведение (times), разность (minus), проекция (projection), сокращение (reduction) и переименование (rename).


### Как создать кортеж?

Кортеж используется исключительно как элемент отношения. Объявить его можно так:

`TUPLE { age integer 18, name char "Ivan" }`

Для удобства, ему можно дать имя:

`boy := TUPLE { age integer 16, name char "Alexander" }`

Затем можно прочитать переменную:

`boy`


### Компоненты и атрибуты кортежа

В записи `{ age integer 18, name char "Ivan" }` два компонента (разделены запятой). Компонент `age integer 18` делится, в свою очередь, на атрибут и значение. Атрибут - это пара имя-тип - `age integer`, значение - 18.


### Как создать отношение?

Можно создать отношение с одним элементом (кортежом):

`RELATION { TUPLE { age integer 18, name char "Ivan" } }`

Или:

`RELATION { boy }`

Часто бывает полезно (но не необходимо) дать отношению имя:

`boys := RELATION { TUPLE { age integer 16, name char "Alexander" }, TUPLE { age integer 13, name char "John" } }`

Так же, как и с кортежом, можно обратиться по имени:

`boys`


## Операции


### Объединение (union)

Вход:

`boys := RELATION { TUPLE { age integer 16, name char "Alexander" }, TUPLE { age integer 13, name char "John" } }`

`girls := RELATION { TUPLE { age integer 18, name char "Julia" } }`

`boys UNION girls`

Выход:

`
RELATION { 
         TUPLE { 
                 (age integer 16) 
                 (name char "Alexander") 
               } 
         TUPLE { 
                 (age integer 13) 
                 (name char "John") 
               } 
         TUPLE { 
                 (age integer 18) 
                 (name char "Julia") 
               } 
         }
`


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

## Вложеные запросы

Допустимо объединение нескольких операций в одном запросе:

`RELATION { TUPLE { age integer 16, name char "Alexander" } } UNION RELATION { TUPLE { age integer 13, name char "John" } } WHERE ( age > 13 )` - объединяет два отношения, фильтрует кортежи в соответствим с условием.

## Обработка ошибок

При ошибке обычно выводится сообщение, в управляющую функцию возвращается `nil`, затем программа ожидает новый ввод.

* Проверяется соответствие значения заявленному типу.

* Проверяется что типы всех кортежей, входящих в отношение, одинаковы.

* Проверяется что множество кортежей, входящих в отношение, ненулевое.

* Проверяется, что аргументы в сокращении принадлежат одному типу.

* Проверяется, что оператор в сокращении применим к типу аргументов.

