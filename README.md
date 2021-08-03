# Golang kata

Просто набор kata для души и развлечений с применением подхода TDD.

# Примеры

- [Банкомат](atm/README.md)
- [Боулинг](bowling/README.md)
- [Лампа](lamp/)
- [Весы](libra/README.md)
- [Деньги](money/README.md)
- [DSL](dsl/README.md)
- [Очередь ожидания](wait_queue/README.md)

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