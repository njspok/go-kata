# Двухфазная фиксация транзакции

Облегченное моделирование двухфазной фиксации транзакции.

Допущения:
- связь между узлами идеальна и не дает сбоев
- узел может отказать в prepare операции
- узел всегда принимает операцию commit или abort и обязан ее выполнить