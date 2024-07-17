# Имитация работы банкомата

Проект имитирует работу банкомата. Репозиторий представляет из себя мапу, где ключ это id пользователя, а значение указатель на структуру model.Account



## Эндпоинты
* POST /accounts - создание нового аккаунта.
* POST /accounts/{id}/deposit - пополнение баланса.
* POST /accounts/{id}/withdraw - снятие средств.
* GET /accounts/{id}/balance - проверка баланса.

P.S. В задании методы Deposit, Withdraw, GetBalance оперируют типом с плавающей точкой. Для точных вычислений необходимо использовать две переменные типа инт(условно рубли и копейки или доллары и центы)