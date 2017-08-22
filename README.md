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

Предвариательно нужно создать два отношения:

`a := RELATION { TUPLE { age integer 16, name char "Alexander" }, TUPLE { age integer 13, name char "John" } }`

`b := RELATION { TUPLE { age integer 18, name char "Julia" }, TUPLE { age integer 16, name char "Alexander" } }`


### Объединение (union)

`a UNION b`

```
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
```


### Пересечение (intersect)

`a INTERSECT b`


```
RELATION {
         TUPLE { 
                 (age integer 16)
                 (name char "Alexander")
               } 
         }
```


### Натуральное соединение (join)

Создадим отношение с новыми атрибутами:

`c := RELATION { TUPLE { name char "Julia", gender char "Girl" }, TUPLE { name char "Alexander", gender char "Boy" }}`

Соединение:

`a JOIN c`

```
RELATION { 
         TUPLE { 
                 (age integer 16) 
                 (gender char "Boy") 
                 (name char "Alexander") 
               } 
         }
```

Мы потеряли Джулию, давайте вернём её в строй:

`a UNION b JOIN c`

```
RELATION { 
         TUPLE { 
                 (age integer 16) 
                 (gender char "Boy") 
                 (name char "Alexander") 
               } 
         TUPLE { 
                 (age integer 18) 
                 (gender char "Girl") 
                 (name char "Julia") 
               } 
         }
```


### Декартово произведение (times)

`b TIMES c`

```
RELATION { 
         TUPLE { 
                 (age integer 18) 
                 (gender char "Girl") 
                 (name char "Julia") 
               } 
         TUPLE { 
                 (age integer 18) 
                 (gender char "Boy") 
                 (name char "Julia") 
               } 
         TUPLE { 
                 (age integer 16) 
                 (gender char "Girl") 
                 (name char "Alexander") 
               } 
         TUPLE { 
                 (age integer 16) 
                 (gender char "Boy") 
                 (name char "Alexander") 
               } 
         }
```


### Разность (minus)

`a MINUS b`

```
RELATION { 
         TUPLE { 
                 (age integer 13) 
                 (name char "John") 
               } 
         } 
```

`b MINUS a`

```
RELATION { 
         TUPLE { 
                 (age integer 18) 
                 (name char "Julia") 
               } 
         }
```


### Сокращение (where)

`a WHERE ( age > 13 )`

```
RELATION { 
         TUPLE { 
                 (age integer 16) 
                 (name char "Alexander") 
               } 
         } 
```


### Проекция (project)

`a PROJECT ( name )`

```
RELATION { 
         TUPLE { 
                 (name char "Alexander") 
               } 
         TUPLE { 
                 (name char "John") 
               } 
         } 
```


### Переименование (rename)

`c RENAME ( gender AS sex )`

```
RELATION { 
         TUPLE { 
                 (name char "Julia") 
                 (sex char "Girl") 
               } 
         TUPLE { 
                 (name char "Alexander") 
                 (sex char "Boy") 
               } 
         } 
```


## Ограничения

* Значение должно соответстовать указанному типу. Приведения типов нет и не будет.

* Типы всех кортежей отношения одинаковы.

* Нельзя создать пустое отношение.

* Аргументы в операции сокращения должны быть одного типа.


## Детали реализации

При ошибке обычно выводится сообщение, в управляющую функцию возвращается `nil`, затем программа ожидает новый ввод.

