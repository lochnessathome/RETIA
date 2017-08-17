# RETIA
RETIA - акроним для RElaTIonal Algebra

## Примеры использования

`TUPLE { age integer 18, name char "Ivan" }`

`boy := TUPLE { age integer 16, name char "Alexander" }`

`RELATION { TUPLE { age integer 18, name char "Ivan" } }`

`boys := RELATIONS { TUPLE { age integer 16, name char "Alexander" }, TUPLE { age integer 13, name char "John" } }`

## Обработка ошибок

Проверяется соответствие значения заданному типу. При несоответствии выводит сообщение об ошибке, в управляющую функцию возвращается `nil`, затем программа ожидает новый ввод.


