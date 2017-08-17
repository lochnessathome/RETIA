# RETIA
RETIA - акроним для RElaTIonal Algebra

## Примеры использования

`TUPLE { age integer 18, name char "Ivan" }`

`boy := TUPLE { age integer 16, name char "Alexander" }`

`boy`

`RELATION { TUPLE { age integer 18, name char "Ivan" } }`

`RELATION { TUPLE { age integer 18, name char "Ivan" }, boy }`

`boys := RELATIONS { TUPLE { age integer 16, name char "Alexander" }, TUPLE { age integer 13, name char "John" } }`

`boys`

## Обработка ошибок

При ошибке обычно выводится сообщение, в управляющую функцию возвращается `nil`, затем программа ожидает новый ввод.

* Проверяется соответствие значения заявленному типу.

* Проверяется что типы всех кортежей, входящих в отношение, одинаковы.

* Проверяется что множество кортежей, входящих в отношение, ненулевое.


