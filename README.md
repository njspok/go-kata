# Golang kata

Просто набор kata для души и развлечений с применением подхода TDD.

# Примеры

- [Банкомат](atm/README.md)
- [Боулинг](bowling/README.md)
- [Лампа](lamp/)
- [Весы](libra/README.md)
- [Деньги](money/README.md)
- [DSL](dsl/README.md)

# Путь TDD

1. Добавить небольшой тест
2. Запустить тест и убедиться, что он терпит неудачу
3. Внести небольшое изменение в код
4. Снова запустить тесты и убедиться, что они прошли успешно
5. Устранить дублирование при помощи рефакторинга

Навыки

1. подходы, которые заставят тест работать быстро: заглушка, триангуляция и очевидная реализация
2. формирование дизайна через устранение дублирования между функциональным кодом и тестами
3. способность контролировать ширину шага - когда неуверенность возрастает, 
надо переходить на маленькие шажки, когда все становиться очевидным - увеличивать шаг
   
Ограничения
1. производительность
2. изоляция - тесты не должны влиять друг на друга
