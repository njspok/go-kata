# Golang kata

Просто набор kata для души и развлечений с применением подхода TDD.

# Примеры

- [Банкомат](atm/README.md)
- [Боулинг](bowling/README.md)
- [Лампа](lamp/README.md)
- [Весы](libra/README.md)
- [Деньги](money/README.md)
- [DSL](dsl/README.md)
- [Очередь ожидания](wait_queue/README.md)
- [Печать отчета](report/README.md)
- [Транзакция](transaction/README.md)
- [Большие числа](big_numbers/README.md)
- [Саги](sagas/README.md)

# Путь TDD

Алгоритм

1. Добавить небольшой тест
2. Запустить тест и убедиться, что он терпит неудачу
3. Внести небольшое изменение в код
4. Снова запустить тесты и убедиться, что они прошли успешно
5. Устранить дублирование при помощи рефакторинга

Цикл

Красны - Зеленый - Рефакторинг

Навыки

1. подходы, которые заставят тест работать быстро: заглушка, триангуляция и очевидная реализация
2. формирование дизайна через устранение дублирования между функциональным кодом и тестами
3. способность контролировать ширину шага - когда неуверенность возрастает, 
надо переходить на маленькие шажки, когда все становиться очевидным - увеличивать шаг
   
Ограничения
1. производительность
2. изоляция - тесты не должны влиять друг на друга

# Паттерны

Шаблоны разработки через тестирование:

1. Тест
2. Изолированный тест
3. Список тестов
4. Сначала тест
5. Сначала оператор assert
6. Тестовые данные
7. Понятные данные

Шаблоны красной полосы:

1. Тест одного шага
2. Начальный тест
3. Объясняющий тест
4. Тест для изучения
5. Еще один тест
6. Регрессионный тест
7. Перерыв
8. Начать сначала
9. Дешевый стол, хорошие кресла

Шаблоны тестирования:

1. Дочерний тест - чтобы заставить работать большой тест, напишите
сначала маленький тест, представляющий часть большого
2. Mock object
3. Самошунтирование - чтобы убедиться, что один объект правильно 
взаимодействует с другим, заставьте контролируемый объект взаимодействовать
с тестом
4. Строка-журнал
5. Тестирование обработок ошибок
6. Оставляйте сломанный тест, если вы работаете один и вам будет
легко вспомнить на каком месте вы прервались в прошлый раз
7. Часто выпускаемый код

Шаблоны зеленой полосы:

1. Подделка - самая первая реализация функционала нужна, чтобы проверить сам тест, поэтому код
может быть настолько простым, чтобы проходил тест
2. Триангуляция
3. Очевидная реализация
4. От одного ко многим - для реализации работы с коллекцией, реализуйте сначала работу с одним элементом


# Литература

- Экстремальное программирование, Кент Бек